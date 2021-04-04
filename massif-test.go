package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	/*--  Пустой срез, который наполнить элементами из просмотренной папки. Наполнить срез в цикле. В конце вывести срез и длину среза --*/

	slice := []string{}

	files, err := ioutil.ReadDir("/home/anton/")

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		slice = append(slice, file.Name())
	}

	fmt.Println(slice, len(slice))

	/*--  Массив, длинной равной количеству элементов в папке, после чего заполнить массив элементами, которые . находятся в папке --*/

	array := make([]string, len(files))

	for k, v := range files {
		array[k] = v.Name()
	}

	fmt.Println(array, len(array))

}
