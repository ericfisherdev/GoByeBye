package ui

import (
	"image"
)

// Window represents the main application window
type Window interface {
	Show() error
	Hide() error
	OnDrop(handler func(files []string))
	DisplayImage(img image.Image) error
}

// MainWindow implements the Window interface
type MainWindow struct {
	title  string
	width  int
	height int
}

// NewMainWindow creates a new main window
func NewMainWindow(title string, width, height int) *MainWindow {
	return &MainWindow{
		title:  title,
		width:  width,
		height: height,
	}
}

// Show displays the window
func (w *MainWindow) Show() error {
	// TODO: Implement window display
	return nil
}

// Hide hides the window
func (w *MainWindow) Hide() error {
	// TODO: Implement window hiding
	return nil
}

// OnDrop registers a handler for dropped files
func (w *MainWindow) OnDrop(handler func(files []string)) {
	// TODO: Implement drag-and-drop functionality
}

// DisplayImage shows an image in the window
func (w *MainWindow) DisplayImage(img image.Image) error {
	// TODO: Implement image display
	return nil
}