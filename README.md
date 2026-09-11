# QA Automation

一个面向 QA Agent + 自动化测试执行的项目骨架，核心目标是把 **AI 生成能力** 与 **测试执行框架** 分开。

## 核心架构

```text
qa-automation/
├── agents/                 # Agent / Skill 层
│   └── skills/
│       ├── qa-setup/
│       ├── test-strategy/
│       ├── test-plan/
│       ├── test-case/
│       ├── test-script/
│       └── github-action/
├── automation-framework/   # Go CLI 自动化测试框架
├── tests/                  # 真正执行的测试脚本
├── artifacts/              # Strategy / Plan / Case / Run / Report 等产物
├── configs/                # Runtime 配置
├── .auth/                  # SSO/session 本地状态（gitignore）
└── .github/                # GitHub Actions
```

## 最重要的规则

### Test Script 必须先有 Test Case

```text
Requirement / Evidence
        ↓
    Test Case
        ↓
    Test Script
```

即使用户只要求 `Test Script`，系统内部也必须先生成并校验对应 `Test Case`。

### Go Framework 只负责执行

```text
Agent / Skills = 生成什么、怎么生成
Go Framework   = 怎么执行现有测试
```

Go Framework 不负责生成 Test Case 或 Test Script。

## CLI-first Framework

开发期：

```bash
cd automation-framework
go run ./cmd/qa help
go run ./cmd/qa version
```

编译：

```bash
make build
./bin/qa help
```

支持的初始命令：

```text
qa run
qa auth
qa report
qa version
```

`qa run` 执行现有测试；`qa report` 根据真实执行结果生成报告；`qa auth` 当前只提供安全的 session 元数据骨架，真实企业 SSO/MFA 需要按 IdP 接入 Provider Adapter。

## 自动化测试建议

- API / Backend / Integration：Go
- Web UI：TypeScript + Playwright
- 执行、结果聚合、报告、CI 编排：Go Framework

## BMAD

该仓库按照多 Skill Custom Module 的方向准备了：

- `.claude-plugin/marketplace.json`
- `agents/skills/qa-setup/`
- `module.yaml`
- `module-help.csv`

详见 [docs/BMAD-INTEGRATION.md](docs/BMAD-INTEGRATION.md)。

## SSO

认证属于 Framework 基础设施能力，而不是每一个 Test Script 自己重新登录。详见 [docs/AUTH-SSO.md](docs/AUTH-SSO.md)。

## GitHub Actions 与报告

- CI: [docs/GITHUB-ACTIONS.md](docs/GITHUB-ACTIONS.md)
- Reporting: [docs/REPORTING.md](docs/REPORTING.md)
- Architecture: [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md)
- Skill design: [docs/SKILL-DESIGN.md](docs/SKILL-DESIGN.md)
- Development: [docs/DEVELOPMENT.md](docs/DEVELOPMENT.md)

## 当前阶段

这是一个可编译、可测试的第一版骨架，重点先固定边界和扩展点。真实企业 SSO、Jira/GitHub API integration、Playwright UI runner、JUnit parser、AI failure analysis 等可以在此结构上继续扩展。
