package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	f "operaciones/op"
	pila "tdas/pila"
)

const (
	baseDecimal  = 10
	cantidadBits = 64
)

func main() {
	resultados := procesarEntrada()
	for _, elemento := range resultados {
		fmt.Printf("%s\n", elemento)
	}
}

func procesarEntrada() []string {
	var resultados []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		linea := scanner.Text()
		elementos := strings.Fields(linea)
		resultado, err := calculadoraPolaca(elementos)
		if err != "" {
			resultados = append(resultados, err)
		} else {
			resultados = append(resultados, strconv.FormatInt(resultado, baseDecimal))
		}
	}
	return resultados
}

func calculadoraPolaca(elementos []string) (int64, string) {
	aux := pila.CrearPilaDinamica[int64]()
	for _, token := range elementos {
		num, err := strconv.ParseInt(token, baseDecimal, cantidadBits)
		if err == nil {
			aux.Apilar(num)
		} else {
			var res int64
			var resError string
			switch token {
			case "sqrt":
				res, resError = operarUnNumero(aux)
			case "+", "-", "*", "/", "^", "log":
				res, resError = operarDosNumeros(aux, token)
			case "?":
				res, resError = operarTresNumeros(aux)
			default:
				return f.RetornarError()
			}
			if resError != "" {
				return f.RetornarError()
			}
			aux.Apilar(res)
		}
	}
	if aux.EstaVacia() {
		return f.RetornarError()
	}
	resultado := aux.Desapilar()
	if !aux.EstaVacia() {
		return f.RetornarError()
	}
	return resultado, ""
}

func operarUnNumero(numeros pila.Pila[int64]) (int64, string) {
	if numeros.EstaVacia() {
		return f.RetornarError()
	}
	a := numeros.Desapilar()
	return f.Raiz(a)
}

func operarDosNumeros(numeros pila.Pila[int64], operador string) (int64, string) {
	if numeros.EstaVacia() {
		return f.RetornarError()
	}
	b := numeros.Desapilar()
	if numeros.EstaVacia() {
		return f.RetornarError()
	}
	a := numeros.Desapilar()
	var resultado int64
	var err string
	switch operador {
	case "+":
		resultado = f.Suma(a, b)
	case "-":
		resultado = f.Resta(a, b)
	case "*":
		resultado = f.Producto(a, b)
	case "^":
		resultado, err = f.Exponencial(a, b)
	case "/":
		resultado, err = f.Division(a, b)
	case "log":
		resultado, err = f.Logaritmo(a, b)
	default:
		return f.RetornarError()
	}
	return resultado, err
}

func operarTresNumeros(numeros pila.Pila[int64]) (int64, string) {
	if numeros.EstaVacia() {
		return f.RetornarError()
	}
	c := numeros.Desapilar()
	if numeros.EstaVacia() {
		return f.RetornarError()
	}
	b := numeros.Desapilar()
	if numeros.EstaVacia() {
		return f.RetornarError()
	}
	a := numeros.Desapilar()
	if a != 0 {
		return b, ""
	} else {
		return c, ""
	}
}
