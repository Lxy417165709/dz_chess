package logic

import (
	"encoding/json"
	"fmt"
	"github.com/Lxy417165709/dz_chess/model"
	"testing"
)

func TestGetCards(t *testing.T) {
	stageToCardsMap := model.BuildStageToCardsMap(model.TestCards)
	stageToProbabilityMap := map[model.Stage]model.Probability{
		1: 50,
		2: 30,
		3: 20,
	}
	cards := GetCards(stageToCardsMap, stageToProbabilityMap, 3)
	cardJson, _ := json.MarshalIndent(model.BuildLayerToIdToCardsMap(cards), "", "\t")
	fmt.Printf("%+v", string(cardJson))
}

func TestUpgradeCard(t *testing.T) {
	cards := UpgradeCard([]*model.Card{
		model.DeepCopyCard(model.CardDmxy1),
		model.DeepCopyCard(model.CardDmxy1),
		model.DeepCopyCard(model.CardDmxy1),
		model.DeepCopyCard(model.CardDmxy1),
		model.DeepCopyCard(model.CardDmxy1),
		model.DeepCopyCard(model.CardDmxy1),
		model.DeepCopyCard(model.CardDmxy1),
		model.DeepCopyCard(model.CardDmxy1),
		model.DeepCopyCard(model.CardDmxy1),
		model.DeepCopyCard(model.CardDmxy1),
		model.DeepCopyCard(model.CardDmxy1),
		model.DeepCopyCard(model.CardDmxy1),
		model.DeepCopyCard(model.CardDmxy1),
		model.DeepCopyCard(model.CardDmxy1),
	}, model.IdToLayerToCardMap)
	cardJson, _ := json.MarshalIndent(model.BuildLayerToIdToCardsMap(cards), "", "\t")
	fmt.Printf("%+v", string(cardJson))
}
