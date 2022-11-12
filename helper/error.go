package helper

import (
	"github.com/sirupsen/logrus"
)

func PanicIfError(err error) {
	if err != nil {
		logger := logrus.New()
		logger.Error(err)
		panic(err)
	}
}
