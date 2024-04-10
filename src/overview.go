package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var toggleRain = player(SOUNDTYPE_RAIN)
var toggleRainstorm = player(SOUNDTYPE_RAINSTORM)
var toggleThunderstorm = player(SOUNDTYPE_THUNDERSTORM)
var toggleWater = player(SOUNDTYPE_WATER)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type keyMap struct {
	Up          key.Binding
	Down        key.Binding
	ToggleSound key.Binding
	Help        key.Binding
	Quit        key.Binding
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	ToggleSound: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "toggle sound"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

var footerRender = lipgloss.NewStyle().
	Height(3).
	Width(80).
	Align(lipgloss.Center).
	Foreground(lipgloss.Color("255")).
	PaddingTop(1).
	PaddingBottom(1).
	PaddingLeft(4).
	PaddingRight(4).
	Bold(true).
	Render

type model struct {
	keys  keyMap
	table table.Model
	help  help.Model
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.ToggleSound},
		{k.Help, k.Quit},
	}
}

func getFooter() string {
	return footerRender("Made with ❤️ by Gorilla Moe  🍌, Fox 🦊,  Corenexus  🌊 and Jan D 🗻. Source code available at github.com/mistweaverco/kelele-rain-tui.")
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			cmds = append(cmds, tea.Quit)
		case "?":
			m.help.ShowAll = !m.help.ShowAll
		case "enter":
			selectedIdx := m.table.Cursor()
			selectedRow := m.table.Rows()[selectedIdx]
			switch st := m.table.SelectedRow()[1]; st {
			case "Rain":
				toggleRain()
			case "Rainstorm":
				toggleRainstorm()
			case "Thunderstorm":
				toggleThunderstorm()
			case "Water":
				toggleWater()
			}
			if m.table.SelectedRow()[0] == "" {
				selectedRow[0] = "✔"
			} else {
				selectedRow[0] = ""
			}
			m.table.SetRows(append(m.table.Rows()[:selectedIdx], append([]table.Row{selectedRow}, m.table.Rows()[selectedIdx+1:]...)...))
		}
	}

	m.table, cmd = m.table.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	helpView := m.help.View(m.keys)
	return baseStyle.Render(m.table.View()) + "\n" + helpView + "\n" + getFooter() + "\n"
}

func overview() {
	columns := []table.Column{
		{Title: "Active", Width: 10},
		{Title: "Sound", Width: 15},
	}

	rows := []table.Row{
		{"", "Rain"},
		{"", "Rainstorm"},
		{"", "Thunderstorm"},
		{"", "Water"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(4),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := model{table: t, help: help.New(), keys: keys}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}