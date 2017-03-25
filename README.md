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
	"github.com/tupipa/debugTrace"
	"github.com/Sirupsen/logrus"
	)

func test() int{
	i := 1
	i += i
	//print only current function info
    	debugTrace.Trace("Print the CODE LINE and FUNCTION NAME Here")

	debugTrace.Trace2("Print (up to) 5 LEVELS OF CALLED FUNCTIONS on stack", debugTrace.TraceLevel5)

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
    	debugTrace.Trace2("Print (up to) 6 LEVELS OF CALLED FUNCTIONS on stack", debugTrace.TraceLevel6)
}

```
run ```go build``` in $GOPATH/github.com/yourid/repo

output will be something like:

```
DEBU[0000] Lele: Print the CODE LINE and FUNCTION NAME Here, /home/smeller/lab/gowork/src/github.com/tupipa/maingotest/maingotest2.go:16; funcName: 'main.test' 
DEBU[0000] Lele: Print (up to) 5 LEVELS OF CALLED FUNCTIONS on stack, /home/smeller/lab/gowork/src/github.com/tupipa/maingotest/maingotest2.go:18; funcName: 'main.test' 
DEBU[0000] Lele: Print (up to) 5 LEVELS OF CALLED FUNCTIONS on stack, /home/smeller/lab/gowork/src/github.com/tupipa/maingotest/maingotest2.go:27; funcName: 'main.main' 
DEBU[0000] Lele: Print (up to) 5 LEVELS OF CALLED FUNCTIONS on stack, /home/smeller/programs/go/go/src/runtime/proc.go:194; funcName: 'runtime.main' 
DEBU[0000] Lele: Print (up to) 5 LEVELS OF CALLED FUNCTIONS on stack, /home/smeller/programs/go/go/src/runtime/asm_amd64.s:2198; funcName: 'runtime.goexit' 
get i: 2
DEBU[0000] Lele: Print (up to) 6 LEVELS OF CALLED FUNCTIONS on stack, /home/smeller/lab/gowork/src/github.com/tupipa/maingotest/maingotest2.go:32; funcName: 'main.main' 
DEBU[0000] Lele: Print (up to) 6 LEVELS OF CALLED FUNCTIONS on stack, /home/smeller/programs/go/go/src/runtime/proc.go:194; funcName: 'runtime.main' 
DEBU[0000] Lele: Print (up to) 6 LEVELS OF CALLED FUNCTIONS on stack, /home/smeller/programs/go/go/src/runtime/asm_amd64.s:2198; funcName: 'runtime.goexit' 
```
