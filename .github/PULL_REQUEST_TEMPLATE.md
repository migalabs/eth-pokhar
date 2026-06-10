<!-- PR title must follow Conventional Commits with a scope, e.g. `fix(lido/cmv2): validate endpoint addresses` -->
<!-- Target branch must be `dev`. `main` (production) only accepts release PRs from `dev` or `hotfix/*` branches. -->

## Description

<!-- What does this PR do and why? 1-3 sentences. -->

## Related issue

<!-- e.g. Closes #17 — remove if none -->

## Type of change

- [ ] Bug fix
- [ ] New feature
- [ ] Refactor (no functional change)
- [ ] Performance improvement
- [ ] Documentation
- [ ] Chore / maintenance

## Affected area(s)

<!-- e.g. identify, lido, rocketpool, coinbase, beacon-deposits, db, alchemy, config, docker/ci, docs -->

## How was this tested?

<!-- Commands run and context: network, EL endpoint type, database state (fresh vs synced). -->

## Checklist

- [ ] PR title follows [Conventional Commits](https://www.conventionalcommits.org/) with a scope
- [ ] Database schema changes include a migration in `db/migrations`
- [ ] README updated if CLI flags, commands or setup steps changed
- [ ] No secrets, API keys or private endpoints committed
- [ ] `go build ./...` and `go vet ./...` pass locally
