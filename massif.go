package main

import "fmt"

// задать масив в виде стринг
func main() {
	array := make([]int, 10)

	// array := []int{1, 2, 3}
	fmt.Println(len(array))

	// внесение в масив доп данных
	array = append(array, 321)
	array[2] = 12

	///сортировка масива
	//sort.Ints(array)

	fmt.Println(len(array))

	for _, i := range array {
	}

	// перебрать масив
	fmt.Println(array)

}
