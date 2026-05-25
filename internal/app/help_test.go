package app

import (
	"bytes"
	"strings"
	"testing"

	"github.com/gentleman-programming/gentle-ai/internal/model"
)

func TestHelpContainsAllCommands(t *testing.T) {
	var buf bytes.Buffer
	printHelp(&buf, "v1.0.0-test")
	output := buf.String()

	commands := []string{"install", "uninstall", "sync", "update", "upgrade", "restore", "version"}
	for _, cmd := range commands {
		if !strings.Contains(output, cmd) {
			t.Errorf("help output missing command %q", cmd)
		}
	}
}

func TestHelpContainsVersion(t *testing.T) {
	var buf bytes.Buffer
	printHelp(&buf, "v1.2.3")
	if !strings.Contains(buf.String(), "v1.2.3") {
		t.Error("help output should contain the version string")
	}
}

func TestHelpCommandsHeadingIsAligned(t *testing.T) {
	var buf bytes.Buffer
	printHelp(&buf, "v1.2.3")
	if !strings.Contains(buf.String(), "\nCOMMANDS\n  install") {
		t.Fatalf("help output has inconsistent command indentation:\n%s", buf.String())
	}
}

func TestHelpOmitsUpdateCommandsInCopilotOnlyBuild(t *testing.T) {
	prev := model.BuildFlavor
	model.BuildFlavor = "copilot-only"
	t.Cleanup(func() {
		model.BuildFlavor = prev
	})

	var buf bytes.Buffer
	printHelp(&buf, "v1.2.3")

	output := buf.String()
	if strings.Contains(output, "  update") {
		t.Fatalf("help output should omit update in copilot-only build:\n%s", output)
	}
	if strings.Contains(output, "  upgrade") {
		t.Fatalf("help output should omit upgrade in copilot-only build:\n%s", output)
	}
}
