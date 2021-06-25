package gg

import (
	"fmt"
	"reflect"
	"runtime"
)

func Log(itf ...interface{}) {
	logs(itf...)
	fmt.Printf("\n")
}

func logs(itf ...interface{}) {

	for index, inter := range itf {
		if index != 0 {
			fmt.Printf(" ")
		}

		if inter == nil {
			fmt.Printf("%c[37;1m%s%c[0m", 0x1B, "nil", 0x1B) // white light
			continue
		}
		logs2(reflect.ValueOf(inter))
	}
}

func logs2(inter reflect.Value) {

	kind := inter.Kind()

	// fmt.Println(109, kind)

	switch {
	case kind == reflect.Bool:
		fmt.Printf("%c[33m%t%c[0m", 0x1B, inter, 0x1B) // yellow
	case kind <= reflect.Uintptr: //reflect.Int <= kind &&
		fmt.Printf("%c[33m%d%c[0m", 0x1B, inter, 0x1B) // yellow
	case kind <= reflect.Complex128: //reflect.Float32 <= kind &&
		fmt.Printf("%c[33m%f%c[0m", 0x1B, inter, 0x1B) // yellow
	case kind == reflect.Array || kind == reflect.Slice: //reflect.Float32 <= kind &&
		fmt.Printf("[ ")
		iValue := (inter)
		arrLen := iValue.Len()
		for i := 0; i < arrLen; i++ {
			if i != 0 {
				fmt.Printf(" ")
			}
			logs2(iValue.Index(i))
		}
		fmt.Printf(" ]")
	case kind == reflect.Chan || kind == reflect.Ptr || kind == reflect.UnsafePointer:
		// fmt.Printf("%p", inter)
		fmt.Print(inter)
	case kind == reflect.Func:
		funcName := runtime.FuncForPC((inter).Pointer()).Name()
		fmt.Printf("%c[36m%s%c[0m", 0x1B, "[Function: "+funcName+"]", 0x1B) // cyan

	case kind == reflect.Map: //reflect.Float32 <= kind &&

		// fmt.Println(24, (i).Elem())
		// fmt.Println(25, (i).Len())
		fmt.Printf("{ ")
		iValue := (inter)
		arrLen := iValue.Len()
		// n := iValue.Len()
		childKey := make([]reflect.Value, 0, arrLen)
		childValue := make([]reflect.Value, 0, arrLen)
		iter := iValue.MapRange()
		for iter.Next() {
			childKey = append(childKey, iter.Key())
			childValue = append(childValue, iter.Value())
		}
		for i, _ := range childKey {
			if i != 0 {
				fmt.Printf(", ")
			}
			logs2(childKey[i])
			fmt.Printf(": ")
			logs2(childValue[i])
		}
		fmt.Printf(" }")

	case kind == reflect.String:
		fmt.Printf("%c[32m%s%c[0m", 0x1B, `"`+inter.String()+`"`, 0x1B) // green

	case kind <= reflect.Struct:
		// fmt.Println(163, (inter).Field(1).Kind())
		fmt.Printf("{ ")
		arrLen := (inter).NumField()
		// fmt.Println(166, arrLen)
		for i := 0; i < arrLen; i++ {
			if i != 0 {
				fmt.Printf(", ")
			}
			// logs2(childKey[i])
			// fmt.Println(172, (inter).Field(i))
			fmt.Printf("%s: ", inter.Type().Field(i).Name)
			// fmt.Println(174, inter)
			logs2(inter.Field(i))
			// fmt.Println(176)
		}
		// fmt.Println(178)
		fmt.Printf(" }")
	default:
		fmt.Print(inter)
	}
	// }
}
func testIArr(iArr []interface{}) {
	fmt.Println(59, iArr)

}

// 前景 背景 颜色
// ---------------------------------------
// 30 40 黑色
// 31 41 红色
// 32 42 绿色
// 33 43 黄色
// 34 44 蓝色
// 35 45 紫红色
// 36 46 青蓝色
// 37 47 白色

// 0 终端默认设置
// 1 高亮显示
// 4 使用下划线
// 5 闪烁
// 7 反白显示
// 8 不可见
// %c[{前景色（文字颜色）};{背景色}{;前景色2，可选}m
// fmt.Printf("%c[32;32;1m%s%c[0m\n\n", 0x1B, "testPrintColor", 0x1B)
func logNil() {

	fmt.Printf("%c[37;1m%s%c[0m", 0x1B, "nil", 0x1B)
}
func logFunc(arg interface{}) {
	funcName := runtime.FuncForPC(reflect.ValueOf(arg).Pointer()).Name()
	ret := "[Function: " + funcName + "]"
	fmt.Printf("%c[36m%s%c[0m", 0x1B, ret, 0x1B)
}
func logBool(arg bool) {
	fmt.Printf("%c[33m%t%c[0m", 0x1B, arg, 0x1B)
}
func logInt(arg int) {
	fmt.Printf("%c[33m%d%c[0m", 0x1B, arg, 0x1B)
}
func logStr(arg string) {
	fmt.Printf("%c[32m%s%c[0m", 0x1B, `"`+arg+`"`, 0x1B)
}
func logArrInt(arg []int) {
	fmt.Printf("%s", `[ `)
	for index, item := range arg {
		// logs(key)
		if index != 0 {
			fmt.Printf("%s", `, `)
		}

		logs(item)
	}
	fmt.Printf("%s", ` ]`)
}
func logArrStr(arg []string) {
	fmt.Printf("%s", `[ `)
	for index, item := range arg {
		// logs(key)
		if index != 0 {
			fmt.Printf("%s", `, `)
		}

		logs(item)
	}
	fmt.Printf("%s", ` ]`)
}
func logMapStrStr(arg map[string]string) {
	fmt.Printf("%s", `{ `)
	isFirst := true
	for key, item := range arg {

		if !isFirst {

			fmt.Printf("%s", `, `)
		}
		isFirst = false
		logs(key)
		fmt.Printf("%s", `: `)
		logs(item)
	}
	fmt.Printf("%s", ` }`)
}
func logMapStrInt(arg map[string]int) {
	fmt.Printf("%s", `{ `)
	isFirst := true
	for key, item := range arg {

		if !isFirst {

			fmt.Printf("%s", `, `)
		}
		isFirst = false
		logs(key)
		fmt.Printf("%s", `: `)
		logs(item)
	}
	fmt.Printf("%s", ` }`)
}
func logMapIntStr(arg map[int]string) {
	fmt.Printf("%s", `{ `)
	isFirst := true
	for key, item := range arg {

		if !isFirst {

			fmt.Printf("%s", `, `)
		}
		isFirst = false
		logs(key)
		fmt.Printf("%s", `: `)
		logs(item)
	}
	fmt.Printf("%s", ` }`)
}
