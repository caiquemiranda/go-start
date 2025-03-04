package main

// Soma todos os números passados como argumento
func soma(i ...int) int {
    total := 0
    for _, v := range i {
        total += v
    }
    return total
}

// Subtrai os números sequencialmente
func subtrai(i ...int) int {
    if len(i) == 0 {
        return 0
    }
    total := i[0]
    for _, v := range i[1:] {
        total -= v
    }
    return total
}

// Multiplica todos os números passados como argumento
func multiplica(i ...int) int {
    total := 1
    for _, v := range i {
        total *= v
    }
    return total
}

// Divide os números sequencialmente, evitando divisão por zero
func divide(i ...int) int {
    if len(i) == 0 {
        return 0
    }
    total := i[0]
    for _, v := range i[1:] {
        if v == 0 {
            return 0 // Evita erro de divisão por zero
        }
        total /= v
    }
    return total
}
