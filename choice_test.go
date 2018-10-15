package flagext_test

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/carlmjohnson/flagext"
)

func ExampleChoice() {
	// Bad flag
	{
		fs := flag.NewFlagSet("ExampleChoice", flag.ContinueOnError)
		fs.SetOutput(ioutil.Discard)
		var mode string
		fs.Var(flagext.Choice(&mode, "a", "b"), "mode", "mode to run")

		err := fs.Parse([]string{"-mode", "c"})
		fmt.Println(err) // produces error
	}
	// Good flag
	{
		fs := flag.NewFlagSet("ExampleChoice", flag.ContinueOnError)
		var mode string
		fs.Var(flagext.Choice(&mode, "x", "y"), "mode", "mode to run")

		fs.Parse([]string{"-mode", "x"})
		fmt.Println(mode) // mode is x
	}
	// Default value
	{
		fs := flag.NewFlagSet("ExampleChoice", flag.ContinueOnError)
		mode := "none"
		fs.Var(flagext.Choice(&mode, "yes", "no"), "mode", "mode to run")
		fs.Parse([]string{})
		fmt.Println(mode) // mode is none
	}
	// Output:
	// invalid value "c" for flag -mode: "c" not in a, b
	// x
	// none
}
