package flagext

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

// Reader is an io.ReadCloser that can be set as a flag.Value
type Reader interface {
	io.ReadCloser
	flag.Getter
}

const StdIO = "-" // Pass to set Stdin/Stdout as the default

type reader struct {
	buf     *bufio.Reader
	closer  io.Closer
	path    string
	client  *http.Client
	useFile bool
}

func File(defaultPath string) Reader {
	return &reader{path: defaultPath, useFile: true}
}

func URL(defaultPath string, client *http.Client) Reader {
	if client == nil {
		client = http.DefaultClient
	}
	return &reader{path: defaultPath, client: client}
}

func FileOrURL(defaultPath string, client *http.Client) Reader {
	if client == nil {
		client = http.DefaultClient
	}
	return &reader{path: defaultPath, client: client, useFile: true}
}

func (r *reader) Get() interface{} {
	return r
}

func (r *reader) Set(val string) error {
	r.path = val
	if r.client == nil || r.useFile {
		return nil
	}
	if _, err := url.Parse(val); err != nil {
		return fmt.Errorf("bad URL: %v", err)
	}

	return nil
}

func (r *reader) String() string {
	if r.path == StdIO {
		return "stdin"
	}
	return r.path
}

func (r *reader) Read(p []byte) (n int, err error) {
	if r.buf == nil {
		if err := r.init(); err != nil {
			return 0, err
		}
	}

	return r.buf.Read(p)
}

func (r *reader) Close() error {
	if r.closer == nil {
		return nil
	}
	return r.closer.Close()
}

func (r *reader) init() error {
	if r.path == "" {
		return fmt.Errorf("no path set")
	}

	useURL := !r.useFile
	if r.client != nil {
		if u, err := url.Parse(r.path); err == nil && u.Scheme != "" {
			useURL = true
		}
	}

	if useURL {
		resp, err := r.client.Get(r.path)
		if err != nil {
			return err
		}
		r.buf = bufio.NewReader(resp.Body)
		r.closer = resp.Body
		return nil
	}

	if !r.useFile {
		panic("using URL reader with known bad URL")
	}

	if r.path == StdIO {
		r.buf = bufio.NewReader(os.Stdin)
		r.closer = os.Stdin
		return nil
	}

	f, err := os.Open(r.path)
	if err != nil {
		return err
	}
	r.buf = bufio.NewReader(f)
	r.closer = f
	return nil
}
