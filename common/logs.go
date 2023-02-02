package common

import "go.uber.org/zap"

var (
	GlobalLogger, _ = CreateSugarLogger()
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
