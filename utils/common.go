package utils

import "github.com/sirupsen/logrus"

func CheckAndExit(err error) {
	if err != nil {
		logrus.Fatal(err)
	}
}
