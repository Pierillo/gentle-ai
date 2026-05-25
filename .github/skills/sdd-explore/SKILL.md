---
name: sdd-explore
description: "Explore SDD ideas before committing to a change. Trigger: architecture exploration, requirement clarification, or implementation option analysis."
license: MIT
metadata:
  author: gentleman-programming
  version: "pilot-1.0"
---

## Purpose

Investigate a change request before implementation.  
Return a concise, decision-ready analysis.

## Inputs

- Topic or change name
- Optional constraints (timeline, risk, compatibility)

## Steps

1. Understand request scope and success criteria.
2. Read relevant code paths, tests, and boundaries.
3. Identify affected areas and coupling.
4. Compare viable approaches with tradeoffs.
5. Recommend one approach with explicit risks.

## Output format

```markdown
## Exploration: {topic}

### Current State
...

### Affected Areas
- `path/to/file` - reason

### Approaches
1. **Option A**
   - Pros:
   - Cons:
   - Effort: Low/Medium/High
2. **Option B**
   - Pros:
   - Cons:
   - Effort: Low/Medium/High

### Recommendation
...

### Risks
- ...

### Ready for Proposal
Yes/No (+ what is missing)
```

## Rules

- Do not modify code in this phase.
- Prefer evidence from source over assumptions.
- Keep the report concise and actionable.
