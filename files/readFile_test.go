package main

import (
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("test title", func(t *testing.T) {
		var want = Post{Title: "Hello, TDD world!",
			Description: "First post on our wonderful blog",
			Body:        "Hello world!\n" + "\n" + "The body of posts starts after the `---`",
			Tags:        []string{"tdd", "go"}}
		got, err := GetPost("./helloWorld.md")
		assertNoError(t, err)
		errorHelper(t, got, want)
	})
	t.Run("test longer title", func(t *testing.T) {
		var want = Post{Title: `Hello, TDD world!
This title is longer than the previous one.`,
			Description: `First post on our wonderful blog.
Some long stuff here.`,
			Body: `Hello world!
REadlly long body here.
erwerwerwer

asdasdasd`,
			Tags: []string{"tdd", "go"}}
		got, err := GetPost("./moreComplicated.md")
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

func errorHelper(t testing.TB, got, want Post) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s want %s", got, want)
	}
}
