package proprepro

const singleQuote = '\u0027' // unicode character for single quote

func UnsmartQuotes(s string) string {
	prose := []rune(s)
	for i, char := range prose {
		prose[i] = removeSmartQuotes(char)
	}
	return string(prose)
}

func RemoveTabs(s string) string {
	prose := []rune(s)
	for i, char := range prose {
		if char == '\t' {
			prose = removeItem(prose, i)
		}
	}
	return string(prose)
}

func removeItem(items []rune, i int) []rune {
	if i < 0 || i >= len(items) {
		return items
	}
	return append(items[:i], items[i+1:]...)
}

func removeSmartQuotes(char rune) rune {
	if char == '“' || char == '”' {
		return '"'
	}
	if char == '‘' || char == '’' {
		return singleQuote
	}
	return char
}