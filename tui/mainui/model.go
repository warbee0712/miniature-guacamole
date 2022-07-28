package mainui

import (
	"github.com/A-Daneel/factuurTui/tui/constants"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// SelectMenuMsg the message to change the view to the selected entry
type SelectMsg struct {
	ActiveMenu uint
}

// Model the Menu model definition
type Model struct {
	list list.Model
}

// New initialize the projectui model for your program
func New() tea.Model {
	m := Model{list: list.NewModel(menu(), list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Main Menu thing"
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
		case key.Matches(msg, constants.Keymap.Enter):
			cmd = selectCmd(uint(m.list.Cursor()))
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

func menu() []list.Item {
	menu := []list.Item{
		item{title: "Menu option 1", desc: ""},
		item{title: "Menu option 2", desc: ""},
		item{title: "Menu option 4", desc: ""},
	}
	return menu
}
