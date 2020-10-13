package adtype

import "encoding/json"

func DeepCopy(target, source interface{}) error {
	if byts, err := json.Marshal(source); err != nil {
		return err
	} else {
		return json.Unmarshal(byts, target)
	}
}
