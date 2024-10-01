package terr

import (
	"fmt"
	"strings"
)

func Join(errs ...error) error {
	var messages []string

	for _, err := range errs {
		if err != nil {
			messages = append(messages, err.Error())
		}
	}

	if len(messages) == 0 {
		return nil
	}

	return fmt.Errorf("%d Errors occured:\n%s", len(messages), strings.Join(messages, ";\n"))
}

func Prefix(prefix string, err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("%s: %w", prefix, err)
}
