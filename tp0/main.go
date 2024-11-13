package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	f "tp0/ejercicios"
)

func main() {
	archivo1, err1 := os.Open("./archivo1.in")
	archivo2, err2 := os.Open("./archivo2.in")

	if err1 != nil {
		fmt.Println("Error %v al abrir el archivo %s", "./archivo1.in", err1)
		return
	} else if err2 != nil {
		fmt.Println("Error %v al abrir el archivo %s", "./archivo2.in", err2)
		return
	}

	defer archivo1.Close()
	defer archivo2.Close()
	array1 := make([]int, 0)
	array2 := make([]int, 0)
	s1 := bufio.NewScanner(archivo1)
	s2 := bufio.NewScanner(archivo2)

	for s1.Scan() {
		numero, error := strconv.Atoi(s1.Text())

		if error != nil {
			fmt.Printf("Error %v al leer el archivo %s\n", "./archivo1.in", error)
			return
		}
		array1 = append(array1, numero)

	}

	for s2.Scan() {
		numero, error := strconv.Atoi(s2.Text())

		if error != nil {
			fmt.Printf("Error %v al leer el archivo %s\n", "./archivo2.in", error)
			return
		}
		array2 = append(array2, numero)

	}

	resultado := f.Comparar(array1, array2)
	var mayorArray []int

	if resultado == 1 {
		mayorArray = array1
	} else if resultado == -1 {
		mayorArray = array2
	} else {
		mayorArray = array1
	}

	f.Seleccion(mayorArray)

	for _, valor := range mayorArray {
		fmt.Println(valor)
	}

}
