# QA Automation

一个面向 QA Agent + 自动化测试执行的项目骨架，核心目标是把 **AI 生成能力** 与 **测试执行框架** 分开。

## 核心架构

```text
qa-automation/
├── AGENTS.md               # Codex / Agent 仓库级工作约束
├── agents/                 # Agent / Skill 层
│   └── skills/
│       ├── qa-setup/
│       ├── test-strategy/
│       ├── test-plan/
│       ├── test-case/
│       ├── test-script/
│       └── github-action/
├── automation-framework/   # Go CLI 自动化测试框架
├── tests/                  # 真正执行的 API / UI 测试脚本
├── artifacts/              # 长期 QA 文档：Strategy / Plan / Test Case
├── test-artifacts/         # 运行时临时产物；不提交 Git，由 Actions 上传
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

## UI Test 分层

```text
Test Case
  ↓
tests/ui/cases/
  ↓
tests/ui/pages/
  ↓
tests/ui/framework/
  ↓
Playwright
```

生成 UI Test Script 前必须先搜索已有 `framework/`、`pages/` 和已有测试，能复用就复用，避免重复生成 selector、Page Object、fixture 和 helper。

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

## GitHub Actions Runtime Artifacts

GitHub Actions 每次执行会创建临时目录：

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

运行产物**不会 commit 到 main**。Action 跑完后统一上传为 GitHub Actions Artifact：

```text
qa-<target>-<github.run_id>-<github.run_attempt>
```

即使测试失败，也会先保存 logs / reports / screenshots / traces / videos 等现场，然后再把 Job 标记为失败。

详见 [docs/TEST-ARTIFACTS.md](docs/TEST-ARTIFACTS.md)。

## BMAD

该仓库按照多 Skill Custom Module 的方向准备了：

- `.claude-plugin/marketplace.json`
- `agents/skills/qa-setup/`
- `module.yaml`
- `module-help.csv`

详见 [docs/BMAD-INTEGRATION.md](docs/BMAD-INTEGRATION.md)。

## SSO

认证属于 Framework 基础设施能力，而不是每一个 Test Script 自己重新登录。详见 [docs/AUTH-SSO.md](docs/AUTH-SSO.md)。

## Codex 后续开发

仓库根目录已经提供 `AGENTS.md`，用于告诉 Codex 项目架构、硬依赖、复用规则、认证规则、Artifact 规则以及验证命令。OpenAI 的 Codex 会读取仓库中的 `AGENTS.md` 作为项目级说明。后续接手说明见 [docs/CODEX-HANDOFF.md](docs/CODEX-HANDOFF.md)。

## 文档

- Architecture: [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md)
- Skill design: [docs/SKILL-DESIGN.md](docs/SKILL-DESIGN.md)
- GitHub Actions: [docs/GITHUB-ACTIONS.md](docs/GITHUB-ACTIONS.md)
- Test Artifacts: [docs/TEST-ARTIFACTS.md](docs/TEST-ARTIFACTS.md)
- Reporting: [docs/REPORTING.md](docs/REPORTING.md)
- SSO/Auth: [docs/AUTH-SSO.md](docs/AUTH-SSO.md)
- Development: [docs/DEVELOPMENT.md](docs/DEVELOPMENT.md)
- Codex handoff: [docs/CODEX-HANDOFF.md](docs/CODEX-HANDOFF.md)

## 当前阶段

当前版本已经包含 QA Skill 体系、Go CLI 框架、UI Playwright 分层、复用优先规则、GitHub Actions API/UI 执行入口，以及统一 GitHub Actions Artifact 生命周期。真实企业 SSO Provider、JUnit/Go JSON 统一解析、case-level artifact mapping、统一 Report Aggregation 等可继续由 Codex 实现。
