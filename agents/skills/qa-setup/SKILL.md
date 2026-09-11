---
name: qa-setup
description: Register and configure the multi-skill QA Automation module and its shared paths.
---

# QA Setup

## Purpose

Register the QA Automation module, configure project paths, and expose all module capabilities to BMAD help/discovery.

## When to Use

Use this skill when installing, configuring, reconfiguring, or validating the QA Automation module.

## Responsibilities

1. Register module metadata from `assets/module.yaml`.
2. Merge capability rows from `assets/module-help.csv` into BMAD help registration.
3. Configure artifact, test, framework, and authentication paths.
4. Do not generate Strategy, Plan, Case, Script, workflow, or report content itself.

## Shared Rules

- `test-script` has a hard dependency on `test-case`.
- If the user requests only Test Script and no Test Case exists, generate and validate Test Case first.
- The Go automation framework executes existing tests; it does not generate test files.
- SSO/session secrets must never be committed to source control.
