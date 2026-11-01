package ej5

import "errors"

// Pila de clase (S6). Si ya tienen la suya, pueden reemplazar este archivo.
type Pila[T any] struct {
	items []T
}

func (p *Pila[T]) Push(v T) { p.items = append(p.items, v) }

func (p *Pila[T]) Pop() (T, error) {
	var cero T
	if len(p.items) == 0 {
		return cero, errors.New("pila vacía")
	}
	n := len(p.items) - 1
	v := p.items[n]
	p.items = p.items[:n]
	return v, nil
}

func (p *Pila[T]) Vacia() bool { return len(p.items) == 0 }
