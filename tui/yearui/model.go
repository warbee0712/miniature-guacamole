package yearui

import (
	"time"

	"github.com/A-Daneel/miniature-guacamole/tui/constants"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var cmd tea.Cmd

// BackMsg change state back to project view
type BackMsg bool

type (
	apiMsg  string
	tickMsg time.Time
)

var apiResults string

var resultStyle = lipgloss.NewStyle().Border(lipgloss.DoubleBorder()).BorderForeground(lipgloss.Color("147")).Foreground(lipgloss.Color("102")).Render

const (
	defaultApiCall = "user1"
	FIRST          = 0
)

// Model the Years model definition
type Model struct {
	list list.Model
}

// New initialize the projectui model for your program
func New() tea.Model {
	m := Model{
		list: list.NewModel(yearMenu(), list.NewDefaultDelegate(), 100, 25),
	} // Updated the default width and height to large defaults from 0.
	m.list.Title = "second menu, maybe borked"
	m.list.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			constants.Keymap.Enter,
			constants.Keymap.Back,
		}
	}
	return m
}

// Init run any intial IO on program start
func (m Model) Init() tea.Cmd {
	tea.Tick(time.Second*5, func(t time.Time) tea.Msg {
		return mockApiCall(defaultApiCall) // have some default value or move endpoints calculation to a function you can call here
	})
	return nil
}

func mockApiCall(call string) tea.Msg {
	url := "https://www.somewebsite.com/v1"
	apiResults = url + "/" + call + "\n"
	// TODO: actually query for a response
	return apiMsg(apiResults)
	// What do you want to do with the response? Have it display in the
	// terminal?
}

// Update handle IO and commands
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case apiMsg:
		// TODO: whatever you'd like to happen after the first API call
	case tea.WindowSizeMsg:
		m = m.UpdateWindowSize(msg.Width, msg.Height).(Model)
	case tickMsg:
		cmds = append(cmds, func() tea.Msg { return mockApiCall(defaultApiCall) })
	case tea.KeyMsg:
		switch {
		case msg.String() == "ctrl+c":
			return m, tea.Quit
		case key.Matches(msg, constants.Keymap.Enter):
			// TODO: Maybe they should select a user then do a get request for that user?
			mockApiCall(m.list.SelectedItem().FilterValue())
		case key.Matches(msg, constants.Keymap.Back):
			return m, func() tea.Msg {
				return BackMsg(true)
			}
		default:
			m.list, cmd = m.list.Update(msg)
		}
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

func (m Model) UpdateWindowSize(w int, h int) tea.Model {
	top, right, bottom, left := constants.DocStyle.GetMargin()
	m.list.SetSize(w-left-right, h-top-bottom-1)
	return m
}

// View return the text UI to be output to the terminal
func (m Model) View() string {
	result := apiResults
	if apiResults != "" {
		result = resultStyle(apiResults)
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, m.list.View()+"\n", result)
}

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func yearMenu() []list.Item {
	menu := []list.Item{
		item{title: "do action a", desc: ""},
		item{title: "maybe action b?", desc: ""},
		item{title: "naaah, go back to start", desc: ""},
	}
	return menu
}

func secondMenu() []list.Item {
	menu := []list.Item{
		item{title: "mabye", desc: ""},
		item{title: "charm", desc: ""},
		item{title: "bubbles", desc: ""},
	}
	return menu
}
