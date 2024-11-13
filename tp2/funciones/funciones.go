package auxiliares

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	f "funciones/solicitud"

	TDAHeap "tdas/cola_prioridad"
	TDADiccionario "tdas/diccionario"
	TDALista "tdas/lista"
)

func ProcesarComandoAgregarArchivo(archivo *os.File) {
	solicitudesPorIp := TDADiccionario.CrearHash[int64, []f.Solicitud]()
	sospechosos := TDAHeap.CrearHeap(func(a, b int64) int { return int(b - a) })
	sospechososDetectados := TDADiccionario.CrearHash[int64, bool]()
	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		log := procesarLinea(scanner.Text())
		actualizarSolicitudesPorIp(solicitudesPorIp, log)
		actualizarVisitantes(log)
		actualizarSitiosMasVisitados(log)

		if !sospechososDetectados.Pertenece(log.IpInt) && detectarDoS(solicitudesPorIp.Obtener(log.IpInt)) {
			sospechosos.Encolar(log.IpInt)
			sospechososDetectados.Guardar(log.IpInt, true)
		}
	}

	if !sospechosos.EstaVacia() {
		reportarIpsSospechosas(sospechosos, solicitudesPorIp)
	}
}

func ProcesarComandoVerVisitantes(visitantes TDADiccionario.DiccionarioOrdenado[int64, string], a int64, b int64) {
	iter := visitantes.IteradorRango(&a, &b)
	i := 0

	for iter.HaySiguiente() {
		if i == 0 {
			fmt.Println("Visitantes:")
			i++
		}

		_, ip := iter.VerActual()
		fmt.Printf("\t%s\n", ip)
		iter.Siguiente()
	}
}

func ProcesarComandoMasVistados(sitiosMasVisitados TDADiccionario.Diccionario[string, int], n int) {
	sitiosPorVisitas, visitasOrdenadas := agruparSitiosPorVisitas(sitiosMasVisitados)

	if !visitasOrdenadas.EstaVacia() {
		fmt.Println("Sitios más visitados:")
	}

	i := 0
	for !visitasOrdenadas.EstaVacia() && i < n {
		cantidadVisitas := visitasOrdenadas.Desencolar()
		listaSitios := sitiosPorVisitas.Obtener(cantidadVisitas)

		for iter := listaSitios.Iterador(); iter.HaySiguiente() && i < n; iter.Siguiente() {
			fmt.Printf("\t%s - %d\n", iter.VerActual(), cantidadVisitas)
			i++
		}
	}
}

func procesarLinea(linea string) f.Solicitud {
	lineaSpliteada := strings.Fields(linea)
	ip := lineaSpliteada[0]
	IpInt := IpAInt64(ip)
	tiempoStr := lineaSpliteada[1][:]
	tiempo, _ := time.Parse("2006-01-02T15:04:05-07:00", tiempoStr)
	metodoHttp := lineaSpliteada[2]
	recurso := lineaSpliteada[3]

	return f.Solicitud{IpString: ip, IpInt: IpInt, FechaHora: tiempo, MetodoHttp: metodoHttp, UrlRecurso: recurso}
}

func actualizarSolicitudesPorIp(solicitudesPorIp TDADiccionario.Diccionario[int64, []f.Solicitud], log f.Solicitud) {
	if !solicitudesPorIp.Pertenece(log.IpInt) {
		solicitudesPorIp.Guardar(log.IpInt, []f.Solicitud{log})
	} else {
		logs := solicitudesPorIp.Obtener(log.IpInt)
		logs = append(logs, log)
		solicitudesPorIp.Guardar(log.IpInt, logs)
	}
}

func detectarDoS(solicitudes []f.Solicitud) bool {
	if len(solicitudes) < 5 {
		return false
	}

	for i := 0; i <= len(solicitudes)-5; i++ {
		if solicitudes[i+4].FechaHora.Sub(solicitudes[i].FechaHora).Seconds() < 2 {
			return true
		}
	}

	return false
}

func actualizarVisitantes(log f.Solicitud) {
	if !f.Visitantes.Pertenece(log.IpInt) {
		f.Visitantes.Guardar(log.IpInt, log.IpString)
	}
}

func actualizarSitiosMasVisitados(log f.Solicitud) {
	if !f.SitiosMasVisitados.Pertenece(log.UrlRecurso) {
		f.SitiosMasVisitados.Guardar(log.UrlRecurso, 1)
	} else {
		f.SitiosMasVisitados.Guardar(log.UrlRecurso, f.SitiosMasVisitados.Obtener(log.UrlRecurso)+1)
	}
}

func reportarIpsSospechosas(sospechosos TDAHeap.ColaPrioridad[int64], solicitudesPorIp TDADiccionario.Diccionario[int64, []f.Solicitud]) {
	ips := make([]string, 0, sospechosos.Cantidad())
	ipRevisadas := TDADiccionario.CrearHash[int64, bool]()

	for !sospechosos.EstaVacia() {
		ipInt := sospechosos.Desencolar()

		if !ipRevisadas.Pertenece(ipInt) {
			log := solicitudesPorIp.Obtener(ipInt)
			ips = append(ips, log[0].IpString)
			ipRevisadas.Guardar(ipInt, true)
		}
	}

	for _, ip := range ips {
		fmt.Printf("DoS: %s\n", ip)
	}
}

func agruparSitiosPorVisitas(sitiosMasVisitados TDADiccionario.Diccionario[string, int]) (TDADiccionario.Diccionario[int, TDALista.Lista[string]], TDAHeap.ColaPrioridad[int]) {
	sitiosPorVisitas := TDADiccionario.CrearHash[int, TDALista.Lista[string]]()
	visitasOrdenadas := TDAHeap.CrearHeap(func(a, b int) int { return a - b })

	for iter := sitiosMasVisitados.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		sitio, cantidadVisitas := iter.VerActual()

		if !sitiosPorVisitas.Pertenece(cantidadVisitas) {
			listaSitios := TDALista.CrearListaEnlazada[string]()
			listaSitios.InsertarUltimo(sitio)
			sitiosPorVisitas.Guardar(cantidadVisitas, listaSitios)
		} else {
			listaSitios := sitiosPorVisitas.Obtener(cantidadVisitas)
			listaSitios.InsertarUltimo(sitio)
		}

		visitasOrdenadas.Encolar(cantidadVisitas)
	}

	return sitiosPorVisitas, visitasOrdenadas
}

func LeerArchivo(nombreArchivo string) (*os.File, error) {
	archivo, err := os.Open(nombreArchivo)

	if err != nil {
		return nil, err
	}

	return archivo, nil
}

func ReportarError(comando string) {
	fmt.Fprintf(os.Stderr, "Error en comando %s\n", comando)
}

func TodoOk() {
	fmt.Println("OK")
}

func ComandoLongitud2EsValido(entrada []string) bool {
	return len(entrada) == 2
}

func ComandoLongitud3EsValido(entrada []string) bool {
	if len(entrada) != 3 || entrada[0] != "ver_visitantes" || !esIPValida(entrada[1]) || !esIPValida(entrada[2]) {
		return false
	}

	return true
}

func IpAInt64(ipStr string) int64 {
	ipSpliteada := strings.Split(ipStr, ".")
	var IpInt int64

	for i := 0; i < len(ipSpliteada); i++ {
		ip, _ := strconv.Atoi(ipSpliteada[i])
		IpInt = IpInt<<8 + int64(ip)
	}

	return IpInt
}

func esIPValida(ip string) bool {
	return net.ParseIP(ip) != nil
}
