package logg

import (
	"testing"
	"time"

	"github.com/akawa1/gbox/errr"
)

func TestLog(t *testing.T) {
	SetConsoleLogger('I', DefaultFormat)
	// SetSingleFileLogger('D', "test.log")
	SetFileLogger('D', "test.log", true, true, 50000, 2)
	err := errr.NewTagError("errtag", "err")
	D("ERROR", "%v", err.Position)
	I("ERROR", err)
	D("AAA", "0")
	I("AAA", "%d\t%s", 10, "hello")
	time.Sleep(time.Second)
}
