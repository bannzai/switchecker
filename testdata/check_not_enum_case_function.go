package data

type language1 int

const (
	golang1 language1 = iota
	swift1
	objectivec1
	ruby1
	typescript1
)

func function(x language1) {
	switch x {
	case swift1:
		println("swift1")
	case ruby1:
		println("ruby1")
	default:
		println("default")
	}
}
