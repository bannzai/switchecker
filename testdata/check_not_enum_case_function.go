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
	default:
		println("default")
	}
}
