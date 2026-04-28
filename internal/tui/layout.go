package tui

func editorContentWidth(termWidth int) int {
	if termWidth < 8 {
		return 40
	}
	return max(20, termWidth-2)
}

// layoutJSONEditor sizes the JSON textarea for the full content area below chrome.
// Fixed chrome: title(1) + blank(1) + HttpStatusCode(1) + blank(1) + "Response:"(1) = 5 lines above.
// Fixed chrome below: blank(1) + help(3) = 4 lines below. Total fixed = 9.
// When an error is present 3 additional lines are reserved (blank separator + 2 error lines).
func layoutJSONEditor(m *model) {
	w := editorContentWidth(m.width)
	const fixedChrome = 9
	errorLines := 0
	if m.jsonErr != "" {
		errorLines = 3
	}
	h := m.height - fixedChrome - errorLines
	if h < 6 {
		h = 6
	}
	m.ta.SetWidth(w)
	m.ta.SetHeight(h)
	m.tas.SetHeight(1)
	m.tas.SetWidth(w)
}
