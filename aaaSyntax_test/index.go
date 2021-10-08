package main

import (
	"fmt"
	"strconv"
)

func test(x float32, y string) (float32, string) {
	return x, y
}

func main() {
	// var bbb string = "123"
	// fmt.Println(bbb.type)
	var a [10]int
	for index := 0; index < 5; index++ {
		a[index] = index
	}
	fmt.Println(a[1:3])
	fmt.Println(string([]byte("夏天")))
	// a, b := test(6.5, "Runoob")
	// fmt.Println(fmt.Sprint(a), b, strconv.FormatFloat(123456789.123456789, 'f', -1, 64))

	// testToString()
	testPointer()
	// testStruct()
	testInterface()
}

func toString(x float64) string {
	//float64转string
	return strconv.FormatFloat(x, 'f', -1, 64)
}

func testToString() {
	fmt.Println(toString(1123456789.123456789))
}

func testPointer() {
	var p *int
	var p2 *int
	var I int
	I = 1
	p = &I //取I的地址
	p2 = p
	*p = 8

	fmt.Println("I =", I)
	fmt.Println("*&I =", *&I)
	fmt.Println("p =", p)
	fmt.Println("p2 =", p2)
	fmt.Println("*p =", *p)
	fmt.Println("*p2 =", *p2)
	fmt.Println("&p =", &p)
	fmt.Println("&p2 =", &p2)

	I = 2
	fmt.Println("*p", *p)

}

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

func testStruct() {

	var p1 person
	p1.name = "张一"

	p1.sex = "男"
	p1.age = 18
	// fmt.Println("getPerson", getPerson(p1))
	fmt.Println("p1", p1)
	// fmt.Println("getPersonPointer", *getPersonPointer(&p1))
	// fmt.Println("p1", p1)
	fmt.Println("p1.getName()", p1.getName())
	p1.setName("张二")
	fmt.Println("p1.getName()", p1.getName())

}

func getPerson(n1 person) person {
	n1.name = "张三丰"
	return n1
}
func getPersonPointer(n1 *person) *person {
	n1.name = "李四"
	return n1
}

type iPerson interface {
	getName() string
	setName(name string)
}

func testInterface() {

	var p1 person
	p1 = person{name: "你好"}
	// p1.name = "张三"
	// p1.sex = "男"
	// p1.age = 18
	// var ip1 iPerson
	// ip1 = &p1
	// ip1.setName("李四")
	fmt.Println("getIPerson", getIPerson(&p1))
	fmt.Println("ip1", p1)
	// fmt.Println("getPersonPointer", *getPersonPointer(&p1))
	// fmt.Println("p1", p1)
	// fmt.Println("p1.getName()", p1.getName())
}

func getIPerson(n1 interface{ setName(name string) }) interface{} {
	n1.setName("张三丰")
	return n1
}

// // func getIPersonPointer(n1 *iPerson) *iPerson {
// // 	n1.setName("李四")
// // 	return n1
// // }
