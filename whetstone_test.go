package whetstone_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"

	"github.com/cloudfoundry-incubator/runtime-schema/models/factories"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var (
	diegoEdgeCli string
	tmpDir       string
)

var _ = BeforeSuite(func() {
	tmpDir = os.TempDir()

	var err error
	diegoEdgeCli, err = gexec.Build("github.com/pivotal-cf-experimental/diego-edge-cli")
	Expect(err).ToNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

var _ = Describe("Diego Edge", func() {
	Context("when desiring a docker-based LRP", func() {

		var (
			appName string
			route   string
		)

		BeforeEach(func() {
			appName = fmt.Sprintf("whetstone-%s", factories.GenerateGuid())
			route = fmt.Sprintf("%s.%s", appName, domain)

			targetApi(receptorUrl)
			targetLoggregator(loggregatorAddress)
		})

		AfterEach(func() {
			stopApp(appName)

			Eventually(errorCheckForRoute(route), timeout, 1).Should(HaveOccurred())
		})

		It("eventually runs a docker app", func() {
			startDockerApp(appName)
			Eventually(errorCheckForRoute(route), timeout, 1).ShouldNot(HaveOccurred())

			logsStream := streamLogs(appName)
			Eventually(logsStream.Out, timeout).Should(gbytes.Say("Diego Edge Docker App. Says Hello"))

			scaleApp(appName)

			instanceCountChan := make(chan int, numCpu)
			go countInstances(route, instanceCountChan)
			Eventually(instanceCountChan, timeout).Should(Receive(Equal(3)))

			logsStream.Terminate().Wait()
		})
	})

})

func startDockerApp(appName string) {
	command := command(diegoEdgeCli, "start", appName, "-i", "docker:///diegoedge/diego-edge-docker-app", "-c", "/dockerapp")
	session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)

	Expect(err).ToNot(HaveOccurred())
	Eventually(session).Should(gexec.Exit(0))
}

func streamLogs(appName string) *gexec.Session {
	command := command(diegoEdgeCli, "logs", appName)
	session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
	Expect(err).ToNot(HaveOccurred())

	return session
}

func scaleApp(appName string) {
	command := command(diegoEdgeCli, "scale", appName, "--instances", "3")
	session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)

	Expect(err).ToNot(HaveOccurred())
	Eventually(session).Should(gexec.Exit(0))
}

func stopApp(appName string) {
	command := command(diegoEdgeCli, "stop", appName)
	session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)

	Expect(err).ToNot(HaveOccurred())
	Eventually(session).Should(gexec.Exit(0))
}

func targetApi(receptorUrl string) {
	command := command(diegoEdgeCli, "target", receptorUrl)
	session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)

	Expect(err).ToNot(HaveOccurred())
	Eventually(session).Should(gexec.Exit(0))
}

func targetLoggregator(loggregatorAddress string) {
	command := command(diegoEdgeCli, "target-loggregator", loggregatorAddress)
	session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)

	Expect(err).ToNot(HaveOccurred())
	Eventually(session).Should(gexec.Exit(0))
}

func command(name string, arg ...string) *exec.Cmd {
	command := exec.Command(name, arg...)
	command.Env = []string{fmt.Sprintf("DIEGO_CLI_HOME=%s", tmpDir)}
	return command
}

func errorCheckForRoute(route string) func() error {
	return func() error {
		response, err := makeGetRequestToRoute(route)
		if err != nil {
			return err
		}

		io.Copy(ioutil.Discard, response.Body)
		defer response.Body.Close()

		if response.StatusCode != 200 {
			return fmt.Errorf("Status code %d should be 200", response.StatusCode)
		}

		return nil
	}
}

func countInstances(route string, instanceCountChan chan<- int) {
	defer GinkgoRecover()
	instanceIndexRoute := fmt.Sprintf("%s/instance-index", route)
	instancesSeen := make(map[int]bool)

	instanceIndexChan := make(chan int, numCpu)

	for i := 0; i < numCpu; i++ {
		go pollForInstanceIndices(instanceIndexRoute, instanceIndexChan)
	}

	for {
		instanceIndex := <-instanceIndexChan
		instancesSeen[instanceIndex] = true
		instanceCountChan <- len(instancesSeen)
	}
}

func pollForInstanceIndices(route string, instanceIndexChan chan<- int) {
	defer GinkgoRecover()
	for {
		response, err := makeGetRequestToRoute(route)
		Expect(err).To(BeNil())

		responseBody, err := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		Expect(err).To(BeNil())

		instanceIndex, err := strconv.Atoi(string(responseBody))
		if err != nil {
			continue
		}
		instanceIndexChan <- instanceIndex
	}
}

func makeGetRequestToRoute(route string) (*http.Response, error) {
	routeWithScheme := fmt.Sprintf("http://%s", route)
	resp, err := http.DefaultClient.Get(routeWithScheme)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
