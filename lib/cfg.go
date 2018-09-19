package qniblib

import (
	"github.com/Sirupsen/logrus"
	"github.com/zpatrick/go-config"
	"os"
)

func NewConfig(c string) (cfg *config.Config, err error) {
	_, err =  os.Open(c)
	if err != nil {
		return
	}
	logrus.Infof("Loading Houdini config '%s'", c)
	iniFile := config.NewINIFile(c)
	cfg = config.NewConfig([]config.Provider{iniFile})
	return
}
