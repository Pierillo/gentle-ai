package app

import "strings"

// CommandName is the CLI command name printed in help/errors.
// Default remains "gentle-ai" for backward compatibility.
var CommandName = "gentle-ai"

func effectiveCommandName() string {
	name := strings.TrimSpace(CommandName)
	if name == "" {
		return "gentle-ai"
	}
	return name
}

