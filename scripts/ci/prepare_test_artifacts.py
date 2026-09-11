#!/usr/bin/env python3
from __future__ import annotations

import json
import os
from datetime import datetime, timezone
from pathlib import Path

ROOT = Path(os.getenv("QA_ARTIFACT_ROOT", "test-artifacts"))

for name in (
    "debug_info",
    "logs",
    "reports",
    "telemetry",
    "test_generation_output",
    "cases",
):
    (ROOT / name).mkdir(parents=True, exist_ok=True)

now = datetime.now(timezone.utc).isoformat()
metadata = {
    "runId": os.getenv("GITHUB_RUN_ID", "local"),
    "runAttempt": int(os.getenv("GITHUB_RUN_ATTEMPT", "1")),
    "source": "github-actions" if os.getenv("GITHUB_ACTIONS") == "true" else "local",
    "repository": os.getenv("GITHUB_REPOSITORY", ""),
    "workflow": os.getenv("GITHUB_WORKFLOW", ""),
    "eventName": os.getenv("GITHUB_EVENT_NAME", ""),
    "actor": os.getenv("GITHUB_ACTOR", ""),
    "ref": os.getenv("GITHUB_REF", ""),
    "sha": os.getenv("GITHUB_SHA", ""),
    "target": os.getenv("QA_TEST_TARGET", ""),
    "environment": os.getenv("QA_ENVIRONMENT", ""),
    "startedAt": now,
}
(ROOT / "metadata.json").write_text(json.dumps(metadata, indent=2) + "\n", encoding="utf-8")

telemetry = {
    "runnerOS": os.getenv("RUNNER_OS", ""),
    "runnerArch": os.getenv("RUNNER_ARCH", ""),
    "runnerEnvironment": os.getenv("RUNNER_ENVIRONMENT", ""),
    "job": os.getenv("GITHUB_JOB", ""),
}
(ROOT / "telemetry" / "github-actions.json").write_text(
    json.dumps(telemetry, indent=2) + "\n",
    encoding="utf-8",
)

generation_manifest = {
    "generatedDuringThisRun": False,
    "note": "Reserved for Agent/Skill generation and validation outputs when generation occurs before execution.",
}
(ROOT / "test_generation_output" / "manifest.json").write_text(
    json.dumps(generation_manifest, indent=2) + "\n",
    encoding="utf-8",
)

print(ROOT)
