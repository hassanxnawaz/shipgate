package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	raw, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	var body map[string]any
	if err := json.Unmarshal(raw, &body); err != nil {
		fmt.Fprintln(os.Stderr, "need one JSON object on stdin")
		os.Exit(1)
	}
	row := Row{
		Delta:     numField(body, "delta_pp"),
		Low:       numField(body, "ci_low"),
		High:      numField(body, "ci_high"),
		Guard:     boolField(body, "guardrail_over"),
		Regressed: boolField(body, "task_regressed"),
		Target:    boolField(body, "target_improved"),
	}
	decision := Decide(row)
	fmt.Printf("%s\n%s\n", decision, Fingerprint(row))
	os.Exit(ExitCode(decision))
}

func numField(body map[string]any, key string) float64 {
	v, ok := body[key].(float64)
	if !ok {
		return 0
	}
	return v
}

func boolField(body map[string]any, key string) bool {
	v, _ := body[key].(bool)
	return v
}
