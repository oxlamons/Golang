package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {

	message := []byte("Hello, World!")

	carta := map[string][]string{
		"Documents": {"Test", "Terr"},
		"Downloads": {"aws", "Books"},
	}

	for keys, values := range carta {

		// Создаем директории на основе ключей
		if _, err := os.Stat(keys); os.IsNotExist(err) {
			os.Mkdir(keys, 0755)
			fmt.Println(keys, "Directory Created")
		} else {
			fmt.Println(keys, "Directory already exists")
		}

		// Создаем файлы в директориях

		for _, files := range values {
			path := filepath.Join(keys, files)
			err := ioutil.WriteFile(path, message, 0644)
			if err != nil {
				log.Fatal(err)
			}

		}

	}

}
