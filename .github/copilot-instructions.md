# Copilot CLI Pilot Instructions

This repository is running a pilot migration of Gentle-AI workflows to GitHub Copilot CLI.

## SDD operating mode

- Use **single-mode SDD** in Copilot CLI.
- Before each SDD phase, set model manually with `/model`.
- Suggested baseline:
  - explore/propose/spec/design: capable model
  - tasks/apply/verify: balanced model
  - archive: fast model

## Workflow

1. Use **plan mode** (`Shift+Tab`) for proposal/spec/design work.
2. Use the `sdd-explore` skill from `.github/skills/sdd-explore/SKILL.md` for discovery.
3. Keep outputs structured and phase-scoped.

## Memory bridge

- Treat **Copilot Memory** as short/operational memory for this repo.
- Treat **Engram** as durable, portable memory (cross-tool and cross-machine).
- If Engram MCP is available, persist SDD artifacts with stable topic keys:
  - `sdd/{change}/explore`
  - `sdd/{change}/proposal`
  - `sdd/{change}/spec`
  - `sdd/{change}/design`
  - `sdd/{change}/tasks`
  - `sdd/{change}/apply-progress`
  - `sdd/{change}/verify-report`
  - `sdd/{change}/archive-report`
