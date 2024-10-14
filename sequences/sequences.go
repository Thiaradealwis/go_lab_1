package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) []int {
	slice = append(slice, slice...)
	return slice
}

func mapSlice(f func(a int) int, slice []int) {
	for i, s := range slice {
		s = f(s)
		slice[i] = s
	}
	//fmt.Println(slice)
}

func mapArray(f func(a int) int, array [3]int) {
	for i, a := range array {
		a = f(a)
		array[i] = a
	}
	fmt.Println(array)

}

func main() {
	intsSlice := []int{1, 2, 3}
	mapSlice(addOne, intsSlice)
	fmt.Println(intsSlice)

	intsArray := [3]int{1, 2, 3}
	mapArray(addOne, intsArray)

	newSlice := intsSlice[1:3]
	mapSlice(square, newSlice)
	fmt.Println(newSlice)

	fmt.Println(double(intsSlice))

}
