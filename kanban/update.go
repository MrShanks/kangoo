package kanban

import (
	tea "github.com/charmbracelet/bubbletea"
	"slices"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		m.Loaded = true
		return m, nil
	}

	if m.AddingNew {
		return m.updateForm(msg)
	}
	return m.updateBoard(msg)
}

func (m Model) updateForm(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.FocusIndex == len(m.Inputs)-1 {
				title := m.Inputs[0].Value()
				desc := m.Inputs[1].Value()

				if title != "" {
					if m.Editing {
						m.Lists[m.Focused][m.Cursors[m.Focused]] = Task{Title: title, Description: desc}
					} else {
						m.Lists[Todo] = append(m.Lists[Todo], Task{Title: title, Description: desc})
					}

					m.Lists.Save()

					m.Inputs[0].SetValue("")
					m.Inputs[1].SetValue("")
					m.AddingNew = false
					m.Editing = false
					m.FocusIndex = 0
					return m, nil
				}
				return m, nil
			}
			m.FocusIndex++
			return m.updateFocus()

		case "tab", "shift+tab", "up", "down":
			if msg.String() == "up" || msg.String() == "shift+tab" {
				m.FocusIndex--
			} else {
				m.FocusIndex++
			}

			if m.FocusIndex > len(m.Inputs)-1 {
				m.FocusIndex = 0
			} else if m.FocusIndex < 0 {
				m.FocusIndex = len(m.Inputs) - 1
			}
			return m.updateFocus()

		case "esc":
			m.AddingNew = false
			m.Editing = false
			return m, nil
		}
	}
	cmd := m.updateFocusedInput(msg)
	return m, cmd
}

func (m Model) updateBoard(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.Quitting = true
			return m, tea.Quit

		case "n":
			m.AddingNew = true
			m.Editing = false
			m.FocusIndex = 0
			m.resetFocus()
			return m, nil

		case "e":
			if len(m.Lists[m.Focused]) == 0 {
				return m, nil
			}
			m.AddingNew = true
			m.Editing = true
			m.FocusIndex = 0

			task := m.Lists[m.Focused][m.Cursors[m.Focused]]
			m.Inputs[0].SetValue(task.Title)
			m.Inputs[1].SetValue(task.Description)
			m.resetFocus()
			return m, nil

		case "d":
			if len(m.Lists[m.Focused]) > 0 {
				idx := m.Cursors[m.Focused]
				m.Lists[m.Focused] = slices.Delete(m.Lists[m.Focused], idx, idx+1)
				if m.Cursors[m.Focused] >= len(m.Lists[m.Focused]) && m.Cursors[m.Focused] > 0 {
					m.Cursors[m.Focused]--
				}
				m.Lists.Save()
			}

		case "h":
			if m.Focused > Todo {
				m.Focused--
			}
		case "l":
			if m.Focused < Done {
				m.Focused++
			}
		case "k":
			if m.Cursors[m.Focused] > 0 {
				m.Cursors[m.Focused]--
			}
		case "j":
			if m.Cursors[m.Focused] < len(m.Lists[m.Focused])-1 {
				m.Cursors[m.Focused]++
			}
		case "enter", " ", "ctrl+l":
			m, cmd := m.moveTask("r")
			return m, cmd
		case "backspace", "ctrl+h":
			m, cmd := m.moveTask("l")
			return m, cmd
		}
	}
	return m, nil
}

func (m Model) updateFocus() (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, len(m.Inputs))
	for i := range m.Inputs {
		if i == m.FocusIndex {
			cmds[i] = m.Inputs[i].Focus()
			m.Inputs[i].PromptStyle = FocusedInputStyle
			m.Inputs[i].TextStyle = FocusedInputStyle
		} else {
			m.Inputs[i].Blur()
			m.Inputs[i].PromptStyle = BlurredInputStyle
			m.Inputs[i].TextStyle = BlurredInputStyle
		}
	}
	return m, tea.Batch(cmds...)
}

func (m *Model) updateFocusedInput(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	if m.FocusIndex >= 0 && m.FocusIndex < len(m.Inputs) {
		m.Inputs[m.FocusIndex], cmd = m.Inputs[m.FocusIndex].Update(msg)
	}
	return cmd
}

func (m *Model) resetFocus() {
	m.Inputs[0].Focus()
	m.Inputs[0].PromptStyle = FocusedInputStyle
	m.Inputs[0].TextStyle = FocusedInputStyle

	m.Inputs[1].Blur()
	m.Inputs[1].PromptStyle = BlurredInputStyle
	m.Inputs[1].TextStyle = BlurredInputStyle
}

func (m Model) moveTask(dir string) (tea.Model, tea.Cmd) {
	var nextCol Status
	if dir == "l" {
		if m.Focused == Todo {
			return m, nil
		}
		nextCol = m.Focused - 1
	} else {
		if m.Focused == Done {
			return m, nil
		}
		nextCol = m.Focused + 1
	}

	curr := m.Lists[m.Focused]
	if len(curr) == 0 {
		return m, nil
	}

	idx := m.Cursors[m.Focused]
	task := curr[idx]

	m.Lists[m.Focused] = slices.Delete(curr, idx, idx+1)
	if m.Cursors[m.Focused] >= len(m.Lists[m.Focused]) && m.Cursors[m.Focused] > 0 {
		m.Cursors[m.Focused]--
	}

	m.Lists[nextCol] = append(m.Lists[nextCol], task)
	return m, nil
}
