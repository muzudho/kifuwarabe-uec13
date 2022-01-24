package entities

// プレイアウトする回数（あとで設定されます）
var PlayoutTrialCount = 0

// UCTをループする回数
var UctLoopCount = 4500

// var UctLoopCount = 2500 // 少なめ（ペンキ塗り）
// var UctLoopCount = 4500 // 盤の上の方に集まりがち
// var UctLoopCount = 5000 // ペンキ塗りはしなくなる
// var UctLoopCount = 10000 // 多め（遅い感じ）

// 原始モンテカルロ試行回数（あとで設定されます）
var PrimitiveMonteCalroTrialCount = 0
