package main

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		fakeDictionary := Dictionary{"test": "this is just a test"}
		want := "this is just a test"
		got, err := fakeDictionary.Search("test")
		if err != nil {
			t.Fatal("should find added word:", err)
		}
		assertStringsEqual(t, want, got)
	})
	t.Run("unknown word", func(t *testing.T) {
		fakeDictionary := Dictionary{"test": "this is just a test"}
		_, err := fakeDictionary.Search("invalid")
		assertError(t, err)
		assertStringsEqual(t, err.Error(), ErrNotFound.Error())
	})
}

func TestAdd(t *testing.T) {
	t.Run("adding new word", func(t *testing.T) {
		fakeDictionary := Dictionary{}
		addErr := fakeDictionary.Add("test", "this is just a test")
		want := "this is just a test"
		got, searchErr := fakeDictionary.Search("test")
		assertNoError(t, addErr)
		assertNoError(t, searchErr)
		assertStringsEqual(t, want, got)
	})
	t.Run("trying to add to a nil map", func(t *testing.T) {
		var fakeDictionary Dictionary
		addErr := fakeDictionary.Add("test", "this is just a test")
		assertError(t, addErr)
	})
	t.Run("adding existing word", func(t *testing.T) {
		fakeDictionary := Dictionary{"test": "this is just a test"}
		addErr := fakeDictionary.Add("test", "this is just a test")
		want := ErrWordExists
		got := addErr
		assertStringsEqual(t, want.Error(), got.Error())
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Updating existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		fakeDictionary := Dictionary{word: definition}
		newDefinition := "new definition"
		fakeDictionary.Update(word, newDefinition)

		value, _ := fakeDictionary.Search(word)

		if value != newDefinition {
			t.Errorf("Expected %q but got %q", newDefinition, value)
		}
	})
	t.Run("Updating non-existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		fakeDictionary := Dictionary{word: definition}
		newDefinition := "new definition"
		updateErr := fakeDictionary.Update("invalid", newDefinition)

		if updateErr != ErrNotFound {
			t.Errorf("Expected %q but got %q", ErrNotFound, updateErr)
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("Deleting existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		fakeDictionary := Dictionary{word: definition}
		err := fakeDictionary.Delete(word)
		if err != nil {
			t.Errorf("Expected no error but got %v", err)
		}
		if len(fakeDictionary) != 0 {
			t.Errorf("Expected dictionary to be empty but got %v", fakeDictionary)
		}
	})
	t.Run("Deleting non-existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		fakeDictionary := Dictionary{word: definition}
		err := fakeDictionary.Delete("invalid")

		if err != ErrNotFound {
			t.Errorf("Expected %q but got %q", ErrNotFound, err)
		}
	})
}

func assertStringsEqual(t testing.TB, want, got string) {
	t.Helper()
	if want != got {
		t.Errorf("got %q wanted %q", got, want)
	}
}

func assertError(t testing.TB, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected to get an error")
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		message := fmt.Sprintf("expected not to get an error, got %v", err)
		t.Fatal(message)
	}
}
