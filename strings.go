package flagext

import (
	"strings"
)

// Strings is a slice of strings useful for accepting multiple option values
type Strings []string

// Set implements flag.Value
func (ss *Strings) Set(val string) error {
	if ss == nil {
		*ss = Strings{}
	}
	*ss = append(*ss, val)
	return nil
}

// String implements flag.Value
func (ss *Strings) String() string {
	if ss == nil {
		return ""
	}
	return strings.Join(*ss, ", ")
}

// Get implements flag.Getter
func (ss *Strings) Get() interface{} {
	if ss == nil {
		return []string(nil)
	}
	return []string(*ss)
}
