# GitHub Actions

Two starter workflows are included:

- `framework-ci.yml`: validates the framework and example API suite.
- `qa-tests.yml`: manually runs QA tests and uploads raw JSON output.

## Authentication

Do not place `.auth` session state or interactive SSO credentials in GitHub.
For enterprise CI use the organization's approved non-interactive authentication path (for example OIDC or service identity) and restrict workflow permissions to the minimum required.
