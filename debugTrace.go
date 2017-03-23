package debugTrace

import (
	"github.com/Sirupsen/logrus"
	"runtime"
)

// Trace will print the current function name and the code line number
func Trace(info string) {

	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	//fmt.Printf("Lele: %s, %s:%d %s\n", info, file, line, f.Name())
	if len(info) == 0 {
		logrus.Debugf("Lele: %s:%d; funcName: '%s'", file, line, f.Name())
	}else{
		logrus.Debugf("Lele: %s, %s:%d; funcName: '%s'", info, file, line, f.Name())
	}
}
