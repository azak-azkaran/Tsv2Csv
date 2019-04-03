package main

import (
	"strings"
	"testing"
)

func TestTsvToCsv(t *testing.T) {
	demo := strings.Split("test\tblub\ntest2\tblub2", "\n")

	out := TsvToCsv(demo)

	if strings.ContainsAny(out, "\t") {
		t.Error("Output still contains tab")
	}
}
