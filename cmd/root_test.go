package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sundowndev/covermyass/v2/test"
	"testing"
)

func TestRootCmd_ValidFlags(t *testing.T) {
	rootCmd := NewRootCmd()
	rootCmd.RunE = func(_ *cobra.Command, args []string) error { return nil }

	cases := []struct {
		args []string
	}{
		{args: []string{}},
		{args: []string{"-n", "5"}},
		{args: []string{"--write"}},
		{args: []string{"-z"}},
		{args: []string{"-l", "--no-read-only"}},
		{args: []string{"-f", "/var/log/1/*.log", "--filter", "/var/log/2/*.log"}},
	}

	for _, tt := range cases {
		output, err := test.Execute(rootCmd, tt.args...)
		if output != "" {
			t.Errorf("Unexpected output: %v", output)
		}
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}

func TestRootCmd_InvalidFlags(t *testing.T) {
	rootCmd := NewRootCmd()
	rootCmd.RunE = func(_ *cobra.Command, args []string) error { return nil }

	cases := []struct {
		args     []string
		expected string
	}{
		{args: []string{"-t"}, expected: "unknown shorthand flag: 't' in -t"},
		{args: []string{"-n"}, expected: "flag needs an argument: 'n' in -n"},
		{args: []string{"-f"}, expected: "flag needs an argument: 'f' in -f"},
	}

	for _, tt := range cases {
		_, err := test.Execute(rootCmd, tt.args...)
		if err == nil {
			t.Errorf("Invalid arg should generate error")
		}
		if err.Error() != tt.expected {
			t.Errorf("Expected '%v', got '%v'", tt.expected, err)
		}
	}
}
