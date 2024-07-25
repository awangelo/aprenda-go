package main

import (
	"fmt"
	"log"
)

// Struct para guardar mais informacoes do erro.
type sqrtError struct {
	lat  string
	long string
	err  error
}

// Error implementa a interface de erro para sqrtError,
// permitindo que ele seja tratado como um `erro` pelo Go.
func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// Cria um novo erro com uma mensagem para este erro especifico.
		newError := fmt.Errorf("deu ruim: %v\n", f)
		// sqrtError implementa a interface de erro. Um tipo de erro que guarda mais informações.
		return 0, sqrtError{"50.22N", "99.46W", newError}
	}
	return 42, nil
}
