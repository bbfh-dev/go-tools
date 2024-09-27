package tools

import (
	"fmt"
	"log/slog"
	"os"
)

func AnyErr(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}

func PrefixErr(prefix string, err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("%s: %w", prefix, err)
}

func Assert(condition bool, errMessage string) {
	if condition {
		return
	}

	slog.Error("[ASSERTION FAILED]", "err", errMessage)
	os.Exit(2)
}
