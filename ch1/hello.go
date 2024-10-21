package main

import "fmt"

func main() {

	greeting := greet("en")
	fmt.Println(greeting)
}

// language represnets the language's code
type language string

// phrasebook holds greering for each supported language

var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε",    // Greek
	"en": "Hello world",      // English
	"fr": "Bonjour le monde", // French
	"he": "שלום עולם",        // Hebrew
	"ur": "دﻧﯿﺎ ﮨﯿﻠﻮ",        // Urdu
	"vi": "Xin chào Thế Giới",
}

func greet(l language) string {
	greeting, ok := phrasebook[l]

	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}
	return greeting
}
