package testdata

type language2 int

const (
	golang2 language2 = iota
	swift2
	objectivec2
	ruby2
	typescript2
)

func function2(x language2) {
	switch x {
	case swift2:
		println("swift2")
	case ruby2:
		println("ruby2")
	case golang2:
		println("golang2")
	case typescript2:
		println("typescript2")
	case objectivec2:
		println("objectivec2")
	default:
		println("default")
	}
}
