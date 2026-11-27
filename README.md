# Deepward demo

This repository is a live example of [Deepward](https://github.com/deepward-oss/deepward).
It holds three small projects with **intentionally outdated and vulnerable
dependencies**. Do not use them as a template.

| Directory | Ecosystem | Resolved from |
|-----------|-----------|---------------|
| [`npm/`](npm) | npm (workspaces) | `package-lock.json` |
| [`go/`](go) | Go modules | `go list` |
| [`cargo/`](cargo) | Cargo (workspace) | `Cargo.lock` |

## What to look at

- **Pull requests** opened by Deepward: each one changes only version
  tokens and lockfiles, and was verified before it was opened: the targeted
  findings are gone and the dependency graph still resolves.
- **Actions → deepward → scan**: the HTML report is attached to every run
  as the `deepward-report` artifact.
- **Security → Code scanning**: findings uploaded as SARIF.

The workflow is [`.github/workflows/deepward.yml`](.github/workflows/deepward.yml).
`scan` runs on every push and pull request; `fix` runs only when started by
hand, so the pull requests shown here stay stable. Dependabot updates are
turned off, so every automated pull request comes from Deepward.

## License

MIT, see [LICENSE](LICENSE).
