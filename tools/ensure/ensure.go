package ensure

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/bbfh-dev/go-tools/tools/tmap"
)

const MAP_CONTAINS_KEY_FMT = "Map expected to have key '%v' but only had [%s]"
const NOT_NILL_FMT = "Value expected not to be nil: %s"

var OnFail = func(format string, args ...any) {
	fmt.Println("ASSERT FAILED:\n" + fmt.Sprintf(format, args...))
	os.Exit(1)
}

func MapContainsKey[K comparable, V any](in map[K]V, key K) {
	_, ok := in[key]
	if !ok {
		keys := tmap.Flatten(in, func(key K, _ V) string { return fmt.Sprintf("%v", key) })
		sort.Strings(keys)
		OnFail(MAP_CONTAINS_KEY_FMT, key, strings.Join(keys, ", "))
	}
}

func NotNil(in interface{}, reason string) {
	if in == nil {
		OnFail(NOT_NILL_FMT, reason)
	}
}
