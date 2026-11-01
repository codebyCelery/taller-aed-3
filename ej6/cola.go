// EJERCICIO 6 · Responsable: Celery
package ej6

import "errors"

// Cola es lo que comparten las cuatro versiones (para poder probarlas con la misma tabla).
type Cola interface {
	Encolar(v int)
	Desencolar() (int, error)
}

// ColaPedidos es la versión CORRECTA del enunciado.
type ColaPedidos struct {
	entrada []int
	salida  []int
}

func (c *ColaPedidos) Encolar(v int) {
	c.entrada = append(c.entrada, v)
}

func (c *ColaPedidos) Desencolar() (int, error) {
	if len(c.salida) == 0 {
		for len(c.entrada) > 0 {
			n := len(c.entrada) - 1
			c.salida = append(c.salida, c.entrada[n])
			c.entrada = c.entrada[:n]
		}
	}
	if len(c.salida) == 0 {
		return 0, errors.New("no hay pedidos")
	}
	n := len(c.salida) - 1
	v := c.salida[n]
	c.salida = c.salida[:n]
	return v, nil
}

// ======================= AQUÍ ESCRIBES =======================
// Tres tipos nuevos: ColaX, ColaY y ColaZ. Cada uno es una COPIA de
// ColaPedidos (mismos campos y mismos dos métodos) con UN solo cambio,
// el de la tabla del enunciado:
//   X -> línea 11:  if len(c.entrada) > 0 {
//   Y -> líneas 13 a 15: pasa c.entrada[0] a salida y hace c.entrada = c.entrada[1:]
//   Z -> línea 12:  for len(c.entrada) > 1 {
// Comenta con // la línea que cambiaste en cada copia.
// =============================================================
