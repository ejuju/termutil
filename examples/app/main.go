package main

import (
	"strings"

	"github.com/ejuju/termutil"
	"github.com/ejuju/termutil/pkg/tui"
)

func main() {
	app, err := termutil.NewApp(termutil.AppConfig{
		Version: "0.0.0",
		Commands: map[string]*termutil.Command{
			"ping": {
				Description:                 "Pulls the latest version of a git repos to main branch",
				PromptConfirmationBeforeRun: tui.NewConfirmationPrompt(nil),
				Func: func(ctrlr termutil.Controller, args []string) error {
					if len(args) != 1 {
						return ctrlr.WriteOut("no args provided \n")
					}
					return ctrlr.WriteOut(strings.Join(args, "\n") + "\n")
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}

	err = app.Run(nil)
	if err != nil {
		panic(err)
	}
}
