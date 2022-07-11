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
	StageToCardSetsMap []*CardSet
}

// CardBar 卡栏。
type CardBar struct {
	Cards []Card // 卡栏内的卡片。
}

// CardSet 卡组，卡组内的卡片是相同的。
type CardSet struct {
	Cards []Card // 卡组内的卡片。
}

// Card 卡片。
type Card struct {
	Hero Hero
}

type Hero struct {
	ID     int64   // ID。
	Stage  int8    // 阶段。示例: 1阶、2阶、5阶。
	Worth  int32   // 身价。示例: 1费、2费、5费。
	Layer  int8    // 星级。示例: 1星、2星。
	Skills []int32 // 技能。
}
