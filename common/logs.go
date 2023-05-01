package common

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"os"
	"path"
)

var (
	GlobalLogger, _     = CreateSugarLogger()
	AdministratorLogger = CreateFileLogger(DefaultAdministratorLogFile)
)

func CreateLogger() (*zap.Logger, error) {
	logger, errInit := zap.NewDevelopment()
	if errInit != nil {
		return nil, errInit
	}
	return logger, nil
}

func CreateSugarLogger() (*zap.SugaredLogger, error) {
	logger, errInit := zap.NewDevelopment()
	if errInit != nil {
		return nil, errInit
	}
	return logger.Sugar(), nil
}

func CreateFileLogger(outputFile string) *zap.SugaredLogger {
	outputPath := path.Join(DefaultDataDir(), DefaultLogsDir)
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		err := os.MkdirAll(outputPath, 0700)
		if err != nil {
			panic(err)
		}
	}
	outputPath = path.Join(outputPath, outputFile)
	sampleJSON := []byte(fmt.Sprintf(`{
       "level" : "info",
       "encoding": "json",
       "outputPaths":["stdout", "%s"],
       "errorOutputPaths":["stderr"],
       "encoderConfig": {
           "messageKey":"message",
           "levelKey":"level",
           "levelEncoder":"lowercase"
       }
   }`, outputPath))

	var cfg zap.Config
	if err := json.Unmarshal(sampleJSON, &cfg); err != nil {
		panic(err)
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return logger.Sugar()
}
