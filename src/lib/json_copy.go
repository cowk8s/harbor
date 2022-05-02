package lib

import (
	"encoding/json"
)

// JSONCopy copy from src to dst with json marshal and unmarshal.
// NOTE: copy from one struct to another struct may miss some values depend on
// the json tag in the fields, you should know what will happened when call this function.
func JSONCopy(dst, src interface{}) error {
	data, err := json.Marshal(src)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, dst)
}
