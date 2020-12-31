package mocks

import (
	openapi_v2 "github.com/googleapis/gnostic/openapiv2"
	"github.com/stretchr/testify/mock"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/version"
	restclient "k8s.io/client-go/rest"
)

type DiscoveryMock struct {
	mock.Mock
}

func (m *DiscoveryMock) RESTClient() restclient.Interface {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil
	} else {
		return i.(restclient.Interface)
	}
}

func (m *DiscoveryMock) ServerGroups() (*metav1.APIGroupList, error) {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil, args.Error(1)
	} else {
		return i.(*metav1.APIGroupList), args.Error(1)
	}
}
func (m *DiscoveryMock) ServerResourcesForGroupVersion(groupVersion string) (*metav1.APIResourceList, error) {
	args := m.Called(groupVersion)
	i := args.Get(0)
	if i == nil {
		return nil, args.Error(1)
	} else {
		return i.(*metav1.APIResourceList), args.Error(1)
	}
}
func (m *DiscoveryMock) ServerResources() ([]*metav1.APIResourceList, error) {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil, args.Error(1)
	} else {
		return i.([]*metav1.APIResourceList), args.Error(1)
	}
}
func (m *DiscoveryMock) ServerGroupsAndResources() ([]*metav1.APIGroup, []*metav1.APIResourceList, error) {
	args := m.Called()
	var r0 []*metav1.APIGroup
	i0 := args.Get(0)
	if i0 != nil {
		r0 = i0.([]*metav1.APIGroup)
	}
	var r1 []*metav1.APIResourceList
	i1 := args.Get(1)
	if i1 != nil {
		r1 = i1.([]*metav1.APIResourceList)
	}
	return r0, r1, args.Error(2)
}
func (m *DiscoveryMock) ServerPreferredResources() ([]*metav1.APIResourceList, error) {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil, args.Error(1)
	} else {
		return i.([]*metav1.APIResourceList), args.Error(1)
	}
}
func (m *DiscoveryMock) ServerPreferredNamespacedResources() ([]*metav1.APIResourceList, error) {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil, args.Error(1)
	} else {
		return i.([]*metav1.APIResourceList), args.Error(1)
	}
}
func (m *DiscoveryMock) ServerVersion() (*version.Info, error) {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil, args.Error(1)
	} else {
		return i.(*version.Info), args.Error(1)
	}
}
func (m *DiscoveryMock) OpenAPISchema() (*openapi_v2.Document, error) {
	args := m.Called()
	i := args.Get(0)
	if i == nil {
		return nil, args.Error(1)
	} else {
		return i.(*openapi_v2.Document), args.Error(1)
	}
}
