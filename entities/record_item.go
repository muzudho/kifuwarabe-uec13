package entities

// RecordItem - 棋譜の1手分
type RecordItem struct {
	// Z - 着手
	Z Point
	// Time - 消費時間
	Time float64
}

// SetZ - 着手
func (recItem *RecordItem) SetZ(z Point) {
	recItem.Z = z
}
func (recItem *RecordItem) GetZ() Point {
	return recItem.Z
}

// SetTime - 消費時間
func (recItem *RecordItem) SetTime(time float64) {
	recItem.Time = time
}
func (recItem *RecordItem) GetTime() float64 {
	return recItem.Time
}
