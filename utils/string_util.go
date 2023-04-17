package utils

import (
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type StringWrapper string

// clean the data using simple rules
// 1. Trim the spaces
// 2. Remove the special characters
// 3. Standardize the all the string to with the same case (capital the first letter of each word)
// 4. Normalize the string
func (s StringWrapper) Clean() StringWrapper {
	// 1. Trim the spaces
	s = StringWrapper(strings.TrimSpace(string(s)))
	// 2. Remove the special characters
	tmpString := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsSpace(r) {
			return r
		}
		return -1
	}, string(s))
	s = StringWrapper(tmpString)

	// 3. Normalize the string
	transformer := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	res, _, _ := transform.String(transformer, string(s))
	return StringWrapper(res)
}

func (s StringWrapper) ToUpper() StringWrapper {
	return StringWrapper(strings.ToUpper(s.ToString()))
}

func (s StringWrapper) ToTitleCase() StringWrapper {
	return StringWrapper(cases.Title(language.English, cases.Compact).String(s.ToString()))
}

func (s StringWrapper) ToLower() StringWrapper {
	return StringWrapper(strings.ToLower(s.ToString()))
}

// function that capitalize the first letter of each word after a period
// .e.g "hello world. this is a test" -> "Hello world. This is a test"
func (s StringWrapper) ToUpperAfterPeriod() StringWrapper {
	// split the string into words
	words := strings.Split(s.ToString(), " ")
	// capitalize the first letter of the first word after a period
	for i := 0; i < len(words); i++ {
		if strings.Contains(words[i], ".") {
			words[i+1] = StringWrapper(words[i+1]).ToTitleCase().ToString()
		}
	}

	// join the words back together
	return StringWrapper(strings.Join(words, " "))
}

func (s StringWrapper) Merge(s2 string) StringWrapper {
	// return the longest string
	if len(s.ToString()) > len(s2) {
		return s
	}

	return StringWrapper(s2)
}

func (s StringWrapper) ToString() string {
	return string(s)
}
