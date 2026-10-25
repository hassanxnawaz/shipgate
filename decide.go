package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

type Row struct {
	Delta     float64
	Low       float64
	High      float64
	Guard     bool
	Regressed bool
	Target    bool
}

func Decide(r Row) string {
	if r.Low > 0 {
		if r.Guard || r.Regressed {
			return "CHANGE"
		}
		return "SHIP"
	}
	if r.High < 0 {
		return "STOP"
	}
	if abs(r.Delta) <= 5 && r.Target && !r.Guard && !r.Regressed {
		return "SHIP"
	}
	if r.Delta > 0 && r.Low <= 0 && r.High >= 0 {
		return "RERUN"
	}
	return "STOP"
}

func abs(v float64) float64 {
	if v < 0 {
		return -v
	}
	return v
}

func num(v float64) string {
	if v == float64(int64(v)) {
		return strconv.FormatInt(int64(v), 10)
	}
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func bit(v bool) string {
	if v {
		return "true"
	}
	return "false"
}

func Canonical(r Row) string {
	return fmt.Sprintf(
		`{"ci_high":%s,"ci_low":%s,"delta_pp":%s,"guardrail_over":%s,"target_improved":%s,"task_regressed":%s}`,
		num(r.High), num(r.Low), num(r.Delta), bit(r.Guard), bit(r.Target), bit(r.Regressed),
	)
}

func Fingerprint(r Row) string {
	sum := sha256.Sum256([]byte(Canonical(r)))
	return hex.EncodeToString(sum[:])
}

func ExitCode(decision string) int {
	switch decision {
	case "SHIP":
		return 0
	case "CHANGE":
		return 2
	case "RERUN":
		return 3
	default:
		return 1
	}
}
