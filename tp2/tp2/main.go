package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	aux "funciones"
	f "funciones/solicitud"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		procesarComandos(scanner.Text())
	}
}

func procesarComandos(input string) {
	partes := strings.Fields(input)

	if len(partes) == 0 {
		aux.ReportarError("Input vacio")
		return
	}

	switch partes[0] {
	case "agregar_archivo":

		agregarArchivo(partes)

	case "ver_visitantes":

		verVisitantes(partes)

	case "ver_mas_visitados":

		masVisitados(partes)

	default:
		aux.ReportarError(partes[0])
	}
}

func agregarArchivo(partes []string) {
	if !aux.ComandoLongitud2EsValido(partes) {
		aux.ReportarError(partes[0])
		return
	}

	archivo, err := aux.LeerArchivo(partes[1])

	if err != nil {
		aux.ReportarError(partes[0])
		return
	}

	aux.ProcesarComandoAgregarArchivo(archivo)

	aux.TodoOk()
}

func verVisitantes(partes []string) {
	if !aux.ComandoLongitud3EsValido(partes) {
		aux.ReportarError(partes[0])
		return
	}

	rangoA := aux.IpAInt64(partes[1])
	rangoB := aux.IpAInt64(partes[2])
	aux.ProcesarComandoVerVisitantes(f.Visitantes, rangoA, rangoB)

	aux.TodoOk()
}

func masVisitados(partes []string) {
	if !aux.ComandoLongitud2EsValido(partes) {
		aux.ReportarError(partes[0])
		return
	}

	n, _ := strconv.Atoi(partes[1])
	aux.ProcesarComandoMasVistados(f.SitiosMasVisitados, n)

	aux.TodoOk()
}
