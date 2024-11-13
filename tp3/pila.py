from collections import deque

class Pila:

    def __init__(self):
        self.__deque = deque()

    def esta_vacia(self):
        return len(self.__deque) == 0

    def ver_tope(self):
        if self.esta_vacia():
            raise Exception("La pila esta vacia")
        
        return self.__deque[len(self.__deque)-1]

    def apilar(self, elem):
        self.__deque.append(elem)

    def desapilar(self):
        if self.esta_vacia():
            raise Exception("La pila esta vacia")
        
        return self.__deque.pop()