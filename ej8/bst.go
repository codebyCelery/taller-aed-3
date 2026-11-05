package main

import (
	"fmt"
)

type Nodo struct {
	Nota     int
	Cantidad int //se agrego aca cantidad
	Izq, Der *Nodo
}

func insertar(n *Nodo, v int) *Nodo {
	if n == nil {
		return &Nodo{Nota: v, Cantidad: 1}
	}

	if v < n.Nota {
		n.Izq = insertar(n.Izq, v)
	} else if v > n.Nota {
		n.Der = insertar(n.Der, v)
	} else {
		n.Cantidad++ //esta es la modificacion
	}

	return n
}
func cantidad(n *Nodo, v int) int {
	if n == nil {
		return 0
	}

	if v < n.Nota {
		return cantidad(n.Izq, v)
	} else if v > n.Nota {
		return cantidad(n.Der, v)
	}

	return n.Cantidad
}
func eliminar(n *Nodo, v int) *Nodo {
	if n == nil {
		return nil
	}

	if v < n.Nota {
		n.Izq = eliminar(n.Izq, v)
	} else if v > n.Nota {
		n.Der = eliminar(n.Der, v)
	} else {
		if n.Cantidad > 1 {
			n.Cantidad--
		} else {
			if n.Izq == nil {
				return n.Der
			}
			if n.Der == nil {
				return n.Izq
			}

			var sucesor *Nodo
			n.Der, sucesor = quitarMenor(n.Der)

			n.Nota = sucesor.Nota
			n.Cantidad = sucesor.Cantidad
		}
	}

	return n
}
func quitarMenor(n *Nodo) (*Nodo, *Nodo) { //funcion auxiliar
	if n.Izq == nil {
		return n.Der, n
	}

	var menor *Nodo
	n.Izq, menor = quitarMenor(n.Izq)

	return n, menor
}
func inorden(n *Nodo) {
	if n != nil {
		inorden(n.Izq)

		for i := 0; i < n.Cantidad; i++ {
			fmt.Print(n.Nota, " ")
		}

		inorden(n.Der)
	}
}
