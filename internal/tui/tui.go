package tui

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/DilanHera/mockTP/internal/app"
	"github.com/DilanHera/mockTP/internal/services/im"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

var ApiStates = make(map[string]string)

type screen int

const (
	screenRoot screen = iota
	screenPGZINV
	screenServiceProvisioning
	screenPHX
	screenDT
	screenIM
	screenESB
	screenServiceProvisioningMockJSON
	screenPHXMockJSON
	screenDTMockJSON
	screenIMMockJSON
	screenESBMockJSON
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
	savedDTCursor     int
	savedIMCursor     int
	savedESBCursor    int

	width  int
	height int

	// mock JSON editor: which resource/API mock is being edited; parent screen for Esc/back.
	jsonMockParent   screen
	jsonMockResource string
	ta               textarea.Model
	tas              textarea.Model
	jsonErr          string
	// Shown on the service provisioning list after Ctrl+S save from the JSON editor.
	saveNotice string
}

// Run starts the Bubble Tea program on the main goroutine.
func Run(app *app.App) error {
	p := tea.NewProgram(
		&model{app: app, cursor: 0},
		tea.WithAltScreen(),
		// Windows / ConPTY: read keys from the real console when stdin is piped or wrapped.
		tea.WithInputTTY(),
	)
	InitApiStates(app)
	_, err := p.Run()
	return err
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) itemCount() int {
	switch m.screen {
	case screenRoot:
		return len(Services)
	case screenPGZINV:
		return len(PgzinvApis)
	case screenServiceProvisioning:
		return len(ServiceProvisioningResources)
	case screenPHX:
		return len(PHXApis)
	case screenDT:
		return len(DTApis)
	case screenIM:
		return len(IMApis)
	case screenESB:
		return len(ESBApis)
	default:
		return 0
	}
}

func (m *model) labels() []string {
	switch m.screen {
	case screenRoot:
		return Services
	case screenPGZINV:
		return PgzinvApis
	case screenServiceProvisioning:
		return ServiceProvisioningResources
	case screenPHX:
		return PHXApis
	case screenDT:
		return DTApis
	case screenIM:
		return IMApis
	case screenESB:
		return ESBApis
	default:
		return nil
	}
}

func (m *model) newServiceProvisioningMockTextarea(resourceName, value string) textarea.Model {
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
		content = m.MarshalJSONForPlaceholder(resourceName)
	}
	t.SetValue(content)
	applyTextareaTheme(&t)
	return t
}

func (m *model) newPHXMockTextarea(apiName, value string) textarea.Model {
	t := textarea.New()
	t.ShowLineNumbers = false
	t.Prompt = ""
	t.Placeholder = ""
	t.CharLimit = 256 * 1024
	content := value
	if content == "" {
		content = m.MarshalJSONForPlaceholder(apiName)
	}
	t.SetValue(content)
	applyTextareaTheme(&t)
	return t
}

func (m *model) newHttpCodeTextarea(apiName, value string) textarea.Model {
	t := textarea.New()
	t.ShowLineNumbers = false
	t.Prompt = ""
	t.Placeholder = ""
	t.CharLimit = 3
	content := value
	if content == "" {
		content = m.HttpStatusCodePlaceholder(apiName)
	}
	t.SetValue(content)
	applyTextareaTheme(&t)
	return t
}

func isJSONMockScreen(s screen) bool {
	return s == screenServiceProvisioningMockJSON || s == screenPHXMockJSON || s == screenDTMockJSON || s == screenIMMockJSON || s == screenESBMockJSON
}

func isJSONSubmit(msg tea.KeyMsg) bool {
	return msg.String() == "ctrl+s" || msg.String() == "cmd+s"
}

