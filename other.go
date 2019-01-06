package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config represents conf.json
type Config struct {
	Me      int64
	Updates int64
	Domain	string
	Token   string
	Link    string
	Key     string
	Cx      string
}

func stringInSlice(s, k string, sl []map[string]interface{}) bool {
	for _, v := range sl {
		if v[k] == s {
			return true
		}
	}
	return false
}

func deleteKeys(m map[string]interface{}, k []string) {
	for _, v := range k {
		if _, ok := m[v]; ok {
			delete(m, v)
		}
	}
}

func getConfig(filename string) Config {
	var config Config
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	decoder := json.NewDecoder(file)
	decoder.Decode(&config)
	return config
}
