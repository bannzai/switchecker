package testdata

type language8 int

const (
	golang8 language8 = iota
	swift8
	objectivec8
	ruby8
	typescript8
)

func function8(x language8) {
	switch x {
	case swift8, ruby8, golang8, typescript8:
		println("match case")
	default:
		println("default")
	}
}
