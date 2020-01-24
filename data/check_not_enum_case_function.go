package data

type language1 int

const (
	golang language1 = iota
	swift
	objectivec
	ruby
	typescript
)

func function(x language1) {
	switch x {
	case swift:
		println("swift")
	case ruby:
		println("ruby")
	default:
		println("default")
	}
}
