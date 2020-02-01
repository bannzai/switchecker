package testdata

type language6 int

const (
	golang6 language6 = iota
	swift6
	objectivec6
	ruby6
	typescript6
)

func function6(x language6) {
	switch x {
	case swift6, ruby6, golang6, typescript6, objectivec6:
		println("match case")
	default:
		println("default")
	}
}
