package flagext_test

import (
	"flag"
	"fmt"
	"strings"

	"github.com/carlmjohnson/flagext"
)

func ExampleMustHave_missingFlag() {
	var buf strings.Builder
	defer func() {
		recover()
		fmt.Println(buf.String())
	}()

	fs := flag.NewFlagSet("ExampleMustHave", flag.PanicOnError)
	fs.SetOutput(&buf)
	fs.String("a", "", "this value must be set")
	fs.String("b", "", "this value must be set")
	fs.String("c", "", "this value is optional")
	fs.Parse([]string{"-a", "set"})
	flagext.MustHave(fs, "a", "b")
	// Output:
	// missing required flag: b
	// Usage of ExampleMustHave:
	//   -a string
	//     	this value must be set
	//   -b string
	//     	this value must be set
	//   -c string
	//     	this value is optional
}

func ExampleMustHave_noMissingFlag() {
	fs := flag.NewFlagSet("ExampleMustHave", flag.PanicOnError)
	fs.String("a", "", "this value must be set")
	fs.String("b", "", "this value must be set")
	fs.String("c", "", "this value is optional")
	fs.Parse([]string{"-a", "set", "-b", "set"})
	flagext.MustHave(fs, "a", "b")
}
