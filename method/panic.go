package method
import (
	"fmt"
	"runtime"
)
func Recovery(format string, values ...interface{}) {
	if err := recover(); err != nil {
		buf := make([]byte, 1<<18)
		n := runtime.Stack(buf, false)
		ctx := fmt.Sprintf(format, values...)
		fmt.Sprintf("%d",n)
		fmt.Sprintf("%s",ctx)

		//log.ErrWithNotification("PANIC: "+ctx+" panic: %+v. stack: %s", err, buf[0:n])
	}
}
