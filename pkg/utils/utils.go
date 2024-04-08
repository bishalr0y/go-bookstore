package utils

import (
	"encoding/json"
)

// to unmarshal the json
func Unmarshal(msg string, x interface{}) error {
	err := json.Unmarshal([]byte(msg), &x)

	if err != nil {
		return err
	}
	return nil
}
