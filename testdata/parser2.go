package testdata

type staticlanguage int

const (
	golang staticlanguage = iota
	swift
)

type s struct{}
type i interface{}

func abc() {

}

type dynamiclanguage int

const (
	ruby dynamiclanguage = iota
	python
)
