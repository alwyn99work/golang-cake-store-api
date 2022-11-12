package helper

import "github.com/sirupsen/logrus"

func Log(msg interface{}) {
	logger := logrus.New()
	logger.Info(msg)
}
