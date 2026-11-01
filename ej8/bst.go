// EJERCICIO 8 · Responsable: Leticia
package main

// Nodo del enunciado (todavía sin Cantidad).
type Nodo struct {
	Nota     int
	Izq, Der *Nodo
}

// insertar del enunciado: ignora los valores repetidos.
func insertar(n *Nodo, v int) *Nodo {
	if n == nil {
		return &Nodo{Nota: v}
	}
	if v < n.Nota {
		n.Izq = insertar(n.Izq, v)
	} else if v > n.Nota {
		n.Der = insertar(n.Der, v)
	}
	return n
}

// eliminar de S9 (elimina el nodo completo).
func eliminar(n *Nodo, v int) *Nodo {
	if n == nil {
		return nil
	}
	if v < n.Nota {
		n.Izq = eliminar(n.Izq, v)
	} else if v > n.Nota {
		n.Der = eliminar(n.Der, v)
	} else {
		if n.Izq == nil {
			return n.Der
		}
		if n.Der == nil {
			return n.Izq
		}
		m := n.Der
		for m.Izq != nil {
			m = m.Izq
		}
		n.Nota = m.Nota
		n.Der = eliminar(n.Der, m.Nota)
	}
	return n
}

// inorden de S9 (imprime cada nota una sola vez).
func inorden(n *Nodo, salida *[]int) {
	if n == nil {
		return
	}
	inorden(n.Izq, salida)
	*salida = append(*salida, n.Nota)
	inorden(n.Der, salida)
}

// ======================= AQUÍ ESCRIBES =======================
// Modifica ESTE archivo para aceptar notas repetidas:
//   1. Agrega el campo Cantidad al struct Nodo.
//   2. insertar: si el nodo es nuevo, Cantidad = 1; si la nota ya existe,
//      aumenta Cantidad en lugar de ignorarla.
//   3. eliminar: si Cantidad es mayor que 1, solo disminúyela; el nodo
//      desaparece únicamente cuando llega a cero. Piensa con cuidado
//      el caso de dos hijos: qué cantidad hereda el nodo del sucesor y
//      cómo se borra el sucesor de su lugar original (mira tu P8.PP).
//   4. inorden: repite cada nota tantas veces como indique Cantidad.
// Comenta con // cada cambio.
// =============================================================
