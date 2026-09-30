package main

import "charm.land/lipgloss/v2"

var (
	keyword = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#b8bb26")).
		Render

	paragraph = lipgloss.NewStyle().
			Width(78).
			Padding(0, 0, 0, 2).
			Render
)
