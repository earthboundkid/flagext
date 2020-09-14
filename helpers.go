package flagext

import "flag"

func listVisitedFlagNames(fl *flag.FlagSet) map[string]bool {
	seen := make(map[string]bool)
	fl.Visit(func(f *flag.Flag) {
		seen[f.Name] = true
	})
	return seen
}

func flagOrDefault(fl *flag.FlagSet) *flag.FlagSet {
	if fl == nil {
		return flag.CommandLine
	}
	return fl
}
