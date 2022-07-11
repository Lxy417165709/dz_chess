package logic

import (
	"github.com/Lxy417165709/dz_chess/model"
	"math/rand"
	"sort"
)

func GetCards(stageToCardsMap map[model.Stage][]*model.Card,
	stageToProbabilityMap map[model.Stage]model.Probability, numOfCard int64) []*model.Card {
	cards := make([]*model.Card, 0)
	for i := 0; i < int(numOfCard); i++ {
		stageSeed := RandNum(1, 100)
		stageProbabilityPairs := ToStageProbabilityPairs(stageToProbabilityMap)
		for _, pair := range stageProbabilityPairs {
			//
			if !isInRange(stageSeed, 1, int(pair.Probability)) {
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

func ToStageProbabilityPairs(stageToProbabilityMap map[model.Stage]model.Probability) []*StageProbabilityPair {
	stageProbabilityPairs := make([]*StageProbabilityPair, 0)
	for stage, probability := range stageToProbabilityMap {
		stageProbabilityPairs = append(stageProbabilityPairs, &StageProbabilityPair{
			Stage:       stage,
			Probability: probability,
		})
	}
	sort.Slice(stageProbabilityPairs, func(i, j int) bool {
		return stageProbabilityPairs[i].Stage < stageProbabilityPairs[j].Stage
	})
	return stageProbabilityPairs
}

type StageProbabilityPair struct {
	Stage       model.Stage
	Probability model.Probability
}

// isInRange num 是否位于区间 [leftEndPoint, rightEndPoint] 内。
func isInRange(num, leftEndPoint, rightEndPoint int) bool {
	return num >= leftEndPoint && num <= rightEndPoint
}

// RandNum 返回区间 [leftEndPoint, rightEndPoint] 中任意数。
func RandNum(leftEndPoint, rightEndPoint int) int {
	return rand.Intn(rightEndPoint-leftEndPoint+1) + leftEndPoint
}
