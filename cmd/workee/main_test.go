package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		wantCode   int
		wantStdout string
		wantStderr string
	}{
		{name: "default", wantStdout: "workee\n"},
		{name: "short help", args: []string{"-h"}, wantStderr: "Usage: workee"},
		{name: "long help", args: []string{"--help"}, wantStderr: "Usage: workee"},
		{name: "unknown option", args: []string{"--unknown"}, wantCode: 2, wantStderr: "flag provided but not defined"},
		{name: "unexpected argument", args: []string{"unknown"}, wantCode: 2, wantStderr: `unexpected argument "unknown"`},
		{name: "argument after separator", args: []string{"--", "unknown"}, wantCode: 2, wantStderr: `unexpected argument "unknown"`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			if got := run(tt.args, &stdout, &stderr); got != tt.wantCode {
				t.Errorf("exit code = %d, want %d", got, tt.wantCode)
			}
			if got := stdout.String(); got != tt.wantStdout {
				t.Errorf("stdout = %q, want %q", got, tt.wantStdout)
			}
			if got := stderr.String(); (tt.wantStderr == "" && got != "") || !strings.Contains(got, tt.wantStderr) {
				t.Errorf("stderr = %q, want substring %q", got, tt.wantStderr)
			}
		})
	}
}
