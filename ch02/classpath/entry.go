package classpath

import "os"
import "strings"

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	switch {
	case strings.Contains(path, pathListSeparator):
		return newCompositeEntry(path)
	case strings.HasSuffix(path, "*"):
		return newWildcardEntry(path)
	case strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP"):
		return newWildcardEntry(path)
	default:
		return newDirEntry(path)
	}
}
