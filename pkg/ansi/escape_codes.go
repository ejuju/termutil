package ansi

import "strconv"

func Wrap(original string, ansiEscapeCode ...EscapeCode) string {
	output := original
	for _, escapecode := range ansiEscapeCode {
		output = "\033[" + strconv.Itoa(int(escapecode)) + "m" + output + "\033[m"
	}
	return output
}

type EscapeCode int

const (
	EscapeCodeBold             EscapeCode = 1
	EscapeCodeItalic           EscapeCode = 3
	EscapeCodeUnderline        EscapeCode = 4
	EscapeCodeBlackText        EscapeCode = 30
	EscapeCodeBlackBackground  EscapeCode = 40
	EscapeCodeRedText          EscapeCode = 31
	EscapeCodeRedBackground    EscapeCode = 41
	EscapeCodeGreenText        EscapeCode = 32
	EscapeCodeGreenBackground  EscapeCode = 42
	EscapeCodeGreyText         EscapeCode = 90
	EscapeCodeGreyBackground   EscapeCode = 100
	EscapeCodeYellowText       EscapeCode = 93
	EscapeCodeYellowBackground EscapeCode = 103
	EscapeCodeBlueText         EscapeCode = 94
	EscapeCodeBlueBackground   EscapeCode = 104
)
