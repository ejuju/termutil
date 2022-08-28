package tui

import (
	"fmt"
	"testing"
)

func TestTable(t *testing.T) {
	table, err := NewTable(TableConfig{
		ColumnTitles: []string{"Column A", "Column B"},
		ColumnWidths: []int{10, 10},
		Rows: [][]string{
			{"1A", "1B"},
			{"2A", "2B"},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	expectedPrintOutput := `
┌──────────┬──────────┐
│Column A  │Column B  │
├──────────┼──────────┤
│1A        │1B        │
├──────────┼──────────┤
│2A        │2B        │
└──────────┴──────────┘
`
	t.Run("should render table correctly", func(t *testing.T) {
		result := table.String()
		fmt.Println("result", result)
		if result != expectedPrintOutput {
			t.Fatal(result, expectedPrintOutput)
		}
	})
}
