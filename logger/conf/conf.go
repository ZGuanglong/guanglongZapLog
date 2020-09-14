package conf

type ZapLogConfig struct {
	LogLevel string	`json:"logLevel"` //日志等级
	LogFileName string	`json:"logFileName"` //日志名称
	LogSectionSize int `json:"logSectionSize"` // 日志分片大小 1<< LogSectionSize
}