func clearSaveNoticeAfter(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(time.Time) tea.Msg { return clearSaveNoticeMsg{} })
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case clearSaveNoticeMsg:
		m.saveNotice = ""
		return m, nil
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		if isJSONMockScreen(m.screen) {
			layoutJSONEditor(m)
			// Resize both editors (PHX has two inputs)
			var cmd1, cmd2 tea.Cmd
			m.tas, cmd1 = m.tas.Update(msg)
			m.ta, cmd2 = m.ta.Update(msg)
			return m, tea.Batch(cmd1, cmd2)
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
				m.tas.Reset()
				layoutJSONEditor(m)
				return m, m.tas.Focus()
			case "tab":
				// Mock JSON editors have two inputs: HttpStatusCode and JSON.
				if m.tas.Focused() {
					m.tas.Blur()
					return m, m.ta.Focus()
				}
				m.ta.Blur()
				return m, m.tas.Focus()
			}
			if isJSONSubmit(msg) {
				return m.submitMockJSON()
			}
			// Update focused editors.
			var cmd1, cmd2 tea.Cmd
			m.tas, cmd1 = m.tas.Update(msg)
			m.ta, cmd2 = m.ta.Update(msg)
			return m, tea.Batch(cmd1, cmd2)
		}

		switch msg.String() {
		case "t", "T":
			m.toggleErrorState()
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

func (m model) toggleErrorState() (model, tea.Cmd) {
	switch m.screen {
	case screenServiceProvisioning:
		if m.cursor < 0 || m.cursor >= len(ServiceProvisioningResources) {
			return m, nil
		}
		ToggleApiState(ServiceProvisioningResources[m.cursor], m.app)
		return m, nil
	case screenPHX:
		if m.cursor < 0 || m.cursor >= len(PHXApis) {
			return m, nil
		}
		ToggleApiState(PHXApis[m.cursor], m.app)
		return m, nil
	case screenDT:
		if m.cursor < 0 || m.cursor >= len(DTApis) {
			return m, nil
		}
		ToggleApiState(DTApis[m.cursor], m.app)
		return m, nil
	case screenIM:
		if m.cursor < 0 || m.cursor >= len(IMApis) {
			return m, nil
		}
		ToggleApiState(IMApis[m.cursor], m.app)
		return m, nil
	case screenESB:
		if m.cursor < 0 || m.cursor >= len(ESBApis) {
			return m, nil
		}
		ToggleApiState(ESBApis[m.cursor], m.app)
		return m, nil
	}
	return m, nil
}

func serviceProvisioningStateIndicator(resource string) string {
	if ApiStates[resource] == "C" {
		return styleCustom.Render("C")
	}
	if ApiStates[resource] == "E" {
		return styleErr.Render("E")
	}
	return styleOK.Render("S")
}

func serviceProvisioningLabelWidth() int {
	width := 0
	for _, resource := range ServiceProvisioningResources {
		if len(resource) > width {
			width = len(resource)
		}
	}
	return width
}

func phxStateIndicator(api string) string {
	if ApiStates[api] == "C" {
		return styleCustom.Render("C")
	}
	if ApiStates[api] == "E" {
		return styleErr.Render("E")
	}
	return styleOK.Render("S")
}

func phxLabelWidth() int {
	width := 0
	for _, api := range PHXApis {
		if len(api) > width {
			width = len(api)
		}
	}
	return width
}

func dtStateIndicator(api string) string {
	if ApiStates[api] == "C" {
		return styleCustom.Render("C")
	}
	if ApiStates[api] == "E" {
		return styleErr.Render("E")
	}
	return styleOK.Render("S")
}

func dtLabelWidth() int {
	width := 0
	for _, api := range DTApis {
		if len(api) > width {
			width = len(api)
		}
	}
	return width
}

func imStateIndicator(api string) string {
	if ApiStates[api] == "C" {
		return styleCustom.Render("C")
	}
	if ApiStates[api] == "E" {
		return styleErr.Render("E")
	}
	return styleOK.Render("S")
}

func imLabelWidth() int {
	width := 0
	for _, api := range IMApis {
		if len(api) > width {
			width = len(api)
		}
	}
	return width
}

func esbStateIndicator(api string) string {
	if ApiStates[api] == "C" {
		return styleCustom.Render("C")
	}
	if ApiStates[api] == "E" {
		return styleErr.Render("E")
	}
	return styleOK.Render("S")
}

func esbLabelWidth() int {
	width := 0
	for _, api := range ESBApis {
		if len(api) > width {
			width = len(api)
		}
	}
	return width
}

func (m *model) newDTMockTextarea(apiName, value string) textarea.Model {
	t := textarea.New()
	t.ShowLineNumbers = false
	t.Prompt = ""
	t.Placeholder = ""
	t.CharLimit = 256 * 1024
	content := value
	if content == "" {
		content = m.MarshalJSONForPlaceholder(apiName)
	}
	t.SetValue(content)
	applyTextareaTheme(&t)
	return t
}

func (m *model) newIMMockTextarea(apiName, value string) textarea.Model {
	t := textarea.New()
	t.ShowLineNumbers = false
	t.Prompt = ""
	t.Placeholder = ""
	t.CharLimit = 256 * 1024
	content := value
	if content == "" {
		content = m.MarshalJSONForPlaceholder(apiName)
	}
	t.SetValue(content)
	applyTextareaTheme(&t)
	return t
}

func (m *model) newESBMockTextarea(apiName, value string) textarea.Model {
	t := textarea.New()
	t.ShowLineNumbers = false
	t.Prompt = ""
	t.Placeholder = ""
	t.CharLimit = 256 * 1024
	content := value
	if content == "" {
		content = m.MarshalJSONForPlaceholder(apiName)
	}
	t.SetValue(content)
	applyTextareaTheme(&t)
	return t
}

func (m *model) enter() (*model, tea.Cmd) {
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
		case IndexOf(Services, "DT"):
			m.savedRootCursor = m.cursor
			m.screen = screenDT
			m.cursor = 0
		case IndexOf(Services, "IM"):
			m.savedRootCursor = m.cursor
			m.screen = screenIM
			m.cursor = 0
		case IndexOf(Services, "ESB"):
			m.savedRootCursor = m.cursor
			m.screen = screenESB
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
		m.ta = m.newServiceProvisioningMockTextarea(name, "")
		m.tas = m.newHttpCodeTextarea(name, "")
		layoutJSONEditor(m)
		cmd := m.tas.Focus()
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
		m.ta = m.newPHXMockTextarea(name, "")
		m.tas = m.newHttpCodeTextarea(name, "")
		layoutJSONEditor(m)
		cmd := m.tas.Focus()
		return m, cmd
	case screenDT:
		if m.cursor < 0 || m.cursor >= len(DTApis) {
			return m, nil
		}
		name := DTApis[m.cursor]
		m.screen = screenDTMockJSON
		m.jsonMockParent = screenDT
		m.jsonMockResource = name
		m.saveNotice = ""
		m.jsonErr = ""
		m.ta = m.newDTMockTextarea(name, "")
		m.tas = m.newHttpCodeTextarea(name, "")
		layoutJSONEditor(m)
		cmd := m.tas.Focus()
		return m, cmd
	case screenIM:
		if m.cursor < 0 || m.cursor >= len(IMApis) {
			return m, nil
		}
		name := IMApis[m.cursor]
		m.screen = screenIMMockJSON
		m.jsonMockParent = screenIM
		m.jsonMockResource = name
		m.saveNotice = ""
		m.jsonErr = ""
		m.ta = m.newIMMockTextarea(name, "")
		m.tas = m.newHttpCodeTextarea(name, "")
		layoutJSONEditor(m)
		cmd := m.tas.Focus()
		return m, cmd
	case screenESB:
		if m.cursor < 0 || m.cursor >= len(ESBApis) {
			return m, nil
		}
		name := ESBApis[m.cursor]
		m.screen = screenESBMockJSON
		m.jsonMockParent = screenESB
		m.jsonMockResource = name
		m.saveNotice = ""
		m.jsonErr = ""
		m.ta = m.newESBMockTextarea(name, "")
		m.tas = m.newHttpCodeTextarea(name, "")
		layoutJSONEditor(m)
		cmd := m.tas.Focus()
		return m, cmd
	}
	return m, nil
}

