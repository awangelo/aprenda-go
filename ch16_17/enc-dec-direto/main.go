package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type gun struct {
	Name    string
	MaxAmmo int
	Auto    bool
}

func main() {
	arma1 := gun{
		Name:    "M4A1",
		MaxAmmo: 30,
		Auto:    true,
	}
	arma2 := gun{
		Name:    "r8",
		MaxAmmo: 8,
		Auto:    false,
	}

	e := json.NewEncoder(os.Stdout)

	if err := e.Encode(arma1); err != nil {
		fmt.Println(err)
	}
	if err := e.Encode(arma2); err != nil {
		fmt.Println(err)
	}
}
