package kanban

import "github.com/charmbracelet/lipgloss"

const (
	Margin  = 1
	Padding = 2
)

var (
	ColumnStyle = lipgloss.NewStyle().
			Padding(1, Padding).
			Margin(0, Margin).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62"))

	FocusedStyle = lipgloss.NewStyle().
			Padding(1, Padding).
			Margin(0, Margin).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("169"))

	HelpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))

	FocusedInputStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	BlurredInputStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))

	SelectedItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	DescriptionStyle = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder(), false, false, false, true).
				BorderForeground(lipgloss.Color("240")).
				Foreground(lipgloss.Color("245")).
				PaddingLeft(1).
				Italic(true)
)
