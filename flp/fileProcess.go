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
		return nil, "File not found"
	} else {
		err = json.Unmarshal(f, &i)
		if err != nil {
			return nil, "File empty"
		} else {
			return i, "success"
		}
	}
}

func ReadToStructTest(filename string, i interface{}) string {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return "File not found"
	} else {
		err = json.Unmarshal(f, &i)
		if err != nil {
			return "File empty"
		} else {
			return "success"
		}
	}
}

func ConvertToJSON(i interface{}) []byte {
	jsonData, err := json.Marshal(i)
	if err == nil {
		return jsonData
	} else {
		return nil
	}
}
