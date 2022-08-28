package tui

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type ConfirmationPrompt struct {
	config *ConfirmationPromptConfig
}

type ConfirmationPromptConfig struct {
	StdOut   io.Writer
	StdIn    io.Reader
	Question string
}

func NewConfirmationPrompt(config *ConfirmationPromptConfig) *ConfirmationPrompt {
	if config == nil {
		config = &ConfirmationPromptConfig{}
	}
	if config.StdOut == nil {
		config.StdOut = os.Stdout
	}
	if config.StdIn == nil {
		config.StdIn = os.Stdin
	}
	if config.Question == "" {
		config.Question = "Are you sure?"
	}

	return &ConfirmationPrompt{config: config}
}

// Prompt the user for confirmation (blocks until the user confirms)
func (prompt *ConfirmationPrompt) Exec() (bool, error) {
	_, err := prompt.config.StdOut.Write([]byte(prompt.config.Question + " (y/n) "))
	if err != nil {
		return false, err
	}

	reader := bufio.NewReader(prompt.config.StdIn)
	response, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}
	response = strings.ToLower(strings.TrimSpace(response))
	if response == "y" || response == "yes" {
		return true, nil
	}

	return false, nil
}
