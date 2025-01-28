package blogfiles_test

import (
	blogfiles "go-server/blog_files"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
world!
`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}
	posts, err := blogfiles.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	assertPost(t, posts[0], blogfiles.Post{Title: "Post 1",
		Description: "Description 1", Tags: []string{"tdd", "go"}, Body: `Hello
world!`})
}

func assertPost(t *testing.T, got blogfiles.Post, want blogfiles.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
