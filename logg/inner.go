package logg

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
)

const (
	depth = 3
)

func set() {
	pc, file, line, _ := runtime.Caller(depth)
	shortfile := filepath.Base(file)
	fc := runtime.FuncForPC(pc).Name()

	// pfix := fmt.Sprintf("%s %s:(%d) ", fc, shortfile, line)
	pfix := fmt.Sprintf("%s:%d %s ", shortfile, line, fc)
	l.SetPrefix(pfix)

	log.SetFlags(0)
	log.SetPrefix(pfix)
}

func (l *Logg) println(s string) {
	set()
	log.Println(s)
}

func (l *Logg) printf(formatter string, args ...interface{}) {
	set()
	log.Printf(formatter, args...)
}
