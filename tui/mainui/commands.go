package mainui

import tea "github.com/charmbracelet/bubbletea"

func selectCmd(ActiveMenuID uint) tea.Cmd {
	return func() tea.Msg {
		return SelectMsg{ActiveMenu: ActiveMenuID + 1}
	}
}
