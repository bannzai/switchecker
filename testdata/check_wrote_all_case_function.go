package testdata

type language2 int

const (
	golang language2 = iota
	swift
	objectivec
	ruby
	typescript
)

func function2(x language2) {
	switch x {
	case swift:
		println("swift")
	case ruby:
		println("ruby")
	case golang:
		println("golang")
	case typescript:
		println("typescript")
	case objectivec:
		println("objectivec")
	default:
		println("default")
	}
}
