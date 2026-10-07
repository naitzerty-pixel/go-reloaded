package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

// Compare Process(sampleN.txt) avec expectedN.txt.
func TestSamples(t *testing.T) {
	for i := 1; i <= 4; i++ {
		in, err := os.ReadFile(fmt.Sprintf("testdata/sample%d.txt", i))
		if err != nil {
			t.Fatal(err)
		}
		want, err := os.ReadFile(fmt.Sprintf("testdata/expected%d.txt", i))
		if err != nil {
			t.Fatal(err)
		}
		got := Process(string(in))
		if strings.TrimSpace(got) != strings.TrimSpace(string(want)) {
			t.Errorf("sample%d : obtenu %q, attendu %q", i, got, strings.TrimSpace(string(want)))
		}
	}
}
