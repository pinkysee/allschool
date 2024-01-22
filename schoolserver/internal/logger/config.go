package logger

type Log_config struct {
	Log_level string `toml:"log_level"`
}

func NewConfig() *Log_config {
	return &Log_config{
		Log_level: "debug",
	}
}
