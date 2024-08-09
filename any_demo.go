package main

func AnyMain() {
	var a interface{}
	println(test1(a))
}

func test1(a any) bool {
	return a == nil
}
