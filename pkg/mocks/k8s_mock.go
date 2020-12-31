package mocks

import (
	"github.com/stretchr/testify/mock"
	"k8s.io/client-go/discovery"
	admissionregistrationv1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1"
	admissionregistrationv1beta1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1"
	internalv1alpha1 "k8s.io/client-go/kubernetes/typed/apiserverinternal/v1alpha1"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	appsv1beta1 "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	appsv1beta2 "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	authenticationv1 "k8s.io/client-go/kubernetes/typed/authentication/v1"
	authenticationv1beta1 "k8s.io/client-go/kubernetes/typed/authentication/v1beta1"
	authorizationv1 "k8s.io/client-go/kubernetes/typed/authorization/v1"
	authorizationv1beta1 "k8s.io/client-go/kubernetes/typed/authorization/v1beta1"
	autoscalingv1 "k8s.io/client-go/kubernetes/typed/autoscaling/v1"
	autoscalingv2beta1 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta1"
	autoscalingv2beta2 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta2"
	batchv1 "k8s.io/client-go/kubernetes/typed/batch/v1"
	batchv1beta1 "k8s.io/client-go/kubernetes/typed/batch/v1beta1"
	batchv2alpha1 "k8s.io/client-go/kubernetes/typed/batch/v2alpha1"
	certificatesv1 "k8s.io/client-go/kubernetes/typed/certificates/v1"
	certificatesv1beta1 "k8s.io/client-go/kubernetes/typed/certificates/v1beta1"
	coordinationv1 "k8s.io/client-go/kubernetes/typed/coordination/v1"
	coordinationv1beta1 "k8s.io/client-go/kubernetes/typed/coordination/v1beta1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	discoveryv1alpha1 "k8s.io/client-go/kubernetes/typed/discovery/v1alpha1"
	discoveryv1beta1 "k8s.io/client-go/kubernetes/typed/discovery/v1beta1"
	eventsv1 "k8s.io/client-go/kubernetes/typed/events/v1"
	eventsv1beta1 "k8s.io/client-go/kubernetes/typed/events/v1beta1"
	extensionsv1beta1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	flowcontrolv1alpha1 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1"
	flowcontrolv1beta1 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta1"
	networkingv1 "k8s.io/client-go/kubernetes/typed/networking/v1"
	networkingv1beta1 "k8s.io/client-go/kubernetes/typed/networking/v1beta1"
	nodev1 "k8s.io/client-go/kubernetes/typed/node/v1"
	nodev1alpha1 "k8s.io/client-go/kubernetes/typed/node/v1alpha1"
	nodev1beta1 "k8s.io/client-go/kubernetes/typed/node/v1beta1"
	policyv1beta1 "k8s.io/client-go/kubernetes/typed/policy/v1beta1"
	rbacv1 "k8s.io/client-go/kubernetes/typed/rbac/v1"
	rbacv1alpha1 "k8s.io/client-go/kubernetes/typed/rbac/v1alpha1"
	rbacv1beta1 "k8s.io/client-go/kubernetes/typed/rbac/v1beta1"
	schedulingv1 "k8s.io/client-go/kubernetes/typed/scheduling/v1"
	schedulingv1alpha1 "k8s.io/client-go/kubernetes/typed/scheduling/v1alpha1"
	schedulingv1beta1 "k8s.io/client-go/kubernetes/typed/scheduling/v1beta1"
	storagev1 "k8s.io/client-go/kubernetes/typed/storage/v1"
	storagev1alpha1 "k8s.io/client-go/kubernetes/typed/storage/v1alpha1"
	storagev1beta1 "k8s.io/client-go/kubernetes/typed/storage/v1beta1"
)

type K8sMock struct {
	mock.Mock
}

func (m *K8sMock) Discovery() discovery.DiscoveryInterface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(discovery.DiscoveryInterface)
	}
}

func (m *K8sMock) AdmissionregistrationV1() admissionregistrationv1.AdmissionregistrationV1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(admissionregistrationv1.AdmissionregistrationV1Interface)
	}
}

func (m *K8sMock) AdmissionregistrationV1beta1() admissionregistrationv1beta1.AdmissionregistrationV1beta1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(admissionregistrationv1beta1.AdmissionregistrationV1beta1Interface)
	}
}

func (m *K8sMock) InternalV1alpha1() internalv1alpha1.InternalV1alpha1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(internalv1alpha1.InternalV1alpha1Interface)
	}
}

func (m *K8sMock) AppsV1() appsv1.AppsV1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(appsv1.AppsV1Interface)
	}
}

func (m *K8sMock) AppsV1beta1() appsv1beta1.AppsV1beta1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(appsv1beta1.AppsV1beta1Interface)
	}
}

func (m *K8sMock) AppsV1beta2() appsv1beta2.AppsV1beta2Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(appsv1beta2.AppsV1beta2Interface)
	}
}

func (m *K8sMock) AuthenticationV1() authenticationv1.AuthenticationV1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(authenticationv1.AuthenticationV1Interface)
	}
}

func (m *K8sMock) AuthenticationV1beta1() authenticationv1beta1.AuthenticationV1beta1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(authenticationv1beta1.AuthenticationV1beta1Interface)
	}
}

func (m *K8sMock) AuthorizationV1() authorizationv1.AuthorizationV1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(authorizationv1.AuthorizationV1Interface)
	}
}

func (m *K8sMock) AuthorizationV1beta1() authorizationv1beta1.AuthorizationV1beta1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(authorizationv1beta1.AuthorizationV1beta1Interface)
	}
}

