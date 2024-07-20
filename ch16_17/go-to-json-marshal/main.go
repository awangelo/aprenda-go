package main

import (
	"encoding/json"
	"fmt"
)

type Gun struct {
	Name    string
	MaxAmmo int
	Auto    bool
}

func main() {
	ak47 := Gun{Name: "AK-47", MaxAmmo: 30, Auto: true}
	m4a1s := Gun{"M4-A1s", 24, false}
	awp := Gun{"AWP", 10, false}

	ak47JSON := ak47.getJSON()
	m4a1sJSON := m4a1s.getJSON()
	awpJSON := awp.getJSON()

	fmt.Println(ak47JSON)
	fmt.Println(m4a1sJSON)
	fmt.Println(awpJSON)
}

func (g Gun) getJSON() string {
	ej, err := json.Marshal(g)
	if err != nil {
		return err.Error()
	}

	return string(ej)
}
