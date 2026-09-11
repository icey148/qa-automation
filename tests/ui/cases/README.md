# UI Cases

This directory contains business-facing Playwright test cases generated from canonical Test Cases.

Rules:

- Keep business scenario logic here.
- Reuse existing Page Objects from `../pages/`.
- Reuse shared fixtures, components, assertions, and helpers from `../framework/`.
- Do not duplicate SSO/login flows inside case files.
- Keep the canonical Test Case ID in each Playwright test title or metadata.
- If a reusable abstraction is missing, create it in `pages/` or `framework/` before duplicating selectors/helpers across cases.
