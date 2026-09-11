---
name: test-strategy
description: Generate a structured, traceable test strategy from authoritative requirement evidence.
---

# Test Strategy

## Purpose
Define overall testing scope, risks, approaches, test types, automation intent, and entry/exit criteria.

## Inputs
At least one authoritative source: Jira requirement, requirement document, acceptance criteria, specification, or baseline evidence.

## Workflow
1. Resolve authoritative evidence.
2. Identify scope and exclusions.
3. Identify product and delivery risks.
4. Define test levels/types and approach.
5. Define environments, data needs, automation direction, entry/exit criteria.
6. Validate traceability and unsupported assumptions.
7. Render with `templates/test-strategy.md`.

## Rules
- Do not invent requirements.
- Do not turn Strategy into step-by-step Test Cases.
- Existing approved artifacts may be used as evidence but are not mandatory.
