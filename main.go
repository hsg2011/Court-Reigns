package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Stats struct {
	Finances int 
	Fans     int
	Morale   int
	Fitness  int
}

func (s *Stats) applyEffect(e Effect) bool {
	s.Finances += e.Finances
	s.Fans += e.Fans
	s.Morale += e.Morale
	s.Fitness += e.Fitness

	// Check Stats and print appropriate messages
	if s.Finances <= 0 {
		fmt.Println("Franchise is in financial trouble! You're FIRED!")
		return false
	}
	if s.Fans <= 0 {
		fmt.Println("Franchise has lost significant fans! You're FIRED!")
		return false
	}
	if s.Morale <= 0 {
		fmt.Println("Franchise morale is too low! You're FIRED!")
		return false
	}
	if s.Fitness <= 0 {
		fmt.Println("Franchise fitness is too low! You're FIRED!")
		return false
	}

	if s.Finances >= 100 {
		fmt.Println("Franchise is hording cash! Commissioner wants a word!")
		s.Finances = 5
	}

	if s.Fans >= 100 {
		fmt.Println("Fan control of the franchise is too high! Players feel extreme pressure performance declines!")
		s.Fans = 5
	}

	if s.Morale >= 100 {
		fmt.Println("Morale is too high! Players are getting complacent and performance declines!")
		s.Morale = 5
	}
	if s.Fitness >= 100 {
		fmt.Println("Fitness is too high! Players are getting overworked and performance declines!")
		s.Fitness = 5
	}
	return true
}

func (s Stats) print() {
	fmt.Printf("Finances: %d, Fans: %d, Morale: %d, Fitness: %d\n", s.Finances, s.Fans, s.Morale, s.Fitness)
} 

func main() {
  cards, err := LoadCards("cards.json")  // assuming cards.json sits next to main.go
  if err != nil {
    fmt.Println("Error loading cards:", err)
    os.Exit(1)
  }

  stats := Stats{50, 50, 50, 50}           
  reader := bufio.NewReader(os.Stdin)

  for i, card := range cards {
    fmt.Printf("\n[%d/%d] %s\n", i+1, len(cards), card.Text)
    stats.print()
    fmt.Print("Choose (l/r): ")

    input, _ := reader.ReadString('\n')
    choice := strings.TrimSpace(strings.ToLower(input))
    if choice != "l" && choice != "r" {
      fmt.Println("  ▶ please type 'l' or 'r'")
      i--  // retry
      continue
    }

    var e Effect
    if choice == "l" {
      e = card.Left
    } else {
      e = card.Right
    }


    if alive := stats.applyEffect(e); !alive {
      fmt.Println("\n💀 Game Over! Final Stats →", stats)
      return
    }

  }

  fmt.Println("\n🎉 You survived! Final Stats →", stats)
}