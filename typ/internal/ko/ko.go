package ko

// Type
type Type int8

const (
	노말 Type = iota
	불꽃
	물
	풀
	전기
	얼음
	격투
	독
	땅
	비행
	에스퍼
	벌레
	바위
	고스트
	드래곤
	악
	강철
	페어리
)

//go:generate stringer -type=Type
