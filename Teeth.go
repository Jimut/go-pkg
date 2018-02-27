package Alligator

import (
	"encoding/json"
	"io/ioutil"
)

type Teeth struct {
	Enamel string
}

func GetTeeth() string {
	return "No Teeth"
}

func GetMeta() map[string]int {
	buff, err := ioutil.ReadFile("data/meta.json")
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
