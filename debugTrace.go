package debugTrace

import (
	"fmt"
	"runtime"
)

func Trace() {

	pc := make ([] uintptr, 10)
	runtime.Callers (2,pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	fmt.Printf("Lele: %s:%d %s\n", file, line, f.Name())
}



