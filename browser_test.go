package browser

import (
	"errors"
	"os/exec"
	"testing"
)

// TestRunCmdMissingBinary exercises the one piece of logic shared by every
// platform's openBrowser: the binary lookup fails before anything is
// executed, so this is safe and deterministic to run in CI without a
// display or a real browser.
func TestRunCmdMissingBinary(t *testing.T) {
	err := runCmd("browser-test-nonexistent-binary")
	if !errors.Is(err, exec.ErrNotFound) {
		t.Fatalf("runCmd() error = %v, want wrapped %v", err, exec.ErrNotFound)
	}
}
