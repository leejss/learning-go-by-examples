package main

import (
	"flag"
	"fmt"
)

// --name and --age 옵션을 사용하여 프로그램을 실행할 수 있도록 합니다.

// 구조체로 관리하여 인자를 파싱하고 출력할 수 있도록 합니다.
type Args struct {
	Name string
	Age  int
}

func (a *Args) Parse() {
	flag.StringVar(&a.Name, "name", "world", "name")
	flag.IntVar(&a.Age, "age", 20, "age")
	flag.Parse()
}

func (a *Args) Print() {
	fmt.Println(a.Name, a.Age)
}

func main() {

	// flag 변수 선언
	// var (
	// 	name string
	// 	age  int
	// )

	// flag 설정
	// flag.StringVar(&name, "name", "world", "name")
	// flag.IntVar(&age, "age", 20, "age")

	// flag parsing
	// flag.Parse()

	var args Args
	args.Parse()

	// print result
	args.Print()

}
