package main

import (
	"github.com/1stgg/gg"
)

type Page struct {
	num     int
	content string
}

type Books struct {
	title string
	id    int
	page  Page
}

func main() {
	// gg.Log(nil, false)
	gg.Log(
		1,
		int8(2),
		int16(3),
		int32(4),
		int64(5),
		uint(6),
		uint8(6),
		uint16(6),
		uint32(6),
		uint64(6),
		uintptr(6),
	)
	gg.Log(
		float32(6),
		float64(6),
		complex(1, 2),
		complex64(1),
		complex128(1),
	)
	// fmt.Println(Animal)

	// var a interface{}
	// a = "6"
	book := Books{"标题", 1, Page{1, "content"}}

	gg.Log(
		// []bool{true, false},
		[2]string{"1", "false"},
		make(chan int),
		func() {},
		// a,
		// map[string]string{"a": "b", "c": "d"},
		// map[int]string{1: "b"},
		// map[string]int{"b": 1},
		book,
		// [2]int{1, 2},
		// map[string]int{"a": 1, "b": 2},
		// make(map[string]int),
		// []int{1, 2, 3},
		// []float32{1, 2, 3},
	)

	// for i := 0; i < 10; i++ {
	// 	switch i {
	// 	case 1:
	// 		fallthrough
	// 	case 2:
	// 		fmt.Println(1, 2)
	// 	case 3:
	// 		fmt.Println(18)
	// 	default:
	// 		fmt.Println(i)
	// 	}
	// }

}
