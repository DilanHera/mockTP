package tui

func editorContentWidth(termWidth int) int {
	if termWidth < 8 {
		return 40
	}
	return max(20, termWidth-2)
}

// layoutJSONEditor sizes the JSON textarea for the full content area below chrome.
func layoutJSONEditor(m *model) {
	w := editorContentWidth(m.width)
	const chromeLines = 9
	h := m.height - chromeLines
	if h < 6 {
		h = 6
	}
	m.ta.SetWidth(w)
	m.ta.SetHeight(h)
	m.tas.SetHeight(1)
	m.tas.SetWidth(w)
}
