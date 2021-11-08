package test

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
)

func CompareString(a, b string) bool {
	return a == b
}

func CompareBytes(a, b []byte) bool {
	return string(a) == string(b)
}

func CompareJSON(a, b interface{}) (bool, error) {
	jsonA, err := json.Marshal(a)
	if err != nil {
		return false, err
	}
	jsonB, err := json.Marshal(b)
	if err != nil {
		return false, err
	}
	return string(jsonA) == string(jsonB), nil
}

func CompareYaml(a, b []byte) (bool, error) {
	var marshalDataA, marshalDataB map[string]interface{}
	err := yaml.Unmarshal(a, &marshalDataA)
	if err != nil {
		return false, err
	}
	err = yaml.Unmarshal(b, &marshalDataB)
	if err != nil {
		return false, err
	}
	return CompareJSON(marshalDataA, marshalDataB)
}