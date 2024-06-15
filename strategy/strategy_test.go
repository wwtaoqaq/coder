package strategy

import (
	"fmt"
	"log"
	_ "strings"
	"testing"
)

func TestStrategy(t *testing.T) {
	req := &UploadOperationRequest{HandleType: 2, Data: "Hello"}
	context, err := NewTaskContext(req)
	if err != nil {
		log.Fatal("err is noy null")
	}
	context.Update()
}

func TestSlice(t *testing.T) {
	var list1 = []int{1, 2, 3, 4, 5}
	list2 := list1[1:2]
	log.Println("len:", len(list2), "cap:", cap(list2))
}

func TestMap(t *testing.T) {
	var myMap map[int8]string
	myMap[1] = "hello"
	log.Println(myMap)
	t.Errorf("test map err")
}

func TestChan(t *testing.T) {
	c := make(chan int, 2)
	c <- 10
	print("data from chan: ", <-c)
}
func TestNilSlice(t *testing.T) {
	var list = []int{1}
	arr := [1]int{12}
	fmt.Printf("list type:%T, arr type : %T", list, arr)

}
