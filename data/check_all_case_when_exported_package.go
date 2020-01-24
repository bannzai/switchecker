package data

import "./x"

func functionForX(f x.Fruit) {
	switch f {
	case x.Apple:
		println("x.Apple")
	case x.Orange:
		println("x.Orange")
	case x.Cherry:
		println("x.Cherry")
	default:
		println("default")
	}
}
