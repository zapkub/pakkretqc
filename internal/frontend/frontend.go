package frontend

import "encoding/json"

func EncodeJSON(p interface{}) (string, error) {
	b, err := json.Marshal(p)
	return string(b), err
}
