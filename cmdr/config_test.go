package cmdr

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfigShouldReturnErrorIfCannotReadFile(t *testing.T) {
	readFileFn = func(filename string) ([]byte, error) {
		return nil, errors.New("Some error")
	}

	var c = NewConfig()
	_, e := c.LoadConfig("")

	assert.Equal(t, "Some error", e.Error(), "they should be equal")
}

func TestLoadConfigShouldReturnMapIfReadFileSuccess(t *testing.T) {
	readFileFn = func(filename string) ([]byte, error) {
		data := []byte(`{"api_token": "abcd", "ip": "0.0.0.0", "port": "1234", "commands": {"aaa":"ccc","ddd":"1234"}}`)
		return data, nil
	}

	var c = NewConfig()
	var expected = &ConfigType{
		APIToken: "abcd",
		IP:       "0.0.0.0",
		Port:     "1234",
		Commands: map[string]string{
			"aaa": "ccc",
			"ddd": "1234",
		},
	}
	result, e := c.LoadConfig("")

	assert.Equal(t, expected, result, "they should be equal")
	assert.Equal(t, nil, e, "they should be equal")
}
