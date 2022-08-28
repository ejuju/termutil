package main

import (
	"fmt"

	"github.com/ejuju/fake"
	"github.com/ejuju/termutil/pkg/ansi"
	"github.com/ejuju/termutil/pkg/stringutil"
	"github.com/ejuju/termutil/pkg/tui"
)

func main() {
	// generate sample users
	usersDataRows := make([][]string, 20)
	for i := range usersDataRows {
		user := fake.User(nil)
		usersDataRows[i] = []string{user.ID, user.EmailAddress.Address, user.FirstName, user.LastName}
	}

	userTable, err := tui.NewTable(tui.TableConfig{
		ColumnTitles:          []string{"ID", "Email address", "First name", "Last name"},
		Rows:                  usersDataRows,
		ColumnWidths:          []int{36, 24, 20, 20},
		CellHorizontalPadding: 1,
		ColumnTitleStyle:      []ansi.EscapeCode{ansi.EscapeCodeBold},
		RowStyle:              []ansi.EscapeCode{},
		CharacterSet:          stringutil.BoxDrawingSetSingleLine,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(userTable.String())
}
