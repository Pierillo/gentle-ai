package copilotcli

import (
	"context"
	"errors"
	"path/filepath"
	"testing"

	"github.com/gentleman-programming/gentle-ai/internal/model"
	"github.com/gentleman-programming/gentle-ai/internal/system"
)

func TestDetect(t *testing.T) {
	tests := []struct {
		name            string
		lookPathPath    string
		lookPathErr     error
		wantInstalled   bool
		wantBinaryPath  string
		wantConfigPath  string
		wantConfigFound bool
	}{
		{
			name:            "binary found",
			lookPathPath:    "/usr/local/bin/copilot",
			wantInstalled:   true,
			wantBinaryPath:  "/usr/local/bin/copilot",
			wantConfigPath:  "",
			wantConfigFound: true,
		},
		{
			name:            "binary missing",
			lookPathErr:     errors.New("missing"),
			wantInstalled:   false,
			wantBinaryPath:  "",
			wantConfigPath:  "",
			wantConfigFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				lookPath: func(string) (string, error) {
					return tt.lookPathPath, tt.lookPathErr
				},
			}

			installed, binaryPath, configPath, configFound, err := a.Detect(context.Background(), "/tmp/home")
			if err != nil {
				t.Fatalf("Detect() error = %v, want nil", err)
			}
			if installed != tt.wantInstalled {
				t.Fatalf("Detect() installed = %v, want %v", installed, tt.wantInstalled)
			}
			if binaryPath != tt.wantBinaryPath {
				t.Fatalf("Detect() binaryPath = %q, want %q", binaryPath, tt.wantBinaryPath)
			}
			if configPath != tt.wantConfigPath {
				t.Fatalf("Detect() configPath = %q, want %q", configPath, tt.wantConfigPath)
			}
			if configFound != tt.wantConfigFound {
				t.Fatalf("Detect() configFound = %v, want %v", configFound, tt.wantConfigFound)
			}
		})
	}
}

func TestInstallCommandReturnsNotInstallable(t *testing.T) {
	a := NewAdapter()
	_, err := a.InstallCommand(system.PlatformProfile{OS: "windows", PackageManager: "winget"})
	if err == nil {
		t.Fatal("InstallCommand() error = nil, want AgentNotInstallableError")
	}

	var notInstallable AgentNotInstallableError
	if !errors.As(err, &notInstallable) {
		t.Fatalf("InstallCommand() error type = %T, want AgentNotInstallableError", err)
	}
}

func TestConfigPaths(t *testing.T) {
	a := NewAdapter()
	home := "/tmp/home"

	tests := []struct {
		name string
		got  string
		want string
	}{
		{
			name: "GlobalConfigDir",
			got:  a.GlobalConfigDir(home),
			want: filepath.Join(home, ".copilot"),
		},
		{
			name: "SystemPromptDir",
			got:  a.SystemPromptDir(home),
			want: filepath.Join(home, ".copilot"),
		},
		{
			name: "SystemPromptFile",
			got:  a.SystemPromptFile(home),
			want: filepath.Join(home, ".copilot", "copilot-instructions.md"),
		},
		{
			name: "SkillsDir",
			got:  a.SkillsDir(home),
			want: filepath.Join(home, ".copilot", "skills"),
		},
		{
			name: "SettingsPath",
			got:  a.SettingsPath(home),
			want: "",
		},
		{
			name: "MCPConfigPath",
			got:  a.MCPConfigPath(home, "context7"),
			want: filepath.Join(home, ".copilot", "mcp.json"),
		},
		{
			name: "CommandsDir",
			got:  a.CommandsDir(home),
			want: "",
		},
		{
			name: "SubAgentsDir",
			got:  a.SubAgentsDir(home),
			want: "",
		},
		{
			name: "OutputStyleDir",
			got:  a.OutputStyleDir(home),
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Fatalf("%s() = %q, want %q", tt.name, tt.got, tt.want)
			}
		})
	}
}

func TestCapabilities(t *testing.T) {
	a := NewAdapter()

	tests := []struct {
		name string
		got  bool
		want bool
	}{
		{"SupportsAutoInstall", a.SupportsAutoInstall(), false},
		{"SupportsSkills", a.SupportsSkills(), true},
		{"SupportsSystemPrompt", a.SupportsSystemPrompt(), true},
		{"SupportsMCP", a.SupportsMCP(), true},
		{"SupportsSlashCommands", a.SupportsSlashCommands(), false},
		{"SupportsSubAgents", a.SupportsSubAgents(), false},
		{"SupportsOutputStyles", a.SupportsOutputStyles(), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Fatalf("%s() = %v, want %v", tt.name, tt.got, tt.want)
			}
		})
	}
}

func TestAdapterIdentityAndStrategies(t *testing.T) {
	a := NewAdapter()

	if got := a.Agent(); got != model.AgentCopilotCLI {
		t.Fatalf("Agent() = %q, want %q", got, model.AgentCopilotCLI)
	}
	if got := a.Tier(); got != model.TierFull {
		t.Fatalf("Tier() = %q, want %q", got, model.TierFull)
	}
	if got := a.SystemPromptStrategy(); got != model.StrategyInstructionsFile {
		t.Fatalf("SystemPromptStrategy() = %v, want %v", got, model.StrategyInstructionsFile)
	}
	if got := a.MCPStrategy(); got != model.StrategyMCPConfigFile {
		t.Fatalf("MCPStrategy() = %v, want %v", got, model.StrategyMCPConfigFile)
	}
}
