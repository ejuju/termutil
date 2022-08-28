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

type BoxDrawingCharacterSet struct {
	CharHorizontalLine                     string
	CharVerticalLine                       string
	CharTopLeftCornerLine                  string
	CharTopRightCornerLine                 string
	CharBottomLeftCornerLine               string
	CharBottomRightCornerLine              string
	CharVerticalRightIntersectionLine      string
	CharVerticalLeftIntersectionLine       string
	CharHorizontalBottomIntersectionLine   string
	CharHorizontalTopIntersectionLine      string
	CharVerticalHorizontalIntersectionLine string
}

var BoxDrawingSetSingleLine = BoxDrawingCharacterSet{
	CharHorizontalLine:                     charHorizontalSingleLine,
	CharVerticalLine:                       charVerticalSingleLine,
	CharTopLeftCornerLine:                  charTopLeftCornerSingleLine,
	CharTopRightCornerLine:                 charTopRightCornerSingleLine,
	CharBottomLeftCornerLine:               charBottomLeftCornerSingleLine,
	CharBottomRightCornerLine:              charBottomRightCornerSingleLine,
	CharVerticalRightIntersectionLine:      charVerticalRightIntersectionSingleLine,
	CharVerticalLeftIntersectionLine:       charVerticalLeftIntersectionSingleLine,
	CharHorizontalBottomIntersectionLine:   charHorizontalBottomIntersectionSingleLine,
	CharHorizontalTopIntersectionLine:      charHorizontalTopIntersectionSingleLine,
	CharVerticalHorizontalIntersectionLine: charVerticalHorizontalIntersectionSingleLine,
}

var BoxDrawingSetDoubleLine = BoxDrawingCharacterSet{
	CharHorizontalLine:                     charHorizontalDoubleLine,
	CharVerticalLine:                       charVerticalDoubleLine,
	CharTopLeftCornerLine:                  charTopLeftCornerDoubleLine,
	CharTopRightCornerLine:                 charTopRightCornerDoubleLine,
	CharBottomLeftCornerLine:               charBottomLeftCornerDoubleLine,
	CharBottomRightCornerLine:              charBottomRightCornerDoubleLine,
	CharVerticalRightIntersectionLine:      charVerticalRightIntersectionDoubleLine,
	CharVerticalLeftIntersectionLine:       charVerticalLeftIntersectionDoubleLine,
	CharHorizontalBottomIntersectionLine:   charHorizontalBottomIntersectionDoubleLine,
	CharHorizontalTopIntersectionLine:      charHorizontalTopIntersectionDoubleLine,
	CharVerticalHorizontalIntersectionLine: charVerticalHorizontalIntersectionDoubleLine,
}

var BoxDrawingSetHorizontalOnly = BoxDrawingCharacterSet{
	CharHorizontalLine:                     charHorizontalSingleLine,
	CharVerticalLine:                       "",
	CharTopLeftCornerLine:                  "",
	CharTopRightCornerLine:                 "",
	CharBottomLeftCornerLine:               "",
	CharBottomRightCornerLine:              "",
	CharVerticalRightIntersectionLine:      "",
	CharVerticalLeftIntersectionLine:       "",
	CharHorizontalBottomIntersectionLine:   charHorizontalSingleLine,
	CharHorizontalTopIntersectionLine:      charHorizontalSingleLine,
	CharVerticalHorizontalIntersectionLine: charHorizontalSingleLine,
}

var BoxDrawingSetNone = BoxDrawingCharacterSet{}
