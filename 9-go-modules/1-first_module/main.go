// Exercise 09.01 – creating and using your first module

package main

import "bookutils/author"

func main() {
	// Create an author instance
	authorInstance := author.NewAuthor("Jane Doe", "jane@example.com")

	// Write and review a chapter
	chapterTitle := "Introduce to Go Modules"
	chapterContent := "Go modules provide a structured way to manage dependencies and improve code maintainability."
	authorInstance.WriteChapter(chapterTitle, chapterContent)
	authorInstance.ReviewChapter(chapterTitle, "This chapter looks great, but let's add some more examples.")
	authorInstance.FinalizeChapter(chapterTitle)
}
