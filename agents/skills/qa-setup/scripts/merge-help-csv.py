#!/usr/bin/env python3
"""Minimal helper for merging QA module help entries by module name."""
from __future__ import annotations
import csv
import sys
from pathlib import Path

if len(sys.argv) != 3:
    raise SystemExit("usage: merge-help-csv.py <source.csv> <target.csv>")

source, target = map(Path, sys.argv[1:])
with source.open(newline="", encoding="utf-8") as f:
    src_rows = list(csv.DictReader(f))
if not src_rows:
    raise SystemExit("source has no rows")
fieldnames = list(src_rows[0].keys())
module = src_rows[0]["module"]
existing = []
if target.exists():
    with target.open(newline="", encoding="utf-8") as f:
        existing = [r for r in csv.DictReader(f) if r.get("module") != module]
target.parent.mkdir(parents=True, exist_ok=True)
with target.open("w", newline="", encoding="utf-8") as f:
    w = csv.DictWriter(f, fieldnames=fieldnames)
    w.writeheader()
    w.writerows(existing + src_rows)
print(target)
