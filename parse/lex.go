package lisp

import "unicode"

// This makes sure we don't append tokens that are empty
// This happens when there's whitespace because that
// normally isn't saved as part of the token
// This decision makes this lisp whitespace insensitive
func shouldAppendToken(token string) bool {
	return len(token) > 0
}

func Lex(str string) []string {
	result := []string{}

	var currentToken string = ""

	for _, rune := range str {
		t := string(rune)
		if unicode.IsSpace(rune) {
			if shouldAppendToken(currentToken) {
				result = append(result, string(currentToken))
				currentToken = ""
			}
		} else if t == "(" || t == ")" {
			if shouldAppendToken(currentToken) {
				result = append(result, string(currentToken))
				currentToken = ""
			}

			result = append(result, t)
		} else {
			currentToken += string(rune)
		}
	}

	if shouldAppendToken(currentToken) {
		result = append(result, string(currentToken))
	}

	return result
}
