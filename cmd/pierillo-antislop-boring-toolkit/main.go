package main

import (
	"fmt"
	"os"

	"github.com/gentleman-programming/gentle-ai/internal/app"
	"github.com/gentleman-programming/gentle-ai/internal/model"
)

// version is set by GoReleaser via ldflags at build time.
var version = "dev"

func main() {
	app.CommandName = "pierillo-antislop-boring-toolkit"
	model.BuildFlavor = "copilot-only"
	app.Version = app.ResolveVersion(version)

	if err := app.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
