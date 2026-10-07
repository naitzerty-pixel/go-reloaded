package main

import "testing"

func TestHexToDec(t *testing.T) {
	tests := []struct{ in, want string }{
		{"1E", "30"},
		{"42", "66"},
	}
	for _, tt := range tests {
		got, err := HexToDec(tt.in)
		if err != nil || got != tt.want {
			t.Errorf("HexToDec(%q) = %q, %v ; attendu %q", tt.in, got, err, tt.want)
		}
	}
}

func TestBinToDec(t *testing.T) {
	tests := []struct{ in, want string }{
		{"10", "2"},
	}
	for _, tt := range tests {
		got, err := BinToDec(tt.in)
		if err != nil || got != tt.want {
			t.Errorf("BinToDec(%q) = %q, %v ; attendu %q", tt.in, got, err, tt.want)
		}
	}
}
