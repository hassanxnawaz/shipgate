# shipgate

The CI gate for an agent change.

One JSON object on stdin. A decision on stdout. Exit 0 only on SHIP. CHANGE exits 2. RERUN exits 3. STOP exits 1.

Go, because this is the binary a pipeline runs.

Same rule as [tracecard](https://github.com/hassanxnawaz/tracecard). The public Qwen case (+11 points, interval -7 to 30, a guardrail over budget) is RERUN, not a win.

```bash
go run . < run.json
```

```json
{"delta_pp":11,"ci_low":-7,"ci_high":30,"guardrail_over":true,"task_regressed":false,"target_improved":false}
```

## GitHub Actions

The [composite action](action.yml) runs the gate against a JSON file in your repository. In an Ubuntu job with Go 1.23 or newer on PATH, add these four lines under `steps:` (install Go first with `actions/setup-go` if needed):

```yaml
- uses: actions/checkout@v4
- uses: hassanxnawaz/shipgate@main
  with:
    file: run.json
```

Commit `run.json` to your repository, or generate it before the shipgate step. The path is relative to the checked-out repository root. For production, replace `@main` with a reviewed full commit SHA.

The step prints the decision and its SHA-256 fingerprint. Only SHIP passes; CHANGE, RERUN, and STOP fail the job. The action uses `go run`, which returns 1 for any nonzero program exit; the distinct exit codes above apply to the compiled binary.

To try the published Qwen case, copy [examples/qwen.json](examples/qwen.json) into your repository and set `file` to its path. It deliberately fails with RERUN: the interval crosses zero, so the reported +11 is not a confirmed win.

Built by Hassan Nawaz.

MIT licensed.
