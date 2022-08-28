package stringutil

import (
	"unicode/utf8"
)

func NumUTF8CharsInStringSlice(strings []string) int {
	sum := 0
	for _, str := range strings {
		sum += utf8.RuneCountInString(str)
	}
	return sum
}

// Adds "..." if the provided string is strictly longer than the defined limit
func ApplyEllipsis(str string, limit int) string {
	strLen := utf8.RuneCountInString(str)

	// dont modify the string if it doesnt exceed the limit
	if strLen <= limit {
		return str
	}

	// handle short text length limit
	if limit <= 3 {
		return Repeat(".", limit)
	}

	// truncate string
	diff := strLen - limit
	for i := 0; i < 3+diff; i++ {
		str = RemoveLastChar(str)
	}
	return str + "..."
}

func AlignLeft(str string, width int) string {
	strLen := utf8.RuneCountInString(str)

	if strLen == width {
		return str
	}

	if strLen > width {
		return ApplyEllipsis(str, width)
	}

	// add padding to the right if needed
	diff := width - strLen
	return str + Repeat(" ", diff)
}

func Repeat(str string, numTimes int) string {
	newstr := ""
	for i := 0; i < numTimes; i++ {
		newstr += str
	}
	return newstr
}

func RemoveLastChar(str string) string {
	if str == "" {
		return ""
	}

	_, size := utf8.DecodeLastRuneInString(str)
	return str[:len(str)-size]
}
