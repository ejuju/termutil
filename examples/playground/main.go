package main

import (
	"fmt"

	"github.com/ejuju/termutil/pkg/tui"
)

func main() {
	ok, err := tui.NewConfirmationPrompt(nil).Exec()
	if err != nil {
		panic(err)
	}

	fmt.Println("ok?", ok)
}
