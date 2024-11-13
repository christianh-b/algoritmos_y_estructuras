class Grafo:

    def __init__(self):
        self.__vertices = {}

    def existe_vertice(self, v):
        if v in self.__vertices:
            return True
        return False

    def agregar_vertice(self, v):
        if not self.existe_vertice(v):
            self.__vertices[v] = {}

    def borrar_vertice(self, v):
        if self.existe_vertice(v):
            del self.__vertices[v]

    def existe_arista(self, v, w):
        if self.existe_vertice(v):
            if w in self.__vertices[v]:
                return True
        return False

    def agregar_arista(self, v, w):
        if self.existe_vertice(v) and self.existe_vertice(w):
            self.__vertices[v][w] = 1
    
    def borrar_arista(self, v, w):
        if self.existe_arista(v, w):
            del self.__vertices[v][w]

    def vertices(self):
        return list(self.__vertices.keys())

    def adyacentes(self, v):
        if not self.existe_vertice(v):
            raise Exception("El vertice no existe")

        return list(self.__vertices[v].keys())
    
    def cant_adyacentes(self, v):
        if not self.existe_vertice(v):
            raise Exception("El vertice no existe")
        
        return len(self.__vertices[v])