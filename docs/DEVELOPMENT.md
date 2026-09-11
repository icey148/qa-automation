# Development

## Prerequisites

- Go 1.23+
- Git
- Optional: BMAD for module installation/validation

## Validate locally

```bash
make test
make build
./bin/qa help
```

## CLI

```bash
# run existing tests
go run ./automation-framework/cmd/qa run --config configs/dev.json

# initialize local auth metadata (not real IdP login)
go run ./automation-framework/cmd/qa auth --session-dir .auth --env dev --role qa-user

# build report from a previously saved execution result
go run ./automation-framework/cmd/qa report --input artifacts/result.json
```
