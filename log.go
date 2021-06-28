package gg

import (
	"fmt"
	"reflect"
	"runtime"
)

func Log(itf ...interface{}) {
	// logs(itf...)
	for index, inter := range itf {
		if index != 0 {
			fmt.Printf(" ")
		}

		if inter == nil {
			fmt.Printf("%c[37;1m%s%c[0m", 0x1B, "nil", 0x1B) // white light
			continue
		}
		logs(reflect.ValueOf(inter))
	}
	fmt.Printf("\n")
}

func logs(inter reflect.Value) {

	kind := inter.Kind()

	// fmt.Println(109, kind)

	switch kind {
	case reflect.Bool:
		fmt.Printf("%c[33m%t%c[0m", 0x1B, inter, 0x1B) // yellow
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr: //reflect.Int <= kind &		fmt.Printf("%c[33m%d%c[0m", 0x1B, inter, 0x1B) // yellow
		fmt.Printf("%c[33m%d%c[0m", 0x1B, inter, 0x1B) // yellow
	case reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128: //reflect.Float32 <= kind &&
		fmt.Printf("%c[33m%f%c[0m", 0x1B, inter, 0x1B) // yellow
	case reflect.Array, reflect.Slice: //reflect.Float32 <= kind &&
		fmt.Printf("[ ")
		iValue := (inter)
		arrLen := iValue.Len()
		for i := 0; i < arrLen; i++ {
			if i != 0 {
				fmt.Printf(" ")
			}
			logs(iValue.Index(i))
		}
		fmt.Printf(" ]")
	case reflect.Chan, reflect.Ptr, reflect.UnsafePointer:
		fmt.Print(inter)
	case reflect.Func:
		funcName := runtime.FuncForPC((inter).Pointer()).Name()
		fmt.Printf("%c[36m%s%c[0m", 0x1B, "[Function: "+funcName+"]", 0x1B) // cyan

	case reflect.Map: //reflect.Float32 <= kind &&
		fmt.Printf("{ ")
		iValue := (inter)
		arrLen := iValue.Len()
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
			logs(childKey[i])
			fmt.Printf(": ")
			logs(childValue[i])
		}
		fmt.Printf(" }")

	case reflect.String:
		fmt.Printf("%c[32m%s%c[0m", 0x1B, `"`+inter.String()+`"`, 0x1B) // green

	case reflect.Struct:
		fmt.Printf("{ ")
		arrLen := (inter).NumField()
		// fmt.Println(166, arrLen)
		for i := 0; i < arrLen; i++ {
			if i != 0 {
				fmt.Printf(", ")
			}
			fmt.Printf("%s: ", inter.Type().Field(i).Name)
			logs(inter.Field(i))
		}
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
