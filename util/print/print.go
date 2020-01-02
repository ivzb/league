package print

import (
	"encoding/json"
	"fmt"
)

func Pretty(v interface{}) error {
	bytes, err := json.MarshalIndent(v, "", "    ")

	if err != nil {
		return nil
	}

	fmt.Println(string(bytes))

	return nil
}
