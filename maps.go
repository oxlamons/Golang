package main

import "fmt"

func main() {

	//пустой срез стринг
	keys := []string{}

	// обявление карти с масивом 1 срез 1 и 2

	carta := map[string][]int{

		"1": {1, 2},
		"2": {2, 10},
	}

	// срез карти и ввиод ключа перебор списка

	for i, key := range carta {

		//	в срез добавляються ключи

		// key = append(keys, key)

		// в ключь i записалось значение 1 2
		fmt.Println(i, key)
	}

	//вівод всех ключей
	fmt.Println(keys)
}
