package main

import "testing"

// func ExampleMain() {
// 	main()
// 	// Output:
// 	// Hello world

// }

// func TestGreet(t *testing.T) {
// 	want := "Hello world"

// 	got := greet()

// 	if got != want {
// 		// mark this test as failed
// 		t.Errorf("expected: %q, got: %q", got, want)
// 	}
// }

// func TestGreet_English(t *testing.T) {

// 	lang := language("en")
// 	want := "Hello world"

// 	got := greet(lang)

// 	if got != want {
// 		// mark this test as failed
// 		t.Errorf("expected: %q, got: %q", want, got)
// 	}
// }
// func TestGreet_French(t *testing.T) {

// 	lang := language("fr")
// 	want := "Bonjour le monde"

// 	got := greet(lang)

// 	if got != want {
// 		// mark this test as failed
// 		t.Errorf("expected: %q, got: %q", want, got)
// 	}
// }
// func TestGreet_Akkadian(t *testing.T) {

// 	lang := language("akk")
// 	want := ""

// 	got := greet(lang)

// 	if got != want {
// 		// mark this test as failed
// 		t.Errorf("expected: %q, got: %q", want, got)
// 	}
// }

func TestGreet(t *testing.T) {
	type testCase struct {
		lang language
		want string
	}

	var tests = map[string]testCase{
		"English": {
			lang: "en",
			want: "Hello world",
		},
		"French":                  {lang: "fr", want: "Bonjour le monde"},
		"Akkadian, not supported": {lang: "akk", want: `unsupported language: "akk"`},
		"Greek":                   {lang: "el", want: "Χαίρετε Κόσμε"},
		"Hebrew":                  {lang: "he", want: "שלום עולם"},
		"Urdu":                    {lang: "ur", want: "دﻧﯿﺎ ﮨﯿﻠﻮ"},
		"Vietnamese":              {lang: "vi", want: "Xin chào Thế Giới"},
		"Empty":                   {lang: "", want: `unsupported language: ""`},
	}

	// range over all the scenarios

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := greet(tc.lang)

			if got != tc.want {
				t.Errorf("expected: %q, got: %q", tc.want, got)
			}
		})
	}

}
