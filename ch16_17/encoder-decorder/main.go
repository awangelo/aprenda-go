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
	sb1 := []byte(`{"name": "AK-47", "max_ammo": 30, "auto": true}`)
	sb2 := []byte(`{"name": "M4A1", "max_ammo": 27, "auto": true}`)
	sb3 := []byte(`{"name": "Desert Eagle", "max_ammo": 7, "auto": false}`)

	var gun1, gun2, gun3 gun

	e := json.NewEncoder(os.Stdout)

	err := json.Unmarshal(sb1, &gun1)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(sb2, &gun2)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(sb3, &gun3)
	if err != nil {
		panic(err)
	}

	err = e.Encode(gun1)
	if err != nil {
		fmt.Printf("Erro no encode gun1:")
		panic(err)
	}
	err = e.Encode(gun2)
	if err != nil {
		fmt.Printf("Erro no encode gun2:")
		panic(err)
	}
	err = e.Encode(gun3)
	if err != nil {
		fmt.Printf("Erro no encode gun3:")
		panic(err)
	}
}
