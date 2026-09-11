# Development

## Prerequisites

- Go 1.23+
- Git
- Node.js 22+ for Playwright UI tests
- Optional: BMAD for module installation/validation

## Validate locally

```bash
make test
make build
./bin/qa help
```

For UI test structure/TypeScript discovery:

```bash
cd tests/ui
npm install
npx playwright test --list
```

Run actual UI tests only when the target application/base URL and required auth context are available.

## CLI

```bash
# run existing tests
go run ./automation-framework/cmd/qa run --config configs/dev.json

# initialize local auth metadata (not real IdP login)
go run ./automation-framework/cmd/qa auth --session-dir .auth --env dev --role qa-user

# build report from a previously saved execution result
go run ./automation-framework/cmd/qa report --input test-artifacts/reports/result.json
```

## Runtime artifacts

Do not place runtime output under `tests/`.

Use `test-artifacts/` as the local/CI temporary workspace. It is ignored by Git and uploaded by GitHub Actions after execution.

See [TEST-ARTIFACTS.md](TEST-ARTIFACTS.md).

## Codex

The repository root contains `AGENTS.md`. Codex should use it as the repository-wide working contract.

See [CODEX-HANDOFF.md](CODEX-HANDOFF.md) for current implementation state and recommended next work.
