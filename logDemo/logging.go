package logDemo

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func LoggerToFile() {

}

// 一个简单的输出
func NormalLog() {
	log.WithFields(log.Fields{
		"k": "value",
	}).Info("xxx")
}

// 设置日志格式
func SetLogFormatter() {
	// 设置日志输出格式
	log.SetFormatter(&log.JSONFormatter{})
	// 设置日志输出位置
	log.SetOutput(os.Stdout)
	// 只用warn 之上的日志会被输出
	log.SetLevel(log.WarnLevel)
	log.WithFields(log.Fields{
		"k": "value",
	}).Info("xxx")
	log.WithFields(log.Fields{
		"k": "value",
	}).Warn("xxx")
	log.Errorf("xcx %s", "adsf")
}

// 如何使用 log hook

type DefaultFileHook struct {
}

func (d *DefaultFileHook) Levels() []log.Level {
	return []log.Level{log.ErrorLevel}
}

func (d *DefaultFileHook) Fire(entry *log.Entry) error {
	entry.Data["maintainer"] = "ian.liu"
	return nil
}

func UseHook() {
	log.AddHook(&DefaultFileHook{})
	log.Errorf("%s", "error hook user")
	log.Warnf("%s", "warn hook user")
}

// 通过hook 分隔日志
// 不同日志级别单独出来
func NewFileHook() log.Hook {
	writer, err := rotatelogs.New("config"+".%Y%m%d%H",
		// 为日志建立连接
		rotatelogs.WithLinkName("xx"),
		// 每24小时分割一次
		rotatelogs.WithRotationTime(24*time.Hour),
		// 文件最多保留个数
		rotatelogs.WithRotationCount(50),
	)
	if err != nil {
		log.Errorf("config local file sys for logger error: %v", err)
	}
	errWriter, err := rotatelogs.New("error"+".%Y%m%d%H",
		// 为日志建立连接
		rotatelogs.WithLinkName("xlog"),
		// 每24小时分割一次
		rotatelogs.WithRotationTime(24*time.Hour),
		// 文件最多保留个数
		rotatelogs.WithRotationCount(50),
	)
	if err != nil {
		log.Errorf("config local file sys for logger error: %v", err)
	}

	divHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer,
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: errWriter,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.TextFormatter{DisableColors: true})
	return divHook
}

// 调用
func UseDivLogFile() {
	log.AddHook(NewFileHook())
	log.Info("Info")
	log.Debug("Debug")
	log.Warn("Warn")
	log.Error("Error")

}
