package parsers

import (
	"Gman/gman"
	"bytes"
	"os"
	"testing"
)

// Mock CommandContext to isolate the test behavior

// Testing handleSourceCommand
func TestHandleSourceCommand(t *testing.T) {

	tests := []struct {
		args      []string
		expectErr bool
	}{
		{[]string{"1", "2", "N"}, false},
		{[]string{"1", "2"}, true},            // Too few arguments
		{[]string{"1", "invalid", "N"}, true}, // Invalid point arguments
		{[]string{"1", "2", "BB"}, true},      // invalid directoin args
	}

	for _, tt := range tests {
		ctx := CommandContext{}
		err := ctx.handleSourceCommand(tt.args)
		if tt.expectErr && err == nil {
			t.Errorf("expected error, but got none for args %v", tt.args)
		}
		if !tt.expectErr && err != nil {
			t.Errorf("did not expect error, but got %v for args %v", err, tt.args)
		}
	}
}

// Testing handleDestinationCommand
func TestHandleDestinationCommand(t *testing.T) {

	tests := []struct {
		args      []string
		expectErr bool
	}{
		{[]string{"3", "4"}, false},
		{[]string{"3"}, true},            // Too few arguments
		{[]string{"invalid", "4"}, true}, // Invalid point
	}

	for _, tt := range tests {
		ctx := &CommandContext{
			gman: &gman.Gman{}, // mock gman setup
		}
		err := ctx.handleDestinationCommand(tt.args)
		if tt.expectErr && err == nil {
			t.Errorf("expected error, but got none for args %v", tt.args)
		}
		if !tt.expectErr && err != nil {
			t.Errorf("did not expect error, but got %v for args %v", err, tt.args)
		}
	}
}

func TestHandlePrintPowerCommand(t *testing.T) {

	// mocking
	ctx := &CommandContext{
		gman: &gman.Gman{Power: 100},
	}

	var buf bytes.Buffer

	oldStdout := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	// reading
	err := ctx.handlePrintPowerCommand()

	// closing the
	w.Close()
	os.Stdout = oldStdout
	buf.ReadFrom(r)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedOutput := "POWER  100"
	if buf.String() != expectedOutput {
		t.Errorf("expected output %q, got %q", expectedOutput, buf.String())
	}
}
