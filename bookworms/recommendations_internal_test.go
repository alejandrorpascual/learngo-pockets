package main

import (
	"reflect"
	"testing"
)

var (
	theHelp   = Book{Author: "Katheryn Stockett", Title: "The Help"}
	fairyTale = Book{Author: "Stephen King", Title: "Fairy Tale"}
)

func TestRecommendOtherBooks(t *testing.T) {
	type testCase struct {
		bookworms []Bookworm
		want      []Bookworm
	}

	tt := map[string]testCase{
		"No common books": {
			bookworms: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, janeEyre}},
			},
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{}},
				{Name: "Peggy", Books: []Book{}},
			},
		},
		"Two bookworms with a common book and several books on their respective shelves": {
			bookworms: []Bookworm{
				{Name: "Peggy", Books: []Book{oryxAndCrake, janeEyre, ilPrincipe}},
				{Name: "Did", Books: []Book{janeEyre, theBellJar}},
			},
			want: []Bookworm{
				{Name: "Peggy", Books: []Book{theBellJar}},
				{Name: "Did", Books: []Book{oryxAndCrake, ilPrincipe}},
			},
		},
		"Two bookworms with a common book and one has only one book on shelf": {
			bookworms: []Bookworm{
				{Name: "Peggy", Books: []Book{janeEyre}},
				{Name: "Did", Books: []Book{janeEyre, theBellJar}},
			},
			want: []Bookworm{
				{Name: "Peggy", Books: []Book{theBellJar}},
				{Name: "Did", Books: []Book{}},
			},
		},
		"Three bookworms with two common books and several books on their respective shelves": {
			bookworms: []Bookworm{
				{Name: "Peggy", Books: []Book{theHelp, janeEyre}},
				{Name: "Did", Books: []Book{janeEyre, theHelp, fairyTale}},
				{Name: "Ali", Books: []Book{janeEyre, ilPrincipe, theHelp}},
			},
			want: []Bookworm{
				{Name: "Peggy", Books: []Book{ilPrincipe, fairyTale}},
				{Name: "Did", Books: []Book{ilPrincipe}},
				{Name: "Ali", Books: []Book{fairyTale}},
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := recommendOtherBooks(tc.bookworms)
			if !equals(t, got, tc.want) {
				t.Fatalf("recommendOtherBooks() = %v, want %v", got, tc.want)
			}
		})
	}
}

func equals(t *testing.T, got []Bookworm, want []Bookworm) bool {
	t.Helper()

	for i := range got {
		if !reflect.DeepEqual(got[i], want[i]) {
			return false
		}
	}

	return true
}
