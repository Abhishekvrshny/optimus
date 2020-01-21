package util

import (
	"io"
	"io/ioutil"
)

func ReadBody(r io.ReadCloser) ([]byte, error){
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return body, nil
}
