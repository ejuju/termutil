package termutil

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/ejuju/termutil/pkg/tui"
)

// App represents an app that runs in the CLI.
// You can instanciate an app with the NewApp function
type App struct {
	config AppConfig
}

type AppConfig struct {
	Controller Controller
	Commands   map[string]*Command
	Version    string
}

func NewApp(config AppConfig) (*App, error) {
	// set default controller if needed
	if config.Controller == nil {
		config.Controller = &DefaultController{
			StdIn:          os.Stdin,
			StdOut:         os.Stdout,
			StdErr:         os.Stderr,
			WriteOutPrefix: "",
			WriteErrPrefix: "error: ",
		}
	}

	// fail if no commands are provided
	if len(config.Commands) == 0 {
		return nil, errors.New("no commands were provided")
	}

	// If the version is defined, add the corresponding command
	if config.Version != "" {
		config.Commands["version"] = &Command{
			Description: "Prints the version of the app",
			Func: func(ctrlr Controller, args []string) error {
				return ctrlr.WriteOut(config.Version + "\n")
			},
		}
	}

	return &App{config: config}, nil
}

// args will default to os.Args[:1] when set to nil
func (app *App) Run(args []string) error {
	// set default args if needed
	if args == nil {
		args = os.Args[1:]
	}

	command, ok := app.config.Commands[args[0]]
	if !ok {
		return fmt.Errorf("command \"%s\"not found", args[0])
	}

	if command.PromptConfirmationBeforeRun != nil {
		ok, err := command.PromptConfirmationBeforeRun.Exec()
		if err != nil {
			return fmt.Errorf("unable to prompt confirmation before run: %w", err)
		}
		if !ok {
			return nil
		}
	}

	err := command.Func(app.config.Controller, args[1:])
	return fmt.Errorf("failed to run command \"%s\": %w", args[0], err)
}

type Command struct {
	Description                 string
	PromptConfirmationBeforeRun *tui.ConfirmationPrompt
	Func                        func(ctrlr Controller, args []string) error
	Flags                       *flag.FlagSet
}

type Controller interface {
	ReadIn() (string, error)
	WriteOut(string) error
	WriteErr(string) error
}

type DefaultController struct {
	WriteOutPrefix string
	WriteErrPrefix string
	StdIn          io.Reader
	StdOut         io.Writer
	StdErr         io.Writer
}

func (ctrlr *DefaultController) ReadIn() (string, error) {
	reader := bufio.NewReader(ctrlr.StdIn)
	response, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return response, nil
}

func (ctrlr *DefaultController) WriteOut(str string) error {
	_, err := ctrlr.StdOut.Write([]byte(ctrlr.WriteOutPrefix + str))
	return err
}

func (ctrlr *DefaultController) WriteErr(str string) error {
	_, err := ctrlr.StdErr.Write([]byte(ctrlr.WriteErrPrefix + str))
	return err
}
