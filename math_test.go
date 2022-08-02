package main 

import "testing"
import "Soma"

funct TestSoma(t *testing.T) {
	total := Soma(15,15)

	if total != 30{
		t.Errorf("Inválido: Resultado %d. Esperado %d", total, 30)
	}
}