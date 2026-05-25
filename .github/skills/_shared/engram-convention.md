# Engram Convention (Pilot)

Use deterministic keys for SDD artifacts:

- `sdd/{change}/explore`
- `sdd/{change}/proposal`
- `sdd/{change}/spec`
- `sdd/{change}/design`
- `sdd/{change}/tasks`
- `sdd/{change}/apply-progress`
- `sdd/{change}/verify-report`
- `sdd/{change}/archive-report`

When available, prefer:

1. `mem_search` -> find candidate observation
2. `mem_get_observation` -> retrieve full content
3. `mem_save` or `mem_update` -> persist current phase output
