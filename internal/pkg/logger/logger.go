package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
)

func New() {
	op := "internal.pkg.logger.logger.New"

	logger := slog.New(slog.NewTextHandler(os.Stdout,
		&slog.HandlerOptions{
			Level: global.LoggerLevel, // Устанавливаем минимальный уровень для логов
		}))

	slog.SetDefault(logger)

	StandardInfo(context.Background(), "created logger", op)
}

func StandardInfo(ctx context.Context, msg string, op string) {
	reqID := extractReqID(ctx)
	slog.Info(standardInput(msg, op, reqID))
}

func StandardInfoF(ctx context.Context, op string, format string, a ...any) {
	StandardInfo(ctx, fmt.Sprintf(format, a...), op)
}

func StandardWarn(ctx context.Context, msg string, op string) {
	reqID := extractReqID(ctx)
	slog.Warn(standardInput(msg, op, reqID))
}

func StandardWarnF(ctx context.Context, op string, format string, a ...any) {
	StandardWarn(ctx, fmt.Sprintf(format, a...), op)
}

func StandardError(ctx context.Context, msg string, op string) {
	reqID := extractReqID(ctx)
	slog.Error(standardInput(msg, op, reqID))
}

func Debug(ctx context.Context, msg string, args ...any) {
	reqID := extractReqID(ctx)
	slog.Debug(fmt.Sprintf("[ReqID: %s] %s", reqID, msg), args...)
}

func StandardDebug(ctx context.Context, op string, msg string) {
	reqID := extractReqID(ctx)
	Debug(ctx, standardInput(msg, op, reqID))
}

func StandardDebugF(ctx context.Context, op string, msg string, a ...any) {
	StandardDebug(ctx, op, fmt.Sprintf(msg, a...))
}

func standardInput(msg string, op string, reqID string) string {
	return fmt.Sprintf("[ReqID: %s] {%v}                 | in %v", reqID, msg, op)
}

func StandardResponse(ctx context.Context, msg string, status int, host string, op string) {
	StandardInfoF(ctx, op, "Response sent, status ='%v', message={%v} to host=%v", status, msg, host)
}

func StandardSendModel(ctx context.Context, msg, op string) {
	StandardInfoF(ctx, op, "Sent model '%v'", msg)
}

// Функция для извлечения reqID из контекста
func extractReqID(ctx context.Context) string {
	if ctx == nil {
		return "unknown"
	}
	reqID, ok := ctx.Value("request_id").(string)
	if !ok {
		return "unknown"
	}
	return reqID
}
