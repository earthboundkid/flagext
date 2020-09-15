package flagext_test

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/carlmjohnson/flagext"
)

func ExampleParseEnv() {
	fs := flag.NewFlagSet("ExampleParseEnv", flag.PanicOnError)
	a := fs.Int("a", 0, "")
	b := fs.Int("b", 0, "")
	fs.Parse([]string{"-a", "1"})

	os.Setenv("TEST_ENV_A", "2")
	os.Setenv("TEST_ENV_B", "3")
	flagext.ParseEnv(fs, "test-env")

	// Does not override existing values
	fmt.Println("a", *a)
	// Does get new values from env
	fmt.Println("b", *b)
	// Output:
	// a 1
	// b 3
}

func assert(t *testing.T, assertion bool, format string, args ...interface{}) {
	t.Helper()
	if !assertion {
		t.Fatalf(format, args...)
	}
}

func assertNil(t *testing.T, err error) {
	t.Helper()
	assert(t, err == nil, "want nil error; got: %v", err)
}

func TestParseEnv(t *testing.T) {
	// Don't override
	fs := flag.NewFlagSet("ExampleParseEnv", flag.ContinueOnError)
	var buf strings.Builder
	fs.SetOutput(&buf)
	a := fs.Int("a", 0, "")
	err := fs.Parse([]string{"-a", "1"})
	assertNil(t, err)
	assert(t, *a == 1, "expected a == 1; got %d", *a)
	// Does not override existing values
	os.Setenv("TEST_ENV_A", "y")
	err = flagext.ParseEnv(fs, "TEST_ENV")
	assertNil(t, err)
	output := buf.String()
	assert(t, output == "", "expected no output; got %q", output)

	// Convert kebabs
	fs = flag.NewFlagSet("ExampleParseEnv", flag.ContinueOnError)
	buf.Reset()
	fs.SetOutput(&buf)
	kebab := fs.Int("a-b-c", 0, "")
	os.Setenv("TEST_ENV_A_B_C", "1")
	err = flagext.ParseEnv(fs, "TEST_ENV")
	assertNil(t, err)
	assert(t, *kebab == 1, "expected kebab == 1; got %d", *kebab)
	output = buf.String()
	assert(t, output == "", "expected no output; got %q", output)

	// With error
	fs = flag.NewFlagSet("ExampleParseEnv", flag.ContinueOnError)
	buf.Reset()
	fs.SetOutput(&buf)
	b := fs.Int("b", 0, "")
	err = fs.Parse(nil)
	assertNil(t, err)
	assert(t, *b == 0, "expected b == 0; got %d", *b)
	os.Setenv("TEST_ENV_B", "y")
	err = flagext.ParseEnv(fs, "TEST_ENV")
	assert(t, err != nil, "expected err; got nil")
	output = buf.String()
	expected := "invalid value \"y\" for flag -b: parse error\nUsage of ExampleParseEnv:\n  -b int\n    \t\n"
	assert(t, output == expected, "expected usage message; got %q", output)
}
