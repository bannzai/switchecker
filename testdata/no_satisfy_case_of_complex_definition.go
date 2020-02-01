package testdata

type language9 int

const (
	golang9 language9 = iota
	swift9
	objectivec9
	ruby9
	typescript9
)

func function9(x language9) {
	switch x {
	case swift9, ruby9, golang9:
		println("match case")
	case typescript9:
		println("typescript9")
	default:
		println("default")
	}
}
