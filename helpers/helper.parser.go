package helpers

import (
	"encoding/json"
	"io"

	"github.com/sirupsen/logrus"
)

func Stringify(data any) any {
	res, err := json.Marshal(data)
	if err != nil {
		logrus.Errorf("Create json format error: %v", err)
	}

	err = json.Unmarshal(res, data)

	if err != nil {
		logrus.Errorf("Create json format error: %v", err)
	}

	return data
}

func BodyParser(reader io.Reader, body any) {
	err := json.NewDecoder(reader).Decode(body)

	if err != nil {
		logrus.Errorf("Parse body data error: %v", err)
	}
}
