package day14

import (
	"testing"
)

func TestLoadExample(t *testing.T) {
	template, rules := loadRules("example-input.txt")
	if template != "NNCB" {
		t.Errorf("bad template: %s", template)
	}
	if len(rules) != 16 {
		t.Errorf("expected 16 rules, got %d", len(rules))
	}
}

func TestPart1(t *testing.T) {
	quantity := processExpansion("example-input.txt", 10)
	if quantity != 1588 {
		t.Errorf("expected 1588, got %d", quantity)
	}
}

func TestPart2(t *testing.T) {
	quantity := processExpansion("example-input.txt", 40)
	if quantity != 2188189693529 {
		t.Errorf("expected 2188189693529, got %d", quantity)
	}
}