func (m *model) leaveJSONEditor() *model {
	m.screen = m.jsonMockParent
	m.jsonMockResource = ""
	m.jsonErr = ""
	m.ta.Blur()
	return m
}

func (m *model) submitMockJSON() (tea.Model, tea.Cmd) {
	switch m.screen {
	case screenServiceProvisioningMockJSON:
		return m.submitServiceProvisioningMockJSON()
	case screenPHXMockJSON:
		return m.submitPHXMockJSON()
	case screenDTMockJSON:
		return m.submitDTMockJSON()
	case screenIMMockJSON:
		return m.submitIMMockJSON()
	case screenESBMockJSON:
		return m.submitESBMockJSON()
	default:
		return m, nil
	}
}

func (m *model) submitServiceProvisioningMockJSON() (tea.Model, tea.Cmd) {
	raw := json.RawMessage(strings.TrimSpace(m.ta.Value()))
	httpCode, ok := parseHTTPStatusCode(m.tas.Value())
	if !ok {
		m.jsonErr = "invalid HttpStatusCode (expected 100-999)"
		m.screen = screenServiceProvisioningMockJSON
		m.ta = m.newServiceProvisioningMockTextarea(m.jsonMockResource, m.ta.Value())
		m.tas = m.newHttpCodeTextarea(m.jsonMockResource, m.tas.Value())
		layoutJSONEditor(m)
		cmd := m.tas.Focus()
		return m, cmd
	}
	_ = m.app.ApiInfoStore.UpdateHttpCode(m.jsonMockResource, httpCode)

	err := m.SetCustomResponse(m.jsonMockResource, raw)
	if err != nil {
		m.jsonErr = err.Error()
		m.screen = screenServiceProvisioningMockJSON
		m.ta = m.newServiceProvisioningMockTextarea(m.jsonMockResource, m.ta.Value())
		m.tas = m.newHttpCodeTextarea(m.jsonMockResource, m.tas.Value())
		layoutJSONEditor(m)
		cmd := m.tas.Focus()
		return m, cmd
	}

	m.screen = screenServiceProvisioning
	m.jsonMockResource = ""
	m.jsonErr = ""
	m.saveNotice = "Saved successfully."
	m.ta.Blur()
	m.tas.Blur()
	return m, clearSaveNoticeAfter(2 * time.Second)
}

