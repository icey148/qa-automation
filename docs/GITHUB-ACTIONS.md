# GitHub Actions

Two starter workflows are included:

- `framework-ci.yml`: validates the Go automation framework and example API suite.
- `qa-tests.yml`: manually runs API or UI tests and uploads one unified runtime artifact bundle.

## QA test workflow

`qa-tests.yml` accepts:

- `target`: `api` or `ui`
- `environment`: logical environment name such as `dev`

Each job performs this sequence:

```text
checkout
  -> prepare test-artifacts/
  -> execute tests
  -> collect logs/reports/case evidence
  -> upload GitHub Actions Artifact
  -> restore failed job status when tests failed
```

The artifact name includes the GitHub Actions run identity:

```text
qa-<target>-<github.run_id>-<github.run_attempt>
```

This preserves multiple executions and reruns without committing runtime files to Git.

## Failure handling

Artifact upload uses `if: always()`.

The actual test step temporarily uses `continue-on-error: true`, which allows the workflow to collect screenshots, traces, logs, JUnit and other failure evidence. After upload, a final step returns the job to `failure` when the test step failed.

## Runtime artifact layout

See [TEST-ARTIFACTS.md](TEST-ARTIFACTS.md).

## Authentication

Do not place `.auth` session state or interactive SSO credentials in GitHub.

For enterprise CI use the organization's approved non-interactive authentication path, for example OIDC or a service/test identity, and restrict workflow permissions to the minimum required.

Interactive MFA/PingID should not be embedded in unattended GitHub Actions jobs.
