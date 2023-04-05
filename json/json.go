package cxfilejson


import (
  "github.com/valyala/fastjson"
	"io/ioutil"
	"os"
)

func ParseJsonFile(file string, config *Config) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	//config := &Config{}
	if err := fastjson.Unmarshal(data, &config); err != nil {
		return err
	}

	return nil
}
