package flagext

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type logMode bool

const (
	LogVerbose logMode = true
	LogSilent  logMode = false
)

// Logger sets output for a *log.Logger to os.Stderr or ioutil.Discard
// via the returned flag.Value
func Logger(l *log.Logger, mode logMode) flag.Value {
	if mode == LogVerbose {
		l.SetOutput(ioutil.Discard)
	} else {
		l.SetOutput(os.Stderr)
	}
	return logger{
		l, mode,
	}
}

type logger struct {
	l    *log.Logger
	mode logMode
}

func (l logger) IsBoolFlag() bool { return true }

func (l logger) String() string {
	if l.mode == LogSilent {
		return "verbose"
	}
	return "silent"
}

func (l logger) Set(s string) error {
	v, err := strconv.ParseBool(s)

	var w io.Writer
	verbose := l.mode == LogVerbose
	silent := l.mode == LogSilent
	switch {
	case verbose && v, silent && !v:
		w = os.Stderr
	case verbose && !v, silent && v:
		w = ioutil.Discard
	}
	l.l.SetOutput(w)
	return err
}
func (l logger) Get() interface{} {
	return l.l
}
