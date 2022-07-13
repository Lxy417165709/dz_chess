package model

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCardGroupsToCards(t *testing.T) {
	cardJson, _ := json.MarshalIndent(CardGroupsToCards([]*CardGroup{
		{
			Card: &Card{
				Hero: &Hero{
					ID:     1,
					Name:   "德玛西亚",
					Stage:  1,
					Worth:  1,
					Layer:  1,
					Skills: nil,
				},
			},
			NumOfCard: 5,
		},
	}), "", "\t")
	fmt.Printf("%+v", string(cardJson))
}
