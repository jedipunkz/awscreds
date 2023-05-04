package cmd

import (
	"os"
	"strings"
	"testing"
)

func Test_getShell(t *testing.T) {
	// Skip the test if the environment is not bash
	if strings.Contains(os.Getenv("SHELL"), "fish") {
		t.Skip("Skipping test in fish environment")
	} else if strings.Contains(os.Getenv("SHELL"), "zsh") {
		t.Skip("Skipping test in zsh environment")
	}

	tests := []struct {
		name      string
		wantShell string
	}{
		{
			name:      "normal1",
			wantShell: "bash", // GitHub Actions shell is bash
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotShell := getShell(); gotShell != tt.wantShell {
				t.Errorf("getShell() = %v, want %v", gotShell, tt.wantShell)
			}
		})
	}
}
