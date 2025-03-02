package maps

import (
	"errors"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}

	got := Search(dictionary, "test")
	want := "this is just a test"

	assertStrings(t, got, want)
}

func TestSearchUpdated(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known key", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown key", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new key", func(t *testing.T) {
		dictionary := Dictionary{}
		newKey := "test"
		newValue := "this is just a test"
		dictionary.Add(newKey, newValue)

		assertDefinition(t, dictionary, newKey, newValue)
	})

	t.Run("existing key", func(t *testing.T) {
		newKey := "test"
		newValue := "this is just a test"
		dictionary := Dictionary{newKey: newValue}
		err := dictionary.Add(newKey, "new test")

		assertError(t, err, ErrKeyExists)
		assertDefinition(t, dictionary, newKey, newValue)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("existing key", func(t *testing.T) {
		existingKey := "test"
		existingValue := "this is just a test"
		dictionary := Dictionary{existingKey: existingValue}

		newValue := "new definition"
		err := dictionary.Update(existingKey, newValue)

		assertError(t, err, nil)

		assertDefinition(t, dictionary, existingKey, newValue)
	})

	t.Run("new key", func(t *testing.T) {
		existingKey := "test"
		existingValue := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(existingKey, existingValue)

		assertError(t, err, ErrKeyDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing key", func(t *testing.T) {
		existingKey := "test"
		existingValue := "this is just a test"

		dictionary := Dictionary{existingKey: existingValue}

		err := dictionary.Delete(existingKey)

		assertError(t, err, nil)

		_, err = dictionary.Search(existingKey)

		assertError(t, err, ErrNotFound)
	})

	t.Run("unknown key", func(t *testing.T) {
		key := "test"

		dictionary := Dictionary{}
		err := dictionary.Delete(key)
		assertError(t, err, ErrKeyDoesNotExist)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if !errors.Is(got, want) {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, key, definition string) {
	t.Helper()

	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("should find added key:", err)
	}
	assertStrings(t, got, definition)
}
