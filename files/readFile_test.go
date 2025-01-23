package main

import (
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("test title", func(t *testing.T) {
		want := "Hello, TDD world!"
		got, err := ParseTitle("./helloWorld.md")
		assertNoError(t, err)
		errorHelper(t, got, want)
	})
	t.Run("test longer title", func(t *testing.T) {
		want := `Hello, TDD world!
This title is longer than the previous one.`
		got, err := ParseTitle("./moreComplicated.md")
		assertNoError(t, err)
		errorHelper(t, got, want)
	})
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func errorHelper(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
