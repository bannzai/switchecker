package testdata

type language3 int

const (
	golang3 language3 = iota
	swift3
	objectivec3
	ruby3
	typescript3
)

func function3(x language3) {
	switch x {
	case swift3:
		println("swift3")
	case ruby3:
		println("ruby3")
	case golang3:
		println("golang3")
	case typescript3:
		println("typescript3")
	case objectivec3:
		println("objectivec3")
	default:
		println("default")
	}
}

func function3_x(x language3) {
	switch x {
	case swift3:
		println("swift3")
	case ruby3:
		println("ruby3")
	default:
		println("default")
	}
}
