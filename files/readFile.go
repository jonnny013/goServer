package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func GetPost(filePath string) (Post, error) {
	var post = Post{
		Title:       "",
		Description: "",
		Body:        "",
		Tags:        []string{},
	}

	title, err := ParseSection(filePath, "Title: ")
	if err != nil {
		return post, err
	}
	post.Title = title
	fmt.Print("\n -----> \n", post)
	return post, nil
}

var postKeys = []string{"Title: ", "Description: ", "Body: ", "Tags: "}

func ParseSection(filePath, prefix string) (string, error) {
	var itemToReturn string
	file, err := os.Open(filePath)

	if err != nil {
		return "", err
	}
	defer file.Close()

	stillTitle := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, prefix) {
			itemToReturn = itemToReturn + strings.TrimPrefix(line, prefix)
			stillTitle = true
		} else if stillTitle {
			canContinue := true
			for _, item := range postKeys {
				if strings.HasPrefix(line, item) {
					canContinue = false
					stillTitle = false
				}
			}
			if canContinue {
				itemToReturn = itemToReturn + "\n" + line
			}
		}

	}
	return itemToReturn, nil
}
