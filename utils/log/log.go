package log

import (
	"fmt"
	"ginexample/config"
	"os"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger        *zap.SugaredLogger
	once          sync.Once
	filePath      string
	errorFilePath string
)

var devOptions = []zap.Option{
	zap.WithCaller(true),
	zap.AddCallerSkip(1),
}

func fullPathEncodeCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(caller.FullPath())
}

func makeDevConfig() zap.Config {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.DisableStacktrace = true
	config.Level = zap.NewAtomicLevelAt(getLogLevel())
	config.EncoderConfig.EncodeCaller = fullPathEncodeCaller
	return config
}

func makeDevConfigWriteOutputToFile() zap.Config {
	zipConfig := zap.NewDevelopmentConfig()
	zipConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zipConfig.DisableStacktrace = true
	zipConfig.Level = zap.NewAtomicLevelAt(getLogLevel())
	zipConfig.EncoderConfig.EncodeCaller = fullPathEncodeCaller
	if config.GetEnv("ERROROUTPUT") != "" {
		zipConfig.Level = zap.NewAtomicLevelAt(zapcore.ErrorLevel) // prints only error level messages
		zipConfig.OutputPaths = []string{errorFilePath}
	} else {
		zipConfig.OutputPaths = []string{filePath}
	}
	return zipConfig
}

var prodOptions = []zap.Option{
	// Add production log settings here(if any)
}

func makeProdConfig() zap.Config {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(getLogLevel())
	return config
}

func getLogLevel() zapcore.Level {
	switch strings.ToLower(config.GetEnv("LOG_LEVEL")) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warning":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

func GetLogger() *zap.SugaredLogger {
	Init()
	return logger
}

func createLogFile() error {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting work directory path")
		return err
	}

	dirPath := wd + "/" + "log_output"

	err = os.MkdirAll(dirPath, 0750)
	if err != nil {
		fmt.Println("error creating directory")
		return err
	}
	filePath = fmt.Sprintf("%v/%v.log", dirPath, config.GetEnv("OUTPUT"))
	errorFilePath = fmt.Sprintf("%v/%v.log", dirPath, config.GetEnv("ERROROUTPUT"))

	if _, err = os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644); err != nil {
		fmt.Println("error opening file, err: ", err)
		return err
	}

	if config.GetEnv("ERROROUTPUT") != "" {
		if _, err = os.OpenFile(errorFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644); err != nil {
			fmt.Println("error opening error file, err: ", err)
			return err
		}
	}

	return nil
}

func makeLogger() (*zap.Logger, error) {
	if strings.ToLower(config.GetEnv("ENV")) == "production" {
		return makeProdConfig().Build(prodOptions...)
	}
	// "local" or "staging"
	if config.GetEnv("OUTPUT") != "" {
		err := createLogFile()
		if err == nil {
			return makeDevConfigWriteOutputToFile().Build(devOptions...)
		}
	}
	return makeDevConfig().Build(devOptions...)
}

func makeProductionLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(getLogLevel())
	return config.Build()
}

func Init() {
	once.Do(func() {
		zap_logger, err := makeLogger()
		if err != nil {
			panic(err)
		}

		logger = zap_logger.Sugar()
		logger.Infof("Initialized logger at log level: %s", zap_logger.Level())
		go func() {
			ticker := time.NewTicker(time.Second)
			for range ticker.C {
				_ = logger.Sync()
			}
		}()
	})
}

func Debugf(template string, args ...interface{}) {
	GetLogger().Debugf("\n"+template, args...)
}

func Infof(template string, args ...interface{}) {
	GetLogger().Infof("\n"+template, args...)
}

func Warnf(template string, args ...interface{}) {
	GetLogger().Warnf("\n"+template, args...)
}

func Errorf(template string, args ...interface{}) {
	GetLogger().Errorf("\n"+template, args...)
}

func Panicf(template string, args ...interface{}) {
	GetLogger().Panicf("\n"+template, args...)
}

func Fatalf(template string, args ...interface{}) {
	GetLogger().Fatalf("\n"+template, args...)
}
