package testdata

func a() {
	type language int

	const (
		golang language = iota
		swift
		objectivec
		ruby
		typescript
	)
}
