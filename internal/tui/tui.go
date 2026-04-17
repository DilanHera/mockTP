package tui

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/DilanHera/mockTP/internal/app"
	"github.com/DilanHera/mockTP/internal/services/pgzinv/serviceprovisioning"
	"github.com/DilanHera/mockTP/internal/services/phx"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type screen int

const (
	screenRoot screen = iota
	screenPGZINV
	screenServiceProvisioning
	screenPHX
	screenServiceProvisioningMockJSON
	screenPHXMockJSON
)

// clearSaveNoticeMsg dismisses the post-save confirmation after a short delay.
type clearSaveNoticeMsg struct{}

type model struct {
	app    *app.App
	screen screen
	cursor int
	// Remember cursor when drilling down so Esc restores the parent list position.
	savedRootCursor   int
	savedPGZINVCursor int

	width  int
	height int

	// mock JSON editor: which resource/API mock is being edited; parent screen for Esc/back.
	jsonMockParent   screen
	jsonMockResource string
	ta               textarea.Model
	jsonErr          string
	// Shown on the service provisioning list after Ctrl+S save from the JSON editor.
	saveNotice string
}

// Run starts the Bubble Tea program on the main goroutine.
func Run(app *app.App) error {
	p := tea.NewProgram(
		model{app: app, cursor: 0},
		tea.WithAltScreen(),
		// Windows / ConPTY: read keys from the real console when stdin is piped or wrapped.
		tea.WithInputTTY(),
	)
	_, err := p.Run()
	return err
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) itemCount() int {
	switch m.screen {
	case screenRoot:
		return len(Services)
	case screenPGZINV:
		return len(PgzinvApis)
	case screenServiceProvisioning:
		return len(ServiceProvisioningResources)
	case screenPHX:
		return len(PHXApis)
	default:
		return 0
	}
}

func (m model) labels() []string {
	switch m.screen {
	case screenRoot:
		return Services
	case screenPGZINV:
		return PgzinvApis
	case screenServiceProvisioning:
		return ServiceProvisioningResources
	case screenPHX:
		return PHXApis
	default:
		return nil
	}
}

func newServiceProvisioningMockTextarea(resourceName, value string) textarea.Model {
	t := textarea.New()
	t.ShowLineNumbers = false
	t.Prompt = ""
	// Do not set Placeholder to the JSON template. Bubbles draws Placeholder only when Value()
	// is empty; that text lives outside the real buffer, so it cannot be selected or deleted
	// like normal content. The initial template is loaded via SetValue below.
	t.Placeholder = ""
	t.CharLimit = 256 * 1024
	content := value
	if content == "" {
		content = ServiceProvisioningMockPlaceholder(resourceName)
	}
	t.SetValue(content)
	applyTextareaTheme(&t)
	return t
}

func newPHXMockTextarea(apiName, value string) textarea.Model {
	t := textarea.New()
	t.ShowLineNumbers = false
	t.Prompt = ""
	t.Placeholder = ""
	t.CharLimit = 256 * 1024
	content := value
	if content == "" {
		content = PHXMockPlaceholder(apiName)
	}
	t.SetValue(content)
	applyTextareaTheme(&t)
	return t
}

func isJSONMockScreen(s screen) bool {
	return s == screenServiceProvisioningMockJSON || s == screenPHXMockJSON
}

func isJSONSubmit(msg tea.KeyMsg) bool {
	return msg.String() == "ctrl+s" || msg.String() == "cmd+s"
}

func clearSaveNoticeAfter(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(time.Time) tea.Msg { return clearSaveNoticeMsg{} })
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case clearSaveNoticeMsg:
		m.saveNotice = ""
		return m, nil
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		if isJSONMockScreen(m.screen) {
			layoutJSONEditor(&m)
			var cmd tea.Cmd
			m.ta, cmd = m.ta.Update(msg)
			return m, cmd
		}
		return m, nil

	case tea.KeyMsg:
		if isJSONMockScreen(m.screen) {
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit
			case "esc":
				return m.leaveJSONEditor(), nil
			case "f4":
				m.jsonErr = ""
				m.ta.Reset()
				layoutJSONEditor(&m)
				return m, m.ta.Focus()
			}
			if isJSONSubmit(msg) {
				return m.submitMockJSON()
			}
			var cmd tea.Cmd
			m.ta, cmd = m.ta.Update(msg)
			return m, cmd
		}

		switch msg.String() {
		case "ctrl+t", "cmd+t":
			if m.app.ResponseState == "SUCCESS" {
				m.app.ResponseState = "ERROR"
			} else {
				m.app.ResponseState = "SUCCESS"
			}
			return m, nil
		case "ctrl+c", "q":
			return m, tea.Quit
		case "esc":
			if m.screen == screenRoot {
				return m, tea.Quit
			}
			return m.goBack(), nil
		case "up", "k":
			n := m.itemCount()
			if n == 0 {
				return m, nil
			}
			m.cursor--
			if m.cursor < 0 {
				m.cursor = n - 1
			}
			return m, nil
		case "down", "j":
			n := m.itemCount()
			if n == 0 {
				return m, nil
			}
			m.cursor++
			if m.cursor >= n {
				m.cursor = 0
			}
			return m, nil
		case "enter":
			next, cmd := m.enter()
			return next, cmd
		}
	}
	return m, nil
}

