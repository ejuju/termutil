package stringutil

import (
	"testing"
	"unicode/utf8"
)

func TestApplyEllipsis(t *testing.T) {
	t.Run("should apply ellipsis correctly", func(t *testing.T) {
		textTooLong := "i'm too long"
		limit := 5
		output := ApplyEllipsis(textTooLong, limit)
		outputLen := utf8.RuneCountInString(output)

		// output should not exceed limit
		if outputLen > limit {
			t.Fatal(output)
		}

		// output should finish with three dots when limit is longer or equal to 3 characters (and string is longer or equal to 3 characters)
		if outputLen >= 3 && output[outputLen-3:] != "..." && limit > 3 {
			t.Fatal(output)
		}
	})
}
