package testdata

type language7 int

const (
	golang7 language7 = iota
	swift7
	objectivec7
	ruby7
	typescript7
)

func function7(x language7) {
	switch x {
	case swift7, ruby7, golang7, typescript7:
		println("match case")
	case objectivec7:
		println("objectivec7")
	default:
		println("default")
	}
}
