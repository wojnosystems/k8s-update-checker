package helpers

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CollectNamespaces(ctx context.Context, k kubernetes.Interface) (namespaces []string, err error) {
	recordSetListOptions := metav1.ListOptions{
		Limit: 1000,
	}
	for {
		var namespaceRecordSet *corev1.NamespaceList
		namespaceRecordSet, err = k.CoreV1().Namespaces().List(ctx, recordSetListOptions)
		if err != nil {
			return
		}
		for _, item := range namespaceRecordSet.Items {
			namespaces = append(namespaces, item.Namespace)
		}
		if len(namespaceRecordSet.Continue) == 0 {
			return
		}
		recordSetListOptions.Continue = namespaceRecordSet.Continue
	}
	return
}
