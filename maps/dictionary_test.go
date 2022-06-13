package maps

import "testing"

// Maps permitem armazenar items assim como um dicionário, com chave -> valor.

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "just a test"}

	result := dictionary.Search(dictionary, "test")
	expect := "just a test"

	CompareStrings(t, result, expect)
}

func CompareStrings(t *testing.T, result, expect string) {
	t.Helper()

	if result != expect {
		t.Errorf("resultado '%s', esperado '%s', dado '%s'", result, expect, "test")
	}
}
