package main

import "fmt"

type set struct {
	key, data int
}

var array = new([3]set)

func insert(slice *[]int) {

	*slice = append(*slice, 3)
	h(slice)
}

func h(slice *[]int) {

	*slice = append(*slice, 5)

	fmt.Println(slice)
}

func main() {

	for i := 0; i < 3; i++ {

		array[i].key = 0
		array[i].data = 0
	}

	fmt.Println(array)

	slice := []int{}

	insert(&slice)

	fmt.Println(slice)

}
