package solicitud

import (
	"time"

	TDADiccionario "tdas/diccionario"
)

var (
	Visitantes         = TDADiccionario.CrearABB[int64, string](func(a, b int64) int { return int(a - b) })
	SitiosMasVisitados = TDADiccionario.CrearHash[string, int]()
)

type Solicitud struct {
	IpString   string
	IpInt      int64
	FechaHora  time.Time
	MetodoHttp string
	UrlRecurso string
}