func (m *model) submitPHXMockJSON() (tea.Model, tea.Cmd) {
	raw := json.RawMessage(strings.TrimSpace(m.ta.Value()))
	// Persist HttpStatusCode (PHX only).
	httpCodeRaw := strings.TrimSpace(m.tas.Value())
	if httpCodeRaw == "" {
		httpCodeRaw = "200"
	}
	var httpCode int
	_, scanErr := fmt.Sscanf(httpCodeRaw, "%d", &httpCode)
	if scanErr != nil || httpCode < 100 || httpCode > 999 {
		m.jsonErr = "invalid HttpStatusCode (expected 100-999)"
		m.screen = screenPHXMockJSON
		m.ta = m.newPHXMockTextarea(m.jsonMockResource, m.ta.Value())
		m.tas = m.newHttpCodeTextarea(m.jsonMockResource, m.tas.Value())
		layoutJSONEditor(m)
		cmd := m.tas.Focus()
		return m, cmd
	}
	_ = m.app.ApiInfoStore.UpdateHttpCode(m.jsonMockResource, httpCode)

	err := m.SetCustomResponse(m.jsonMockResource, raw)
	if err != nil {
		m.jsonErr = err.Error()
		m.screen = screenPHXMockJSON
		m.ta = m.newPHXMockTextarea(m.jsonMockResource, m.ta.Value())
		m.tas = m.newHttpCodeTextarea(m.jsonMockResource, m.tas.Value())
		layoutJSONEditor(m)
		cmd := m.ta.Focus()
		return m, cmd
	}

	m.screen = screenPHX
	m.jsonMockResource = ""
	m.jsonErr = ""
	m.saveNotice = "Saved successfully."
	m.ta.Blur()
	m.tas.Blur()
	return m, clearSaveNoticeAfter(2 * time.Second)
}

