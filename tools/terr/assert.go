package terr

import (
	"os"

	"github.com/bbfh-dev/go-tools/tools/tlog"
)

func Assert(condition bool, message string) {
	if condition {
		return
	}

	tlog.Error("[ The program reached a state it should have never reached ]")
	tlog.Error("ASSERTION FAILED: %s\n", message)
	os.Exit(1)
}
