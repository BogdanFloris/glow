package ui

import (
	"image/color"

	"charm.land/lipgloss/v2"
)

// Styles contains all styles used in the UI. They are resolved from light and
// dark color variants according to the terminal's background color, which
// Bubble Tea reports via tea.BackgroundColorMsg. Until then, we default to
// a dark background.
type Styles struct {
	lightDark lipgloss.LightDarkFunc

	fuchsia color.Color

	dimNormalFg      func(...string) string
	brightGrayFg     func(...string) string
	dimBrightGrayFg  func(...string) string
	grayFg           func(...string) string
	midGrayFg        func(...string) string
	darkGrayFg       lipgloss.Style
	greenFg          func(...string) string
	semiDimGreenFg   func(...string) string
	dimGreenFg       func(...string) string
	fuchsiaFg        func(...string) string
	dimFuchsiaFg     func(...string) string
	dullFuchsiaFg    func(...string) string
	dimDullFuchsiaFg func(...string) string
	redFg            func(...string) string

	tabStyle         lipgloss.Style
	selectedTabStyle lipgloss.Style
	errorTitleStyle  lipgloss.Style
	subtleStyle      lipgloss.Style
	paginationStyle  lipgloss.Style

	statusBarScrollPosStyle        func(...string) string
	statusBarNoteStyle             func(...string) string
	statusBarHelpStyle             func(...string) string
	statusBarMessageStyle          func(...string) string
	statusBarMessageScrollPosStyle func(...string) string
	statusBarMessageHelpStyle      func(...string) string
	helpViewStyle                  func(...string) string
	lineNumberStyle                func(...string) string

	dividerDot            lipgloss.Style
	dividerBar            lipgloss.Style
	logoStyle             lipgloss.Style
	stashSpinnerStyle     lipgloss.Style
	stashInputPromptStyle lipgloss.Style
}

// newStyles builds all styles for the given terminal background.
func newStyles(isDark bool) Styles {
	s := Styles{lightDark: lipgloss.LightDark(isDark)}

	// Colors
	normalDim := s.adaptive("#7c6f64", "#a89984")
	gray := s.adaptive("#928374", "#928374")
	midGray := s.adaptive("#a89984", "#7c6f64")
	darkGray := s.adaptive("#d5c4a1", "#504945")
	brightGray := s.adaptive("#665c54", "#d5c4a1")
	dimBrightGray := s.adaptive("#bdae93", "#665c54")
	cream := s.adaptive("#fbf1c7", "#fbf1c7")
	yellowGreen := s.adaptive("#98971a", "#b8bb26")
	s.fuchsia = s.adaptive("#d65d0e", "#fe8019")
	dimFuchsia := s.adaptive("#d79921", "#fabd2f")
	dullFuchsia := s.adaptive("#d65d0e", "#fe8019")
	dimDullFuchsia := s.adaptive("#bdae93", "#a89984")
	green := lipgloss.Color("#b8bb26")
	red := s.adaptive("#cc241d", "#fb4934")
	semiDimGreen := s.adaptive("#98971a", "#98971a")
	dimGreen := s.adaptive("#689d6a", "#689d6a")

	// Pager colors
	mintGreen := s.adaptive("#fbf1c7", "#fbf1c7")
	darkGreen := s.adaptive("#689d6a", "#689d6a")
	lineNumberFg := s.adaptive("#928374", "#7c6f64")
	statusBarNoteFg := s.adaptive("#504945", "#d5c4a1")
	statusBarBg := s.adaptive("#ebdbb2", "#3c3836")

	// Render-func styles
	s.dimNormalFg = lipgloss.NewStyle().Foreground(normalDim).Render
	s.brightGrayFg = lipgloss.NewStyle().Foreground(brightGray).Render
	s.dimBrightGrayFg = lipgloss.NewStyle().Foreground(dimBrightGray).Render
	s.grayFg = lipgloss.NewStyle().Foreground(gray).Render
	s.midGrayFg = lipgloss.NewStyle().Foreground(midGray).Render
	s.darkGrayFg = lipgloss.NewStyle().Foreground(darkGray)
	s.greenFg = lipgloss.NewStyle().Foreground(green).Render
	s.semiDimGreenFg = lipgloss.NewStyle().Foreground(semiDimGreen).Render
	s.dimGreenFg = lipgloss.NewStyle().Foreground(dimGreen).Render
	s.fuchsiaFg = lipgloss.NewStyle().Foreground(s.fuchsia).Render
	s.dimFuchsiaFg = lipgloss.NewStyle().Foreground(dimFuchsia).Render
	s.dullFuchsiaFg = lipgloss.NewStyle().Foreground(dullFuchsia).Render
	s.dimDullFuchsiaFg = lipgloss.NewStyle().Foreground(dimDullFuchsia).Render
	s.redFg = lipgloss.NewStyle().Foreground(red).Render

	// Named styles
	s.tabStyle = lipgloss.NewStyle().Foreground(s.adaptive("#928374", "#928374"))
	s.selectedTabStyle = lipgloss.NewStyle().Foreground(s.adaptive("#3c3836", "#ebdbb2"))
	s.errorTitleStyle = lipgloss.NewStyle().Foreground(cream).Background(red).Padding(0, 1)
	s.subtleStyle = lipgloss.NewStyle().Foreground(s.adaptive("#928374", "#665c54"))
	s.paginationStyle = s.subtleStyle

	// Pager styles
	s.statusBarScrollPosStyle = lipgloss.NewStyle().
		Foreground(s.adaptive("#7c6f64", "#a89984")).
		Background(statusBarBg).
		Render

	s.statusBarNoteStyle = lipgloss.NewStyle().
		Foreground(statusBarNoteFg).
		Background(statusBarBg).
		Render

	s.statusBarHelpStyle = lipgloss.NewStyle().
		Foreground(statusBarNoteFg).
		Background(s.adaptive("#d5c4a1", "#504945")).
		Render

	s.statusBarMessageStyle = lipgloss.NewStyle().
		Foreground(mintGreen).
		Background(darkGreen).
		Render

	s.statusBarMessageScrollPosStyle = lipgloss.NewStyle().
		Foreground(mintGreen).
		Background(darkGreen).
		Render

	s.statusBarMessageHelpStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#282828")).
		Background(green).
		Render

	s.helpViewStyle = lipgloss.NewStyle().
		Foreground(statusBarNoteFg).
		Background(s.adaptive("#f2e5bc", "#1d2021")).
		Render

	s.lineNumberStyle = lipgloss.NewStyle().
		Foreground(lineNumberFg).
		Render

	// Stash styles
	s.dividerDot = s.darkGrayFg.SetString(" • ")
	s.dividerBar = s.darkGrayFg.SetString(" │ ")

	s.logoStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#282828")).
		Background(s.fuchsia).
		Bold(true)

	s.stashSpinnerStyle = lipgloss.NewStyle().
		Foreground(gray)
	s.stashInputPromptStyle = lipgloss.NewStyle().
		Foreground(yellowGreen).
		MarginRight(1)

	return s
}

// adaptive returns a color appropriate for the current terminal background.
func (s Styles) adaptive(light, dark string) color.Color {
	return s.lightDark(lipgloss.Color(light), lipgloss.Color(dark))
}
