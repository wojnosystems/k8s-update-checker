package rule_testing

import (
	"errors"
	"fmt"
	"github.com/wojnosystems/k8s-update-checker/pkg/k8s"
	"io"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/json"
	"os/exec"
)

const (
	miniKubeExitCodeFound    = 0
	miniKubeExitCodePaused   = 2
	miniKubeExitCodeNotFound = 85

	defaultMiniKubeName = "wojnosystems-k8s-updatechecker-ruletest"
)

type miniKube struct {
	profile string
}

func NewMiniKube(profile string) (h *miniKube, err error) {

	// Create K8s
	// see if it's running
	_, err = miniKubeStartOrUnPauseCluster("")
	if err != nil {
		return
	}

	// use the current context in kubeconfig
	//config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	//if err != nil {
	//	panic(err.Error())
	//}
	//
	//
	//// create the clientset
	//clientset, err := kubernetes.NewForConfig(config)
	//if err != nil {
	//	panic(err.Error())
	//}

	return &miniKube{
		profile: profile,
	}, nil
}

func profileArg(profile string) string {
	if profile == "" {
		profile = defaultMiniKubeName
	}
	return fmt.Sprintf(`--profile=%s`, profile)
}

type miniKubeStatus struct {
	Name       string
	Host       string
	Kubelet    string
	APIServer  string
	Kubeconfig string
	Worker     bool
	TimeToStop string
}

func (m miniKubeStatus) isAPIServerRunning() bool {
	return m.APIServer == "Running"
}

// Harness that uses MiniKube to run the tests against the cluster.
func miniKubeStartOrUnPauseCluster(profile string) (k k8s.Type, err error) {
	statusCmd := exec.Command("minikube", "status",
		profileArg(profile),
		"-o", "json")
	var statusOutput io.ReadCloser
	statusOutput, err = statusCmd.StdoutPipe()
	if err != nil {
		return
	}

	err = statusCmd.Start()
	if err != nil {
		_ = statusCmd.Wait()

		var exitErr *exec.ExitError
		if !errors.As(err, &exitErr) {
			// black hole all errors other than exit error
			return
		}

		switch exitErr.ExitCode() {
		case miniKubeExitCodeFound:
			startCmd := exec.Command("minikube", "start",
				profileArg(profile),
				fmt.Sprintf(`--kubernetes-version=%s`, "v"+k8s.CurrentVersion.String()))
			err = startCmd.Run()
			if err != nil {
				return
			}
		case miniKubeExitCodePaused:
			startCmd := exec.Command("minikube", "unpause",
				profileArg(profile))
			err = startCmd.Run()
			return
		default:
			err = fmt.Errorf("unknown return code from minikube: %d", exitErr.ExitCode())
		}
	} else {
		defer func() {
			_ = statusCmd.Wait()
		}()
		var response []byte
		response, err = ioutil.ReadAll(statusOutput)
		if err != nil {
			return
		}
		var status miniKubeStatus
		err = json.Unmarshal(response, &status)
		if err != nil {
			return
		}
	}

	return
}

func (k *miniKube) K8s() k8s.Type {
	return nil
}

func (k *miniKube) Close() error {
	stopCmd := exec.Command("minikube", "stop",
		profileArg(k.profile))
	return stopCmd.Run()
}