func (m model) enter() (model, tea.Cmd) {
	switch m.screen {
	case screenRoot:
		switch m.cursor {
		case IndexOf(Services, "PGZINV"):
			m.savedRootCursor = m.cursor
			m.screen = screenPGZINV
			m.cursor = 0
		case IndexOf(Services, "PHX"):
			m.savedRootCursor = m.cursor
			m.screen = screenPHX
			m.cursor = 0
		}
	case screenPGZINV:
		if m.cursor == IndexOf(PgzinvApis, "serviceProvisioning") {
			m.savedPGZINVCursor = m.cursor
			m.screen = screenServiceProvisioning
			m.cursor = 0
		}
	case screenServiceProvisioning:
		if m.cursor < 0 || m.cursor >= len(ServiceProvisioningResources) {
			return m, nil
		}
		name := ServiceProvisioningResources[m.cursor]
		m.screen = screenServiceProvisioningMockJSON
		m.jsonMockParent = screenServiceProvisioning
		m.jsonMockResource = name
		m.saveNotice = ""
		m.jsonErr = ""
		m.ta = newServiceProvisioningMockTextarea(name, "")
		layoutJSONEditor(&m)
		cmd := m.ta.Focus()
		return m, cmd
	case screenPHX:
		if m.cursor < 0 || m.cursor >= len(PHXApis) {
			return m, nil
		}
		name := PHXApis[m.cursor]
		m.screen = screenPHXMockJSON
		m.jsonMockParent = screenPHX
		m.jsonMockResource = name
		m.saveNotice = ""
		m.jsonErr = ""
		m.ta = newPHXMockTextarea(name, "")
		layoutJSONEditor(&m)
		cmd := m.ta.Focus()
		return m, cmd
	}
	return m, nil
}

func (m model) leaveJSONEditor() model {
	m.screen = m.jsonMockParent
	m.jsonMockResource = ""
	m.jsonErr = ""
	m.ta.Blur()
	return m
}

func (m model) submitMockJSON() (tea.Model, tea.Cmd) {
	switch m.screen {
	case screenServiceProvisioningMockJSON:
		return m.submitServiceProvisioningMockJSON()
	case screenPHXMockJSON:
		return m.submitPHXMockJSON()
	default:
		return m, nil
	}
}

func (m model) submitServiceProvisioningMockJSON() (tea.Model, tea.Cmd) {
	raw := json.RawMessage(strings.TrimSpace(m.ta.Value()))
	sp := serviceprovisioning.NewServiceProvisioning(m.app)
	var err error
	switch m.jsonMockResource {
	case "lockNumberByCriteriaPrepaid":
		err = sp.SetUserLockNumberByCriteriaPrepaid(raw)
	case "lockNumberByCriteriaPostpaid":
		err = sp.SetUserLockNumberByCriteriaPostpaid(raw)
	case "lockNumberByMobilePrepaid":
		err = sp.SetUserLockNumberByMobilePrepaid(raw)
	case "lockNumberByMobilePostpaid":
		err = sp.SetUserLockNumberByMobilePostpaid(raw)
	case "clearNumberPreparationPrepaid":
		err = sp.SetUserClearNumberPreparationPrepaid(raw)
	case "clearNumberPreparationPostpaid":
		err = sp.SetUserClearNumberPreparationPostpaid(raw)
	case "querySimInfo":
		err = sp.SetUserQuerySimInfo(raw)
	case "requestPrepNoPrepaid":
		err = sp.SetUserRequestPrepNoPrepaid(raw)
	case "requestPrepNoPostpaid":
		err = sp.SetUserRequestPrepNoPostpaid(raw)
	case "confirmPreparationPrepaid":
		err = sp.SetUserConfirmPreparationPrepaid(raw)
	case "confirmPreparationPostpaid":
		err = sp.SetUserConfirmPreparationPostpaid(raw)
	default:
		m.jsonErr = "internal: unknown serviceProvisioning resource"
		return m, nil
	}
	if err != nil {
		m.jsonErr = err.Error()
		m.screen = screenServiceProvisioningMockJSON
		m.ta = newServiceProvisioningMockTextarea(m.jsonMockResource, m.ta.Value())
		layoutJSONEditor(&m)
		cmd := m.ta.Focus()
		return m, cmd
	}
	m.screen = screenServiceProvisioning
	m.jsonMockResource = ""
	m.jsonErr = ""
	m.saveNotice = "Saved successfully."
	m.ta.Blur()
	return m, clearSaveNoticeAfter(2 * time.Second)
}

