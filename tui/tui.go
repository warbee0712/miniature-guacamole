package tui

import (
	"fmt"
	"log"
	"os"

	"github.com/A-Daneel/miniature-guacamole/tui/mainui"
	"github.com/A-Daneel/miniature-guacamole/tui/yearui"
	tea "github.com/charmbracelet/bubbletea"
)

var p *tea.Program

type sessionState int

const (
	menuView sessionState = iota
	yearView
	monthView
	loadingView
)

type MainModel struct {
	state        sessionState
	main         tea.Model
	year         tea.Model
	ActiveMenuID uint
	windowSize   tea.WindowSizeMsg
}

// StartTea the entry point for the UI. Initializes the model.
func StartTea() {
	if f, err := tea.LogToFile("debug.log", "help"); err != nil {
		fmt.Println("Couldn't open a file for logging:", err)
		os.Exit(1)
	} else {
		defer func() {
			err = f.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	m := New()
	p = tea.NewProgram(m)
	p.EnterAltScreen()
	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

// New initialize the main model for your program
func New() MainModel {
	return MainModel{
		state: sessionState(0),
		main:  mainui.New(),
		year:  yearui.New(),
	}
}

// Init run any intial IO on program start
func (m MainModel) Init() tea.Cmd {
	return nil
}

// Update handle IO and commands
func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowSize = msg
	case yearui.BackMsg:
		m.state = menuView
	case mainui.SelectMsg:
		m.state = yearView
	}

	switch m.state {
	case menuView:
		newProject, newCmd := m.main.Update(msg)
		newModel, ok := newProject.(mainui.Model)
		if !ok {
			panic("could not perform assertion on mainui model")
		}
		m.main = newModel
		cmd = newCmd
	case yearView:
		newProject, newCmd := m.year.Update(msg)
		newModel, ok := newProject.(yearui.Model)
		if !ok {
			panic("could not perform assertion on yearui model")
		}
		m.year = newModel
		cmd = newCmd
	}
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

// View return the text UI to be output to the terminal
func (m MainModel) View() string {
	switch m.state {
	case menuView:
		return m.main.View()
	case yearView:
		return m.year.View()
	default:
		return m.main.View()
	}
}
