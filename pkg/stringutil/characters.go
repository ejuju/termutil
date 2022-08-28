package stringutil

const (
	// General purpose
	CharWhiteSpace = " "
	CharLineReturn = "\n"

	// Box drawing - single line
	charHorizontalSingleLine                     = "─"
	charVerticalSingleLine                       = "│"
	charTopLeftCornerSingleLine                  = "┌"
	charBottomLeftCornerSingleLine               = "└"
	charTopRightCornerSingleLine                 = "┐"
	charBottomRightCornerSingleLine              = "┘"
	charHorizontalTopIntersectionSingleLine      = "┴"
	charVerticalHorizontalIntersectionSingleLine = "┼"
	charHorizontalBottomIntersectionSingleLine   = "┬"
	charVerticalRightIntersectionSingleLine      = "├"
	charVerticalLeftIntersectionSingleLine       = "┤"

	// Box drawing - bold line
	charHorizontalBoldLine                     = "━"
	charVerticalBoldLine                       = "┃"
	charTopLeftCornerBoldLine                  = "┏"
	charTopRightCornerBoldLine                 = "┓"
	charBottomLeftCornerBoldLine               = "┗"
	charBottomRightCornerBoldLine              = "┛"
	charVerticalRightIntersectionBoldLine      = "┣"
	charVerticalLeftIntersectionBoldLine       = "┫"
	charHorizontalBottomIntersectionBoldLine   = "┳"
	charHorizontalTopIntersectionBoldLine      = "┻"
	charVerticalHorizontalIntersectionBoldLine = "╋"

	// Box drawing - double line
	charHorizontalDoubleLine                     = "═"
	charVerticalDoubleLine                       = "║"
	charTopLeftCornerDoubleLine                  = "╔"
	charTopRightCornerDoubleLine                 = "╗"
	charBottomLeftCornerDoubleLine               = "╚"
	charBottomRightCornerDoubleLine              = "╝"
	charVerticalRightIntersectionDoubleLine      = "╠"
	charVerticalLeftIntersectionDoubleLine       = "╣"
	charHorizontalBottomIntersectionDoubleLine   = "╦"
	charHorizontalTopIntersectionDoubleLine      = "╩"
	charVerticalHorizontalIntersectionDoubleLine = "╬"
)

type BoxDrawingCharacters struct {
	CharHorizontalLine                   string
	CharVerticalLine                     string
	CharTopLeftCornerLine                string
	CharTopRightCornerLine               string
	CharBottomLeftCornerLine             string
	CharBottomRightCornerLine            string
	CharVerticalRightIntersectionLine    string
	CharVerticalLeftIntersectionLine     string
	CharHorizontalBottomIntersectionLine string
	CharHorizontalTopIntersectionLine    string
	CharCrossIntersectionLine            string
}

var BoxDrawingCharactersSingleLine = BoxDrawingCharacters{
	CharHorizontalLine:                   charHorizontalSingleLine,
	CharVerticalLine:                     charVerticalSingleLine,
	CharTopLeftCornerLine:                charTopLeftCornerSingleLine,
	CharTopRightCornerLine:               charTopRightCornerSingleLine,
	CharBottomLeftCornerLine:             charBottomLeftCornerSingleLine,
	CharBottomRightCornerLine:            charBottomRightCornerSingleLine,
	CharVerticalRightIntersectionLine:    charVerticalRightIntersectionSingleLine,
	CharVerticalLeftIntersectionLine:     charVerticalLeftIntersectionSingleLine,
	CharHorizontalBottomIntersectionLine: charHorizontalBottomIntersectionSingleLine,
	CharHorizontalTopIntersectionLine:    charHorizontalTopIntersectionSingleLine,
	CharCrossIntersectionLine:            charVerticalHorizontalIntersectionSingleLine,
}

var BoxDrawingCharactersBoldLine = BoxDrawingCharacters{
	CharHorizontalLine:                   charHorizontalBoldLine,
	CharVerticalLine:                     charVerticalBoldLine,
	CharTopLeftCornerLine:                charTopLeftCornerBoldLine,
	CharTopRightCornerLine:               charTopRightCornerBoldLine,
	CharBottomLeftCornerLine:             charBottomLeftCornerBoldLine,
	CharBottomRightCornerLine:            charBottomRightCornerBoldLine,
	CharVerticalRightIntersectionLine:    charVerticalRightIntersectionBoldLine,
	CharVerticalLeftIntersectionLine:     charVerticalLeftIntersectionBoldLine,
	CharHorizontalBottomIntersectionLine: charHorizontalBottomIntersectionBoldLine,
	CharHorizontalTopIntersectionLine:    charHorizontalTopIntersectionBoldLine,
	CharCrossIntersectionLine:            charVerticalHorizontalIntersectionBoldLine,
}

var BoxDrawingCharactersDoubleLine = BoxDrawingCharacters{
	CharHorizontalLine:                   charHorizontalDoubleLine,
	CharVerticalLine:                     charVerticalDoubleLine,
	CharTopLeftCornerLine:                charTopLeftCornerDoubleLine,
	CharTopRightCornerLine:               charTopRightCornerDoubleLine,
	CharBottomLeftCornerLine:             charBottomLeftCornerDoubleLine,
	CharBottomRightCornerLine:            charBottomRightCornerDoubleLine,
	CharVerticalRightIntersectionLine:    charVerticalRightIntersectionDoubleLine,
	CharVerticalLeftIntersectionLine:     charVerticalLeftIntersectionDoubleLine,
	CharHorizontalBottomIntersectionLine: charHorizontalBottomIntersectionDoubleLine,
	CharHorizontalTopIntersectionLine:    charHorizontalTopIntersectionDoubleLine,
	CharCrossIntersectionLine:            charVerticalHorizontalIntersectionDoubleLine,
}

var BoxDrawingCharactersHorizontalOnly = BoxDrawingCharacters{
	CharHorizontalLine:                   charHorizontalSingleLine,
	CharVerticalLine:                     "",
	CharTopLeftCornerLine:                "",
	CharTopRightCornerLine:               "",
	CharBottomLeftCornerLine:             "",
	CharBottomRightCornerLine:            "",
	CharVerticalRightIntersectionLine:    "",
	CharVerticalLeftIntersectionLine:     "",
	CharHorizontalBottomIntersectionLine: charHorizontalSingleLine,
	CharHorizontalTopIntersectionLine:    charHorizontalSingleLine,
	CharCrossIntersectionLine:            charHorizontalSingleLine,
}

var BoxDrawingCharactersWhiteSpace = BoxDrawingCharacters{
	CharHorizontalLine:                   " ",
	CharVerticalLine:                     "",
	CharTopLeftCornerLine:                "",
	CharTopRightCornerLine:               "",
	CharBottomLeftCornerLine:             "",
	CharBottomRightCornerLine:            "",
	CharVerticalRightIntersectionLine:    "",
	CharVerticalLeftIntersectionLine:     "",
	CharHorizontalBottomIntersectionLine: " ",
	CharHorizontalTopIntersectionLine:    " ",
	CharCrossIntersectionLine:            " ",
}
