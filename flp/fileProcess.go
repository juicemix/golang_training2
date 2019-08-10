package flp

import (
	"encoding/json"
	"io/ioutil"
)

func WriteData(i interface{}, filename string) string {
	err := ioutil.WriteFile(filename, ConvertToJSON(i), 0777)
	if err != nil {
		return err.Error()
	} else {
		return "success"
	}
}

func ReadToStruct(filename string, i interface{}) (interface{}, string) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, "failed"
	} else {
		err = json.Unmarshal(f, &i)
		if err != nil {
			return nil, "failed parsing"
		} else {
			return i, "success"
		}
	}
}

func ConvertToJSON(i interface{}) []byte {
	jsonData, err := json.Marshal(i)
	if err != nil {
		return jsonData
	} else {
		return nil
	}
}
