package zlog

import "testing"

func TestSugar(t *testing.T) {
	initSugar()
	sugarLogger.Debug("debug")
	sugarLogger.Infof("infof")
}
