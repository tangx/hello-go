package logg

func Println(s string) {
	l.println(s)
}

func Printf(formatter string, args ...interface{}) {
	l.printf(formatter, args...)
}
