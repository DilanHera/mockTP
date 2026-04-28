package tui

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/DilanHera/mockTP/internal/app"
	"github.com/DilanHera/mockTP/internal/services/im"
	"github.com/atotto/clipboard"
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
	screenEOS
	screenIDS
	screenSMIS
	screenMyChannel
	screenServiceProvisioningMockJSON
	screenPHXMockJSON
	screenDTMockJSON
	screenIMMockJSON
	screenESBMockJSON
	screenEOSMockJSON
	screenIDSMockJSON
	screenSMISMockJSON
	screenMyChannelMockJSON
)

// clearSaveNoticeMsg dismisses the post-save confirmation after a short delay.
type clearSaveNoticeMsg struct{}

type model struct {
	app    *app.App
	screen screen
	cursor int
	// Remember cursor when drilling down so Esc restores the parent list position.
	savedRootCursor      int
	savedPGZINVCursor    int
	savedDTCursor        int
	savedIMCursor        int
	savedESBCursor       int
	savedEOSCursor       int
	savedIDSCursor       int
	savedSMISCursor      int
	savedMyChannelCursor int

	width  int
	height int

	// mock JSON editor: which resource/API mock is being edited; parent screen for Esc/back.
	jsonMockParent   screen
	jsonMockResource string
	ta               textarea.Model
	tas              textarea.Model
	jsonErr          string
	jsonPlaceholder  string
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
	case screenEOS:
		return len(EOSApis)
	case screenIDS:
		return len(IDSApis)
	case screenSMIS:
		return len(SMISApis)
	case screenMyChannel:
		return len(MyChannelApis)
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
	case screenEOS:
		return EOSApis
	case screenIDS:
		return IDSApis
	case screenSMIS:
		return SMISApis
	case screenMyChannel:
		return MyChannelApis
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
	return s == screenServiceProvisioningMockJSON || s == screenPHXMockJSON || s == screenDTMockJSON ||
		s == screenIMMockJSON || s == screenESBMockJSON ||
		s == screenEOSMockJSON || s == screenIDSMockJSON || s == screenSMISMockJSON || s == screenMyChannelMockJSON
}

func isJSONSubmit(msg tea.KeyMsg) bool {
	return msg.String() == "ctrl+s" || msg.String() == "cmd+s"
}

func clearSaveNoticeAfter(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(time.Time) tea.Msg { return clearSaveNoticeMsg{} })
}

