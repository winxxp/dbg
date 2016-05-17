package dbg

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func Printf(format string, a ...interface{}) {
	buf := header(0)
	fmt.Fprintf(buf, format, a...)
	io.Copy(os.Stdout, buf)
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Println(a ...interface{}) {
	buf := header(0)
	fmt.Fprintln(buf, a...)
	io.Copy(os.Stdout, buf)
}

func header(depth int) *bytes.Buffer {
	fn := "???"
	pc, file, line, ok := runtime.Caller(2 + depth)
	if !ok {
		file = "???"
		line = 1
	} else {
		p := runtime.FuncForPC(pc)
		fn = p.Name()
		slash := strings.LastIndex(fn, ".")
		if slash >= 0 {
			fn = fn[slash+1:]
		}
		slash = strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}

	buf := &bytes.Buffer{}

	fmt.Fprintf(buf, "[%d-%s-%s-%d] ", time.Now().Unix(), file, fn, line)
	return buf
}
