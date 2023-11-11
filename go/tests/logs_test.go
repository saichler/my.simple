package tests

import (
	"fmt"
	"github.com/saichler/my.simple/go/utils/logs"
	"os"
	"testing"
	"time"
)

const (
	logFmt = "" +
		"Tr -> trace trace\n" +
		" Dg -> debug debug\n" +
		"  In -> info info\n" +
		"   Wr -> warning warning\n" +
		"    Er -> error error\n"
)

func TestLog(t *testing.T) {
	file, err := os.Create("/tmp/logtest.log")
	if err != nil {
		fmt.Println("Unable to open file:", err)
	}
	oldFile := os.Stdout
	defer func() { os.Stdout = oldFile }()
	defer func() { os.Remove("/tmp/logtest.log") }()
	os.Stdout = file
	logs.Trace("trace", "trace")
	logs.Debug("debug", "debug")
	logs.Info("info", "info")
	logs.Warning("warning", "warning")
	err = logs.Error("error", "error")
	for !logs.Empty() {
		time.Sleep(time.Millisecond * 50)
	}
	os.Stdout = oldFile
	if err.Error() != "    Er -> error error" {
		fmt.Println("'"+err.Error()+"'", "is not", "'    Er -> error error'")
		t.Fail()
		return
	}
	data, err := os.ReadFile("/tmp/logtest.log")
	if string(data) != logFmt {
		fmt.Println("'"+string(data)+"'", "\nis not eq to \n'"+logFmt+"'")
		t.Fail()
	}
}
