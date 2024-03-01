package benchmark

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"testing"
)

type book struct {
	title string
	year  int
}

var testCases = []struct {
	size int
}{
	{size: 100},
	{size: 1000},
	{size: 10000},
	{size: 1000000},
}

func bookGenerator(n int) []book {
	books := make([]book, n)
	for i := 0; i < n; i++ {
		books[i] = book{
			title: fmt.Sprintf("Book %d", i),
			year:  i,
		}
	}
	return books
}

//go:noinline
func sortUsingSort(books []book) {
	sort.Slice(books, func(a, b int) bool { return books[a].year < books[b].year })
}

//go:noinline
func sortUsingSlices(books []book) {
	slices.SortFunc(books, func(a, b book) int { return a.year - b.year })
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func BenchmarkSortUsingSort(b *testing.B) {
	for _, testCase := range testCases {
		b.Run(fmt.Sprintf("input size %d", testCase.size), func(b *testing.B) {
			for range b.N {
				b.StopTimer()
				books := bookGenerator(testCase.size)
				b.StartTimer()
				sortUsingSort(books)
				b.StopTimer()
			}
		})
	}
}

func BenchmarkSortUsingSlices(b *testing.B) {
	for _, testCase := range testCases {
		b.Run(fmt.Sprintf("input size %d", testCase.size), func(b *testing.B) {
			for range b.N {
				b.StopTimer()
				books := bookGenerator(testCase.size)
				b.StartTimer()
				sortUsingSlices(books)
			}
		})
	}
}
