# BMAD Integration

This repository is designed as a multi-skill custom BMAD module/plugin.

## Relevant files

```text
.claude-plugin/marketplace.json
agents/skills/qa-setup/SKILL.md
agents/skills/qa-setup/assets/module.yaml
agents/skills/qa-setup/assets/module-help.csv
```

`qa-setup` is the module registration/configuration skill. Other skills stay focused on their artifact responsibilities.

## Local development install

Use BMAD's custom-source/local-path installation flow and point it at this repository (or the skill source path required by the installed BMAD version).

After install, validate module discovery and registration using BMAD Builder's Validate Module capability.

## Dependency declaration

`module-help.csv` registers `test-script` as preceded by `test-case`, matching the project's hard dependency rule.
