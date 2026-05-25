# gentle-copilot (GitHub Copilot CLI port)

Este repositorio es un **port de Gentle-AI** reducido para usarlo con **GitHub Copilot CLI**.

---

## Alcance del port

| Área | Estado |
| --- | --- |
| Agente soportado | `copilot-cli` (único) |
| Binario/comando | `gentle-copilot` |
| Superficies gestionadas | `~/.copilot/copilot-instructions.md`, `~/.copilot/skills/*`, `~/.copilot/mcp.json` |
| `update` / `upgrade` | Deshabilitados en este flavor |

---

## Instalación (Windows / Scoop)

Instalar directo desde el manifest del release:

```powershell
scoop install https://github.com/Pierillo/gentle-ai/releases/download/v1.30.11-copilotcli.3/gentle-copilot.json
```

Verificar:

```powershell
gentle-copilot --version
```

---

## Quickstart (propio de este port)

1. Instala `gentle-copilot` con el comando de arriba.
2. Ejecuta un dry-run:

```powershell
gentle-copilot install --dry-run
```

3. Aplica instalación:

```powershell
gentle-copilot install
```

4. Refresca configuración gestionada cuando corresponda:

```powershell
gentle-copilot sync
```

---

## Cómo usarlo con GitHub Copilot CLI

Después de instalar:

1. Abre tu proyecto con Copilot CLI.
2. Mantén la configuración de agente en `~/.copilot` (este port no toca otros agentes).
3. Cuando cambies skills/reglas del repo, actualiza el registro:

```powershell
gentle-copilot skill-registry refresh
```

Comandos disponibles en este port:

- `gentle-copilot install`
- `gentle-copilot sync`
- `gentle-copilot uninstall`
- `gentle-copilot restore`
- `gentle-copilot version`
- `gentle-copilot skill-registry refresh`

Comandos bloqueados en este port:

- `gentle-copilot update`
- `gentle-copilot upgrade`

---

## Release

- Release actual:  
  https://github.com/Pierillo/gentle-ai/releases/tag/v1.30.11-copilotcli.3
- PR de documentación/alcance del port (mergeado):  
  https://github.com/Pierillo/gentle-ai/pull/1

---

## Licencia

MIT (ver `LICENSE`).
