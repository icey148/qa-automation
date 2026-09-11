---
name: test-case
description: Generate canonical, traceable test cases from authoritative evidence; Strategy and Plan are optional upstream inputs.
---

# Test Case

## Purpose
Generate the canonical human-readable definition of WHAT must be tested.

## Inputs
Required: at least one authoritative requirement source.
Preferred when available: Test Strategy, Test Plan, existing Test Cases.
Supporting: OpenAPI, source code, PR/diff, baseline behavior.

## Workflow
1. Resolve evidence and existing coverage.
2. Extract testable conditions.
3. Apply suitable design techniques.
4. Generate deterministic preconditions, steps, test data, and expected results.
5. Assign traceability metadata.
6. Validate no duplicate or fabricated behavior.
7. Render with `templates/test-case.md`.

## Rules
- Test Strategy and Test Plan are NOT hard dependencies.
- Do not generate missing upstream artifacts merely to make the chain look complete.
- Every case must be independently understandable and deterministic.
