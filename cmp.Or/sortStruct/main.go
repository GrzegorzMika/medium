package main

import (
	"cmp"
	"fmt"
	"slices"
)

type book struct {
	title  string
	author string
	year   int
	pages  int
}

func printBooks(books []book) {
	for _, book := range books {
		fmt.Printf("%s %s %d %d\n", book.title, book.author, book.year, book.pages)
	}
	fmt.Println("")
}

func sortBooks(books []book) []book {
	sortedBooks := books
	slices.SortFunc(sortedBooks, func(a, b book) int {
		return cmp.Or(
			cmp.Compare(a.year, b.year),
			cmp.Compare(a.pages, b.pages),
			cmp.Compare(a.author, b.author),
			cmp.Compare(a.title, b.title),
		)
	})
	return sortedBooks
}

func main() {
	books := []book{
		{"Domain-Driven Design: Tackling Complexity in the Heart of Software", "Eric Evans", 2003, 560},
		{"Clean Architecture: A Craftsman's Guide to Software Structure and Design", "Robert Martin", 2017, 432},
		{"Designing Data-Intensive Applications: The Big Ideas Behind Reliable, Scalable, and Maintainable Systems", "Martin Kleppmann", 2017, 590},
		{"Clean Code: A Handbook of Agile Software Craftsmanship", "Robert Martin", 2008, 464},
	}

	fmt.Println("Books before sorting:")
	printBooks(books)

	fmt.Println("Books after sorting:")
	printBooks(sortBooks(books))

}
