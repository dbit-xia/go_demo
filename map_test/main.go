package main

import "fmt"

func main() {
	var a = make(map[string]string)
	func(a map[string]string) {
		a = map[string]string{"a": "a"}
		//a["b"] = "2";
		fmt.Println(a)
	}(a)
	fmt.Println(&a)
	fmt.Println(a)
}
