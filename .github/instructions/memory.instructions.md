# Memory Instructions (Copilot + Engram)

## Policy

- Use **Copilot Memory** for repository conventions and recurring preferences.
- Use **Engram** for durable workflow artifacts and cross-tool portability.

## Practical rule

- If information must survive outside Copilot CLI sessions/tools, persist to Engram.
- If information is local guidance for this repo/user workflow, Copilot Memory is enough.

## SDD artifact persistence

For SDD phases, save outputs under stable Engram keys (`sdd/{change}/{artifact}`) when MCP is configured.
