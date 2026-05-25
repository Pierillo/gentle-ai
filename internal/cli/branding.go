package cli

import "github.com/gentleman-programming/gentle-ai/internal/model"

func effectiveCommandName() string {
	if model.IsCopilotOnlyBuild() {
		return "pierillo-antislop-boring-toolkit"
	}
	return "gentle-ai"
}
