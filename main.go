package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const Version = "1.0.0"

func main() {
	application := app.New()

	// Create window - FIXED: added title argument
	w := application.NewWindow("Cache Wipe - System Optimizer")
	w.Resize(fyne.NewSize(900, 700))

	// Create main UI
	mainContent := createMainUI()
	w.SetContent(mainContent)

	w.ShowAndRun()
}

func createMainUI() *fyne.Container {
	// Header
	title := canvas.NewText("Cache Wipe", nil)
	title.TextSize = 28
	title.TextStyle.Bold = true

	subtitle := canvas.NewText("v"+Version+" - Clear cache and optimize your system", nil)
	subtitle.TextSize = 12

	header := container.NewVBox(title, subtitle)

	// Statistics
	statsLabel := widget.NewLabel("System Status: Ready")
	cacheLabel := widget.NewLabel("Total Cache: Scanning...")

	// Buttons
	scanBtn := widget.NewButton("📊 Scan Cache", func() {
		log.Println("Scan started")
		cacheLabel.SetText("Total Cache: Scanning in progress...")
	})
	scanBtn.Importance = widget.HighImportance

	cleanBtn := widget.NewButton("🗑️ Clean Cache", func() {
		log.Println("Cleanup started")
		cacheLabel.SetText("Total Cache: Cleanup in progress...")
	})
	cleanBtn.Importance = widget.HighImportance

	settingsBtn := widget.NewButton("⚙️ Settings", func() {
		log.Println("Settings opened")
	})

	// Main layout
	mainContent := container.NewVBox(
		header,
		widget.NewSeparator(),
		statsLabel,
		cacheLabel,
		widget.NewSeparator(),
		container.NewHBox(scanBtn, cleanBtn, settingsBtn),
	)

	return mainContent
}
