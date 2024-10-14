package ipprotocols

import (
	"testing"
)

func TestGetName(t *testing.T) {
	tests := []struct {
		number int
		name   string
		err    bool
	}{
		{6, "TCP", false},
		{17, "UDP", false},
		{47, "GRE", false},
		{255, "Reserved", false},
		{999, "", true}, // Invalid protocol number
	}

	for _, tt := range tests {
		name, err := GetProtocolName(tt.number)
		if tt.err && err == nil {
			t.Errorf("expected an error for number %d", tt.number)
		}
		if !tt.err && name != tt.name {
			t.Errorf("expected %s for number %d, got %s", tt.name, tt.number, name)
		}
	}
}

func TestGetNumber(t *testing.T) {
	tests := []struct {
		name   string
		number int
		err    bool
	}{
		{"TCP", 6, false},
		{"UDP", 17, false},
		{"GRE", 47, false},
		{"Reserved", 255, false},
		{"INVALID", 0, true}, // Invalid protocol name
	}

	for _, tt := range tests {
		number, err := GetProtocolNumber(tt.name)
		if tt.err && err == nil {
			t.Errorf("expected an error for name %s", tt.name)
		}
		if !tt.err && number != tt.number {
			t.Errorf("expected %d for name %s, got %d", tt.number, tt.name, number)
		}
	}
}