func (m *K8sMock) AutoscalingV1() autoscalingv1.AutoscalingV1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(autoscalingv1.AutoscalingV1Interface)
	}
}

func (m *K8sMock) AutoscalingV2beta1() autoscalingv2beta1.AutoscalingV2beta1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(autoscalingv2beta1.AutoscalingV2beta1Interface)
	}
}

func (m *K8sMock) AutoscalingV2beta2() autoscalingv2beta2.AutoscalingV2beta2Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(autoscalingv2beta2.AutoscalingV2beta2Interface)
	}
}

func (m *K8sMock) BatchV1() batchv1.BatchV1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(batchv1.BatchV1Interface)
	}
}

func (m *K8sMock) BatchV1beta1() batchv1beta1.BatchV1beta1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(batchv1beta1.BatchV1beta1Interface)
	}
}
func (m *K8sMock) BatchV2alpha1() batchv2alpha1.BatchV2alpha1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(batchv2alpha1.BatchV2alpha1Interface)
	}
}
func (m *K8sMock) CertificatesV1() certificatesv1.CertificatesV1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(certificatesv1.CertificatesV1Interface)
	}
}
func (m *K8sMock) CertificatesV1beta1() certificatesv1beta1.CertificatesV1beta1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(certificatesv1beta1.CertificatesV1beta1Interface)
	}
}
func (m *K8sMock) CoordinationV1beta1() coordinationv1beta1.CoordinationV1beta1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(coordinationv1beta1.CoordinationV1beta1Interface)
	}
}
func (m *K8sMock) CoordinationV1() coordinationv1.CoordinationV1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(coordinationv1.CoordinationV1Interface)
	}
}
func (m *K8sMock) CoreV1() corev1.CoreV1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(corev1.CoreV1Interface)
	}
}
func (m *K8sMock) DiscoveryV1alpha1() discoveryv1alpha1.DiscoveryV1alpha1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(discoveryv1alpha1.DiscoveryV1alpha1Interface)
	}
}
func (m *K8sMock) DiscoveryV1beta1() discoveryv1beta1.DiscoveryV1beta1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(discoveryv1beta1.DiscoveryV1beta1Interface)
	}
}
func (m *K8sMock) EventsV1() eventsv1.EventsV1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(eventsv1.EventsV1Interface)
	}
}
func (m *K8sMock) EventsV1beta1() eventsv1beta1.EventsV1beta1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(eventsv1beta1.EventsV1beta1Interface)
	}
}
func (m *K8sMock) ExtensionsV1beta1() extensionsv1beta1.ExtensionsV1beta1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(extensionsv1beta1.ExtensionsV1beta1Interface)
	}
}
func (m *K8sMock) FlowcontrolV1alpha1() flowcontrolv1alpha1.FlowcontrolV1alpha1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(flowcontrolv1alpha1.FlowcontrolV1alpha1Interface)
	}
}
func (m *K8sMock) FlowcontrolV1beta1() flowcontrolv1beta1.FlowcontrolV1beta1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(flowcontrolv1beta1.FlowcontrolV1beta1Interface)
	}
}
func (m *K8sMock) NetworkingV1() networkingv1.NetworkingV1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(networkingv1.NetworkingV1Interface)
	}
}
func (m *K8sMock) NetworkingV1beta1() networkingv1beta1.NetworkingV1beta1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(networkingv1beta1.NetworkingV1beta1Interface)
	}
}
func (m *K8sMock) NodeV1() nodev1.NodeV1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(nodev1.NodeV1Interface)
	}
}
func (m *K8sMock) NodeV1alpha1() nodev1alpha1.NodeV1alpha1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(nodev1alpha1.NodeV1alpha1Interface)
	}
}
func (m *K8sMock) NodeV1beta1() nodev1beta1.NodeV1beta1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(nodev1beta1.NodeV1beta1Interface)
	}
}
func (m *K8sMock) PolicyV1beta1() policyv1beta1.PolicyV1beta1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(policyv1beta1.PolicyV1beta1Interface)
	}
}
func (m *K8sMock) RbacV1() rbacv1.RbacV1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(rbacv1.RbacV1Interface)
	}
}
func (m *K8sMock) RbacV1beta1() rbacv1beta1.RbacV1beta1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(rbacv1beta1.RbacV1beta1Interface)
	}
}
func (m *K8sMock) RbacV1alpha1() rbacv1alpha1.RbacV1alpha1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(rbacv1alpha1.RbacV1alpha1Interface)
	}
}
func (m *K8sMock) SchedulingV1alpha1() schedulingv1alpha1.SchedulingV1alpha1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(schedulingv1alpha1.SchedulingV1alpha1Interface)
	}
}
func (m *K8sMock) SchedulingV1beta1() schedulingv1beta1.SchedulingV1beta1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(schedulingv1beta1.SchedulingV1beta1Interface)
	}
}
func (m *K8sMock) SchedulingV1() schedulingv1.SchedulingV1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(schedulingv1.SchedulingV1Interface)
	}
}
func (m *K8sMock) StorageV1beta1() storagev1beta1.StorageV1beta1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(storagev1beta1.StorageV1beta1Interface)
	}
}
func (m *K8sMock) StorageV1() storagev1.StorageV1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(storagev1.StorageV1Interface)
	}
}
func (m *K8sMock) StorageV1alpha1() storagev1alpha1.StorageV1alpha1Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(storagev1alpha1.StorageV1alpha1Interface)
	}
}