func (m model) submitPHXMockJSON() (tea.Model, tea.Cmd) {
	raw := json.RawMessage(strings.TrimSpace(m.ta.Value()))
	px := phx.NewPhx(m.app)
	var err error
	switch m.jsonMockResource {
	case "requestESIM":
		err = px.SetUserRequestESIM(raw)
	case "newRegistration":
		err = px.SetUserNewRegistration(raw)
	default:
		m.jsonErr = "internal: unknown PHX API"
		return m, nil
	}
	if err != nil {
		m.jsonErr = err.Error()
		m.screen = screenPHXMockJSON
		m.ta = newPHXMockTextarea(m.jsonMockResource, m.ta.Value())
		layoutJSONEditor(&m)
		cmd := m.ta.Focus()
		return m, cmd
	}
	m.screen = screenPHX
	m.jsonMockResource = ""
	m.jsonErr = ""
	m.saveNotice = "Saved successfully."
	m.ta.Blur()
	return m, clearSaveNoticeAfter(2 * time.Second)
}

func (m model) goBack() model {
	switch m.screen {
	case screenServiceProvisioningMockJSON, screenPHXMockJSON:
		return m.leaveJSONEditor()
	case screenServiceProvisioning:
		m.screen = screenPGZINV
		m.cursor = m.savedPGZINVCursor
	case screenPGZINV:
		m.screen = screenRoot
		m.cursor = m.savedRootCursor
	case screenPHX:
		m.screen = screenRoot
		m.cursor = m.savedRootCursor
	}
	return m
}

func (m model) breadcrumb() string {
	switch m.screen {
	case screenRoot:
		return "mockTP"
	case screenPGZINV:
		return "PGZINV"
	case screenServiceProvisioning:
		return "PGZINV > ServiceProvisioning"
	case screenPHX:
		return "PHX"
	case screenServiceProvisioningMockJSON:
		return "PGZINV > ServiceProvisioning > " + m.jsonMockResource + " [JSON]"
	case screenPHXMockJSON:
		return "PHX > " + m.jsonMockResource + " [JSON]"
	default:
		return "mockTP"
	}
}

func (m model) View() string {
	if isJSONMockScreen(m.screen) {
		var b strings.Builder
		b.WriteString(styleTitle.Render(m.breadcrumb()))
		b.WriteString("\n\n")
		b.WriteString(m.ta.View())
		if m.jsonErr != "" {
			b.WriteString("\n\n")
			b.WriteString(styleErr.Render("Error: " + m.jsonErr))
		}
		b.WriteString("\n\n")
		b.WriteString(styleHelp.Render("Esc cancel · Ctrl+S save mock · Enter = new line · Ctrl+C quit"))
		b.WriteString("\n")
		b.WriteString(styleHelp.Render("F4 clear editor · Ctrl+Home / Ctrl+End jump document"))
		b.WriteString("\n")
		b.WriteString(styleHelp.Render("(clear editor + save) to reset mock to default"))
		return b.String()
	}

	var b strings.Builder
	b.WriteString(styleTitle.Render(m.breadcrumb()))
	if (m.screen == screenServiceProvisioning || m.screen == screenPHX) && m.saveNotice != "" {
		b.WriteString("\n\n")
		b.WriteString(styleOK.Render(m.saveNotice))
	}
	b.WriteString("\n\n")

	for i, label := range m.labels() {
		line := "  " + label
		if i == m.cursor {
			line = styleMenuSel.Render("▸ " + label)
		} else {
			line = styleMenu.Render(line)
		}
		b.WriteString(line)
		b.WriteByte('\n')
	}

	b.WriteString("\n")
	if m.app.ResponseState == "SUCCESS" {
		b.WriteString(styleOK.Render("Response State: " + m.app.ResponseState))
	} else {
		b.WriteString("Response State: " + styleErr.Render(m.app.ResponseState))
	}
	b.WriteString("\n")
	b.WriteString(styleHelp.Render("↑/↓ · Enter open · Ctrl+T toggle response state · Esc back (root: quit) · q quit"))
	return b.String()
}
