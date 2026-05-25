package app

import (
	"fmt"
	"io"
	"strings"

	"github.com/gentleman-programming/gentle-ai/internal/model"
)

func printHelp(w io.Writer, version string) {
	command := effectiveCommandName()
	productLabel := "Gentle-AI: Ecosystem, Frameworks, Workflows"
	docsURL := "https://github.com/Gentleman-Programming/gentle-ai"
	if model.IsCopilotOnlyBuild() {
		productLabel = "Pierillo AntiSlop Boring Toolkit for GitHub Copilot CLI"
		docsURL = "https://github.com/Pierillo/pierillo-antislop-boring-toolkit"
	}
	commands := `  install      Configure AI coding agents on this machine
  uninstall    Remove managed files from this machine
  sync         Sync agent configs and skills to current version
  skill-registry refresh
               Refresh .atl/skill-registry.md with cache-hit fast path
  restore      Restore a config backup
  version      Print version`
	if !model.IsCopilotOnlyBuild() {
		commands = strings.Join([]string{
			commands,
			"  update       Check for available updates",
			"  upgrade      Apply updates to managed tools",
		}, "\n")
	}

	fmt.Fprintf(w, `%s — %s (%s)

USAGE
  %s                     Launch interactive TUI
  %s <command> [flags]

COMMANDS
%s

FLAGS
  --help, -h    Show this help

Run '%s help' for this message.
Documentation: %s
`, command, productLabel, version, command, command, commands, command, docsURL)
}
