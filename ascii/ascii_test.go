package ascii

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	result, err := Generate("hello", "standard")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	t.Log("Generated ascii: ", result)
	if result == "" {
		t.Errorf("expected output, got empty string")
	}
}
