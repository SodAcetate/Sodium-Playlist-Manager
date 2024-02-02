package ui_terminal

import (
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func getManualFrame() *tview.Frame {
	manstr, _ := os.ReadFile("config/manual")

	var manualText = tview.NewTextView().SetText(string(manstr)).
		SetTextColor(tcell.ColorGreen)
	var manualFrame = tview.NewFrame(manualText)

	manualFrame.SetBorder(true)

	manualFrame.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' {
			uiBody.RemoveItem(manualFrame)
		}
		return event
	})

	return manualFrame
}
