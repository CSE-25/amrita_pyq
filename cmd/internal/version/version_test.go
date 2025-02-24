package version

import (
	"bytes"
	"testing"
)

func TestVersionCmd(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "SuccessVersion",
			args: nil,
			want: "Amrita Previous Year Questions v1.0.2\n",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var buf bytes.Buffer
			versionCmd.SetOut(&buf)
			versionCmd.SetErr(&buf)
			versionCmd.Run(versionCmd, tc.args)
			if got := buf.String(); got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}
