package main

import (
	"encoding/json"

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





