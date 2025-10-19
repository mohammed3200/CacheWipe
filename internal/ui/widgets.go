package ui

import (
	"cache-wipe/internal/cache"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type CacheItemWidget struct {
	widget.BaseWidget
	item *cache.CacheItem
}

func NewCacheItemWidget(item *cache.CacheItem) *CacheItemWidget {
	w := &CacheItemWidget{item: item}
	w.ExtendBaseWidget(w)
	return w
}

func (w *CacheItemWidget) CreateRenderer() fyne.WidgetRenderer {
	checkbox := widget.NewCheck(fmt.Sprintf("%s (%s)", w.item.Path, formatBytes(w.item.Size)), func(b bool) {
		w.item.Selected = b
	})
	checkbox.SetChecked(w.item.Selected)

	details := canvas.NewText(fmt.Sprintf("Files: %d", w.item.FileCount), nil)
	details.TextSize = 10

	return widget.NewSimpleRenderer(container.NewVBox(checkbox, details))
}

type ProgressCard struct {
	widget.BaseWidget
	progress *cache.CleanupProgress
}

func NewProgressCard(progress *cache.CleanupProgress) *ProgressCard {
	w := &ProgressCard{progress: progress}
	w.ExtendBaseWidget(w)
	return w
}

func (w *ProgressCard) CreateRenderer() fyne.WidgetRenderer {
	progressBar := widget.NewProgressBar()
	progressBar.Max = float64(w.progress.Total)
	progressBar.Value = float64(w.progress.Current)

	statusText := canvas.NewText(w.progress.Status, nil)
	statusText.TextSize = 12

	deletedText := canvas.NewText(
		fmt.Sprintf("Deleted: %s (%d files)", formatBytes(w.progress.DeletedSize), w.progress.DeletedCount),
		nil,
	)
	deletedText.TextSize = 10

	return widget.NewSimpleRenderer(container.NewVBox(
		statusText,
		progressBar,
		deletedText,
	))
}