func renderHelpKeys(s string) string {
	// Render common keyboard tokens with a dedicated style so they stand out.
	// Keep this simple (string replacement) since the help copy is static.
	r := strings.NewReplacer(
		"Tab", styleKey.Render("Tab"),
		"Esc", styleKey.Render("Esc"),
		"Enter", styleKey.Render("Enter"),
		"Ctrl+S", styleKey.Render("Ctrl+S"),
		"Ctrl+C", styleKey.Render("Ctrl+C"),
		"Ctrl+X", styleKey.Render("Ctrl+X"),
		"Ctrl+Home", styleKey.Render("Ctrl+Home"),
		"Ctrl+End", styleKey.Render("Ctrl+End"),
		"F4", styleKey.Render("F4"),
		"↑/↓", styleKey.Render("↑/↓"),
	)
	out := r.Replace(s)
	// Single-letter tokens must only match whole words; otherwise we'd color every "t" in text.
	out = regexp.MustCompile(`\bt\b`).ReplaceAllString(out, styleKey.Render("t"))
	out = regexp.MustCompile(`\bq\b`).ReplaceAllString(out, styleKey.Render("q"))
	return out
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
				clipboard.WriteAll(m.jsonPlaceholder)
			case "ctrl+x":
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
		case "ctrl+x", "q":
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
	case screenEOS:
		if m.cursor < 0 || m.cursor >= len(EOSApis) {
			return m, nil
		}
		ToggleApiState(EOSApis[m.cursor], m.app)
		return m, nil
	case screenIDS:
		if m.cursor < 0 || m.cursor >= len(IDSApis) {
			return m, nil
		}
		ToggleApiState(IDSApis[m.cursor], m.app)
		return m, nil
	case screenSMIS:
		if m.cursor < 0 || m.cursor >= len(SMISApis) {
			return m, nil
		}
		ToggleApiState(SMISApis[m.cursor], m.app)
		return m, nil
	case screenMyChannel:
		if m.cursor < 0 || m.cursor >= len(MyChannelApis) {
			return m, nil
		}
		ToggleApiState(MyChannelApis[m.cursor], m.app)
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

func eosStateIndicator(api string) string {
	if ApiStates[api] == "C" {
		return styleCustom.Render("C")
	}
	if ApiStates[api] == "E" {
		return styleErr.Render("E")
	}
	return styleOK.Render("S")
}

func eosLabelWidth() int {
	width := 0
	for _, api := range EOSApis {
		if len(api) > width {
			width = len(api)
		}
	}
	return width
}

func idsStateIndicator(api string) string {
	if ApiStates[api] == "C" {
		return styleCustom.Render("C")
	}
	if ApiStates[api] == "E" {
		return styleErr.Render("E")
	}
	return styleOK.Render("S")
}

func idsLabelWidth() int {
	width := 0
	for _, api := range IDSApis {
		if len(api) > width {
			width = len(api)
		}
	}
	return width
}

func smisStateIndicator(api string) string {
	if ApiStates[api] == "C" {
		return styleCustom.Render("C")
	}
	if ApiStates[api] == "E" {
		return styleErr.Render("E")
	}
	return styleOK.Render("S")
}

func smisLabelWidth() int {
	width := 0
	for _, api := range SMISApis {
		if len(api) > width {
			width = len(api)
		}
	}
	return width
}

func myChannelStateIndicator(api string) string {
	if ApiStates[api] == "C" {
		return styleCustom.Render("C")
	}
	if ApiStates[api] == "E" {
		return styleErr.Render("E")
	}
	return styleOK.Render("S")
}

func myChannelLabelWidth() int {
	width := 0
	for _, api := range MyChannelApis {
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

func (m *model) newEOSMockTextarea(apiName, value string) textarea.Model {
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

func (m *model) newIDSMockTextarea(apiName, value string) textarea.Model {
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

func (m *model) newSMISMockTextarea(apiName, value string) textarea.Model {
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

func (m *model) newMyChannelMockTextarea(apiName, value string) textarea.Model {
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
		case IndexOf(Services, "EOS"):
			m.savedRootCursor = m.cursor
			m.screen = screenEOS
			m.cursor = 0
		case IndexOf(Services, "IDS"):
			m.savedRootCursor = m.cursor
			m.screen = screenIDS
			m.cursor = 0
		case IndexOf(Services, "SMIS"):
			m.savedRootCursor = m.cursor
			m.screen = screenSMIS
			m.cursor = 0
		case IndexOf(Services, "MYCHANNEL"):
			m.savedRootCursor = m.cursor
			m.screen = screenMyChannel
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
		m.jsonPlaceholder = PgzinvMockPlaceholder(name)
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
		m.jsonPlaceholder = PhxMockPlaceholder(name)
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
		m.jsonPlaceholder = DtMockPlaceholder(name)
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
		m.jsonPlaceholder = ImMockPlaceholder(name)
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
		m.jsonPlaceholder = EsbMockPlaceholder(name)
		m.ta = m.newESBMockTextarea(name, "")
		m.tas = m.newHttpCodeTextarea(name, "")
		layoutJSONEditor(m)
		cmd := m.tas.Focus()
		return m, cmd
	case screenEOS:
		if m.cursor < 0 || m.cursor >= len(EOSApis) {
			return m, nil
		}
		name := EOSApis[m.cursor]
		m.screen = screenEOSMockJSON
		m.jsonMockParent = screenEOS
		m.jsonMockResource = name
		m.saveNotice = ""
		m.jsonErr = ""
		m.jsonPlaceholder = EosMockPlaceholder(name)
		m.ta = m.newEOSMockTextarea(name, "")
		m.tas = m.newHttpCodeTextarea(name, "")
		layoutJSONEditor(m)
		cmd := m.tas.Focus()
		return m, cmd
	case screenIDS:
		if m.cursor < 0 || m.cursor >= len(IDSApis) {
			return m, nil
		}
		name := IDSApis[m.cursor]
		m.screen = screenIDSMockJSON
		m.jsonMockParent = screenIDS
		m.jsonMockResource = name
		m.saveNotice = ""
		m.jsonErr = ""
		m.jsonPlaceholder = IdsMockPlaceholder(name)
		m.ta = m.newIDSMockTextarea(name, "")
		m.tas = m.newHttpCodeTextarea(name, "")
		layoutJSONEditor(m)
		cmd := m.tas.Focus()
		return m, cmd
	case screenSMIS:
		if m.cursor < 0 || m.cursor >= len(SMISApis) {
			return m, nil
		}
		name := SMISApis[m.cursor]
		m.screen = screenSMISMockJSON
		m.jsonMockParent = screenSMIS
		m.jsonMockResource = name
		m.saveNotice = ""
		m.jsonErr = ""
		m.jsonPlaceholder = SmisMockPlaceholder(name)
		m.ta = m.newSMISMockTextarea(name, "")
		m.tas = m.newHttpCodeTextarea(name, "")
		layoutJSONEditor(m)
		cmd := m.tas.Focus()
		return m, cmd
	case screenMyChannel:
		if m.cursor < 0 || m.cursor >= len(MyChannelApis) {
			return m, nil
		}
		name := MyChannelApis[m.cursor]
		m.screen = screenMyChannelMockJSON
		m.jsonMockParent = screenMyChannel
		m.jsonMockResource = name
		m.saveNotice = ""
		m.jsonErr = ""
		m.jsonPlaceholder = MyChannelMockPlaceholder(name)
		m.ta = m.newMyChannelMockTextarea(name, "")
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
	m.jsonPlaceholder = ""
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
	case screenEOSMockJSON:
		return m.submitEOSMockJSON()
	case screenIDSMockJSON:
		return m.submitIDSMockJSON()
	case screenSMISMockJSON:
		return m.submitSMISMockJSON()
	case screenMyChannelMockJSON:
		return m.submitMyChannelMockJSON()
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

func (m *model) submitEOSMockJSON() (tea.Model, tea.Cmd) {
	raw := json.RawMessage(strings.TrimSpace(m.ta.Value()))
	httpCode, ok := parseHTTPStatusCode(m.tas.Value())
	if !ok {
		m.jsonErr = "invalid HttpStatusCode (expected 100-999)"
		m.screen = screenEOSMockJSON
		m.ta = m.newEOSMockTextarea(m.jsonMockResource, m.ta.Value())
		m.tas = m.newHttpCodeTextarea(m.jsonMockResource, m.tas.Value())
		layoutJSONEditor(m)
		cmd := m.tas.Focus()
		return m, cmd
	}
	_ = m.app.ApiInfoStore.UpdateHttpCode(m.jsonMockResource, httpCode)

	err := m.SetCustomResponse(m.jsonMockResource, raw)
	if err != nil {
		m.jsonErr = err.Error()
		m.screen = screenEOSMockJSON
		m.ta = m.newEOSMockTextarea(m.jsonMockResource, m.ta.Value())
		m.tas = m.newHttpCodeTextarea(m.jsonMockResource, m.tas.Value())
		layoutJSONEditor(m)
		cmd := m.tas.Focus()
		return m, cmd
	}

	m.screen = screenEOS
	m.jsonMockResource = ""
	m.jsonErr = ""
	m.saveNotice = "Saved successfully."
	m.ta.Blur()
	m.tas.Blur()
	return m, clearSaveNoticeAfter(2 * time.Second)
}

func (m *model) submitIDSMockJSON() (tea.Model, tea.Cmd) {
	raw := json.RawMessage(strings.TrimSpace(m.ta.Value()))
	httpCode, ok := parseHTTPStatusCode(m.tas.Value())
	if !ok {
		m.jsonErr = "invalid HttpStatusCode (expected 100-999)"
		m.screen = screenIDSMockJSON
		m.ta = m.newIDSMockTextarea(m.jsonMockResource, m.ta.Value())
		m.tas = m.newHttpCodeTextarea(m.jsonMockResource, m.tas.Value())
		layoutJSONEditor(m)
		cmd := m.tas.Focus()
		return m, cmd
	}
	_ = m.app.ApiInfoStore.UpdateHttpCode(m.jsonMockResource, httpCode)

	err := m.SetCustomResponse(m.jsonMockResource, raw)
	if err != nil {
		m.jsonErr = err.Error()
		m.screen = screenIDSMockJSON
		m.ta = m.newIDSMockTextarea(m.jsonMockResource, m.ta.Value())
		m.tas = m.newHttpCodeTextarea(m.jsonMockResource, m.tas.Value())
		layoutJSONEditor(m)
		cmd := m.tas.Focus()
		return m, cmd
	}

	m.screen = screenIDS
	m.jsonMockResource = ""
	m.jsonErr = ""
	m.saveNotice = "Saved successfully."
	m.ta.Blur()
	m.tas.Blur()
	return m, clearSaveNoticeAfter(2 * time.Second)
}

func (m *model) submitSMISMockJSON() (tea.Model, tea.Cmd) {
	raw := json.RawMessage(strings.TrimSpace(m.ta.Value()))
	httpCode, ok := parseHTTPStatusCode(m.tas.Value())
	if !ok {
		m.jsonErr = "invalid HttpStatusCode (expected 100-999)"
		m.screen = screenSMISMockJSON
		m.ta = m.newSMISMockTextarea(m.jsonMockResource, m.ta.Value())
		m.tas = m.newHttpCodeTextarea(m.jsonMockResource, m.tas.Value())
		layoutJSONEditor(m)
		cmd := m.tas.Focus()
		return m, cmd
	}
	_ = m.app.ApiInfoStore.UpdateHttpCode(m.jsonMockResource, httpCode)

	err := m.SetCustomResponse(m.jsonMockResource, raw)
	if err != nil {
		m.jsonErr = err.Error()
		m.screen = screenSMISMockJSON
		m.ta = m.newSMISMockTextarea(m.jsonMockResource, m.ta.Value())
		m.tas = m.newHttpCodeTextarea(m.jsonMockResource, m.tas.Value())
		layoutJSONEditor(m)
		cmd := m.tas.Focus()
		return m, cmd
	}

	m.screen = screenSMIS
	m.jsonMockResource = ""
	m.jsonErr = ""
	m.saveNotice = "Saved successfully."
	m.ta.Blur()
	m.tas.Blur()
	return m, clearSaveNoticeAfter(2 * time.Second)
}

func (m *model) submitMyChannelMockJSON() (tea.Model, tea.Cmd) {
	raw := json.RawMessage(strings.TrimSpace(m.ta.Value()))
	httpCode, ok := parseHTTPStatusCode(m.tas.Value())
	if !ok {
		m.jsonErr = "invalid HttpStatusCode (expected 100-999)"
		m.screen = screenMyChannelMockJSON
		m.ta = m.newMyChannelMockTextarea(m.jsonMockResource, m.ta.Value())
		m.tas = m.newHttpCodeTextarea(m.jsonMockResource, m.tas.Value())
		layoutJSONEditor(m)
		cmd := m.tas.Focus()
		return m, cmd
	}
	_ = m.app.ApiInfoStore.UpdateHttpCode(m.jsonMockResource, httpCode)

	err := m.SetCustomResponse(m.jsonMockResource, raw)
	if err != nil {
		m.jsonErr = err.Error()
		m.screen = screenMyChannelMockJSON
		m.ta = m.newMyChannelMockTextarea(m.jsonMockResource, m.ta.Value())
		m.tas = m.newHttpCodeTextarea(m.jsonMockResource, m.tas.Value())
		layoutJSONEditor(m)
		cmd := m.tas.Focus()
		return m, cmd
	}

	m.screen = screenMyChannel
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
	case screenServiceProvisioningMockJSON, screenPHXMockJSON, screenDTMockJSON, screenIMMockJSON, screenESBMockJSON,
		screenEOSMockJSON, screenIDSMockJSON, screenSMISMockJSON, screenMyChannelMockJSON:
		return m.leaveJSONEditor()
	case screenServiceProvisioning:
		m.screen = screenPGZINV
		m.cursor = m.savedPGZINVCursor
	case screenPGZINV:
		m.screen = screenRoot
		m.cursor = m.savedRootCursor
	case screenPHX, screenDT, screenIM, screenESB, screenEOS, screenIDS, screenSMIS, screenMyChannel:
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
	case screenEOS:
		return "EOS"
	case screenIDS:
		return "IDS"
	case screenSMIS:
		return "SMIS"
	case screenMyChannel:
		return "MYCHANNEL"
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
	case screenEOSMockJSON:
		return "EOS > " + m.jsonMockResource + " [JSON]"
	case screenIDSMockJSON:
		return "IDS > " + m.jsonMockResource + " [JSON]"
	case screenSMISMockJSON:
		return "SMIS > " + m.jsonMockResource + " [JSON]"
	case screenMyChannelMockJSON:
		return "MYCHANNEL > " + m.jsonMockResource + " [JSON]"
	default:
		return "mockTP"
	}
}

func (m *model) View() string {
	if isJSONMockScreen(m.screen) {
		var b strings.Builder
		b.WriteString(styleTitle.Render(m.breadcrumb()))
		b.WriteString("\n\n")
		// b.WriteString("Placeholder: " + m.jsonPlaceholder + "\n")
		b.WriteString("HttpStatusCode: " + m.tas.View())
		b.WriteString("\n\n")
		b.WriteString("Response:")
		b.WriteString("\n")
		b.WriteString(m.ta.View())
		if m.jsonErr != "" {
			w := editorContentWidth(m.width)
			errText := "Error: " + m.jsonErr
			// Wrap onto at most 2 lines; truncate the second line with … if still too long.
			line1, line2 := errText, ""
			if len(errText) > w && w > 0 {
				line1 = errText[:w]
				line2 = errText[w:]
				if len(line2) > w && w > 1 {
					line2 = line2[:w-1] + "…"
				}
			}
			b.WriteString("\n")
			b.WriteString(styleErr.Render(line1))
			b.WriteString("\n")
			b.WriteString(styleErr.Render(line2))
			b.WriteString("\n")
		} else {
			b.WriteString("\n\n\n")
		}
		b.WriteString(styleHelp.Render(renderHelpKeys("Tab switch input · Esc cancel · Ctrl+S save mock · Enter = new line · Ctrl+C copy placeholder · Ctrl+X quit")))
		b.WriteString("\n")
		b.WriteString(styleHelp.Render(renderHelpKeys("F4 clear editor · Ctrl+Home / Ctrl+End jump document")))
		b.WriteString("\n")
		b.WriteString(styleHelp.Render("(clear editor + save) to reset mock to default"))
		return b.String()
	}

	var b strings.Builder
	b.WriteString(styleTitle.Render(m.breadcrumb()))
	if (m.screen == screenServiceProvisioning || m.screen == screenPHX || m.screen == screenDT ||
		m.screen == screenIM || m.screen == screenESB || m.screen == screenEOS ||
		m.screen == screenIDS || m.screen == screenSMIS || m.screen == screenMyChannel) && m.saveNotice != "" {
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
		if m.screen == screenEOS {
			padding := eosLabelWidth() - len(label)
			if padding < 0 {
				padding = 0
			}
			displayLabel = label + strings.Repeat(" ", padding+1) + eosStateIndicator(label)
		}
		if m.screen == screenIDS {
			padding := idsLabelWidth() - len(label)
			if padding < 0 {
				padding = 0
			}
			displayLabel = label + strings.Repeat(" ", padding+1) + idsStateIndicator(label)
		}
		if m.screen == screenSMIS {
			padding := smisLabelWidth() - len(label)
			if padding < 0 {
				padding = 0
			}
			displayLabel = label + strings.Repeat(" ", padding+1) + smisStateIndicator(label)
		}
		if m.screen == screenMyChannel {
			padding := myChannelLabelWidth() - len(label)
			if padding < 0 {
				padding = 0
			}
			displayLabel = label + strings.Repeat(" ", padding+1) + myChannelStateIndicator(label)
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
		b.WriteString(styleHelp.Render(renderHelpKeys("↑/↓ · Enter open JSON · t toggle selected API state · Esc back (root: quit) · q quit")))
		b.WriteString("\n")
		b.WriteString(styleHelp.Render("Legend: " + styleOK.Render("S") + " = Success · " + styleErr.Render("E") + " = Error · " + styleCustom.Render("C") + " = Custom"))
		b.WriteString("\n")
		b.WriteString(styleHelp.Render(renderHelpKeys("Note: t cycles through " + styleOK.Render("S") + " → " + styleErr.Render("E") + " → " + styleCustom.Render("C") + " → " + styleOK.Render("S") + ".")))
	case screenPHX:
		b.WriteString(styleHelp.Render(renderHelpKeys("↑/↓ · Enter open JSON · t toggle selected API state · Esc back (root: quit) · q quit")))
		b.WriteString("\n")
		b.WriteString(styleHelp.Render("Legend: " + styleOK.Render("S") + " = Success · " + styleErr.Render("E") + " = Error · " + styleCustom.Render("C") + " = Custom"))
		b.WriteString("\n")
		b.WriteString(styleHelp.Render(renderHelpKeys("Note: t cycles through " + styleOK.Render("S") + " → " + styleErr.Render("E") + " → " + styleCustom.Render("C") + " → " + styleOK.Render("S") + ".")))
	case screenDT:
		b.WriteString(styleHelp.Render(renderHelpKeys("↑/↓ · Enter open JSON · t toggle selected API state · Esc back (root: quit) · q quit")))
		b.WriteString("\n")
		b.WriteString(styleHelp.Render("Legend: " + styleOK.Render("S") + " = Success · " + styleErr.Render("E") + " = Error · " + styleCustom.Render("C") + " = Custom"))
		b.WriteString("\n")
		b.WriteString(styleHelp.Render(renderHelpKeys("Note: t cycles through " + styleOK.Render("S") + " → " + styleErr.Render("E") + " → " + styleCustom.Render("C") + " → " + styleOK.Render("S") + ".")))
	case screenIM:
		b.WriteString(styleHelp.Render(renderHelpKeys("↑/↓ · Enter open JSON · t toggle selected API state · Esc back (root: quit) · q quit")))
		b.WriteString("\n")
		b.WriteString(styleHelp.Render("Legend: " + styleOK.Render("S") + " = Success · " + styleErr.Render("E") + " = Error · " + styleCustom.Render("C") + " = Custom"))
		b.WriteString("\n")
		b.WriteString(styleHelp.Render(renderHelpKeys("Note: t cycles through " + styleOK.Render("S") + " → " + styleErr.Render("E") + " → " + styleCustom.Render("C") + " → " + styleOK.Render("S") + ".")))
	case screenESB, screenEOS, screenIDS, screenSMIS, screenMyChannel:
		b.WriteString(styleHelp.Render(renderHelpKeys("↑/↓ · Enter open JSON · t toggle selected API state · Esc back (root: quit) · q quit")))
		b.WriteString("\n")
		b.WriteString(styleHelp.Render("Legend: " + styleOK.Render("S") + " = Success · " + styleErr.Render("E") + " = Error · " + styleCustom.Render("C") + " = Custom"))
		b.WriteString("\n")
		b.WriteString(styleHelp.Render(renderHelpKeys("Note: t cycles through " + styleOK.Render("S") + " → " + styleErr.Render("E") + " → " + styleCustom.Render("C") + " → " + styleOK.Render("S") + ".")))
	default:
		b.WriteString(styleHelp.Render(renderHelpKeys("↑/↓ · Enter open · Esc back (root: quit) · q quit")))
	}
	return b.String()
}
