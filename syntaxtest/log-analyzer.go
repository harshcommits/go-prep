package syntaxtest

import (
	"fmt"
	"os"
)

/*
Idea is to process logs like these:

INFO User logged in
ERROR Database timeout
INFO Request completed
WARN High memory usage
ERROR Connection refused

and generate output like this:
ERROR: 2
INFO: 2
WARN: 1
*/

type LogAnalysis struct {
	error int
	info int
	warn int
}

func AnalyzeLogs(logFile string) {

	dataFile, err := os.ReadFile(logFile)
	if err != nil {
		fmt.Println("failed to read log file:", err)
		return
	}

	_ = dataFile
}