package matcher_test

import (
	"search-engine/internal/matcher"
	"sort"
	"testing"
)

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Strings(a)
	sort.Strings(b)

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestSuggestions(t *testing.T) {
	m := matcher.NewMatcher()
	words := []string{"cat", "car", "cart", "dog", "dot"}
	m.InsertAll(words)

	input := "crt"
	maxDistance := 1

	suggestions := m.Suggestions(input, maxDistance)

	expected := []string{"cart", "cat"}
	if !equal(suggestions, expected) {
		t.Errorf("Expected %v, got %v", expected, suggestions)
	}
}

func TestMatch(t *testing.T) {
	m := matcher.NewMatcher()
	words := []string{"cat", "car", "cart", "dog", "dot"}
	m.InsertAll(words)

	input := "ca"

	suggestions := m.Match(input)

	expected := []string{"cat", "cart", "car"}
	if !equal(suggestions, expected) {
		t.Errorf("Expected %v, got %v", expected, suggestions)
	}

	input = "do"

	suggestions = m.Match(input)

	expected = []string{"dog", "dot"}

	if !equal(suggestions, expected) {
		t.Errorf("Expected %v, got %v", expected, suggestions)
	}
}
