package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Numero de valores a serem enviados ou seja, nro de goroutines
	go send(ch1, 30)
	go diverge(ch1, ch2)

	for v := range ch2 {
		fmt.Println(v)
	}
}

func send(c chan int, n int) {
	for v := range n { // Envia valores de 0 a n para o canal c
		c <- v
	}
	close(c) // Fecha o canal c após enviar todos os valores
}

func diverge(ch1, ch2 chan int) {
	var wg sync.WaitGroup // Inicializa um WaitGroup para sincronizar goroutines

	// Esse loop lê valores do canal ch1 e cria uma nova goroutine para cada valor
	// Vai terminar quando o canal ch1 for fechado (na funcao send)
	for v := range ch1 {
		wg.Add(1)
		go func(v int) { // Cada goroutine executa a função worker com o valor v, ou seja, cada goroutine envia apenas um valor
			ch2 <- worker(v) // Envia o resultado da função worker para o canal ch2
			wg.Done()
		}(v)
	}

	wg.Wait()  // Espera todas as goroutines terminarem com o seu valor
	close(ch2) // Fecha o canal final
}

// Apenas simula um worker real
func worker(n int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
	return n
}
