package main

import "log"

type person struct {
	name string
	sex  string
	age  int
}

func (p *person) getName() string {
	return p.name
}
func (p *person) setName(name string) {
	p.name = name
	return
}

func main() {
	// type Person struct {
	// 	name string
	// 	age  int
	// }

	// func (p *Person) setName(name string) {
	// 	p.name = name
	// 	return
	// }
	// func (p *Person) getName() string {
	// 	return p.name
	// }

	type inface interface {
		setName(name string)
		getName() string
	}
	var p person
	p = person{name: "张三", age: 18}
	var i inface
	i = &p
	log.Println(p)
	log.Println(&p)
	log.Println(i == &p)

}
