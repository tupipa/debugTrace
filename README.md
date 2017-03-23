# Debug Tracer helper lib for go program
This is a helper lib for printing the call stack of a go program

## install 
```
git clone https://github.com/tupipa/debugTrace $GOPATH/src/github.com/tupipa/debugTrace
git clone https://github.com/Sirupsen/logrus $GOPATH/src/github.com/Sirupsen/logrus
go install -v github.com/tupipa/debugTrace
go install -v github.com/Sirupsen/logrus
```

## Usage Sample
create your samplecode.go in $GOPATH/github.com/yourid/repo

```
package main

import (
	"os"
	"fmt"
	debugTrace "github.com/tupipa/debugTrace"
	"github.com/Sirupsen/logrus"
	)

func test() int{
	i := 1
	i += i
	//print only current function info
    debugTrace.Trace("")
	
	debugTrace.Trace2("5 Levels on top of the call stack", debugTrace.TraceLevel5)
    return i
}

func main(){
    //set up the debug io directions and format
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)

    //call functions
	c := test()
	fmt.Printf("get i: %d\n", c)
	
    //print call stack info of current function(main)
    debugTrace.Trace("main")
}

```
run ```go build``` in $GOPATH/github.com/yourid/repo

output will be something like:

```
DEBU[0000] Lele: $GOPATH/github.com/yourid/repo/samplecode.go:16; funcName: 'main.test' 
DEBU[0000] Lele: 5 Levels on top of the call stack, $GOPATH/github.com/yourid/repo/samplecode.go:17; funcName: 'main.test' 
DEBU[0000] Lele: 5 Levels on top of the call stack, $GOPATH/github.com/yourid/repo/samplecode.go:26; funcName: 'main.main' 
DEBU[0000] Lele: 5 Levels on top of the call stack, $GOROOT/src/runtime/proc.go:194; funcName: 'runtime.main' 
DEBU[0000] Lele: 5 Levels on top of the call stack,  $GOROOTsrc/runtime/asm_amd64.s:2198; funcName: 'runtime.goexit' 
get i: 2
DEBU[0000] Lele: main, $GOPATH/github.com/yourid/repo/samplecode.go:31; funcName: 'main.main'
```