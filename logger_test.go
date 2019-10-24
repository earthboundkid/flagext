package flagext_test

import (
    "flag"
    "log"
    "os"

    "github.com/carlmjohnson/flagext"
)

func ExampleLogger() {
    // Redirect Stderr for test
    stderr := os.Stderr
    os.Stderr = os.Stdout

    {
        fs := flag.NewFlagSet("ExampleLogger", flag.ContinueOnError)
        l := log.New(nil, "myapp ", 00)
        fs.Var(
            flagext.Logger(l, flagext.LogVerbose),
            "verbose",
            `log output`,
        )
        if err := fs.Parse([]string{"-verbose"}); err != nil {
            l.Fatal(err)
        }
        l.Print("hello log 1")
    }
    {
        fs := flag.NewFlagSet("ExampleLogger", flag.ContinueOnError)
        l := log.New(nil, "myapp ", 00)
        fs.Var(
            flagext.Logger(l, flagext.LogSilent),
            "silent",
            `don't log output`,
        )
        if err := fs.Parse([]string{"-silent=false"}); err != nil {
            l.Fatal(err)
        }
        l.Print("hello log 2")
    }
    {
        fs := flag.NewFlagSet("ExampleLogger", flag.ContinueOnError)
        l := log.New(nil, "myapp ", 00)
        fs.Var(
            flagext.Logger(l, flagext.LogVerbose),
            "verbose",
            `log output`,
        )
        if err := fs.Parse([]string{"-verbose=false"}); err != nil {
            l.Fatal(err)
        }
        l.Print("silenced!")
    }
    {
        fs := flag.NewFlagSet("ExampleLogger", flag.ContinueOnError)
        l := log.New(nil, "myapp ", 00)
        fs.Var(
            flagext.Logger(l, flagext.LogSilent),
            "silent",
            `don't log output`,
        )
        if err := fs.Parse([]string{"-silent=1"}); err != nil {
            l.Fatal(err)
        }
        l.Print("silenced!")
    }
    os.Stderr = stderr
    // Output:
    // myapp hello log 1
    // myapp hello log 2
}
