package main

import (
	"encoding/json"
	"fmt"
)

type gun struct {
	Name    string `json:"name"`
	MaxAmmo int    `json:"max_ammo"`
	Auto    bool   `json:"auto"`
}

func main() {
	sb1 := []byte(`{"name": "AK-47", "max_ammo": 30, "auto": true}`)
	sb2 := []byte(`{"name": "M4A1", "max_ammo": 27, "auto": true}`)
	sb3 := []byte(`{"name": "Desert Eagle", "max_ammo": 7, "auto": false}`)

	var gun1, gun2, gun3 gun

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

	fmt.Println("Gun 1:", gun1)
	fmt.Println("Gun 2:", gun2)
	fmt.Println("Gun 3:", gun3)
}
