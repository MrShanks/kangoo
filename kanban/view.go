package kanban

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	if m.Quitting {
		return "Goodbye, your majesty.\n"
	}
	if !m.Loaded {
		return "Loading..."
	}

	colWidth := (m.Width / 3) - (Margin * 2) - 2

	col1 := m.renderColumn(Todo, "TODO", colWidth)
	col2 := m.renderColumn(InProgress, "IN PROGRESS", colWidth)
	col3 := m.renderColumn(Done, "DONE", colWidth)

	board := lipgloss.JoinHorizontal(lipgloss.Top, col1, col2, col3)

	help := HelpStyle.Render("\nPress 'n' to add • 'e' to edit • 'd' to delete • 'q' to quit")

	if m.AddingNew {
		return lipgloss.JoinVertical(lipgloss.Left, board, m.Inputs[0].View(), m.Inputs[1].View())
	}

	return lipgloss.JoinVertical(lipgloss.Left, board, help)
}

func (m Model) renderColumn(status Status, title string, width int) string {
	var tasks string
	currentList := m.Lists[status]

	for i, t := range currentList {
		cursor := " "
		isSelected := m.Focused == status && m.Cursors[status] == i

		if isSelected {
			cursor = SelectedItemStyle.Render(">")
			tasks += fmt.Sprintf("%s %s\n", cursor, SelectedItemStyle.Render(t.Title))

			if t.Description != "" {
				descWidth := width - 6
				desc := DescriptionStyle.Width(descWidth).MarginLeft(2).Render(t.Description)
				tasks += fmt.Sprintf("%s\n", desc)
			}
		} else {
			tasks += fmt.Sprintf("%s %s\n", cursor, t.Title)
		}
	}

	style := ColumnStyle.Width(width).Height(m.Height - 5)
	if m.Focused == status {
		style = FocusedStyle.Width(width).Height(m.Height - 5)
	}

	return style.Render(title + "\n" + tasks)
}
