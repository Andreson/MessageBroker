package log

//camada wrapper apra deixar mais resumida a chamada dos logs
func Info(msg string) {
	Log.Info(msg)
}

func Infof(msg string, i ...interface{}) {
	Log.Infof(msg, i)
}

func Error(msg string) {
	Log.Error(msg)
}
func Errorf(msg string, i ...interface{}) {
	Log.Errorf(msg, i)
}

func Warn(msg string) {
	Log.Warn(msg)
}

func Warnf(msg string, i ...interface{}) {
	Log.Warnf(msg, i)
}
func Debug(msg string) {
	Log.Debug(msg)
}
func Trace(msg string) {
	Log.Trace(msg)
}
