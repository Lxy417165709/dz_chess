package model

var TestCardGroups = []*CardGroup{
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
	{
		Card: &Card{
			Hero: &Hero{
				ID:     2,
				Name:   "大肿",
				Stage:  1,
				Worth:  1,
				Layer:  1,
				Skills: nil,
			},
		},
		NumOfCard: 5,
	},
	{
		Card: &Card{
			Hero: &Hero{
				ID:     101,
				Name:   "凯丽",
				Stage:  2,
				Worth:  2,
				Layer:  1,
				Skills: nil,
			},
		},
		NumOfCard: 5,
	},
	{
		Card: &Card{
			Hero: &Hero{
				ID:     1001,
				Name:   "狐狸",
				Stage:  3,
				Worth:  3,
				Layer:  1,
				Skills: nil,
			},
		},
		NumOfCard: 5,
	},
}

var IdToLayerToCardMap = map[int64]map[Layer]*Card{
	1: {
		1: &Card{
			Hero: &Hero{
				ID:     1,
				Name:   "德玛西亚",
				Stage:  1,
				Worth:  1,
				Layer:  1,
				Skills: nil,
				HP:     100,
				MP:     100,
			},
		},
		2: &Card{
			Hero: &Hero{
				ID:     1,
				Name:   "德玛西亚",
				Stage:  1,
				Worth:  3,
				Layer:  2,
				Skills: nil,
				HP:     1000,
				MP:     100,
			},
		},
		3: &Card{
			Hero: &Hero{
				ID:     1,
				Name:   "德玛西亚",
				Stage:  1,
				Worth:  9,
				Layer:  3,
				Skills: nil,
				HP:     10000,
				MP:     100,
			},
		},
	},
}

var TestCards = CardGroupsToCards(TestCardGroups)

func BuildStageToCardsMap(cards []*Card) map[Stage][]*Card {
	stageToCardsMap := make(map[Stage][]*Card)
	for _, card := range cards {
		stageToCardsMap[card.Hero.Stage] = append(stageToCardsMap[card.Hero.Stage], card)
	}
	return stageToCardsMap
}

func BuildLayerToIdToCardsMap(cards []*Card) map[Stage]map[int64][]*Card {
	layerToCardsMap := BuildLayerToCardsMap(cards)
	return BuildLayerToIdToCardsMapByLayerToCardsMap(layerToCardsMap)
}

func BuildLayerToCardsMap(cards []*Card) map[Stage][]*Card {
	layerToCardsMap := make(map[Layer][]*Card)
	for _, card := range cards {
		layerToCardsMap[card.Hero.Layer] = append(layerToCardsMap[card.Hero.Layer], card)
	}
	return layerToCardsMap
}

func BuildLayerToIdToCardsMapByLayerToCardsMap(layerToCardsMap map[Layer][]*Card) map[Layer]map[int64][]*Card {
	layerToIdToCardsMap := make(map[Layer]map[int64][]*Card)
	for layer, cards := range layerToCardsMap {
		if layerToIdToCardsMap[layer] == nil {
			layerToIdToCardsMap[layer] = make(map[int64][]*Card)
		}
		for _, card := range cards {
			layerToIdToCardsMap[layer][card.Hero.ID] = append(layerToIdToCardsMap[layer][card.Hero.ID], card)
		}
	}
	return layerToIdToCardsMap
}
