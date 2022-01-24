package config_obj

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/pelletier/go-toml"
)

// Config - Tomlファイル
type Config struct {
	Nngs Nngs
	Game Game
}

// Nngs - [Nngs] テーブル
type Nngs struct {
	Server string
	Port   uint16
	User   string
	Pass   string
}

// Game - [Game] テーブル
type Game struct {
	Komi      float32
	BoardSize int8
	MaxMoves  int16
	BoardData string
}

// GetBoardArray - 盤上の石の色の配列
func (config Config) GetBoardArray() []int {
	// 最後のカンマを削除しないと、要素数が 1 多くなってしまいます
	var s = strings.TrimRight(config.Game.BoardData, ",")
	var nodes = strings.Split(s, ",")
	var array = make([]int, len(nodes))
	for i, s := range nodes {
		var s = strings.Trim(s, " ")
		var color, _ = strconv.Atoi(s)
		array[i] = color
	}

	return array
}

// BoardSize - 何路盤か
func (config Config) BoardSize() int {
	return int(config.Game.BoardSize)
}

// SentinelBoardArea - 枠付きの盤上の交点の数
func (config Config) SentinelBoardArea() int {
	// Width - 枠込み。
	var Width = int(config.Game.BoardSize) + 2
	// BoardMax - 枠込み盤の配列サイズ
	return Width * Width
}

// Komi - float 32bit で足りるが、実行速度優先で float 64bit に変換して返します
func (config Config) Komi() float64 {
	return float64(config.Game.Komi)
}

// MaxMovesNum - 最大手数
func (config Config) MaxMovesNum() int {
	return int(config.Game.MaxMoves)
}

// LoadGameConf - ゲーム設定ファイルを読み込みます
func LoadGameConf(
	path string,
	onFatal func(string)) Config {

	// ファイル読込
	var fileData, err = ioutil.ReadFile(path)
	if err != nil {
		onFatal(fmt.Sprintf("path=%s", path))
		panic(err)
	}

	// Toml解析
	var binary = []byte(string(fileData))
	var config = Config{}
	toml.Unmarshal(binary, &config)

	return config
}
