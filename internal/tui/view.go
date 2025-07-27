package tui

func (m Model) View() string {
	s := ""

	if len(m.episodes) > 0 {
		for _, ep := range m.episodes {
			s = s + ep.Title + "\n"
		}
	}
	s = s + "q to quit\n"

	return s
}
