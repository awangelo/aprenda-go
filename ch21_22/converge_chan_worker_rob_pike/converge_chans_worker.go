package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// "Junta" dois canais em um
	canal := converge(worker("primeiro_job"), worker("segundo_jb"))

	// Recebe mensagens do canal convergido
	for v := range canal {
		fmt.Println(v)
	}
}

// Eh usado para criar um canal que envia mensagens
func worker(m string) chan string {
	ch := make(chan string)

	// Envia mensagens para o canal
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%v diz: %v", m, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}()
	return ch
}

// Converge dois canais em um e retorna o canal convergido
func converge(c1, c2 chan string) chan string {
	new_ch := make(chan string) // Cria um novo canal para convergência

	go func() {
		for {
			new_ch <- <-c1
		}
	}()
	go func() {
		for {
			new_ch <- <-c2
		}
	}()

	return new_ch
}
