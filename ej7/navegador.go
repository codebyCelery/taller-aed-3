// EJERCICIO 7 · Responsable: Celery
package main

// maxHistorial es el límite de páginas en atras, contando la actual.
const maxHistorial = 5

// Navegador con el historial de atrás limitado a 5 páginas.
type Navegador struct {
	atras    Deque[string] // CAMBIO: antes era Pila. Frente = más antigua, final = actual
	adelante Pila[string]
}

func (n *Navegador) Visitar(url string) {
	n.atras.PushBack(url)
	// CAMBIO: si pasamos de 5, descartamos la más antigua (por el frente).
	if n.atras.Len() > maxHistorial {
		n.atras.PopFront()
	}
	n.adelante = Pila[string]{} // una visita nueva borra el "adelante"
}

func (n *Navegador) Atras() {
	// Solo se puede retroceder si hay más de una página (la actual no cuenta).
	if n.atras.Len() > 1 {
		url, _ := n.atras.PopBack() // CAMBIO: se saca por el final del deque
		n.adelante.Push(url)
	}
}

func (n *Navegador) Adelante() {
	if url, err := n.adelante.Pop(); err == nil {
		n.atras.PushBack(url) // CAMBIO: se vuelve a agregar por el final
	}
}

func (n *Navegador) Actual() string {
	return n.atras.items[len(n.atras.items)-1] // la actual es el último del deque
}

// Estado devuelve atras (de la más antigua a la actual) y adelante
// (del fondo al tope) como copias, solo para poder imprimirlas.
func (n *Navegador) Estado() ([]string, []string) {
	return append([]string(nil), n.atras.items...),
		append([]string(nil), n.adelante.items...)
}
