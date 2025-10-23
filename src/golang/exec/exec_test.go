package exec

import (
	"os/exec"
	"testing"
)

func Test_Exec(t *testing.T) {
	cmd := exec.Command("sleep", "10")

	startError := cmd.Start()
	if startError != nil {
		t.Fatalf("cmd start error: %v", startError)
	}

	waitError := cmd.Wait()
	if waitError != nil {
		t.Fatalf("cmd wait error: %v", waitError)
	}
}
