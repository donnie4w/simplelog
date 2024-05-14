package logging

import (
	"log/slog"
	"path/filepath"
	"strconv"
	"testing"
	"time"
)

func Test_Log(t *testing.T) {
	SetRollingDaily(`D:\cfoldTest`, "log2.txt")
	//控制台打印
	// SetConsole(false)
	Debug("0000000000") //默认格式：[DEBUG]2023/07/10 18:40:49 logging_test.go:12: 00000000000

	SetFormat(FORMAT_NANO) //设置格式(无格式化)：111111111111
	Debug("111111111111")

	SetFormat(FORMAT_LONGFILENAME) //设置格式(长文件路径) ：[INFO]d:/github.com/simplelog/logging/logging_test.go:14: 22222222
	Info("22222222")

	SetFormat(FORMAT_DATE | FORMAT_SHORTFILENAME) //设置格式(日期+短文件路径) ：[WARN]2023/07/10 logging_test.go:16: 333333333
	Warn("333333333")

	SetFormat(FORMAT_DATE | FORMAT_TIME) //设置格式 ：[ERROR]2023/07/10 18:35:19 444444444
	// SetLevel(FATAL) //设置为FATAL后，下面Error()级别小于FATAL,将不打印出来
	Error("444444444")

	SetFormat(FORMAT_SHORTFILENAME) //设置格式 ：[FATAL]logging_test.go:21: 5555555555
	Fatal("5555555555")

	SetFormat(FORMAT_TIME) //设置格式 ：[FATAL]18:35:19 66666666666
	Fatal("66666666666")
}

/*设置日志文件*/
func Test_LogOne(t *testing.T) {
	/*获取全局log单例，单日志文件项目日志建议使用单例*/
	//log := GetStaticLogger()
	/*获取新的log实例，要求不同日志文件时，使用多实例对象*/
	log := NewLogger()
	/*按日期分割日志文件，也是默认设置值*/
	// log.SetRollingDaily(`D:\cfoldTest`, "log.txt")
	log.SetRollingByTime(`D:\cfoldTest`, "log.txt", MODE_HOUR)
	/*按日志文件大小分割日志文件*/
	// log.SetRollingFile("", "log1.txt", 3, KB)
	// log.SetRollingFileLoop(`D:\cfoldTest`, "log1.txt", 3, KB, 5)
	/* 设置打印级别 OFF,DEBUG,INFO,WARN,ERROR,FATAL
	log.SetLevel(OFF) 设置OFF后，将不再打印后面的日志 默认日志级别为ALL，打印级别*/
	/* 日志写入文件时，同时在控制台打印出来，设置为false后将不打印在控制台，默认值true*/
	// log.SetConsole(false)
	log.SetFormat(FORMAT_NANO)
	log.Debug("aaaaaaaaaaaaaaaaaaaaaaaa")
	log.SetFormat(FORMAT_LONGFILENAME) //设置后将打印出文件全部路径信息
	log.Info("bbbbbbbbbbbbbbbbbbbbbbbb")
	log.SetFormat(FORMAT_MICROSECNDS | FORMAT_SHORTFILENAME) //设置日志格式，时间+短文件名
	log.Warn("ccccccccccccccccccccccc")
	log.SetLevel(LEVEL_FATAL) //设置为FATAL后，下面Error()级别小于FATAL,将不打印出来
	log.Error("ddddddddddddddddddddddd")
	log.Fatal("eeeeeeeeeeeeeeeeeeeeeee")
	time.Sleep(2 * time.Second)
}

func TestSerialLog(b *testing.T) {
	// SetRollingFile(`D:\cfoldTest`, "log.txt", 100, KB)
	SetRollingFileLoop(`D:\cfoldTest`, "log.txt", 2000, KB, 10)
	SetGzipOn(true)
	SetConsole(false)
	for i := 0; i < 100000; i++ {
		Debug(i, ">>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
		Info(i, ">>>bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
		Warn(i, ">>>cccccccccccccccccccccccccccccccccccc")
		Error(i, ">>>dddddddddddddddddddddddddddddddddddd")
		Fatal(i, ">>>eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee")
	}
}

func TestTimeLog(b *testing.T) {
	SetRollingByTime(`D:\cfoldTest`, "log.txt", MODE_DAY)
	SetGzipOn(true)
	SetConsole(true)
	for i := 0; i < 1000; i++ {
		Debug(i, ">>>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
		Info(i, ">>>bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
		Warn(i, ">>>cccccccccccccccccccccccccccccccccccc")
		Error(i, ">>>dddddddddddddddddddddddddddddddddddd")
		Fatal(i, ">>>eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee")
	}
}

func TestSlog(t *testing.T) {
	replace := func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.SourceKey {
			source := a.Value.Any().(*slog.Source)
			source.File = filepath.Base(source.File)
		}
		return a
	}
	loggingFile := NewLogger()
	loggingFile.SetRollingFile("./1", "slogfile.txt", 100, KB)
	h := slog.NewJSONHandler(loggingFile, &slog.HandlerOptions{AddSource: true, ReplaceAttr: replace})
	log := slog.New(h)
	for i := 0; i < 1000; i++ {
		log.Info(">>>aaaaaaaaaaaaaaaaaaaaaaa:" + strconv.Itoa(i))
	}

}