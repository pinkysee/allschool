package logger

import "github.com/sirupsen/logrus"

func Start(cfg *Log_config) error {
	level, err := logrus.ParseLevel(cfg.Log_level)
	logrus.SetLevel(level)
	logrus.Info(level)
	if err != nil {
		return err
	}
	return nil
}
