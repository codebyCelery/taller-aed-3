package ej5

import "testing"

func TestBalanceada(t *testing.T) {
    casos := []struct {
        nombre   string
        formula  string
        esperado bool
    }{

{"Cierre correcto", "()", true},
{"Simbolos correctos", "{([])}", true},
{"Sin cerrar", "((", false},
{"Cierre incorrecto", "{[(])}", false},
{"Con letras", "=SUMA(A1:B4)", true},
{"Formula vacia", "", true},
{"Parentesis sin cerrar", "(", false},
{"Letra sola", "a", true},

    }

    for _, caso := range casos {
        t.Run(caso.nombre, func(t *testing.T) {
            obtenido := Balanceada(caso.formula)

            if obtenido != caso.esperado {
                t.Errorf("Esperaba %v, pero salio %v",
                    caso.esperado, obtenido)
            }
        })
    }
}
