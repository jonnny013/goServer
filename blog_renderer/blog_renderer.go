package main

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

type Post struct {
	Title, SanitizedTitle, Description, Body string
	Tags                                     []string
}
type PostRenderer struct {
	templ    *template.Template
	mdParser *parser.Parser
}
type PostViewModel struct {
	Title, SanitizedTitle, Description, Body string
	Tags                                     []string
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	return &PostRenderer{templ: templ, mdParser: parser}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", newPostVM(p, r))
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}

type postViewModel struct {
	Post
	HTMLBody template.HTML
}

func newPostVM(p Post, r *PostRenderer) postViewModel {
	vm := postViewModel{Post: p}
	vm.HTMLBody = template.HTML(markdown.ToHTML([]byte(p.Body), r.mdParser, nil))
	return vm
}


func main() {
	post := Post{Title: "Test", Description: "something goes here", Body: "something longer goes here", Tags: []string{"1", "2"}}
	renderer, err := NewPostRenderer()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	err = renderer.Render(os.Stdout, post)
	if err != nil {
		fmt.Println("Rendering error:", err)
	}
}