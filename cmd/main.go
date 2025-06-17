package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    "github.com/hsg2011/court-reigns/game"
)

func main() {
    g, err := game.NewGame("../cards.json")
    if err != nil {
        fmt.Println("Error loading cards:", err)
        os.Exit(1)
    }

    reader := bufio.NewReader(os.Stdin)
    for {
        card := g.CurrentCard()
        fmt.Printf("\n[%d/%d] %s\n", g.Position+1, len(g.Cards), card.Text)
        fmt.Println(g.Stats.String())
        fmt.Print("Choose (l/r): ")

        input, _ := reader.ReadString('\n')
        choice := strings.TrimSpace(strings.ToLower(input))
        if choice!="l" && choice!="r" {
            fmt.Println("Type 'l' or 'r'")
            continue
        }
        if !g.Apply(choice) {
            fmt.Println("\n💀 Game Over!", g.Stats.String())
            return
        }
        if g.Position >= len(g.Cards) {
            fmt.Println("\n🎉 You survived!", g.Stats.String())
            return
        }
    }
}
