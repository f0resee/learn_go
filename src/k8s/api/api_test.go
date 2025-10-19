package api

import (
	"fmt"
	"testing"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Test_api_Deployment(t *testing.T) {
	dep := appsv1.Deployment{
		Spec: appsv1.DeploymentSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					NodeName: "test-node",
				},
			},
		},
	}
	fmt.Printf("deployment: %#v\n", dep)
}

func Test_api_machi(t *testing.T) {
	meta := metav1.TypeMeta{
		Kind: "pod",
	}
	fmt.Printf("type meta: %#v\n", meta)
}
