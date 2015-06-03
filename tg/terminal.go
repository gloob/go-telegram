package tg

import (
	"github.com/jroimartin/gocui"

	"log"
)

// A window is similar to a page in a MCV environment. It's a set of Layout,
// meaning the definition of the display composition and a controller that
// handle the behaviour.
type Window struct {
	Layout     string
	Controller string
}

func NewWindow() *Window {
	return &Window{Layout: "as", Controller: "es"}
}

func (w *Window) Destroy() {
	// Release resources and memory.
}

type Terminal struct {
	Gui     *gocui.Gui
	Windows []*Window
}

func NewTerminal() *Terminal {
	return &Terminal{Gui: gocui.NewGui()}
}

func (term *Terminal) Init() {
	if err := term.Gui.Init(); err != nil {
		log.Panicln(err)
	}

	term.Gui.SetLayout(welcomeLayout)

	if err := keybindings(term.Gui); err != nil {
		log.Panicln(err)
	}

	term.Gui.SelBgColor = gocui.ColorGreen
	term.Gui.SelFgColor = gocui.ColorBlack
	term.Gui.ShowCursor = true
}

func (term *Terminal) Loop() {
	var err error = term.Gui.MainLoop()
	if err != nil && err != gocui.Quit {
		log.Panicln(err)
	}
}

func (term *Terminal) AddWindow(w *Window) {
	// Add window to list of windows.
	term.Windows[0] = w
	// term.Gui.SetLayout(layout)
}

func (term *Terminal) Destroy() {
	term.Gui.Close()
}

// Define main window.
func welcomeLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("main", 30, -1, maxX, maxY); err != nil {
		if err != gocui.ErrorUnkView {
			return err
		}
		v.Editable = true
		v.Wrap = true
		if err := g.SetCurrentView("main"); err != nil {
			return err
		}
	}

	return nil
}

// Define initial credentials window.

// Define welcome window.

// Define global keybindings.
func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.Quit
}
