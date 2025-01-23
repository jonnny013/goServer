package main

import (
	"bufio"
	"os"
	"strings"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func GetPost(filePath string) Post {
	var post = Post{
		Title: "",
		Description: "",
		Body: "",
		Tags: []string{},
	}
	return post
}

func ParseTitle(filePath string) (string, error) {
	var title string
	file, err := os.Open(filePath)

	if err != nil {
		return "", err
	}
	defer file.Close()

	stillTitle := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Title:") {
			title = title + strings.TrimPrefix(line, "Title: ")
			stillTitle = true
		} else if stillTitle {
			if !strings.HasPrefix(line, "Description:") {
				title = title + "\n" + line
			} else {
				stillTitle = false
			}

		}

	}
	return title, nil
}
