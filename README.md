# pierillo-antislop-boring-toolkit (GitHub Copilot CLI port)

Este repositorio es un **port de Gentle-AI** reducido para usarlo con **GitHub Copilot CLI**.

---

## Alcance del port

| Área | Estado |
| --- | --- |
| Agente soportado | `copilot-cli` (único) |
| Binario/comando | `pierillo-antislop-boring-toolkit` |
| Superficies gestionadas | `~/.copilot/copilot-instructions.md`, `~/.copilot/skills/*`, `~/.copilot/mcp.json` |
| `update` / `upgrade` | Deshabilitados en este flavor |

---

## Instalación (Windows / Scoop)

Instalar directo desde el manifest del release:

```powershell
scoop install https://github.com/Pierillo/pierillo-antislop-boring-toolkit/releases/download/v1.30.11-antislop.1/pierillo-antislop-boring-toolkit.json
```

Verificar:

```powershell
pierillo-antislop-boring-toolkit --version
```

---

## Quickstart (propio de este port)

1. Instala `pierillo-antislop-boring-toolkit` con el comando de arriba.
2. Ejecuta un dry-run:

```powershell
pierillo-antislop-boring-toolkit install --dry-run
```

3. Aplica instalación:

```powershell
pierillo-antislop-boring-toolkit install
```

4. Refresca configuración gestionada cuando corresponda:

```powershell
pierillo-antislop-boring-toolkit sync
```

---

## Cómo usarlo con GitHub Copilot CLI

Después de instalar:

1. Abre tu proyecto con Copilot CLI.
2. Mantén la configuración de agente en `~/.copilot` (este port no toca otros agentes).
3. Cuando cambies skills/reglas del repo, actualiza el registro:

```powershell
pierillo-antislop-boring-toolkit skill-registry refresh
```

Comandos disponibles en este port:

- `pierillo-antislop-boring-toolkit install`
- `pierillo-antislop-boring-toolkit sync`
- `pierillo-antislop-boring-toolkit uninstall`
- `pierillo-antislop-boring-toolkit restore`
- `pierillo-antislop-boring-toolkit version`
- `pierillo-antislop-boring-toolkit skill-registry refresh`

Comandos bloqueados en este port:

- `pierillo-antislop-boring-toolkit update`
- `pierillo-antislop-boring-toolkit upgrade`

---

## Release

- Release actual:  
  https://github.com/Pierillo/pierillo-antislop-boring-toolkit/releases/tag/v1.30.11-antislop.1
- PR de documentación/alcance del port (mergeado):  
  https://github.com/Pierillo/pierillo-antislop-boring-toolkit/pull/1

---

## Licencia

MIT (ver `LICENSE`).
