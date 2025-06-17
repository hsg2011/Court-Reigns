package main

import (
	"fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/layout"
    "fyne.io/fyne/v2/widget"

    "github.com/hsg2011/court-reigns/game"
)

func main() {
    g, err := game.NewGame("cards.json")
    if err != nil {
        panic(err)
    }

    a := app.New()
    w := a.NewWindow("Court Reigns")

    cardLabel := widget.NewLabel("")
    statsLabel := widget.NewLabel("")
    leftBtn := widget.NewButton("← Swipe Left", nil)
    rightBtn := widget.NewButton("Swipe Right →", nil)

    // update UI helper
    update := func() {
        card := g.CurrentCard()
        cardLabel.SetText(card.Text)
        statsLabel.SetText(g.Stats.String())
    }

    // button callbacks
    leftBtn.OnTapped = func() {
        if !g.Apply("l") {
            cardLabel.SetText("💀 Game Over! " + g.Stats.String())
            leftBtn.Disable(); rightBtn.Disable()
            return
        }
        update()
    }
    rightBtn.OnTapped = func() {
        if !g.Apply("r") {
            cardLabel.SetText("💀 Game Over! " + g.Stats.String())
            leftBtn.Disable(); rightBtn.Disable()
            return
        }
        update()
    }

    // layout
    content := container.NewVBox(
        cardLabel,
        statsLabel,
        layout.NewSpacer(),
        container.NewHBox(leftBtn, layout.NewSpacer(), rightBtn),
    )
    w.SetContent(content)
    w.Resize(fyne.NewSize(400, 300))

    update()      // show first card
    w.ShowAndRun()
}
