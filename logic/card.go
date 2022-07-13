package logic

import (
	"github.com/Lxy417165709/dz_chess/model"
	"github.com/Lxy417165709/dz_chess/util"
	"math/rand"
	"time"
)

// GetCards 卡池抽卡。
func GetCards(stageToCardsMap map[model.Stage][]*model.Card,
	stageToProbabilityMap map[model.Stage]model.Probability, numOfCard int64) []*model.Card {

	rand.Seed(time.Now().Unix())

	cards := make([]*model.Card, 0)
	for i := 0; i < int(numOfCard); i++ {
		stageSeed := util.RandNum(1, 100)
		stageProbabilityPairs := model.ToStageProbabilityPairs(stageToProbabilityMap)
		for _, pair := range stageProbabilityPairs {
			//
			if !util.IsInRange(stageSeed, 1, int(pair.Probability)) {
				stageSeed -= int(pair.Probability)
				continue
			}

			//
			meetStage := pair.Stage
			stageCards := stageToCardsMap[meetStage]

			//
			if len(stageCards) == 0 {
				break
			}

			//
			cardSeed := rand.Intn(len(stageCards))
			meetCard := stageCards[cardSeed]
			cards = append(cards, meetCard)

			//
			stageToCardsMap[meetStage][len(stageToCardsMap[meetStage])-1], stageToCardsMap[meetStage][cardSeed] =
				stageToCardsMap[meetStage][cardSeed], stageToCardsMap[meetStage][len(stageToCardsMap[meetStage])-1]
			stageToCardsMap[meetStage] = stageToCardsMap[meetStage][:len(stageToCardsMap[meetStage])-1]

			break
		}
	}
	return cards
}

func UpgradeCard(cards []*model.Card) []*model.Card {
	//layerToIdToCardsMap := model.BuildLayerToIdToCardsMap(cards)
	//upgradedCards := make([]*model.Card,0)
	//for layer, idToCardsMap := range layerToIdToCardsMap {
	//	for id,cards := range idToCardsMap{
	//
	//	}
	//}
	return nil
}

func getHexadecimalString(num int) string {
	return ""
}
