package k8s

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os/exec"
	"regexp"
)

const (
	miniKubeExitCodeFound    = 0
	miniKubeExitCodePaused   = 2
	miniKubeExitCodeNotFound = 85

	defaultMiniKubeName = "wojnosystems-k8s-updatechecker-ruletest"
)

type k8sStatuses uint8

const (
	k8sStatusInvalid k8sStatuses = iota
	k8sStatusNotExists
	k8sStatusNotRunning
	k8sStatusPaused
	k8sStatusRunning
	k8sStatusProfileNotFound
	k8sStatusEnd
)

type miniKube struct {
	profile string
	k8s     kubernetes.Interface
}

func NewMiniKube() (mk *miniKube, err error) {
	mk = &miniKube{
		profile: "",
	}
	err = mk.miniKubeStartOrUnPauseCluster()
	return
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

type miniKubeData struct {
	CurrentStep string `json:"currentstep"`
	Message     string `json:"message"`
	Name        string `json:"name"`
	TotalSteps  string `json:"totalsteps"`
}

var profileNotFoundRegexp = regexp.MustCompile(`Profile "[^"]+" not found.`)

func (m miniKubeData) isProfileNotFound() bool {
	return profileNotFoundRegexp.MatchString(m.Message)
}

type miniKubeError struct {
	Data            miniKubeData `json:"data"`
	DataContentType string       `json:"datacontenttype"`
	Id              string       `json:"id"`
	Source          string       `json:"source"`
	SpecVersion     string       `json:"specversion"`
	Type            string       `json:"type"`
}

func (s miniKubeStatus) isAPIServerRunning() bool {
	return s.APIServer == "Running"
}

// Harness that uses MiniKube to run the tests against the cluster.
func (m *miniKube) miniKubeStartOrUnPauseCluster() (err error) {
	status, err := m.currentK8sStatus(m.profile)
	if err != nil {
		return
	}
	switch status {
	case k8sStatusNotRunning:
		startCmd := exec.Command("minikube", "start",
			profileArg(m.profile))
		err = startCmd.Run()
	case k8sStatusNotExists, k8sStatusProfileNotFound:
		startCmd := exec.Command("minikube", "start",
			profileArg(m.profile),
			fmt.Sprintf(`--kubernetes-version=v%s`, CurrentVersion.String()))
		err = startCmd.Run()
	case k8sStatusPaused:
		startCmd := exec.Command("minikube", "unpause",
			profileArg(m.profile))
		err = startCmd.Run()
	case k8sStatusRunning:
		// do nothing
	default:
		err = errors.New("unknown kubernetes state, i don't know what to do")
	}

	_, m.k8s, err = getKubeClient(m.profile)
	return
}

func (m *miniKube) currentK8sStatus(profile string) (status k8sStatuses, err error) {
	statusCmd := exec.Command("minikube", "status",
		profileArg(profile),
		"-o", "json")
	var statusOutput io.ReadCloser
	statusOutput, err = statusCmd.StdoutPipe()
	if err != nil {
		return
	}

	err = statusCmd.Start()
	defer func() {
		_ = statusCmd.Wait()
	}()
	if err != nil {
		var exitErr *exec.ExitError
		if !errors.As(err, &exitErr) {
			// black hole all errors other than exit error
			return
		}

		switch exitErr.ExitCode() {
		case miniKubeExitCodeFound:
			status = k8sStatusNotRunning
		case miniKubeExitCodePaused:
			status = k8sStatusPaused
		default:
			err = fmt.Errorf("unknown return code from minikube: %d", exitErr.ExitCode())
			return
		}
	} else {
		var fullResponse []byte
		fullResponse, err = ioutil.ReadAll(statusOutput)
		if err != nil {
			return
		}
		bufferedOutput := bytes.NewReader(fullResponse)
		statusMsg := miniKubeStatus{}
		decoder := json.NewDecoder(bufferedOutput)
		decoder.DisallowUnknownFields()
		err = decoder.Decode(&statusMsg)
		if err == nil {
			switch statusMsg.APIServer {
			case "Running":
				// do nothing
			case "Stopped":
				status = k8sStatusNotRunning
			case "Paused":
				status = k8sStatusPaused
			}
		} else {
			bufferedOutput = bytes.NewReader(fullResponse)
			decoder := json.NewDecoder(bufferedOutput)
			decoder.DisallowUnknownFields()
			mkError := miniKubeError{}
			err = decoder.Decode(&mkError)
			if err != nil {
				return
			}
			if mkError.Data.isProfileNotFound() {
				status = k8sStatusNotExists
			}
		}
	}
	return
}

// GetKubeClient creates a Kubernetes config and client for a given kubeconfig context.
func getKubeClient(context string) (*rest.Config, kubernetes.Interface, error) {
	config, err := configForContext(context)
	if err != nil {
		return nil, nil, err
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, nil, fmt.Errorf("could not get Kubernetes client: %s", err)
	}
	return config, client, nil
}

// configForContext creates a Kubernetes REST client configuration for a given kubeconfig context.
func configForContext(context string) (*rest.Config, error) {
	config, err := getConfig(context).ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("could not get Kubernetes config for context %q: %s", context, err)
	}
	return config, nil
}

// getConfig returns a Kubernetes client config for a given context.
func getConfig(context string) clientcmd.ClientConfig {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	rules.DefaultClientConfig = &clientcmd.DefaultClientConfig

	overrides := &clientcmd.ConfigOverrides{
		ClusterDefaults: clientcmd.ClusterDefaults,
		Timeout:         "30s",
	}

	if context != "" {
		overrides.CurrentContext = context
	}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, overrides)
}

func (m *miniKube) K8s() kubernetes.Interface {
	return m.k8s
}

func (m *miniKube) Stop() error {
	stopCmd := exec.Command("minikube", "stop",
		profileArg(m.profile))
	return stopCmd.Run()
}
