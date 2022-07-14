package model

import (
	"encoding/json"
	"sort"
)

// World 世界。
type World struct {
	Players  []*Player // 玩家。
	CardPool CardPool  // 卡池。
}

type Player struct {
}

type Chessboard struct {
}

// CardPool 卡池。
type CardPool struct {
	StageToIdToCardsMap map[int8]map[int64][]*Card
}

// CardBar 卡栏。
type CardBar struct {
	Cards []Card // 卡栏内的卡片。
}

// Card 卡片。
type Card struct {
	Hero *Hero
}

// Hero 英雄。
type Hero struct {
	ID     int64   // ID。
	Name   string  // 名称。
	Stage  Stage   // 阶段。
	Worth  int32   // 身价。示例: 1费、2费、5费。
	Layer  Layer   // 星级。
	Skills []int32 // 技能。
	HP     int64   // 血量。
	MP     int64   // 蓝量。
}

// CardGroup 卡组。
type CardGroup struct {
	Card      *Card
	NumOfCard uint32
}

type StageProbabilityPair struct {
	Stage       Stage
	Probability Probability
}

func ToStageProbabilityPairs(stageToProbabilityMap map[Stage]Probability) []*StageProbabilityPair {
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

func CardGroupsToCards(groups []*CardGroup) []*Card {
	cards := make([]*Card, 0)
	for _, group := range groups {
		for i := 0; i < int(group.NumOfCard); i++ {
			cards = append(cards, DeepCopyCard(group.Card))
		}
	}
	return cards
}

func DeepCopyCard(card *Card) *Card {
	objJson, _ := json.Marshal(card)
	var c Card
	_ = json.Unmarshal(objJson, &c)
	return &c
}
