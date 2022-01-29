package entities

// RecordItem - 棋譜の1手分
type RecordItem struct {
	// Z - 着手
	Z int
	// Time - 消費時間
	Time float64
}

// SetZ - 着手
func (recItem *RecordItem) SetZ(z int) {
	recItem.Z = z
}
func (recItem *RecordItem) GetZ() int {
	return recItem.Z
}

// SetTime - 消費時間
func (recItem *RecordItem) SetTime(time float64) {
	recItem.Time = time
}
func (recItem *RecordItem) GetTime() float64 {
	return recItem.Time
}
