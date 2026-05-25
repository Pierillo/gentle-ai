# Copilot CLI Port

← [Back to README](../README.md)

---

This repository is a **port of Gentle-AI** narrowed to a **Copilot CLI-only** experience.

## Scope

| Area | Port behavior |
| --- | --- |
| Supported agent | `copilot-cli` only |
| Command name | `pierillo-antislop-boring-toolkit` |
| Config scan | `~/.copilot` only |
| Main managed files | `~/.copilot/copilot-instructions.md`, `~/.copilot/skills/*`, `~/.copilot/mcp.json` |

## Commands in this port

- Supported: `install`, `sync`, `uninstall`, `restore`, `version`, `skill-registry refresh`
- Disabled by design: `update`, `upgrade` (use internal release channel)

## Runtime defaults changed for copilot-only

- Install/sync defaults exclude non-Copilot surfaces (for example GGA/theme-related components).
- Uninstall defaults are constrained to Copilot-relevant managed components.
- Dependency profile is reduced to optional checks for `git` and `engram`.

## Release and installation

- Current release:  
  https://github.com/Pierillo/pierillo-antislop-boring-toolkit/releases/tag/v1.30.11-antislop.1

- Scoop install (direct manifest):

```powershell
scoop install https://github.com/Pierillo/pierillo-antislop-boring-toolkit/releases/download/v1.30.11-antislop.1/pierillo-antislop-boring-toolkit.json
```

---

If you need the full multi-agent ecosystem (Claude, OpenCode, Cursor, etc.), use upstream Gentle-AI instead of this port.
