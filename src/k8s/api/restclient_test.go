package api

import (
	"context"
	"fmt"
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

func Test_RestClient(t *testing.T) {
	config := &rest.Config{
		Host:    "127.0.0.1:8080",
		APIPath: "api",
		ContentConfig: rest.ContentConfig{
			GroupVersion:         &corev1.SchemeGroupVersion,
			NegotiatedSerializer: scheme.Codecs,
		},
	}

	restClient, err := rest.RESTClientFor(config)

	if err != nil {
		panic(err.Error())
	}

	result := &corev1.PodList{}

	namespace := "kube-system"
	err = restClient.Get().
		Namespace(namespace).
		Resource("pods").
		VersionedParams(&metav1.ListOptions{Limit: 100}, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(result)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("namespace\t status\t\t name\n")

	// 每个pod都打印namespace、status.Phase、name三个字段
	for _, d := range result.Items {
		fmt.Printf("%v\t %v\t %v\n",
			d.Namespace,
			d.Status.Phase,
			d.Name)
	}
}
