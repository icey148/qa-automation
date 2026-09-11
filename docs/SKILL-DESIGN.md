# Skill Design

The project uses independent but composable artifact skills:

- `test-strategy`: no QA artifact hard dependency.
- `test-plan`: no QA artifact hard dependency; may reuse existing Strategy.
- `test-case`: no QA artifact hard dependency; may reuse existing Strategy/Plan.
- `test-script`: hard dependency on Test Case.
- `github-action`: expects executable Test Script(s) and repository/runtime context.

## Rule: no artificial upstream generation

Do not create Strategy or Plan only to make a Test Case chain look complete. Use them only when requested or already available.

## Rule: Test Script

Test Case defines WHAT to test. Code/OpenAPI/framework context defines HOW to automate it. A Test Script must not invent extra behavior.
