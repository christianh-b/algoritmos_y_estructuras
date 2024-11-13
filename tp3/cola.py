from collections import deque

class Cola:

    def __init__(self):
        self.__deque = deque()

    def esta_vacia(self):
        return len(self.__deque) == 0

    def ver_primero(self):
        if self.esta_vacia():
            raise Exception("La cola esta vacia")
        
        return self.__deque[0]

    def encolar(self, elem):
        self.__deque.append(elem)

    def desencolar(self):
        if self.esta_vacia():
            raise Exception("La cola esta vacia")
        
        return self.__deque.popleft()