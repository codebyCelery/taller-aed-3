// EJERCICIO 6 · Responsable: Celery
package ej6

import "errors"

// Cola es lo que comparten las cuatro versiones, para poder
// probarlas todas con la misma tabla de pruebas.
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
	// Solo se traspasa cuando salida está vacía (línea 11).
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

// ---------------------------------------------------------------
// COPIA X: cambia la línea 11.
// Traspasa SIEMPRE que entrada tenga algo, aunque salida todavía
// tenga pedidos más antiguos. Los pedidos nuevos quedan encima de
// los viejos y salen primero: se rompe el orden de llegada (FIFO).
// ---------------------------------------------------------------
type ColaX struct {
	entrada []int
	salida  []int
}

func (c *ColaX) Encolar(v int) {
	c.entrada = append(c.entrada, v)
}

func (c *ColaX) Desencolar() (int, error) {
	if len(c.entrada) > 0 { // CAMBIO (línea 11): antes era len(c.salida) == 0
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

// ---------------------------------------------------------------
// COPIA Y: cambia las líneas 13 a 15.
// Pasa entrada[0] (el primero) en vez del último. El traspaso ya no
// invierte el orden, y al sacar del final de salida sale el ÚLTIMO
// pedido en llegar: se comporta como una pila (LIFO).
// ---------------------------------------------------------------
type ColaY struct {
	entrada []int
	salida  []int
}

func (c *ColaY) Encolar(v int) {
	c.entrada = append(c.entrada, v)
}

func (c *ColaY) Desencolar() (int, error) {
	if len(c.salida) == 0 {
		for len(c.entrada) > 0 {
			c.salida = append(c.salida, c.entrada[0]) // CAMBIO (líneas 13-15)
			c.entrada = c.entrada[1:]
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

// ---------------------------------------------------------------
// COPIA Z: cambia la línea 12.
// El for se detiene cuando queda 1 elemento en entrada, así que ese
// elemento nunca pasa a salida. Con un solo pedido encolado, Desencolar
// dice "no hay pedidos" aunque sí hay uno.
// ---------------------------------------------------------------
type ColaZ struct {
	entrada []int
	salida  []int
}

func (c *ColaZ) Encolar(v int) {
	c.entrada = append(c.entrada, v)
}

func (c *ColaZ) Desencolar() (int, error) {
	if len(c.salida) == 0 {
		for len(c.entrada) > 1 { // CAMBIO (línea 12): antes era > 0
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
