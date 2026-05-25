package cli

import (
	"bytes"
	"strings"
	"testing"

	"github.com/gentleman-programming/gentle-ai/internal/model"
)

func TestRunUninstallWithoutFlagsPrintsUsage(t *testing.T) {
	var buf bytes.Buffer
	_, err := RunUninstall(nil, &buf)
	if err == nil {
		t.Fatal("expected error when no flags provided")
	}
}

func TestPromptUninstallConfirmUsesCopilotCommandNameInCopilotOnlyBuild(t *testing.T) {
	prev := model.BuildFlavor
	model.BuildFlavor = "copilot-only"
	t.Cleanup(func() {
		model.BuildFlavor = prev
	})

	var out bytes.Buffer
	confirmed, err := promptUninstallConfirm(UninstallFlags{All: true}, &out, strings.NewReader("yes\n"))
	if err != nil {
		t.Fatalf("promptUninstallConfirm() error = %v", err)
	}
	if !confirmed {
		t.Fatal("promptUninstallConfirm() should confirm on yes input")
	}
	if !strings.Contains(out.String(), "pierillo-antislop-boring-toolkit managed configuration") {
		t.Fatalf("prompt should mention pierillo-antislop-boring-toolkit branding; got:\n%s", out.String())
	}
}
