package config

import (
	"encoding/json"
	"io/ioutil"
)

func Load(path string, obj interface{}) (err error) {
	rawBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	err = json.Unmarshal(rawBytes, obj)
	if err != nil {
		return
	}

	return
}
