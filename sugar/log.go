package sugar

import (
	"fmt"
	"net/http"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.SugaredLogger

func InitLogger() {
	writeSyncer := getLogWriter()
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)
	var writes = []zapcore.WriteSyncer{zapcore.AddSync(writeSyncer)}
	encoder := getEncoder()
	writes = append(writes, zapcore.AddSync(os.Stdout))
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writes...), zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller(), zap.Development())
	Log = logger.Sugar()
}

func InitLogger1(path string) {
	writeSyncer := getLogWriter1(path)
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)
	var writes = []zapcore.WriteSyncer{zapcore.AddSync(writeSyncer)}
	encoder := getEncoder()
	writes = append(writes, zapcore.AddSync(os.Stdout))
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writes...), zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller(), zap.Development())
	Log = logger.Sugar()
}
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
func getLogWriter() *lumberjack.Logger {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    50,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return lumberJackLogger
}
func getLogWriter1(path string) *lumberjack.Logger {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    50,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return lumberJackLogger
}

//test simple get http
func simpleHttpGet(url string) {
	Log.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		Log.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		Log.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}

func Rebot() {
	fmt.Println("  This is a rebot ~~~~")
}

func CreateFolder() {
	os.Mkdir("abc", os.ModePerm)
}
