package game

import (
    "encoding/json"
    "fmt"
    "os"
)

// Effect describes stat changes
type Effect struct {
    Finances int `json:"finances"`
    Morale   int `json:"morale"`
    Fitness  int `json:"fitness"`
    Fans     int `json:"fans"`
}

// Card bundles a prompt + its two outcomes
type Card struct {
    Text  string `json:"text"`
    Left  Effect `json:"left"`
    Right Effect `json:"right"`
}

// Stats holds your four bars
type Stats struct {
    Finances int
    Fans     int
    Morale   int
    Fitness  int
}

// Game holds the deck, stats, and position
type Game struct {
    Cards    []Card
    Stats    Stats
    Position int
}

// LoadCards parses your JSON into Go structs
func LoadCards(path string) ([]Card, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }
    var cards []Card
    if err := json.Unmarshal(data, &cards); err != nil {
        return nil, err
    }
    return cards, nil
}

// NewGame constructs a Game loaded from a JSON file
func NewGame(path string) (*Game, error) {
    cards, err := LoadCards(path)
    if err != nil {
        return nil, err
    }
    return &Game{
        Cards: cards,
        Stats: Stats{50, 50, 50, 50},
    }, nil
}

// CurrentCard returns the active card
func (g *Game) CurrentCard() Card {
    return g.Cards[g.Position]
}

// Apply records a left/right choice and advances the game.
// Returns false if the game is over.
func (g *Game) Apply(choice string) bool {
    var e Effect
    if choice == "l" {
        e = g.CurrentCard().Left
    } else {
        e = g.CurrentCard().Right
    }
    alive := g.Stats.applyEffect(e)
    g.Position++
    return alive && g.Position < len(g.Cards)
}

// String renders stats for logging or UI
func (s Stats) String() string {
    return fmt.Sprintf(
        "Finances: %d | Fans: %d | Morale: %d | Fitness: %d",
        s.Finances, s.Fans, s.Morale, s.Fitness,
    )
}

// applyEffect mutates Stats and returns false on terminal conditions
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

