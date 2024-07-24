package main

import "fmt"

type Logger interface {
	Log(m string)
}

func mostrarLog(l Logger, m string) {
	l.Log(m)
}

type ConsoleLogger struct {
}

func (c ConsoleLogger) Log(m string) {
	fmt.Println("Log no console:", m)
}

type FileLogger struct {
}

func (f FileLogger) Log(m string) {
	fmt.Println("Log no arquivo:", m)
}

func main() {
	c := ConsoleLogger{}
	f := FileLogger{}

	qualquerLogAleatoria := "Kernel panic - Fatal exception"

	fmt.Println("Logs:")
	mostrarLog(c, qualquerLogAleatoria)
	mostrarLog(f, qualquerLogAleatoria)
}
