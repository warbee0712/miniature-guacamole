package yearui

import (
	"github.com/A-Daneel/factuurTui/tui/constants"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

var cmd tea.Cmd

// BackMsg change state back to project view
type BackMsg bool

// Model the Years model definition
type Model struct {
	list list.Model
}

// New initialize the projectui model for your program
func New() tea.Model {
	m := Model{list: list.NewModel(yearMenu(), list.NewDefaultDelegate(), 0, 0)}
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

// Update handle IO and commands
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		top, right, bottom, left := constants.DocStyle.GetMargin()
		m.list.SetSize(msg.Width-left-right, msg.Height-top-bottom-1)
	case tea.KeyMsg:
		switch {
		case msg.String() == "ctrl+c":
			return m, tea.Quit
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

// View return the text UI to be output to the terminal
func (m Model) View() string {
	return constants.DocStyle.Render(m.list.View() + "\n")
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
