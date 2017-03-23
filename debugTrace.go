package debugTrace

import (
	"github.com/Sirupsen/logrus"
	"runtime"
)

// print the function name and code line from the call stack of Trace()/Trace2() function.
// TraceLevel1=0 is for the caller of Trace func; 
// TraceLevel2=1 is the caller of caller of Trace func, etc..
// up to 10 levels on the stack; also limited by the stack depth.

const (
 TraceLevel1 int = 0
 TraceLevel2 int = 1
 TraceLevel3 int = 2
 TraceLevel4 int = 3
 TraceLevel5 int = 4
 TraceLevel6 int = 5
 TraceLevel7 int = 6
 TraceLevel8 int = 7
 TraceLevel9 int = 8
 TraceLevel10 int = 9
 MAXLEVEL int = 10
)


// Trace(str) will print the caller function name and the code line number
// only print current level 0 start from the caller.
func Trace(info string) {
	pc := make([]uintptr, 1)

	// get call stack, ommitting 2 elements on stack, which is runtime.Callers and Trace2, respectively
	runtime.Callers(2 , pc)
	f := runtime.FuncForPC(pc[0])
	// get the file and line number of the instruction immediately following the call
	file, line := f.FileLine(pc[0])

	if len(info) == 0 {
		logrus.Debugf("Lele: %s:%d; funcName: '%s'", file, line, f.Name())
	}else{
		logrus.Debugf("Lele: %s, %s:%d; funcName: '%s'", info, file, line, f.Name())
	}

}

// Trace(str) will print calling stack info from current Level 0 to the given level.
func Trace2(info string, level int) {

	pc := make([]uintptr, MAXLEVEL)

	//get call stack, ommitting 2 elements on stack, which is runtime.Callers and Trace2, respectively
	funcsNum := runtime.Callers(2 , pc)

	for i := 0; i< funcsNum && i <= level && i <= MAXLEVEL; i++{

		f := runtime.FuncForPC(pc[i])
		// get the file and line number of the instruction immediately following the call
		file, line := f.FileLine(pc[i])

		if len(info) == 0 {
			logrus.Debugf("Lele: %s:%d; funcName: '%s'", file, line, f.Name())
		}else{
			logrus.Debugf("Lele: %s, %s:%d; funcName: '%s'", info, file, line, f.Name())
		}
	}

	if (level > MAXLEVEL){
		logrus.Warnf("Backtracing level is up to %d; cannot backtracing to %d level.",
					MAXLEVEL, level)
	}

}
