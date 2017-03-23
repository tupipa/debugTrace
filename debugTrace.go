package debugTrace

import (
	"fmt"
	"runtime"
)

// Trace will print the current function name and the code line number
func Trace(info string) {

	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	fmt.Printf("Lele: %s, %s:%d %s\n", info, file, line, f.Name())
}
