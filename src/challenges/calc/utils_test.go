package main

import "testing"

func TestSoma(t *testing.T) {
    resultado := soma(3, 2, 1)
    esperado := 6
    if resultado != esperado {
        t.Errorf("Esperado: %d, Obtido: %d", esperado, resultado)
    }
}

func TestSomaErrado(t *testing.T) {
    resultado := soma(3, 2, 1)
    esperado := 7
    if resultado != esperado {
        t.Errorf("Esperado: %d, Obtido: %d", esperado, resultado)
    }
}

func TestMultiplica(t *testing.T) {
    resultado := multiplica(10, 10)
    esperado := 100
    if resultado != esperado {
        t.Errorf("Esperado: %d, Obtido: %d", esperado, resultado)
    }
}

func TestMultiplicaErrado(t *testing.T) {
    resultado := multiplica(10, 10)
    esperado := 2560
    if resultado != esperado {
        t.Errorf("Esperado: %d, Obtido: %d", esperado, resultado)
    }
}

func TestSubtrai(t *testing.T) {
    resultado := subtrai(10, 5)
    esperado := 5
    if resultado != esperado {
        t.Errorf("Esperado: %d, Obtido: %d", esperado, resultado)
    }
}

func TestSubtraiErrado(t *testing.T) {
    resultado := subtrai(10, 10)
    esperado := 5
    if resultado != esperado {
        t.Errorf("Esperado: %d, Obtido: %d", esperado, resultado)
    }
}

func TestDivide(t *testing.T) {
    resultado := divide(10, 2)
    esperado := 5
    if resultado != esperado {
        t.Errorf("Esperado: %d, Obtido: %d", esperado, resultado)
    }
}

func TestDivideErrado(t *testing.T) {
    resultado := divide(10, 2)
    esperado := 3
    if resultado != esperado {
        t.Errorf("Esperado: %d, Obtido: %d", esperado, resultado)
    }
}
