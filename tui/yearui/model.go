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

var apiResults string

// Model the Years model definition
type Model struct {
	list list.Model
}

// New initialize the projectui model for your program
func New() tea.Model {
	m := Model{list: list.NewModel(yearMenu(), list.NewDefaultDelegate(), 100, 25)} //Updated the default width and height to large defaults from 0.
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
	return nil
}

func mockApiCall(call string) {
	url := "https://www.somewebsite.com/v1"
	apiResults += url + "/" + call + "\n"
	time.Sleep(5 * time.Second)
}

// Update handle IO and commands
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m = m.UpdateWindowSize(msg.Width, msg.Height).(Model)
	case tea.KeyMsg:
		switch {
		case msg.String() == "ctrl+c":
			return m, tea.Quit
		case key.Matches(msg, constants.Keymap.Enter):
			endpoints := [4]string{"user1", "user2", "user3", "user4"}
			for _, call := range endpoints {
				mockApiCall(call)
			}
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
	//return constants.DocStyle.Render(m.list.View() + "\n")
	return lipgloss.JoinHorizontal(lipgloss.Top, m.list.View()+"\n", apiResults)
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
