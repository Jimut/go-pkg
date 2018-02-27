package Alligator

import (
	"encoding/json"
)

type Teeth struct {
	Enamel string
}

func GetTeeth() string {
	return "No Teeth"
}

func GetMeta() map[string]int {
	buff, err := Asset("data/meta.json")
	check(err)

	data := make(map[string]int)
	json.Unmarshal(buff, &data)
	return data
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
