package social

import (
	"fmt"
	"go.uber.org/zap"
)

type SocialLogger struct {
	internalLogger *zap.SugaredLogger
}

func (s *SocialLogger) Named(name string) *SocialLogger {
	return &SocialLogger{s.internalLogger.Named(name)}
}

func (s *SocialLogger) ContextLogger(reqId string, methoName string) *SocialLogger {
	return &SocialLogger{s.internalLogger.With("reqId", reqId, "method", methoName)}
}

func (s *SocialLogger) MethodLogger(methodName string) *SocialLogger {
	return &SocialLogger{s.internalLogger.With("method", methodName)}
}

func (s *SocialLogger) Debug(args ...interface{}) {
	s.internalLogger.Debug(args)
}
func (s *SocialLogger) Debugf(template string, args ...interface{}) {
	s.internalLogger.Debugf(template, args)
}

func (s *SocialLogger) Info(args ...interface{}) {
	s.internalLogger.Info(args)
}
func (s *SocialLogger) Infof(template string, args ...interface{}) {
	s.internalLogger.Infof(template, args)
}

func (s *SocialLogger) Warn(args ...interface{}) {
	s.internalLogger.Warn(args)
}

func (s *SocialLogger) Error(args ...interface{}) {
	s.internalLogger.Error(args)
}
func (s *SocialLogger) Errorf(template string, args ...interface{}) {
	s.internalLogger.Errorf(template, args)
}

func (s *SocialLogger) DPanic(args ...interface{}) {
	s.internalLogger.DPanic(args)
}

func (s *SocialLogger) Panic(args ...interface{}) {
	s.internalLogger.Panic(args)
}

func (s *SocialLogger) DPanicf(template string, args ...interface{}) {
	s.internalLogger.DPanic(template, args)
}

func (s *SocialLogger) Panicf(template string, args ...interface{}) {
	s.internalLogger.Panicf(template, args)
}

func (s *SocialLogger) Sync() error {
	return s.internalLogger.Sync()
}

func InitLogger(cfg *zap.Config) *SocialLogger {
	logger, err := cfg.Build()
	if err != nil {
		panic(fmt.Sprintf("Logger config is corrupted. Error: %s", err.Error()))
	}

	return &SocialLogger{logger.Sugar()}
}
