package main

import (
	"strings"
	"testing"
	"time"
)

const host = "0.beevik-ntp.pool.ntp.org"

// проверка ошибок и их обработка
func isNil(t *testing.T, err error) bool {
	switch {
	case err == nil:
		return true
	case strings.Contains(err.Error(), "timeout"):
		t.Errorf("[%s] Query timeout: %s", host, err)
		return false
	default:
		t.Errorf("[%s] Query failed: %s", host, err)
		return false
	}
}

func TestExactTime(t *testing.T) {
	tm, err := exactTime(host)
	now := time.Now()
	if isNil(t, err) {
		t.Logf("Local Time %v\n", now)
		t.Logf("~True Time %v\n", tm)
		t.Logf("Offset %v\n", tm.Sub(now))
	}
}
