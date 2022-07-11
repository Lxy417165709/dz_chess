package logic

import (
	"encoding/json"
	"fmt"
	"github.com/Lxy417165709/dz_chess/model"
	"testing"
)

func TestGetCards(t *testing.T) {
	stageToCardsMap := map[model.Stage][]*model.Card{
		1: {
			{
				Hero: model.Hero{
					ID:    1,
					Stage: 1,
				},
			},
			{
				Hero: model.Hero{
					ID:    1,
					Stage: 1,
				},
			},
			{
				Hero: model.Hero{
					ID:    1,
					Stage: 1,
				},
			},
		},
		2: {
			{
				Hero: model.Hero{
					ID:    2,
					Stage: 2,
				},
			},
			{
				Hero: model.Hero{
					ID:    2,
					Stage: 2,
				},
			},
			{
				Hero: model.Hero{
					ID:    2,
					Stage: 2,
				},
			},
		},
		3: {
			{
				Hero: model.Hero{
					ID:    3,
					Stage: 3,
				},
			},
			{
				Hero: model.Hero{
					ID:    3,
					Stage: 3,
				},
			},
			{
				Hero: model.Hero{
					ID:    3,
					Stage: 3,
				},
			},
		},
	}
	stageToProbabilityMap := map[model.Stage]model.Probability{
		1: 33,
		2: 33,
		3: 34,
	}
	cards := GetCards(stageToCardsMap, stageToProbabilityMap, 3)
	cardJson, _ := json.MarshalIndent(cards, "", "\t")
	fmt.Printf("%+v", string(cardJson))
}
