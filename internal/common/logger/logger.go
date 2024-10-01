package logger

import (
	"fmt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/common/global"
	"log/slog"
	"os"
)

func New() {
	op := "internal.common.logger.logger.New"

	logger := slog.New(slog.NewTextHandler(os.Stdout,
		&slog.HandlerOptions{
			Level: global.LoggerLevel, // Устанавливаем минимальный уровень для логов
		}))

	slog.SetDefault(logger)

	StandardInfo("created logger", op)
}

func Info(msg string, args ...any) {
	slog.Info(msg, args...)
}

func Debug(msg string, args ...any) {
	slog.Info(msg, args...)
}

func Error(msg string, args ...any) {
	slog.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	slog.Info(msg, args...)
}

func StandardInfo(msg string, op string) {
	slog.Info(standardInput(msg, op))
}

func StandardWarn(msg string, op string) {
	slog.Warn(standardInput(msg, op))
}

func StandardError(msg string, op string) {
	slog.Error(standardInput(msg, op))
}

func StandardDebug(msg string, op string) {
	slog.Debug(standardInput(msg, op))
}

func standardInput(msg string, op string) string {
	return fmt.Sprintf("{%v}                 | in %v", msg, op)
}

func StandardResponse(msg string, status int, host string, op string) {
	StandardInfo(fmt.Sprintf("Response sent, status ='%v', message={%v} to host=%v", status, msg, host), op)
}

func StandardSendModel(msg, op string) {
	StandardInfo(fmt.Sprintf("Sent model '%v'", msg), op)
}
