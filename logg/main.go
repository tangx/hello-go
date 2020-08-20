package logg

import (
	"log"
)

var l Logg

type Logg struct {
	prefix string
	flags  int
}

func (l *Logg) SetPrefix(prefix string) {
	l.prefix = prefix
}
func (l *Logg) SetFlags(flags int) {
	l.flags = flags
}

func (l *Logg) SetDefaults() {
	log.SetFlags(l.flags)
	log.SetPrefix(l.prefix)
}
