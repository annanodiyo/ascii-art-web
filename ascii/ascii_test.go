package ascii

import (
	"testing"
	// "ascii-art-web/ascii"
)

func TestGenerate(t *testing.T) {
	result, err := Generate("A", "standard")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result == "" {
		t.Errorf("expected output, got empty string")
	}
}
