package log

type XLogInterface interface {
	NewLogFromConfiguration(c *LogConfiguration) error
	Init(config map[string]string) error
	ReOpen() error
	SetLevel(level string)
	SetSkip(skip int)

	Warn(format string, a ...interface{})
	Fatal(format string, a ...interface{})
	Notice(format string, a ...interface{})
	Trace(format string, a ...interface{})
	Debug(format string, a ...interface{})

	Warnx(logId, format string, a ...interface{})
	Fatalx(logId, format string, a ...interface{})
	Noticex(logId, format string, a ...interface{})
	Tracex(logId, format string, a ...interface{})
	Debugx(logId, format string, a ...interface{})

	Close()
}
