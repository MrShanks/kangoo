package kanban

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Lists   Board
	Focused Status
	Cursors [3]int

	Quitting  bool
	AddingNew bool
	Editing   bool

	Inputs     []textinput.Model
	FocusIndex int

	Loaded bool
	Width  int
	Height int
}

func New() Model {
	lists, found := Load()

	m := Model{
		Focused: Todo,
		Inputs:  make([]textinput.Model, 2),
		Lists:   lists,
	}

	m.Inputs[0] = textinput.New()
	m.Inputs[0].Placeholder = "Task Title"
	m.Inputs[0].CharLimit = 20
	m.Inputs[0].Width = 20
	m.Inputs[0].PromptStyle = FocusedInputStyle
	m.Inputs[0].TextStyle = FocusedInputStyle
	m.Inputs[0].Focus()

	m.Inputs[1] = textinput.New()
	m.Inputs[1].Placeholder = "Description"
	m.Inputs[1].CharLimit = 100
	m.Inputs[1].Width = 30
	m.Inputs[1].PromptStyle = BlurredInputStyle
	m.Inputs[1].TextStyle = BlurredInputStyle

	if !found {
		m.Lists.Save()
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}
