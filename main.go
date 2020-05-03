package main

import "fmt"

func swap(m1, m2 *int) {
	var temp int
	temp = *m2
	*m2 = *m1
	*m1 = temp
}

func testswap() {
	m1, m2 := 2,3
	fmt.Println(m1, m2)
	swap(&m1, &m2) // use the address since the function accepts pointers
	fmt.Println(m1, m2)
}

func declare() {
	// variables
	// var (
	// 	i1 int
	// 	i2 int32
	// 	i3 int64
	// 	u1 uint
	// 	u2 uint32
	// 	u3 uint64
	// 	arr1 []int
	// 	s1 string
	// )
	// arr1 = []int{1, 2, 3, 4}
	
	// alternate var declaration :=
	arr2 := []string{"my", "name", "is", "sample"}
	fmt.Println(arr2)

	// pointers
	v1 := 2
	ptr := &v1 // use & to get the address / pointer
	v2 := *ptr // use * to get the value in the address
	fmt.Println(v1, ptr, v2)
}

func main() {
	testswap()
}