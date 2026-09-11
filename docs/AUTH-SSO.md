# SSO / Authentication Design

Authentication is framework infrastructure, not Test Case business logic.

## Principles

1. Test Cases declare the required identity/role, not login implementation details.
2. Test Scripts consume an authenticated context/session where possible.
3. SSO cookies, tokens, storage state, passwords, and MFA secrets must not be committed.
4. `.auth/` is gitignored except for `.gitkeep`.
5. CI should avoid interactive MFA/PingID and prefer an approved non-interactive identity mechanism such as OIDC/service identity/test machine account.

## Scaffold status

The initial Go CLI includes auth session metadata lifecycle (`qa auth`) only. It intentionally does **not** automate or bypass enterprise SSO/MFA. Add provider adapters for the actual enterprise IdP/browser stack.
