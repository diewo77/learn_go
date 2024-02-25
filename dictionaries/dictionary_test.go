package main

import (
	"testing"
)

func TestDictionary(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("Known", func(t *testing.T) {
		got, _ := dictionary.Search("test")

		assertString(t, got, "this is just a test")
	})
	t.Run("Unknown", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrNotFound

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("Add new", func(t *testing.T) {
		dictionary := Dictionary{}

		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("Existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update existing", func(t *testing.T) {

		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Update(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("Update not existing", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Update("nop", definition)

		assertError(t, err, ErrUpdateWordDoesNotExist)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete existing", func(t *testing.T) {

		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Delete(word)

		assertError(t, err, nil)
		assertDeleted(t, dictionary, word)
	})
	t.Run("Delete not existing", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Delete("nop")

		assertError(t, err, ErrDeleteWordDoesNotExist)
		assertDefinition(t, dictionary, word, definition)
	})
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, defintion string) {
	t.Helper()

	got, err := dictionary.Search("test")

	if err != nil {
		t.Fatal("should find added world:", err)
	}
	assertString(t, got, defintion)
}

func assertDeleted(t testing.TB, dictionary Dictionary, word string) {
	t.Helper()

	_, err := dictionary.Search("test")

	if err == nil {
		t.Fatal("should find added world:", err)
	}
}
