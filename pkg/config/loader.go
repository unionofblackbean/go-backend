package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func Load(path string, obj interface{}) error {
	rawBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to open config file -> %v", err)
	}

	err = json.Unmarshal(rawBytes, obj)
	if err != nil {
		return fmt.Errorf("failed to unmarshal content in config file -> %v", err)
	}

	return nil
}
