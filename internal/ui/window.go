package ui

import (
	"cache-wipe/internal/cache"
	"fmt"
	"image/color"
	"log"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type AppState struct {
	scanner    *cache.Scanner
	cleaner    *cache.Cleaner
	scanResult *cache.ScanResult
	mu         sync.Mutex
}

var appState = &AppState{}

func SetupMainWindow(w fyne.Window, app fyne.App) error {
	w.SetTitle("Cache Cleaner - Optimize Your System")
	w.Resize(fyne.NewSize(900, 700))

	// Header
	headerBox := createHeader()

	// Statistics Panel
	statsPanel := createStatsPanel()

	// Tabs
	tabs := container.NewAppTabs(
		container.NewTabItem("Dashboard", createDashboardTab()),
		container.NewTabItem("Scan Results", createScanResultsTab(w)),
		container.NewTabItem("Settings", createSettingsTab()),
	)

	// Main layout
	mainContent := container.NewVBox(
		headerBox,
		widget.NewSeparator(),
		statsPanel,
		widget.NewSeparator(),
		tabs,
	)

	scrollContainer := container.NewScroll(mainContent)
	w.SetContent(scrollContainer)

	return nil
}

func createHeader() *fyne.Container {
	titleText := canvas.NewText("Cache Cleaner", color.White)
	titleText.TextSize = 28
	titleText.TextStyle.Bold = true

	subtitleText := canvas.NewText("Optimize your system by removing unnecessary files", color.White)
	subtitleText.TextSize = 12

	return container.NewVBox(
		titleText,
		subtitleText,
	)
}

func createStatsPanel() *fyne.Container {
	totalCacheLabel := widget.NewLabel("Total Cache: Scanning...")
	filesLabel := widget.NewLabel("Files: -")
	statusLabel := widget.NewLabel("Status: Ready")

	scanButton := widget.NewButton("📊 Scan System", func() {
		log.Println("Scanning cache...")
		totalCacheLabel.SetText("Total Cache: Scanning...")
		filesLabel.SetText("Files: -")
		statusLabel.SetText("Status: Scanning...")

		go scanCache(totalCacheLabel, filesLabel, statusLabel)
	})

	return container.NewVBox(
		container.NewHBox(scanButton),
		widget.NewSeparator(),
		totalCacheLabel,
		filesLabel,
		statusLabel,
	)
}

func createDashboardTab() *fyne.Container {
	cleanButton := widget.NewButton("🗑️  Clean Cache Now", func() {
		log.Println("Cleanup started")
	})
	cleanButton.Importance = widget.HighImportance

	dryRunButton := widget.NewButton("Preview Cleanup", func() {
		log.Println("Dry run started")
	})

	cancelButton := widget.NewButton("Cancel", func() {
		if appState.scanner != nil {
			appState.scanner.Stop()
		}
		if appState.cleaner != nil {
			appState.cleaner.Stop()
		}
	})

	progressBar := widget.NewProgressBar()

	return container.NewVBox(
		widget.NewCard("Cache Cleanup", "", container.NewVBox(
			widget.NewRichTextFromMarkdown("### Quick Cleanup\nRemove all detected cache files from your system."),
			progressBar,
			container.NewHBox(cleanButton, dryRunButton, cancelButton),
		)),
	)
}

func createScanResultsTab(w fyne.Window) *fyne.Container {
	resultsList := widget.NewList(
		func() int { return 0 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i widget.ListItemID, o fyne.CanvasObject) {},
	)

	expandButton := widget.NewButton("Expand All", func() {
		log.Println("Expand all results")
	})

	collapseButton := widget.NewButton("Collapse All", func() {
		log.Println("Collapse all results")
	})

	return container.NewVBox(
		container.NewHBox(expandButton, collapseButton),
		widget.NewSeparator(),
		resultsList,
	)
}

func createSettingsTab() *fyne.Container {
	autoCleanCheckbox := widget.NewCheck("Enable Auto-Cleanup", func(b bool) {
		log.Printf("Auto-cleanup: %v", b)
	})

	dryRunCheckbox := widget.NewCheck("Dry Run Mode", func(b bool) {
		log.Printf("Dry run mode: %v", b)
	})

	logCleanupCheckbox := widget.NewCheck("Log Cleanup Activities", func(b bool) {
		log.Printf("Log cleanup: %v", b)
	})

	saveButton := widget.NewButton("💾 Save Settings", func() {
		log.Println("Settings saved")
	})

	return container.NewVBox(
		widget.NewCard("Preferences", "", container.NewVBox(
			autoCleanCheckbox,
			dryRunCheckbox,
			logCleanupCheckbox,
		)),
		saveButton,
	)
}

func scanCache(totalLabel, filesLabel, statusLabel *widget.Label) {
	appState.mu.Lock()
	appState.scanner = cache.NewScanner()
	appState.mu.Unlock()

	result, err := appState.scanner.Scan(func(progress int64) {
		statusLabel.SetText(fmt.Sprintf("Status: Scanning... %d%%", progress))
	})

	if err != nil {
		statusLabel.SetText(fmt.Sprintf("Status: Error - %v", err))
		return
	}

	appState.mu.Lock()
	appState.scanResult = result
	appState.mu.Unlock()

	totalLabel.SetText(fmt.Sprintf("Total Cache: %s", formatBytes(result.TotalSize)))
	filesLabel.SetText(fmt.Sprintf("Files: %d", result.FileCount))
	statusLabel.SetText("Status: Scan Complete")
}

func formatBytes(bytes int64) string {
	units := []string{"B", "KB", "MB", "GB", "TB"}
	size := float64(bytes)

	for i, unit := range units {
		if size < 1024 || i == len(units)-1 {
			return fmt.Sprintf("%.2f %s", size, unit)
		}
		size /= 1024
	}

	return fmt.Sprintf("%.2f %s", size, "TB")
}
