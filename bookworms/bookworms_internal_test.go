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
