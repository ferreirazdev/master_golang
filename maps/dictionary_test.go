package maps

import "testing"

// Maps permitem armazenar items assim como um dicionário, com chave -> valor.

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "just a test"}

	t.Run("known word", func(t *testing.T) {
		result, _ := dictionary.Search("test")
		expect := "just a test"

		CompareError(t, result, expect)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err != nil {
			t.Fatal("expected to get an error.")
		}
	})
}

func CompareError(t *testing.T, result, expect string) {
	t.Helper()

	if result != expect {
		t.Errorf("resultado '%s', esperado '%s', dado '%s'", result, expect, "test")
	}
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	key := "test"
	value := "just a test"

	dictionary.Add(key, value)

	compareValue(t, dictionary, key, value)
}

func compareValue(t *testing.T, dictionary Dictionary, key, value string) {
	t.Helper()

	result, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("word not found", err)
	}

	if result != value {
		t.Errorf("resultado '%s', esperado '%s'", result, value)
	}

}
