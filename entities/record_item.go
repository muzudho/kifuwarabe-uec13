package entities

type IRecordItemV01 interface {
	SetZ(z int) // 着手
	GetZ() int
}

type IRecordItemV02 interface {
	IRecordItemV01
	SetTime(z float64) // 消費時間
	GetTime() float64
}

// RecordItemV01 - 棋譜の1手分
type RecordItemV01 struct {
	// 着手
	Z int
}

func (recItem *RecordItemV01) SetZ(z int) {
	recItem.Z = z
}
func (recItem *RecordItemV01) GetZ() int {
	return recItem.Z
}

// RecordItemV02 - 棋譜の1手分
type RecordItemV02 struct {
	RecordItemV01
	// 消費時間
	Time float64
}

func (recItem *RecordItemV02) SetTime(time float64) {
	recItem.Time = time
}
func (recItem *RecordItemV02) GetTime() float64 {
	return recItem.Time
}
