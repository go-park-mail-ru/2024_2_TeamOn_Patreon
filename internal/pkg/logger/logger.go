package logger

import (
	"fmt"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"log/slog"
	"os"
)

func New() {
	op := "internal.pkg.logger.logger.New"

	logger := slog.New(slog.NewTextHandler(os.Stdout,
		&slog.HandlerOptions{
			Level: global.LoggerLevel, // Устанавливаем минимальный уровень для логов
		}))

	slog.SetDefault(logger)

	StandardInfo("created logger", op)
}

func StandardInfo(msg string, op string) {
	slog.Info(standardInput(msg, op))
}

func StandardInfoF(op string, format string, a ...any) {
	StandardInfo(fmt.Sprintf(format, a...), op)
}

func StandardWarn(msg string, op string) {
	slog.Warn(standardInput(msg, op))
}

func StandardWarnF(op string, format string, a ...any) {
	StandardWarn(fmt.Sprintf(format, a...), op)
}

func StandardError(msg string, op string) {
	slog.Error(standardInput(msg, op))
}

func Debug(msg string, args ...any) {
	slog.Debug(msg, args...)
}

func StandardDebug(msg string, op string) {
	Debug(standardInput(msg, op))
}

func StandardDebugF(op string, msg string, a ...any) {
	StandardDebug(op, fmt.Sprintf(msg, a...))
}

func standardInput(msg string, op string) string {
	return fmt.Sprintf("{%v}                 | in %v", msg, op)
}

func StandardResponse(msg string, status int, host string, op string) {
	StandardInfoF(op, "Response sent, status ='%v', message={%v} to host=%v", status, msg, host)
}

func StandardSendModel(msg, op string) {
	StandardInfoF(op, "Sent model '%v'", msg)
}
