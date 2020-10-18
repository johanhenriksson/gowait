package gowait

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"os"
)

const envGzip = "COWAIT_GZIP"
const envTaskdef = "COWAIT_TASK"

func unpackenv(key string) (string, error) {
	data := os.Getenv(key)
	if os.Getenv(envGzip) != "1" {
		return data, nil
	}

	z, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	r, err := gzip.NewReader(bytes.NewReader(z))
	if err != nil {
		return "", err
	}

	result, _ := ioutil.ReadAll(r)
	return string(result), nil
}

func unpackstruct(key string, out interface{}) error {
	str, err := unpackenv(key)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(str), out)
}
