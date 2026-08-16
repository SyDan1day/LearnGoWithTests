package maps

import (
	"errors"
	"testing"
)

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("want %s but got %s", want, got)
	}
}

func assertErrors(t *testing.T, got, want error) {
	t.Helper()

	if !errors.Is(got, want) {
		t.Errorf("want error %s but got error %s", want, got)
	}
}

func assertDefinitions(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word")
	}
	if got != definition {
		t.Errorf("want %s but got %s", definition, got)
	}
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "This is a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "This is a test"

		if got != want {
			t.Errorf("want %s but got %s", want, got)
		}
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get a error but didn't")
		}
		assertErrors(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}

		dictionary.Add("test", "This is a test")

		word := "test"
		definition := "This is a test"

		assertDefinitions(t, dictionary, word, definition)
	})

	t.Run("exsited word", func(t *testing.T) {
		dictionary := Dictionary{"test": "This is a test"}

		err := dictionary.Add("test", "")
		if err == nil {
			t.Fatal("wanted a error but didn't get one")
		}

		assertErrors(t, err, ErrKeyExsited)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "This is a test"}

		dictionary.Update("test", "new definition")

		word := "test"
		definition := "new definition"

		assertDefinitions(t, dictionary, word, definition)
	})

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{"test2": "This is test2"}

		err := dictionary.Update("test", "")
		assertErrors(t, err, ErrNotFound)
	})
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{"test": "This is a test"}

	dictionary.Delete("test")
	_, err := dictionary.Search("test")

	assertErrors(t, err, ErrNotFound)
}
