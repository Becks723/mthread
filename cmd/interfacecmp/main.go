package main

import (
	"fmt"
)

type MyImpl struct {
	Name string
}

func main() {
	// nil
	var nili1 interface{} = nil
	var nili2 interface{} = nil
	fmt.Println("- nil(non-typed) == nil(non-typed): ", nili1 == nili2) // true

	// typed nil
	var typedi *int = nil
	nili2 = typedi
	fmt.Println("- nil(non-typed) == nil(*int): ", nili1 == nili2) // false

	// same runtime type, same value
	var myImpl1 interface{} = MyImpl{"aaa"}
	var myImpl2 interface{} = MyImpl{"aaa"}
	fmt.Println("- abc(MyImpl) == abc(MyImpl): ", myImpl1 == myImpl2) // true

	// same runtime type, different values
	var myImplptr1 interface{} = &MyImpl{"aaa"}
	var myImplptr2 interface{} = &MyImpl{"aaa"}
	fmt.Println("- abc(*MyImpl) == abc(*MyImpl): ", myImplptr1 == myImplptr2) // false

	// uncomparable type
	var slice1 interface{} = []int{1, 2, 3}
	var slice2 interface{} = []int{1, 2, 3}
	fmt.Println("- [1 2 3]([]int) == [1 2 3]([]int): ", slice1 == slice2) // panic
}
