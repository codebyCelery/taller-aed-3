// EJERCICIO 7 · Responsable: Celery
package main

// Navegador del enunciado. Funciona, pero guarda TODO el historial.
type Navegador struct {
	atras    Pila[string] // el tope es la página actual
	adelante Pila[string]
}

func (n *Navegador) Visitar(url string) {
	n.atras.Push(url)
	n.adelante = Pila[string]{}
}

func (n *Navegador) Atras() {
	if len(n.atras.items) > 1 {
		url, _ := n.atras.Pop()
		n.adelante.Push(url)
	}
}

func (n *Navegador) Adelante() {
	if url, err := n.adelante.Pop(); err == nil {
		n.atras.Push(url)
	}
}

func (n *Navegador) Actual() string {
	return n.atras.items[len(n.atras.items)-1]
}

// ======================= AQUÍ ESCRIBES =======================
// Modifica este archivo para cumplir el requisito nuevo: atras guarda
// como máximo las 5 últimas páginas (contando la actual). Al visitar
// una sexta, se descarta la más antigua.
//   1. Cambia el tipo de atras por la estructura que respondiste en P7.E
//      (escríbela en un archivo nuevo, por ejemplo deque.go).
//   2. Cambia Visitar, Atras, Adelante y Actual para usarla.
//   3. Escribe una función que devuelva el estado de atras y adelante
//      para poder imprimirlo en main.go.
// Comenta con // qué hace cada cambio.
// =============================================================
