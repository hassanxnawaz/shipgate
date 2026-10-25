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

Built by Hassan Nawaz.

MIT licensed.
