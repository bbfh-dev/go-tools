package terr

import (
	"os"

	"github.com/bbfh-dev/go-tools/tools/tlog"
)

func Assert(condition bool, message string) {
	if condition {
		return
	}

	tlog.Error(
		"ASSERTION FAILED: %s\n[This means that the program reached a state it should have never reached]",
		message,
	)
	os.Exit(1)
}
