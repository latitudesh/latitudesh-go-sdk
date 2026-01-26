# GitHub Actions Workflows

This directory contains the GitHub Actions workflows for the latitudesh-go-sdk project.

## Available Workflows

### 1. Integration Tests (`sdk_integration_tests.yaml`)

Runs SDK integration tests automatically.

**Triggers:**
- Push to branches: `main`, `develop`, `feat/**`
- Pull requests to: `main`, `develop`
- Manual dispatch via GitHub interface

**Jobs:**
- **integration-tests**: Runs tests on Go 1.22 and 1.23
  - Runs tests with race detector
  - Generates coverage report
  - Checks coverage threshold (80%)
  - Uploads coverage artifacts

- **lint**: Runs golangci-lint
  - Verifies code quality
  - Applies style rules

- **test-summary**: Summarizes results
  - Consolidates status of all jobs
  - Reports overall success or failure

**Generated artifacts:**
- `coverage.out` - Coverage in Go format
- `coverage.html` - HTML coverage report
- `coverage.txt` - Text coverage summary

**How to view:**
1. Go to the "Actions" tab on GitHub
2. Select the "Integration Tests" workflow
3. Click on a specific run
4. Download artifacts at the bottom of the page

### 2. SDK Generation (`sdk_generation.yaml`)

Speakeasy workflow to automatically generate the SDK.

**Triggers:**
- Manual dispatch
- Daily schedule (00:00 UTC)

**Configuration:**
- Uses Speakeasy API v15
- Mode: Pull Request
- Requires secrets: `SPEAKEASY_API_KEY`

### 3. SDK Publish (`sdk_publish.yaml`)

Workflow to publish the SDK when there are changes.

**Triggers:**
- Push to `main` when `.speakeasy/gen.lock` changes
- Manual dispatch

## Secret Configuration

The following secrets must be configured in the repository:

| Secret | Description | Used by |
|--------|-------------|---------|
| `GITHUB_TOKEN` | Automatic GitHub token | All workflows |
| `SPEAKEASY_API_KEY` | Speakeasy API key | SDK Generation, SDK Publish |

## Permissions

Workflows require the following permissions:

```yaml
permissions:
  contents: read      # Read code
  pull-requests: write # Comment on PRs
  checks: write       # Create check runs
```

## Running Workflows Manually

To run a workflow manually:

1. Go to the "Actions" tab on GitHub
2. Select the desired workflow in the sidebar
3. Click "Run workflow"
4. Select the branch
5. Fill in inputs (if any)
6. Click "Run workflow"

## Cache

Workflows use GitHub Actions cache for:
- Go dependencies (`go.sum`)
- Downloaded Go modules
- Go build cache

This significantly speeds up subsequent runs.

## Monitoring

### Status Badges

Add badges to README to show workflow status:

```markdown
![Integration Tests](https://github.com/latitudesh/latitudesh-go-sdk/actions/workflows/sdk_integration_tests.yaml/badge.svg)
![PR Tests](https://github.com/latitudesh/latitudesh-go-sdk/actions/workflows/sdk_pr_tests.yaml/badge.svg)
```

### Notifications

Configure GitHub notifications to:
- Receive emails when workflows fail
- Desktop/mobile notifications
- Slack integration (via webhooks)

## Debugging

### Detailed Logs

To see detailed logs of a specific step:
1. Click on the step name in the workflow run
2. Expand relevant sections
3. Use Ctrl+F to search for specific errors

### Re-running Workflows

To re-run a failed workflow:
1. Open the failed run
2. Click "Re-run jobs"
3. Choose:
   - "Re-run failed jobs" - failed jobs only
   - "Re-run all jobs" - all jobs

### Logs with Race Detector

Tests are run with the `-race` flag, which detects:
- Race conditions
- Unsynchronized concurrent access
- Thread-safety issues

## Troubleshooting

### Workflow not being triggered

Check:
- Branch is in the trigger list
- Modified files match configured paths
- No YAML syntax errors

### Tests failing in CI but passing locally

Possible causes:
- Environment differences (Go version, OS)
- Race conditions not detected locally
- Uncommitted dependencies

Solution:
- Use same Go version as CI
- Run with `-race` locally
- Check `go.mod` and `go.sum`

### Coverage below threshold

If coverage drops below 80%:
1. Workflow shows a warning (doesn't fail)
2. Add more tests for uncovered code
3. Use `go test -coverprofile` locally to identify gaps

## Maintenance

### Updating Versions

Periodically update:
- Action versions (`actions/checkout@v4`, `actions/setup-go@v5`)
- Go versions in matrix
- Linter versions

### Optimizing Performance

To improve performance:
- Use cache aggressively
- Minimize redundant steps
- Run jobs in parallel when possible
- Use more powerful runners if needed

## Useful Links

- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [Go Testing Package](https://pkg.go.dev/testing)
- [golangci-lint](https://golangci-lint.run/)
- [Speakeasy Docs](https://speakeasy.com/docs)
