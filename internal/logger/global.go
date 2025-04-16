package logger

var global Logger

func InitGlobal(lc Logger) {
	global = lc
}

func L() Logger {
	return global
}
