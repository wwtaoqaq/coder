package main

import "fmt"

type Obj struct {
	field1 []int64
}

func main() {
	obj := &Obj{}
	obj.field1 = make([]int64, 0, 2)
	fmt.Println("obj.list before:", obj.field1)
	fmt.Printf("obj.list before addr: %p\n", obj.field1)
	list := obj.field1
	list = append(list, 1, 2)
	fmt.Printf("list addr: %p\n", list)
	fmt.Printf("field addr: %p\n", obj.field1)
	list = append(list, 3)
	fmt.Printf("list addr: %p\n", list)
	fmt.Printf("field1 addr: %p\n", obj.field1)

	//for i := 1; i < 20; i++ {
	//	list = append(list, int64(i))
	//	fmt.Printf("list addr: %p\n", list)
	//}
	//fmt.Println("obj.list after:", obj.field1)
	//fmt.Printf("obj.list after addr: %p\n", obj.field1)
	//fmt.Printf("list after addr: %p\n", list)
	//fmt.Println("list:", list)
	//test(list)
	//fmt.Println("list:", list)

}

func test(list []int64) {
	for i := 1; i < 200; i++ {
		list = append(list, int64(i))
	}
}
