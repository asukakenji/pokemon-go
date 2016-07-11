package ja

// Type
type Type int8

const (
	なし Type = iota
	ノーマル
	ほのお
	みず
	くさ
	でんき
	こおり
	かくとう
	どく
	じめん
	ひこう
	エスパー
	むし
	いわ
	ゴースト
	ドラゴン
	あく
	はがね
	フェアリー
)

//go:generate stringer -type=Type
