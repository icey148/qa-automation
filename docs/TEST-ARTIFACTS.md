# Test Artifacts

Runtime test output is temporary execution evidence. It is not source code and must not be committed to the repository.

## Runtime workspace

Every GitHub Actions job writes into the same logical workspace:

```text
test-artifacts/
├── metadata.json
├── debug_info/
├── logs/
├── reports/
├── telemetry/
├── test_generation_output/
└── cases/
```

The workflow run itself provides historical isolation. Each upload is named with GitHub `run_id` and `run_attempt`, so reruns do not overwrite earlier artifacts.

Example artifact names:

```text
qa-api-1938123123-1
qa-ui-1938123123-1
qa-ui-1938123123-2
```

## Directory responsibilities

- `metadata.json`: workflow/run identity, commit SHA, ref, target, environment, actor and timestamps.
- `debug_info/`: test outcome and execution diagnostics.
- `logs/`: raw API/Playwright execution logs.
- `reports/`: machine-readable and human-readable reports such as JSON, JUnit XML and Playwright HTML.
- `telemetry/`: runner/job execution metadata that is safe to retain.
- `test_generation_output/`: generation/validation manifests when Agent/Skill generation is part of the run.
- `cases/`: case-level evidence such as Playwright traces, screenshots, videos and per-test output.

## GitHub Actions upload rule

Artifact upload MUST use `if: always()` so failed tests still preserve failure evidence.

The test step may use `continue-on-error: true` temporarily. The workflow must upload artifacts first, then explicitly fail the job when the test step failed.

## Source control rule

`test-artifacts/` is ignored by Git. Runtime output belongs in GitHub Actions Artifacts, not on the main branch.

Long-lived QA documents remain under `artifacts/`:

```text
artifacts/
├── test-strategy.md
├── test-plan.md
└── test-cases.md
```

Executable tests remain under `tests/`.
