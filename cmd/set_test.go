package cmd

import "testing"

func Test_getShell(t *testing.T) {
	tests := []struct {
		name      string
		wantShell string
	}{
		{
			name:      "normal1",
			wantShell: "fish",
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
