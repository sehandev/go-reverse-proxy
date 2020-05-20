package main

import (
	"encoding/json"
	"io/ioutil"
)

func readSettingJSON() (*settingStruct, error) {
	file, err := ioutil.ReadFile(settingJSONPath)
	if err != nil {
		return nil, err
	}
	setting := &settingStruct{}
	err = json.Unmarshal(file, setting)
	return setting, err
}
