package tui

import (
	"errors"
	"strconv"

	"github.com/ejuju/termutil/pkg/ansi"
	"github.com/ejuju/termutil/pkg/stringutil"
)

type Table struct {
	config TableConfig
}

type TableConfig struct {
	ColumnTitles          []string
	ColumnWidths          []int
	Rows                  [][]string
	CellHorizontalPadding int
	ColumnTitleStyle      []ansi.EscapeCode
	RowStyle              []ansi.EscapeCode
	CharacterSet          stringutil.BoxDrawingCharacters
}

const (
	DefaultTableColumnWidth = 14
)

func NewTable(config TableConfig) (*Table, error) {

	// fail if no rows are provided
	if len(config.Rows) == 0 {
		return nil, ErrNoDataRows
	}

	// fail if no column titles are provided
	if len(config.ColumnTitles) == 0 {
		return nil, ErrNoColumnTitles
	}

	// set default column width if needed
	if len(config.ColumnWidths) == 0 {
		config.ColumnWidths = make([]int, len(config.ColumnTitles))
		for i := 0; i < len(config.ColumnTitles); i++ {
			config.ColumnWidths[i] = DefaultTableColumnWidth
		}
	}

	// check for mismatched column width number
	if len(config.ColumnWidths) != len(config.ColumnTitles) {
		return nil, ErrInvalidColumnWidth
	}

	// check that column width are strictly positive and set defaults if needed
	for i, width := range config.ColumnWidths {
		if width <= 0 {
			config.ColumnWidths[i] = DefaultTableColumnWidth
		}
	}

	// check for missing cells
	for i, row := range config.Rows {
		if len(row) != len(config.ColumnTitles) {
			return nil, &ErrMissingCell{
				RowIndex:   i,
				Difference: len(row) - len(config.ColumnTitles),
			}
		}
	}

	return &Table{config: config}, nil
}

var (
	ErrNoColumnTitles     = errors.New("no column titles were provided")
	ErrNoDataRows         = errors.New("no data rows were provided")
	ErrInvalidColumnWidth = errors.New("no data rows were provided")
)

type ErrMissingCell struct {
	RowIndex   int
	Difference int
}

func (err *ErrMissingCell) Error() string {
	return "missing cells (difference: " + strconv.Itoa(err.Difference) + ") at row index: " + strconv.Itoa(err.RowIndex)
}

func (t *Table) String() string {
	numColumns := len(t.config.ColumnWidths)
	output := stringutil.CharLineReturn

	// write first line
	output += t.config.CharacterSet.CharTopLeftCornerLine
	for i, width := range t.config.ColumnWidths {
		output += stringutil.Repeat(t.config.CharacterSet.CharHorizontalLine, width+(t.config.CellHorizontalPadding*2))
		if i >= numColumns-1 {
			break
		}
		output += t.config.CharacterSet.CharHorizontalBottomIntersectionLine
	}
	output += t.config.CharacterSet.CharTopRightCornerLine + stringutil.CharLineReturn

	// write column titles
	output += t.config.CharacterSet.CharVerticalLine
	for i, title := range t.config.ColumnTitles {
		cell := stringutil.AlignLeft(title, t.config.ColumnWidths[i])
		output += stringutil.Repeat(stringutil.CharWhiteSpace, t.config.CellHorizontalPadding)
		output += ansi.Wrap(cell, t.config.ColumnTitleStyle...)
		output += stringutil.Repeat(stringutil.CharWhiteSpace, t.config.CellHorizontalPadding)
		if i >= numColumns-1 {
			break
		}
		output += t.config.CharacterSet.CharVerticalLine
	}
	output += t.config.CharacterSet.CharVerticalLine + stringutil.CharLineReturn

	// write middle line
	divider := t.config.CharacterSet.CharVerticalRightIntersectionLine
	for i, width := range t.config.ColumnWidths {
		divider += stringutil.Repeat(t.config.CharacterSet.CharHorizontalLine, width+(t.config.CellHorizontalPadding*2))
		if i >= numColumns-1 {
			break
		}
		divider += t.config.CharacterSet.CharCrossIntersectionLine
	}
	divider += t.config.CharacterSet.CharVerticalLeftIntersectionLine + stringutil.CharLineReturn

	// write middle line to output
	output += divider

	// write rows
	for rowIndex, row := range t.config.Rows {
		output += t.config.CharacterSet.CharVerticalLine

		// write all cells
		for i, cell := range row {
			cell := stringutil.AlignLeft(cell, t.config.ColumnWidths[i])
			output += stringutil.Repeat(stringutil.CharWhiteSpace, t.config.CellHorizontalPadding)
			output += ansi.Wrap(cell, t.config.RowStyle...)
			output += stringutil.Repeat(stringutil.CharWhiteSpace, t.config.CellHorizontalPadding)
			if i >= numColumns-1 {
				break
			}
			output += t.config.CharacterSet.CharVerticalLine
		}
		output += t.config.CharacterSet.CharVerticalLine + stringutil.CharLineReturn

		// write divider in between rows
		if rowIndex >= len(t.config.Rows)-1 {
			break
		}
		output += divider
	}

	// write last line
	output += t.config.CharacterSet.CharBottomLeftCornerLine
	for i, width := range t.config.ColumnWidths {
		output += stringutil.Repeat(t.config.CharacterSet.CharHorizontalLine, width+(t.config.CellHorizontalPadding*2))
		if i >= numColumns-1 {
			break
		}
		output += t.config.CharacterSet.CharHorizontalTopIntersectionLine
	}
	output += t.config.CharacterSet.CharBottomRightCornerLine + stringutil.CharLineReturn

	return output
}
