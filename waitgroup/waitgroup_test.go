package waitgroup

import (
	"testing"
)

func TestWaitPrintln(t *testing.T) {
	tests := []struct {
		n      int
		input  string
		output string
	}{
		{
			1, "This is the End", "This is the End",
		},
		{
			2, "Hold you breath and count to ten", "Hold you breath and count to ten",
		},
		{
			3, "It's the Skyfall, when you crumble", "It's the Skyfall, when you crumble",
		},
	}

	for _, tc := range tests {
		result := WaitPrintln(tc.input)
		if result != tc.output {
			t.Errorf("expected %s, got %s", tc.output, result)
		}
	}
}

func TestChannelPrintln(t *testing.T) {
	tests := []struct {
		n      int
		input  string
		output string
	}{
		{
			1, "This is the End", "This is the End",
		},
		{
			2, "Hold you breath and count to ten", "Hold you breath and count to ten",
		},
		{
			3, "It's the Skyfall, when you crumble", "It's the Skyfall, when you crumble",
		},
	}

	for _, tc := range tests {
		result := ChannelPrintln(tc.input)
		if result != tc.output {
			t.Errorf("expected %s, got %s", tc.output, result)
		}
	}
}
