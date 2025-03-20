package main

import "testing"

func TestGreet(t *testing.T) {
	type testCase struct {
		lang language
		want string
	}

	var tests = map[string]testCase{
		"English": {
			lang: language("en"),
			want: "Hello world",
		},
		"French": {
			lang: language("fr"),
			want: "Bonjour le monde",
		},
		"Akkadian, not supported": {
			lang: language("akk"),
			want: `unsupported language: "akk"`,
		},
		"Greek": {
			lang: language("el"),
			want: "Χαίρετε Κόσμε",
		},
		"Hebrew": {
			lang: language("he"),
			want: "שלום עולם",
		},
		"Urdu": {
			lang: language("ur"),
			want: "ہیلو دنیا",
		},
		"Vietnamese": {
			lang: language("vi"),
			want: "Xin chào Thế Giới",
		},
		"Empty": {
			lang: language(""),
			want: `unsupported language: ""`,
		},
	}


	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := greet(tc.lang)
			if got != tc.want {
				t.Errorf("expected: %q, got: %q", tc.want, got)
			}
		})
	}
}

func ExampleMain() {
	main()
	// Output:
	// Hello world

}
