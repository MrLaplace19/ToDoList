package core_logger

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)
type Logger struct {
	*zap.Logger
	file *os.File
}

func NewLogger(cfg Config) (*Logger, error) {
	zaplvl := zap.NewAtomicLevel()
	if err := zaplvl.UnmarshalText([]byte(cfg.Lvl)); err != nil{
		return nil, fmt.Errorf("Unmarshal log level: %w", err)
	}

	if err :=os.MkdirAll(cfg.Folder, 0755); err != nil{
		return nil, fmt.Errorf("Create log folder: %w", err)
	}

	timestamp := time.Now().UTC().Format("2006-01-2T15-04-05.00000")
	lofFilePath := filepath.Join(
		cfg.Folder,
		fmt.Sprintf("%v.log", timestamp),
	)

	logFile, err := os.OpenFile(lofFilePath,os.O_CREATE | os.O_WRONLY, 0644 )
	if err != nil{
		return nil, fmt.Errorf("Open log file: %w", err)
	}

	zapConfig := zap.NewDevelopmentEncoderConfig()
	zapConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-2T15:04:05.00000")
	zapEncoder := zapcore.NewConsoleEncoder(zapConfig)
	
	core := zapcore.NewTee(
		zapcore.NewCore(zapEncoder, zapcore.AddSync(os.Stdout), zaplvl),
		zapcore.NewCore(zapEncoder, zapcore.AddSync(logFile), zaplvl),
	)
	
	logger := zap.New(core, zap.AddCaller())
	return &Logger{
		Logger: logger,
		file: logFile,
	}, nil
}

func (l *Logger) With(field ...zap.Field) *Logger{
	return &Logger{
		Logger: l.Logger.With(field...),
		file: l.file,
	}
}

func (l *Logger) Close() {
	if err := l.file.Close(); err != nil {
		fmt.Println("Ошибка закрытия лога", err)
	}
}

func FromContext(ctx context.Context) *Logger{
	log, ok := ctx.Value("log").(*Logger)
	if !ok{
		panic("no logger context")
	}

	return log
}