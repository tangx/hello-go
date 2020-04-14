package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func Test_main(t *testing.T) {
	verbose := 8
	logLevel := logrus.Level(verbose)
	// logrus.SetLevel(logrus.DebugLevel)
	logrus.SetLevel(logLevel)
	main()
}