func (m *model) submitDTMockJSON() (tea.Model, tea.Cmd) {
	raw := json.RawMessage(strings.TrimSpace(m.ta.Value()))
	httpCode, ok := parseHTTPStatusCode(m.tas.Value())
	if !ok {
		m.jsonErr = "invalid HttpStatusCode (expected 100-999)"
		m.screen = screenDTMockJSON
		m.ta = m.newDTMockTextarea(m.jsonMockResource, m.ta.Value())
		m.tas = m.newHttpCodeTextarea(m.jsonMockResource, m.tas.Value())
		layoutJSONEditor(m)
		cmd := m.tas.Focus()
		return m, cmd
	}
	_ = m.app.ApiInfoStore.UpdateHttpCode(m.jsonMockResource, httpCode)

	err := m.SetCustomResponse(m.jsonMockResource, raw)
	if err != nil {
		m.jsonErr = err.Error()
		m.screen = screenDTMockJSON
		m.ta = m.newDTMockTextarea(m.jsonMockResource, m.ta.Value())
		m.tas = m.newHttpCodeTextarea(m.jsonMockResource, m.tas.Value())
		layoutJSONEditor(m)
		cmd := m.tas.Focus()
		return m, cmd
	}

	m.screen = screenDT
	m.jsonMockResource = ""
	m.jsonErr = ""
	m.saveNotice = "Saved successfully."
	m.ta.Blur()
	m.tas.Blur()
	return m, clearSaveNoticeAfter(2 * time.Second)
}

func (m *model) submitIMMockJSON() (tea.Model, tea.Cmd) {
	raw := json.RawMessage(strings.TrimSpace(m.ta.Value()))
	httpCode, ok := parseHTTPStatusCode(m.tas.Value())
	if !ok {
		m.jsonErr = "invalid HttpStatusCode (expected 100-999)"
		m.screen = screenIMMockJSON
		m.ta = m.newIMMockTextarea(m.jsonMockResource, m.ta.Value())
		m.tas = m.newHttpCodeTextarea(m.jsonMockResource, m.tas.Value())
		layoutJSONEditor(m)
		cmd := m.tas.Focus()
		return m, cmd
	}
	_ = m.app.ApiInfoStore.UpdateHttpCode(m.jsonMockResource, httpCode)

	err := m.SetCustomResponse(m.jsonMockResource, raw)
	if err != nil {
		m.jsonErr = err.Error()
		m.screen = screenIMMockJSON
		m.ta = m.newIMMockTextarea(m.jsonMockResource, m.ta.Value())
		m.tas = m.newHttpCodeTextarea(m.jsonMockResource, m.tas.Value())
		layoutJSONEditor(m)
		cmd := m.tas.Focus()
		return m, cmd
	}

	switch m.jsonMockResource {
	case "sendSimSerialNo":
		if im.UserSendSimSerialNo != nil {
			im.UserSendSimSerialNo.HttpStatusCode = httpCode
		}
	}

	m.screen = screenIM
	m.jsonMockResource = ""
	m.jsonErr = ""
	m.saveNotice = "Saved successfully."
	m.ta.Blur()
	m.tas.Blur()
	return m, clearSaveNoticeAfter(2 * time.Second)
}

func (m *model) submitESBMockJSON() (tea.Model, tea.Cmd) {
	raw := json.RawMessage(strings.TrimSpace(m.ta.Value()))
	httpCode, ok := parseHTTPStatusCode(m.tas.Value())
	if !ok {
		m.jsonErr = "invalid HttpStatusCode (expected 100-999)"
		m.screen = screenESBMockJSON
		m.ta = m.newESBMockTextarea(m.jsonMockResource, m.ta.Value())
		m.tas = m.newHttpCodeTextarea(m.jsonMockResource, m.tas.Value())
		layoutJSONEditor(m)
		cmd := m.tas.Focus()
		return m, cmd
	}
	_ = m.app.ApiInfoStore.UpdateHttpCode(m.jsonMockResource, httpCode)

	err := m.SetCustomResponse(m.jsonMockResource, raw)
	if err != nil {
		m.jsonErr = err.Error()
		m.screen = screenESBMockJSON
		m.ta = m.newESBMockTextarea(m.jsonMockResource, m.ta.Value())
		m.tas = m.newHttpCodeTextarea(m.jsonMockResource, m.tas.Value())
		layoutJSONEditor(m)
		cmd := m.tas.Focus()
		return m, cmd
	}

	m.screen = screenESB
	m.jsonMockResource = ""
	m.jsonErr = ""
	m.saveNotice = "Saved successfully."
	m.ta.Blur()
	m.tas.Blur()
	return m, clearSaveNoticeAfter(2 * time.Second)
}

