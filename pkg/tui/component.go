package tui

// A component is a sort of UI component (for ex a table or a warning message)
type Component interface {
	String() (string, error)
}
