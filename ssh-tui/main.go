package main

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	lm "github.com/charmbracelet/wish/logging"
)

var (
	docsDir = "/var/lib/devopstech/docs"
)

func main() {
	if env := os.Getenv("DOCS_DIR"); env != "" {
		docsDir = env
	}

	// Ensure host key directory exists
	if err := os.MkdirAll(".ssh", 0700); err != nil {
		log.Printf("Warning: could not create .ssh directory: %v", err)
	}

	s, err := wish.NewServer(
		wish.WithAddress(":2222"),
		wish.WithHostKeyPath(".ssh/term_info_ed25519"),
		wish.WithMiddleware(
			bm.Middleware(teaHandler),
			lm.Middleware(),
		),
	)
	if err != nil {
		log.Fatalln(err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("Starting SSH server on %s", s.Addr)

	go func() {
		if err = s.ListenAndServe(); err != nil && err != ssh.ErrServerClosed {
			log.Fatalln(err)
		}
	}()

	<-done
	log.Println("Stopping SSH server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, active := s.Pty()
	if !active {
		return nil, nil
	}

	// Initialize model
	m := initialModel(pty.Window.Width, pty.Window.Height)
	return m, []tea.ProgramOption{tea.WithAltScreen()}
}

type item struct {
	title, path string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.path }
func (i item) FilterValue() string { return i.title + " " + i.path }

type model struct {
	list     list.Model
	viewport viewport.Model
	ready    bool
	content  string
	width    int
	height   int
}

func initialModel(w, h int) model {
	items := loadDocs()

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "DevOpsTech Docs"
	l.SetShowHelp(false)

	return model{
		list:   l,
		width:  w,
		height: h,
	}
}

func loadDocs() []list.Item {
	var items []list.Item
	filepath.WalkDir(docsDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".md") {
			rel, _ := filepath.Rel(docsDir, path)
			items = append(items, item{
				title: d.Name(),
				path:  rel,
			})
		}
		return nil
	})
	return items
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		// Layout: 30% list, 70% viewport
		listWidth := int(float64(m.width) * 0.3)
		if listWidth < 20 {
			listWidth = 20
		}
		viewportWidth := m.width - listWidth - 4 // margins

		m.list.SetSize(listWidth, m.height-2)

		if !m.ready {
			m.viewport = viewport.New(viewportWidth, m.height-2)
			m.ready = true
			// Load first item
			if len(m.list.Items()) > 0 {
				m.selectItem(m.list.Items()[0].(item))
			}
		} else {
			m.viewport.Width = viewportWidth
			m.viewport.Height = m.height - 2
			// Re-render content with new width
			m.renderContent()
		}

	case tea.KeyMsg:
		// If list is focused (default), it handles keys.
		// But we want to scroll viewport if list is not focused?
		// For simplicity, let's say:
		// Tab switches focus? Or just use arrows for list and something else for viewport?
		// The user asked for: / search, j/k nav, enter open.
		// List handles / and j/k.
		// We need to handle viewport scrolling.
		// Let's use Ctrl+j/k for viewport or PageUp/Down.

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			if i, ok := m.list.SelectedItem().(item); ok {
				m.selectItem(i)
			}
		}
	}

	// Update list
	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)

	// Update viewport
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m *model) selectItem(i item) {
	fullPath := filepath.Join(docsDir, i.path)
	content, err := os.ReadFile(fullPath)
	if err != nil {
		m.content = fmt.Sprintf("Error reading file: %v", err)
	} else {
		m.content = string(content)
	}
	m.renderContent()
	m.viewport.GotoTop()
}

func (m *model) renderContent() {
	r, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(m.viewport.Width),
	)
	out, err := r.Render(m.content)
	if err != nil {
		m.viewport.SetContent("Error rendering markdown")
		return
	}
	m.viewport.SetContent(out)
}

func (m model) View() string {
	if !m.ready {
		return "Initializing..."
	}

	listStyle := lipgloss.NewStyle().
		Width(m.list.Width()).
		Height(m.height).
		Border(lipgloss.NormalBorder(), false, true, false, false) // Right border

	viewStyle := lipgloss.NewStyle().
		Width(m.viewport.Width).
		Height(m.height).
		PaddingLeft(2)

	return lipgloss.JoinHorizontal(lipgloss.Top,
		listStyle.Render(m.list.View()),
		viewStyle.Render(m.viewport.View()),
	)
}
