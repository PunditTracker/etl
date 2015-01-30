package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

func getStringOrEmpty(i interface{}) string {
	if v, ok := i.(string); ok {
		return v
	} else {
		return ""
	}
}

func parseOldDateFormat(ti interface{}) time.Time {
	form := "2006-01-02 15:04:05"
	t := ti.(string)
	timeCreated, err := time.Parse(form, t)
	if err != nil {
		fmt.Println(err)
		timeCreated = time.Now()
	}
	return timeCreated
}

func toJsonFromFile(filename string) []map[string]interface{} {
	data, err := ioutil.ReadFile(filename)
	data = bytes.Replace(data, []byte("\\'"), []byte(`'`), -1)
	data = bytes.Replace(data, []byte("\t"), []byte(` `), -1)
	data = bytes.Replace(data, []byte("\r"), []byte(` `), -1)
	data = bytes.Replace(data, []byte("\n"), []byte(` `), -1)
	//data = bytes.Replace(data, []byte("/"), []byte(``), -1)
	categories := make([]map[string]interface{}, 0)
	if err != nil {
		fmt.Println("in file: ", filename, err)
		return categories
	}
	err = json.Unmarshal(data, &categories)
	if err != nil {
		fmt.Println("error unmarshalling ", filename, err)
	}
	return categories
}
