package testdata

import "github.com/bannzai/switchecker/testdata/checker/x"

func function5(a x.Language4) {
	switch a {
	case Swift4:
		println("Swift4")
	case Ruby4:
		println("Ruby4")
	case Golang4:
		println("Golang4")
	case Typescript4:
		println("Typescript4")
		//	case Objectivec4:
		//		println("Objectivec4")
	default:
		println("default")
	}
}
