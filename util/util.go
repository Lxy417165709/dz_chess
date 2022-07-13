package util

import "math/rand"

// IsInRange num 是否位于区间 [leftEndPoint, rightEndPoint] 内。
func IsInRange(num, leftEndPoint, rightEndPoint int) bool {
	return num >= leftEndPoint && num <= rightEndPoint
}

// RandNum 返回区间 [leftEndPoint, rightEndPoint] 中任意数。
func RandNum(leftEndPoint, rightEndPoint int) int {
	return rand.Intn(rightEndPoint-leftEndPoint+1) + leftEndPoint
}
