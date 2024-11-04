package config

import (
	"bufio"
	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
	"github.com/pkg/errors"
	"os"
	"strings"
)

func GetEnv(key, def string) string {
	value := os.Getenv(key)
	if value == "" {
		return def
	}
	return value
}

func InitEnv(pathToCommonEnv string, pathToServiceEnv string) {
	err := InitEnvErr(pathToCommonEnv, pathToServiceEnv)
	if err != nil {
		panic(err)
	}
}

func InitEnvErr(pathToCommonEnv string, pathToServiceEnv string) error {
	op := "pkg.global.env.InitEnv"

	wd, _ := os.Getwd()
	logger.StandardDebugF(op, "Current working directory:", wd)

	// Достаем из окружения информацию в докере ли мы
	key := "IN_DOCKER"

	inDocker := os.Getenv(key)
	if inDocker == "true" {
		// В таком случае переменные окружения уже экспортированы докеркомпозом
		return nil
	}

	err := initEnv(pathToCommonEnv)
	if err != nil {
		return errors.Wrap(err, op)
	}

	if pathToServiceEnv != "" {
		err = initEnv(pathToServiceEnv)
		if err != nil {
			return errors.Wrap(err, op)
		}
	}

	// Пример вывода переменных окружения
	logger.StandardDebugF(op, "Переменные окружения установлены:")

	for _, key := range []string{"STATUS", "IN_DOCKER", "DB_HOST", "SERVICE_NAME", "PORT"} {
		logger.StandardDebugF(op, "%s=%s", key, os.Getenv(key))
	}

	return nil
}

func initEnv(pathToEnv string) error {
	op := "pkg.global.env.initEnv"

	// Открываем файл с переменными окружения
	file, err := os.Open(pathToEnv)

	if err != nil {
		logger.StandardDebugF(op, "Didn't open env file, err=%s", err.Error())
		return errors.Wrap(err, op)

	}

	defer file.Close()

	// Считываем файл построчно
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Игнорируем пустые строки и комментарии
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Разделяем строку на ключ и значение
		parts := strings.SplitN(line, "=", 2)

		if len(parts) != 2 {
			logger.StandardDebugF(op, "Неверный формат строки: %v", line)
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// value = strings.TrimSuffix(value, "\n")

		// Устанавливаем переменную окружения
		err := os.Setenv(key, value)

		if err != nil {
			logger.StandardDebugF(op, "Ошибка при установке переменной окружения: %v", err)
			return errors.Wrap(err, op)
		}

	}

	if err := scanner.Err(); err != nil {
		logger.StandardDebugF(op, "Ошибка при чтении файла: %v", err)
		return errors.Wrap(err, op)

	}

	return nil
}
