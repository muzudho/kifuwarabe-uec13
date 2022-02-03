package entities

// Ren - 連
type Ren struct {
	// LibertyArea - 呼吸点の数
	LibertyArea int
	// StoneArea - 石の数
	StoneArea int
	// Color - 石の色
	Color Stone
}

func NewRen(libertyArea int, stoneArea int, color Stone) *Ren {
	var ren = new(Ren)
	ren.LibertyArea = libertyArea
	ren.StoneArea = stoneArea
	ren.Color = color
	return ren
}
