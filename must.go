package flagext

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// MustHave is a convenince method that checks that the named flags
// were set on fl. Missing flags are treated with the policy of
// fl.ErrorHandling(): ExitOnError, ContinueOnError, or PanicOnError.
// Returned errors will have type MissingFlags.
//
// If nil, fl defaults to flag.CommandLine.
func MustHave(fl *flag.FlagSet, names ...string) error {
	if fl == nil {
		fl = flag.CommandLine
	}
	seen := make(map[string]bool)
	fl.Visit(func(f *flag.Flag) {
		seen[f.Name] = true
	})
	var missing MissingFlags
	for _, name := range names {
		if !seen[name] {
			missing = append(missing, name)
		}
	}
	if len(missing) == 0 {
		return nil
	}

	fmt.Fprintln(fl.Output(), missing)
	if fl.Usage != nil {
		fl.Usage()
	}
	switch fl.ErrorHandling() {
	case flag.PanicOnError:
		panic(missing.Error())
	case flag.ExitOnError:
		os.Exit(1)
	}
	return missing
}

// MissingFlags is the error type returned by MustHave.
type MissingFlags []string

func (missing MissingFlags) Error() string {
	if len(missing) == 0 {
		return "MissingFlags<empty>"
	}
	if len(missing) == 1 {
		return fmt.Sprintf("missing required flag: %s", missing[0])
	}
	return fmt.Sprintf("missing %d required flags: %s",
		len(missing), strings.Join(missing, ", "))
}