func parseHTTPStatusCode(raw string) (int, bool) {
	s := strings.TrimSpace(raw)
	if s == "" {
		s = "200"
	}
	var code int
	_, err := fmt.Sscanf(s, "%d", &code)
	if err != nil || code < 100 || code > 999 {
		return 0, false
	}
	return code, true
}

func (m *model) goBack() *model {
	switch m.screen {
	case screenServiceProvisioningMockJSON, screenPHXMockJSON, screenDTMockJSON, screenIMMockJSON, screenESBMockJSON:
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
	case screenDT:
		m.screen = screenRoot
		m.cursor = m.savedRootCursor
	case screenIM:
		m.screen = screenRoot
		m.cursor = m.savedRootCursor
	case screenESB:
		m.screen = screenRoot
		m.cursor = m.savedRootCursor
	}
	return m
}

func (m *model) breadcrumb() string {
	switch m.screen {
	case screenRoot:
		return "mockTP"
	case screenPGZINV:
		return "PGZINV"
	case screenServiceProvisioning:
		return "PGZINV > ServiceProvisioning"
	case screenPHX:
		return "PHX"
	case screenDT:
		return "DT"
	case screenIM:
		return "IM"
	case screenESB:
		return "ESB"
	case screenServiceProvisioningMockJSON:
		return "PGZINV > ServiceProvisioning > " + m.jsonMockResource + " [JSON]"
	case screenPHXMockJSON:
		return "PHX > " + m.jsonMockResource + " [JSON]"
	case screenDTMockJSON:
		return "DT > " + m.jsonMockResource + " [JSON]"
	case screenIMMockJSON:
		return "IM > " + m.jsonMockResource + " [JSON]"
	case screenESBMockJSON:
		return "ESB > " + m.jsonMockResource + " [JSON]"
	default:
		return "mockTP"
	}
}

