package assignments

import (
	"os"
	"runtime"
	"strings"
)

func writeFile(fileName string, fileContent string) {
	defaultDir, _ := os.Getwd()
	handle, err := os.OpenFile(defaultDir+"/Output/"+fileName+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	handle.WriteString(fileContent)
	handle.Close()
}

func getFuncName() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	b := strings.Split(frame.Function, ".")
	return b[1]
}
