package log

import (
	"fmt"
	"icode.baidu.com/baidu/goodcoder/gongyitong/constants"
	"testing"
)
func setLogTest() {
	SetDebugMode(true)
	SetEnv(constants.ENV_DEVBOX)
}
func TestLog(t *testing.T) {
	setLogTest()
	if err := setup();
		err != nil {
		t.Error("lo`g.setup() fail")
	}
	level := fmt.Sprint("debug  log")
	Debug("%s output", level)

	level = fmt.Sprint("warn  log")
	Warning("%s output", level)

	level = fmt.Sprint("Critical  log")
	Critical("%s output", level)

	level = fmt.Sprint("Info log")
	Info("%s output", level)

	level = fmt.Sprint("Err  log")
	Err("%s output ", level)

}