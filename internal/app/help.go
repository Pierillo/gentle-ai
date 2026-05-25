package app

import (
	"fmt"
	"io"
	"strings"

	"github.com/gentleman-programming/gentle-ai/internal/model"
)

func printHelp(w io.Writer, version string) {
	command := effectiveCommandName()
	commands := `  install      Configure AI coding agents on this machine
  uninstall    Remove Gentle AI managed files from this machine
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

	fmt.Fprintf(w, `%s — Gentle-AI: Ecosystem, Frameworks, Workflows (%s)

USAGE
  %s                     Launch interactive TUI
  %s <command> [flags]

COMMANDS
%s

FLAGS
  --help, -h    Show this help

Run '%s help' for this message.
Documentation: https://github.com/Gentleman-Programming/gentle-ai
`, command, version, command, command, commands, command)
}
