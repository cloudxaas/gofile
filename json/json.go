package cxfilejson

import (
  "github.com/valyala/fastjson"
	"io/ioutil"
	"os"
)

func ParseJsonFile(file string) (*Config, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	if err := fastjson.Unmarshal(data, config); err != nil {
		return nil, err
	}

	return config, nil
}
