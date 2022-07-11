package log

import "fmt"

type Logggin interface {
	Info(msg string)
	Infof(format string, v ...interface{})
	Error(msg string)
	Errorf(format string, v ...interface{})
	Warn(msg string)
	Warnf(format string, v ...interface{})
	Debug(msg string)
	Trace(msg string)
}

var Log Logggin

func init() {
	Build()
}

//Realizar a configuração de provider de log de acordo com as variaveis de ambiente
func Build() Logggin {
	fmt.Println("Iniciando logs ***")
	Log = Zerolog{}

	return Log
}
