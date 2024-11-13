from grafo import Grafo
from cola import Cola
from pila import Pila
import random


def bfs(grafo, origen, dest_posibles =[], dist_max =float("inf")):
    """
    Devuelve los diccionarios "padres" y "distancias" y el vertice que sea el destino final del recorrido
        - Si encuentra uno de los `dest_posibles`, el destino final sera el primero que encuentre
            - Si el destino encontrado es el origen, devuelve el padre del mismo
        - Si no encuentra ninguno, el destino final sera el `str` "destino_invalido" que tiene asociado
            el valor float("inf") en el diccionario "distancias"
    """
    padres = {}
    distancias = {}
    cola = Cola()

    dest_invalido = "destino_invalido"
    distancias[dest_invalido] = float("inf")

    padres[origen] = None
    distancias[origen] = 0
    cola.encolar(origen)
    
    while not cola.esta_vacia():
        v = cola.desencolar()
        if distancias[v] == dist_max:
            return padres, distancias, dest_invalido
        for w in grafo.adyacentes(v):
            if w == origen and origen in dest_posibles:
                return padres, distancias, v
            if w not in padres:
                padres[w] = v
                distancias[w] = distancias[v] + 1
                cola.encolar(w)
                if w in dest_posibles:
                    return padres, distancias, w
            
    return padres, distancias, dest_invalido


def page_rank(grafo, n, d =0.85):
    pr = {}
    vertices = grafo.vertices()
    ady_entrantes = obtener_adyacentes_entrantes(grafo)

    for v in vertices:
        pr[v] = (1-d) / len(vertices)

    for _ in range(n):
        random.shuffle(vertices)
        pr_aux = {}
        for w in vertices:
            pr_aux[w] = 0
            for v in ady_entrantes[w]:
                pr_aux[w] += d * pr[v] / grafo.cant_adyacentes(v) # si v es ady a w (v -> w) cant_adyacentes(v) nunca es 0
        for k in pr:
            pr[k] += pr_aux[k]

    return pr


def label_propagation(grafo, n =1):
    labels = {}
    vertices = grafo.vertices()
    ady_entrantes = obtener_adyacentes_entrantes(grafo)

    for v in vertices:
        labels[v] = v

    for _ in range(n):
        random.shuffle(vertices)
        for w in vertices:
            if not ady_entrantes[w]:
                continue
            frecuencias = {}
            for v in ady_entrantes[w]:
                frecuencias[labels[v]] = frecuencias.get(labels[v], 0) + 1 
            labels[w] = max(frecuencias, key=frecuencias.get)

    comunidades = {}
    for v, lab in labels.items():
        comunidades[lab] = comunidades.get(lab, set())
        comunidades[lab].add(v)

    return comunidades


def obtener_adyacentes_entrantes(grafo):
    dicc = {}

    for v in grafo.vertices():
        dicc[v] = set()

    for v in grafo.vertices():
        for w in grafo.adyacentes(v):
            # dicc[w] = dicc.get(w, set())
            dicc[w].add(v)

    return dicc


def dfs_cfc(grafo, v, visitados, orden, mas_bajo, pila, apilados, cfcs, contador_global):
    orden[v] = mas_bajo[v] = contador_global[0]
    contador_global[0] += 1
    visitados.add(v)
    pila.apilar(v)
    apilados.add(v)

    for w in grafo.adyacentes(v):
        if w not in visitados:
            dfs_cfc(grafo, w, visitados, orden, mas_bajo, pila, apilados, cfcs, contador_global)
        if w in apilados:
            mas_bajo[v] = min(mas_bajo[v], mas_bajo[w])

    if orden[v] == mas_bajo[v]:
        nueva_cfc = []
        while True:
            w = pila.desapilar()
            apilados.remove(w)
            nueva_cfc.append(w)
            if w == v:
                break
        cfcs.append(nueva_cfc)


def cfcs_grafo(grafo):
    resultados = []
    visitados = set()

    for v in grafo.vertices():
        if v not in visitados:
            dfs_cfc(grafo, v, visitados, {}, {}, Pila(), set(), resultados, [0])
    
    return resultados
