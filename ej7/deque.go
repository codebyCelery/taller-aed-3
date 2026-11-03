package main

// Deque es una cola de dos extremos: se puede agregar y sacar por
// AMBOS lados. Lo necesitamos porque una Pila solo deja sacar por el
// tope, y la página más antigua está en el fondo.
//
// El frente es la página más antigua; el final es la página actual.
type Deque[T any] struct {
	items []T
}

// PushBack agrega al final (página nueva).
func (d *Deque[T]) PushBack(v T) { d.items = append(d.items, v) }

// PopBack saca del final (retroceder a la página anterior).
func (d *Deque[T]) PopBack() (T, bool) {
	var cero T
	if len(d.items) == 0 {
		return cero, false
	}
	n := len(d.items) - 1
	v := d.items[n]
	d.items = d.items[:n]
	return v, true
}

// PopFront saca del frente (descartar la más antigua). Como el deque
// nunca tiene más de 5 elementos, desplazarlos cuesta O(1).
func (d *Deque[T]) PopFront() (T, bool) {
	var cero T
	if len(d.items) == 0 {
		return cero, false
	}
	v := d.items[0]
	copy(d.items, d.items[1:])
	d.items = d.items[:len(d.items)-1]
	return v, true
}

func (d *Deque[T]) Len() int { return len(d.items) }
