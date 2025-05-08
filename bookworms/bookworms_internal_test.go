package main

import "testing"

type testCase struct {
	bookwormsFile string
	want          []Bookworm
	wantErr       bool
}

var (
	handmaidsTale = Book{Author: "Margaret Atwood", Title: "The Handmaid's Tale"}
	oryxAndCrake  = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar    = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	janeEyre      = Book{Author: "Charlotte Brontë", Title: "Jane Eyre"}
	ilPrincipe    = Book{Author: "Niccolò Machiavelli", Title: "Il Principe"}
)

func TestLoadBookworms_Success(t *testing.T) {
	tests := map[string]testCase{
		"file exists": {
			bookwormsFile: "testdata/bookworms.json",
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			wantErr: false,
		},
		"file doesn't exist": {
			bookwormsFile: "testdata/no_file_here.json",
			want:          nil,
			wantErr:       true,
		},
		"invalid JSON": {
			bookwormsFile: "testdata/invalid.json",
			want:          nil,
			wantErr:       true,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := loadBookworms(testCase.bookwormsFile)

			if testCase.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nothing")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error, got one %s", err.Error())
			}

			if !equalBookworms(t, got, testCase.want) {
				t.Fatalf("different result: got %v, expected %v", got, testCase.want)
			}

		})

	}
}

func TestBooksCount(t *testing.T) {
	tt := map[string]struct {
		input []Bookworm
		want  map[Book]uint
	}{
		"nominal use case": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: map[Book]uint{
				handmaidsTale: 2,
				theBellJar:    1,
				oryxAndCrake:  1,
				janeEyre:      1,
			},
		},
		"no bookworms": {
			input: []Bookworm{},
			want:  map[Book]uint{},
		},
		"bookworm wihout books": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{}},
			},
			want: map[Book]uint{},
		},
		"bookworm with twice the same book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar, theBellJar}},
				{Name: "Peggy", Books: []Book{janeEyre, oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: map[Book]uint{
				handmaidsTale: 2,
				theBellJar:    2,
				janeEyre:      2,
				oryxAndCrake:  1,
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := booksCount(tc.input)
			if !equalBooksCount(t, got, tc.want) {
				t.Fatalf("got a different list of books: %v, expected: %v", got, tc.want)
			}
		})

	}
}

func TestFindCommonBooks(t *testing.T) {
	tt := map[string]struct {
		input []Bookworm
		want  []Book
	}{
		"everyone has read the same books": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, janeEyre}},
			},
			want: nil,
		},
		"one common book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: []Book{handmaidsTale},
		},
		"three bookworms have the same books on their shelves": {
			input: []Bookworm{
				{Name: "Peggy", Books: []Book{oryxAndCrake, ilPrincipe, janeEyre}},
				{Name: "Did", Books: []Book{janeEyre}},
				{Name: "Ali", Books: []Book{janeEyre, ilPrincipe}},
			},
			want: []Book{ilPrincipe, janeEyre},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := findCommonBooks(tc.input)
			if !equalBooks(t, got, sortBooks(tc.want)) {
				t.Fatalf("got different list of books: %v, expected %v", got, tc.want)
			}
		})
	}

}

func equalBookworms(t *testing.T, got, target []Bookworm) bool {
	t.Helper()

	if len(got) != len(target) {
		return false
	}

	for i := range got {
		if got[i].Name != target[i].Name {
			return false
		}

		if !equalBooks(t, got[i].Books, target[i].Books) {
			return false
		}
	}

	return true
}

func equalBooks(t *testing.T, gotBooks, targetBooks []Book) bool {
	t.Helper()

	if len(gotBooks) != len(targetBooks) {
		return false
	}

	for i, book := range gotBooks {
		if book != targetBooks[i] {
			return false
		}
	}

	return true
}

func equalBooksCount(t *testing.T, got, want map[Book]uint) bool {
	t.Helper()

	if len(got) != len(want) {
		return false
	}

	for book, wantCount := range want {
		count, ok := got[book]
		if !ok {
			return false
		}

		if count != wantCount {
			return false
		}
	}

	return true
}
