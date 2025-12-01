package utils

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type PartLogger struct {
	day       string
	dayStyle  lipgloss.Style
	partStyle lipgloss.Style
	msgStyle  lipgloss.Style
}

func NewPartLogger(day string) *PartLogger {
	return &PartLogger{
		day:       day,
		dayStyle:  lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12")), // Light Blue
		partStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("10")),            // Green
		msgStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("9")),             // Red
	}
}

func (pl *PartLogger) PrintPart(part int, msg any) {
	dayStr := pl.dayStyle.Render(pl.day)
	partStr := pl.partStyle.Render(fmt.Sprintf("Part %v", part))
	msgStr := pl.msgStyle.Render(fmt.Sprintf("%v", msg))

	fmt.Printf("[%s: %s] Output: %s\n", dayStr, partStr, msgStr)
}
