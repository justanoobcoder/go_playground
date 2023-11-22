package dictionary

import (
	"errors"
	"go_playground/src/dictionary"
	"testing"
)

func assertStrings(t testing.TB, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("actual %q, expected %q", actual, expected)
	}
}

func assertError(t testing.TB, actual, expected error) {
	t.Helper()
	if !errors.Is(actual, expected) {
		t.Errorf("actual error %q, expected error %q", actual, expected)
	}
}

func assertDefinition(t testing.TB, dic dictionary.Dictionary, key, value string) {
	t.Helper()
	actual, err := dictionary.Search(dic, key)
	expected := value
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, actual, expected)
}

func TestSearch(t *testing.T) {
	dic := dictionary.Dictionary{"test": "This is a test"}
	t.Run("test search when known word should return successfully", func(t *testing.T) {
		actual, _ := dictionary.Search(dic, "test")
		expected := "This is a test"
		assertStrings(t, actual, expected)
	})

	t.Run("test search when unknown word should return error", func(t *testing.T) {
		_, err := dictionary.Search(dic, "unknown")
		expected := dictionary.ErrNotFound
		assertError(t, err, expected)
	})
}

func TestAdd(t *testing.T) {
	dic := dictionary.Dictionary{}
	key := "test"
	value := "this is a test"
	t.Run("test add when new word should be successful", func(t *testing.T) {
		err := dic.Add(key, value)
		assertError(t, err, nil)
		assertDefinition(t, dic, key, value)
	})

	t.Run("test add when existing word should return error", func(t *testing.T) {
		dic["key"] = "value"
		err := dic.Add("key", "value 2")
		assertError(t, err, dictionary.ErrWordExits)
	})
}

func TestUpdate(t *testing.T) {
	key := "test"
	value := "this is a test"
	newValue := "this is a test 2"
	dic := dictionary.Dictionary{key: value}
	t.Run("test update when unknown word should return error", func(t *testing.T) {
		err := dic.Update("unknown", newValue)
		assertError(t, err, dictionary.ErrNotFound)
	})

	t.Run("test update when existing word should be successful", func(t *testing.T) {
		err := dic.Update(key, newValue)
		assertError(t, err, nil)
		assertDefinition(t, dic, key, newValue)
	})
}

func TestDelete(t *testing.T) {
	key := "test"
	value := "this is a test"
	dic := dictionary.Dictionary{key: value}
	t.Run("test delete when existing word should be successful", func(t *testing.T) {
		dic.Delete(key)
		_, err := dictionary.Search(dic, key)
		if err != dictionary.ErrNotFound {
			t.Errorf("Expected %q to be deleted", key)
		}
	})
}
