package main

import (
	"k8s.io/apimachinery/pkg/util/wait"
	"testing"
)

func TestWait(t *testing.T) {
	wait.Forever()
	wait.PollUntil()
	setupsignalhandler
}
