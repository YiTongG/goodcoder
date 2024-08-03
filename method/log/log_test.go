package log

import (
	"testing"
)

func setLogTest() {
	SetDebugMode(true)
	SetEnv(EnvDevbox)
}
func TestLog(t *testing.T) {
	setLogTest()

	level := "debug  log"
	Debug("%v ", level)

	level = "warn  log"
	Warning("output:", level)

	level = "Info log"
	Info("output:", level)

	level = "Err  log"
	Err("output:", level)

}