func (m *model) View() string {
	if isJSONMockScreen(m.screen) {
		var b strings.Builder
		b.WriteString(styleTitle.Render(m.breadcrumb()))
		b.WriteString("\n\n")
		b.WriteString("HttpStatusCode: " + m.tas.View())
		b.WriteString("\n\n")
		b.WriteString("Response:")
		b.WriteString("\n")
		b.WriteString(m.ta.View())
		if m.jsonErr != "" {
			b.WriteString("\n\n")
			b.WriteString(styleErr.Render("Error: " + m.jsonErr))
		}
		b.WriteString("\n\n")
		b.WriteString(styleHelp.Render("Tab switch input · Esc cancel · Ctrl+S save mock · Enter = new line · Ctrl+C quit"))
		b.WriteString("\n")
		b.WriteString(styleHelp.Render("F4 clear editor · Ctrl+Home / Ctrl+End jump document"))
		b.WriteString("\n")
		b.WriteString(styleHelp.Render("(clear editor + save) to reset mock to default"))
		return b.String()
	}

	var b strings.Builder
	b.WriteString(styleTitle.Render(m.breadcrumb()))
	if (m.screen == screenServiceProvisioning || m.screen == screenPHX || m.screen == screenDT || m.screen == screenIM || m.screen == screenESB) && m.saveNotice != "" {
		b.WriteString("\n\n")
		b.WriteString(styleOK.Render(m.saveNotice))
	}
	b.WriteString("\n\n")

	for i, label := range m.labels() {
		displayLabel := label
		if m.screen == screenServiceProvisioning {
			padding := serviceProvisioningLabelWidth() - len(label)
			if padding < 0 {
				padding = 0
			}
			displayLabel = label + strings.Repeat(" ", padding+1) + serviceProvisioningStateIndicator(label)
		}
		if m.screen == screenPHX {
			padding := phxLabelWidth() - len(label)
			if padding < 0 {
				padding = 0
			}
			displayLabel = label + strings.Repeat(" ", padding+1) + phxStateIndicator(label)
		}
		if m.screen == screenDT {
			padding := dtLabelWidth() - len(label)
			if padding < 0 {
				padding = 0
			}
			displayLabel = label + strings.Repeat(" ", padding+1) + dtStateIndicator(label)
		}
		if m.screen == screenIM {
			padding := imLabelWidth() - len(label)
			if padding < 0 {
				padding = 0
			}
			displayLabel = label + strings.Repeat(" ", padding+1) + imStateIndicator(label)
		}
		if m.screen == screenESB {
			padding := esbLabelWidth() - len(label)
			if padding < 0 {
				padding = 0
			}
			displayLabel = label + strings.Repeat(" ", padding+1) + esbStateIndicator(label)
		}

		line := "  " + displayLabel
		if i == m.cursor {
			line = styleMenuSel.Render("▸ " + displayLabel)
		} else {
			line = styleMenu.Render(line)
		}
		b.WriteString(line)
		b.WriteByte('\n')
	}

	b.WriteString("\n")
	switch m.screen {
	case screenServiceProvisioning:
		b.WriteString(styleHelp.Render("↑/↓ · Enter open JSON · t toggle selected API state · Esc back (root: quit) · q quit"))
		b.WriteString("\n")
		b.WriteString(styleHelp.Render("Legend: " + styleOK.Render("S") + " = Success · " + styleErr.Render("E") + " = Error · " + styleCustom.Render("C") + " = Custom"))
		b.WriteString("\n")
		b.WriteString(styleHelp.Render("Note: t cycles through " + styleOK.Render("S") + " → " + styleErr.Render("E") + " → " + styleCustom.Render("C") + " → " + styleOK.Render("S") + "."))
	case screenPHX:
		b.WriteString(styleHelp.Render("↑/↓ · Enter open JSON · t toggle selected API state · Esc back (root: quit) · q quit"))
		b.WriteString("\n")
		b.WriteString(styleHelp.Render("Legend: " + styleOK.Render("S") + " = Success · " + styleErr.Render("E") + " = Error · " + styleCustom.Render("C") + " = Custom"))
		b.WriteString("\n")
		b.WriteString(styleHelp.Render("Note: t cycles through " + styleOK.Render("S") + " → " + styleErr.Render("E") + " → " + styleCustom.Render("C") + " → " + styleOK.Render("S") + "."))
	case screenDT:
		b.WriteString(styleHelp.Render("↑/↓ · Enter open JSON · t toggle selected API state · Esc back (root: quit) · q quit"))
		b.WriteString("\n")
		b.WriteString(styleHelp.Render("Legend: " + styleOK.Render("S") + " = Success · " + styleErr.Render("E") + " = Error · " + styleCustom.Render("C") + " = Custom"))
		b.WriteString("\n")
		b.WriteString(styleHelp.Render("Note: t cycles through " + styleOK.Render("S") + " → " + styleErr.Render("E") + " → " + styleCustom.Render("C") + " → " + styleOK.Render("S") + "."))
	case screenIM:
		b.WriteString(styleHelp.Render("↑/↓ · Enter open JSON · t toggle selected API state · Esc back (root: quit) · q quit"))
		b.WriteString("\n")
		b.WriteString(styleHelp.Render("Legend: " + styleOK.Render("S") + " = Success · " + styleErr.Render("E") + " = Error · " + styleCustom.Render("C") + " = Custom"))
		b.WriteString("\n")
		b.WriteString(styleHelp.Render("Note: t cycles through " + styleOK.Render("S") + " → " + styleErr.Render("E") + " → " + styleCustom.Render("C") + " → " + styleOK.Render("S") + "."))
	case screenESB:
		b.WriteString(styleHelp.Render("↑/↓ · Enter open JSON · t toggle selected API state · Esc back (root: quit) · q quit"))
		b.WriteString("\n")
		b.WriteString(styleHelp.Render("Legend: " + styleOK.Render("S") + " = Success · " + styleErr.Render("E") + " = Error · " + styleCustom.Render("C") + " = Custom"))
		b.WriteString("\n")
		b.WriteString(styleHelp.Render("Note: t cycles through " + styleOK.Render("S") + " → " + styleErr.Render("E") + " → " + styleCustom.Render("C") + " → " + styleOK.Render("S") + "."))
	default:
		b.WriteString(styleHelp.Render("↑/↓ · Enter open · Esc back (root: quit) · q quit"))
	}
	return b.String()
}
