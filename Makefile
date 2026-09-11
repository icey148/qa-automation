.PHONY: test build run auth report

test:
	cd automation-framework && go test ./...
	cd tests/api && go test ./...

build:
	cd automation-framework && go build -o ../bin/qa ./cmd/qa

run:
	cd automation-framework && go run ./cmd/qa run --config ../configs/dev.json

auth:
	cd automation-framework && go run ./cmd/qa auth --session-dir ../.auth --env dev --role qa-user

report:
	cd automation-framework && go run ./cmd/qa report --input ../artifacts/result.json
