package ascii

import (
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	resultDefault, err := Generate("hello", "")
	resultBanner, err2 := Generate("hello", "standard")
	if err != nil || err2 != nil {
		t.Errorf("unexpected error: %v %v", err, err2)
	}
	t.Log("Generated ascii: ", resultDefault)
	t.Log("Generated ascii: ", resultBanner)
	if resultDefault != resultBanner {
		t.Errorf("expected default to match standard")
	}
}

func TestAllSamples(t *testing.T) {
	text := "Xiamara"
	all, err := AllSamples(text)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	standard, _ := Generate(text, "standard")
	shadow, _ := Generate(text, "shadow")
	thinkertoy, _ := Generate(text, "thinkertoy")

	t.Log("sample standard: ", standard, thinkertoy, shadow)

	if !strings.Contains(all, standard) {
		t.Errorf("missing standard output")
	}
	if !strings.Contains(all, shadow) {
		t.Errorf("missing standard output")
	}
	if !strings.Contains(all, thinkertoy) {
		t.Errorf("missing thnkertoy output")
	}
}
