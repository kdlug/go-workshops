package main

// https://segmentfault.com/a/1190000040841232/en
import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	str1 := "hello"
	str2 := "hello"
	//  the same string values are sharing the same backing arrays → efficiency
	fmt.Printf("str1: %q, ptr: %d \n", str1, (*reflect.StringHeader)(unsafe.Pointer(&str1)).Data)
	fmt.Printf("str2: %q, ptr: %d \n", str2, (*reflect.StringHeader)(unsafe.Pointer(&str2)).Data)

	// string slicing doesn’t allocate a new backing array
	s := str1[1:]
	fmt.Printf("s: %q, ptr: %d \n", s, (*reflect.StringHeader)(unsafe.Pointer(&s)).Data)

}
