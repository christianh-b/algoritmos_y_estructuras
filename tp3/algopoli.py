#!/usr/bin/python3

import biblioteca as bib
from grafo import Grafo
import sys
import heapq

ITS_PAGE_RANK = 15
ITS_LBL_PROP = 5

def main():
    ruta_archivo = sys.argv[1]
    grafo = crear_grafo(ruta_archivo)
    pr = {}

    for linea in sys.stdin:
        args = linea.rstrip().split(" ")
        comando = args[0]
        parametros = args[1:]

        if comando == "min_seguimientos":
            origen, destino = int(parametros[0]), int(parametros[1])
            minimos_seguimientos(grafo, origen, destino)

        elif comando == "mas_imp":
            cant = int(parametros[0])
            if pr == {}:
                pr = bib.page_rank(grafo, ITS_PAGE_RANK)
            print(str(top_k(pr, cant)).strip("[]"))

        elif comando == "persecucion":
            delincuentes, k_cant = [int(p) for p in parametros[0].split(",")], int(parametros[1])
            if pr == {}:
                pr = bib.page_rank(grafo, ITS_PAGE_RANK)
            persecucion(grafo, delincuentes, top_k(pr, k_cant))

        elif comando == "comunidades":
            min_miembros = int(parametros[0])
            comunidades = bib.label_propagation(grafo, ITS_LBL_PROP)
            imprimir_comunidades(comunidades, min_miembros)

        elif comando == "divulgar":
            delincuente, dist_max = int(parametros[0]), int(parametros[1])
            divulgacion(grafo, delincuente, dist_max)            

        elif comando == "divulgar_ciclo":
            origen_ciclo = int(parametros[0])
            ciclo_mas_corto(grafo, origen_ciclo)

        elif comando == "cfc":
            cfcs = bib.cfcs_grafo(grafo)
            imprimir_cfcs(cfcs)

        else:
            print("Comando invalido")
        


def crear_grafo(ruta_archivo):
    grafo = Grafo()
    with open(ruta_archivo) as datos:
        for linea in datos:
            campos = linea.rstrip().split("\t")
            v, w = int(campos[0]), int(campos[1])
            grafo.agregar_vertice(v)
            grafo.agregar_vertice(w)
            grafo.agregar_arista(v, w)

    return grafo


def minimos_seguimientos(grafo, origen, destino):
    padres, distancias, dest_final = bib.bfs(grafo, origen, [destino])
    
    if distancias[dest_final] == float("inf"):
        print("Seguimiento imposible")
    else:
        imprimir_recorrido(grafo, padres, dest_final, [destino])


def persecucion(grafo, delincuentes, top_k):
    p_actual, d_actual, df_actual = bib.bfs(grafo, delincuentes[0], top_k)
    for delincuente in delincuentes[1:]:
        p_nuevo, d_nuevo, df_nuevo = bib.bfs(grafo, delincuente, top_k)
        if d_nuevo[df_nuevo] == float("inf"):
            continue
        elif d_nuevo[df_nuevo] < d_actual[df_actual]:
            p_actual, d_actual, df_actual = p_nuevo, d_nuevo, df_nuevo
        elif d_nuevo[df_nuevo] == d_actual[df_actual] and top_k.index(df_nuevo) < top_k.index(df_actual):
            p_actual, d_actual, df_actual = p_nuevo, d_nuevo, df_nuevo

    if d_actual[df_actual] == float("inf"):
        print("Persecucion Imposible")
    else:
        imprimir_recorrido(grafo, p_actual, df_actual, top_k)


def divulgacion(grafo, delincuente, dist_max):
    padres, _, _ = bib.bfs(grafo, delincuente, dist_max= dist_max)
    del padres[delincuente]

    print(str(list(padres.keys())).strip("[]"))


def ciclo_mas_corto(grafo, origen_ciclo):
    padres, distancias, dest_final = bib.bfs(grafo, origen_ciclo, [origen_ciclo])

    if distancias[dest_final] == float("inf"):
        print("No se encontro recorrido")
    else:
        imprimir_recorrido(grafo, padres, dest_final, [origen_ciclo])


class ParKVComparable:
    def __init__(self, clave, valor):
        self.clave = clave
        self.valor = valor

    def __lt__(self, otro):
        return self.valor < otro.valor
    
    def __gt__(self, otro):
        return self.valor > otro.valor
    
def top_k(dicc, k_cant):
    heap_min = []
    resultado = []

    for k, v in dicc.items():
        heapq.heappush(heap_min, ParKVComparable(k, v))
        if len(heap_min) > k_cant:
            heapq.heappop(heap_min)

    for par_kv in sorted(heap_min, reverse=True):
        resultado.append(par_kv.clave)

    return resultado


def imprimir_recorrido(grafo, padres, dest_final, dest_posibles):
    resultado = []

    actual = dest_final
    while padres[actual] != None:
        resultado.append(str(actual))
        actual = padres[actual]   
    resultado.append(str(actual))
    
    resultado = resultado[::-1]
    if actual in dest_posibles and grafo.existe_arista(dest_final, actual): # aca actual == origen
        resultado.append(str(actual))
    
    print(" -> ".join(resultado))


def imprimir_comunidades(comunidades, min_miembros):
    for com, miembros in comunidades.items():
        if len(miembros) >= min_miembros:
            miembros = str(miembros).strip("{}")
            print(f"Comunidad {com}: {miembros}")


def imprimir_cfcs(cfcs):
    for i, cfc in enumerate(cfcs):
        cfc = str(cfc).strip("[]")
        print(f"CFC {i+1}: {cfc}")


main()