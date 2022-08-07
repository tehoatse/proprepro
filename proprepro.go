package proprepro

func unsmartQuotes(s string) string {

	singleQuote := '\u0027' // singleQuote unicode character, I think?
	prose := []rune(s)

	for i, char := range prose {
		switch char {
		case '“':
			prose[i] = '"'
		case '”':
			prose[i] = '"'
		case '‘':
			prose[i] = singleQuote 
		case '’':
			prose[i] = singleQuote
		}
	}
	return string(prose)
}

func removeTabs(s string) string {

	prose := []rune(s)

	for i, char := range prose {
		if char == '\t' {
			prose[i] = ' '
		}
	}

	return string(prose)

}