package model

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
	Hero Hero
}

// Hero 英雄。
type Hero struct {
	ID     int64   // ID。
	Name   string  // 名称。
	Stage  Stage   // 阶段。
	Worth  int32   // 身价。示例: 1费、2费、5费。
	Layer  Layer   // 星级。
	Skills []int32 // 技能。
}
