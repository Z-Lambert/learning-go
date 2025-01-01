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
		assertStrings(t, want, got)
	})
	t.Run("unknown word", func(t *testing.T) {
		fakeDictionary := Dictionary{"test": "this is just a test"}
		_, err := fakeDictionary.Search("invalid")
		assertError(t, err)
		assertStrings(t, err.Error(), ErrNotFound.Error())
	})
}

func TestAdd(t *testing.T) {
	t.Run("adding word", func(t *testing.T) {
		fakeDictionary := Dictionary{}
		addErr := fakeDictionary.Add("test", "this is just a test")
		want := "this is just a test"
		got, searchErr := fakeDictionary.Search("test")
		assertNoError(t, addErr)
		assertNoError(t, searchErr)
		assertStrings(t, want, got)
	})
	t.Run("trying to add to a nil map", func(t *testing.T) {
		var fakeDictionary Dictionary
		addErr := fakeDictionary.Add("test", "this is just a test")
		assertError(t, addErr)
	})
}

func assertStrings(t testing.TB, want, got string) {
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
