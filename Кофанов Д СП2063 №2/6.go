package main

import "fmt"

func collectUniqueTags(postsTags [][]string) map[string]bool {
	uniqueTags := make(map[string]bool)
	for _, tags := range postsTags {
		for _, tag := range tags {
			uniqueTags[tag] = true
		}
	}

	return uniqueTags
}

func main() {
	posts := [][]string{
		{"go", "backend"},
		{"git", "go", "tools"},
	}
	tagsMap := collectUniqueTags(posts)
	fmt.Println("Уникальные теги (из карты):")
	for tag := range tagsMap {
		fmt.Println("-", tag)
	}
}
