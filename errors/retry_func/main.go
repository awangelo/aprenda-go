package main

import (
	"fmt"
	"math/rand"
)

func main() {
	if err := thisMightFail(); err != nil {
		retry()

	} else {
		fmt.Println("Success")
	}

}

func retry() error {
	var err error
	for i := range 3 {
		fmt.Printf("Failed %d times\n", i+1)
		err = thisMightFail()
		if err == nil {
			return nil
		}
	}
	return err
}

func thisMightFail() error {
	if rand.Intn(3) == 0 {
		return nil
	} else {
		return fmt.Errorf("Random error")
	}
}
