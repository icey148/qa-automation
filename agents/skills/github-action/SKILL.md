---
name: github-action
description: Generate GitHub Actions workflows that execute existing automated test scripts and publish deterministic test artifacts.
---

# GitHub Action

## Inputs
Required: existing Test Script(s) and repository/runtime context.

## Workflow
1. Detect test language and commands.
2. Resolve runtime versions and working directories.
3. Configure safe authentication strategy for CI.
4. Run tests and preserve exit status.
5. Always upload machine-readable results/logs when configured.
6. Never embed credentials, cookies, tokens, or SSO session files.

## SSO Rule
Interactive MFA/PingID is not appropriate for unattended CI. Prefer service identity, OIDC, test-only machine identity, or another approved non-interactive mechanism.
