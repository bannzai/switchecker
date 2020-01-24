package testdata

type language int

const (
	golang language = iota
	swift
	objectivec
	ruby
	typescript
)

func function(x language) {
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
