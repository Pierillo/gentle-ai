# SDD Instructions (Copilot CLI Pilot)

Use this when working on Spec-Driven Development phases in this repository.

## Phase sequence

`explore -> propose -> spec -> design -> tasks -> apply -> verify -> archive`

## Manual model rule

Before starting each phase, select model with `/model`.  
Do not assume automatic per-phase routing.

## Required phase output format

For every phase return:

- `status`
- `executive_summary`
- `artifacts`
- `next_recommended`
- `risks` (when relevant)

## Explore phase expectations

- Read real code; do not infer architecture without evidence.
- Identify affected files and constraints.
- Compare at least two approaches when there is real tradeoff.
- End with a clear recommendation and risk list.
