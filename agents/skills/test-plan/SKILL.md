---
name: test-plan
description: Generate a structured test plan from requirements, optionally enriched by an existing test strategy.
---

# Test Plan

## Purpose
Turn requirement scope into concrete test scenarios, priorities, coverage, dependencies, and execution planning.

## Inputs
Required: authoritative requirement evidence.
Optional: approved Test Strategy, existing Plan, risks, API/UI contracts.

## Workflow
1. Resolve requirement scope.
2. Reuse approved Strategy when available; never create one merely to satisfy this skill.
3. Define test scenarios and scenario IDs.
4. Assign priority, execution type, and traceability.
5. Validate coverage and duplication.
6. Render with `templates/test-plan.md`.
