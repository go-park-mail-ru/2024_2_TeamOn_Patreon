package static

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	global "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/global"
	"github.com/pkg/errors"
)

// ExtractFileFromMultipart извлекает файл и его MIME-тип из тела запроса
// В качестве аргументов принимает request и названия поля с файлом
// Пример названия поля: "file"
func ExtractFileFromMultipart(r *http.Request, fieldName string) (multipart.File, string, error) {
	op := "internal.pkg.static.ExtractFileFromMultipart"

	// Получаем файл из multipart form
	file, fileHeader, err := r.FormFile(fieldName)
	if err != nil {
		return nil, "", errors.Wrap(err, op)
	}

	// Получаем MIME-тип файла
	contentType := fileHeader.Header.Get("Content-Type")
	return file, contentType, nil
}

// GetFileExtension извлекает расширение файла из его MIME-типа
// В качестве аргумента принимает MIME-тип
func GetFileExtension(contentType string) (string, error) {
	// Создаем соответствие между MIME-типами и расширениями файлов
	mimeToExt := map[string]string{
		"image/jpeg": ".jpeg",
		"image/jpg":  ".jpg",
		"image/png":  ".png",
		// "image/gif":  ".gif",  // может быть, когда-нибудь
		"audio/mpeg":      ".mp3",
		"video/mp4":       ".mp4",
		"application/pdf": ".pdg",
	}

	// Возвращаем расширение, если MIME-тип найден в карте
	if extension, exists := mimeToExt[contentType]; exists {
		return extension, nil
	}
	return "", global.ErrInvalidFileFormat // Возвращаем пустую строку и 415, если MIME-тип не найден
}

// ConvertMultipartToBytes конвертирует multipart.File в []byte
func ConvertMultipartToBytes(file multipart.File) ([]byte, error) {
	op := "internal.pkg.static.ConvertMultipartToBytes"

	// Читаем содержимое файла в []byte
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		// logger.StandardDebugF(ctx, op, "unable to read file {%v}", err)
		return nil, errors.Wrap(err, op)
	}
	return fileBytes, nil
}

// TODO: Проверить на максимально допустимый размер

// CreateFilePath возвращает созданный путь к файлу для дальнейшего сохранения в БД
// В качестве аргументов принимает путь к папке сохранения и формат файла contentType
// Пример аргументов: ("./static/avatar", ".jpg")
func CreateFilePath(folderPath, fileID, extension string) string {
	// Название файла
	fileName := fileID + extension

	// Путь к файлу
	return filepath.Join(folderPath, fileName)
}

// SaveFile сохраняет файл в файловой системе
// В качестве аргументов принимает байты и путь сохранения
func SaveFile(file []byte, filePath string) error {
	op := "internal.pkg.static.SaveFile"

	// Создание файла
	out, err := os.Create(filePath)
	if err != nil {
		return errors.Wrap(err, op)
	}
	defer out.Close() // Закрываем файл после завершения работы

	// Запись данных в файл
	_, err = out.Write(file)
	if err != nil {
		return err
	}
	return nil
}

// DeleteFile удаляет файл из файловой системы
// В качестве аргументов принимает путь к файлу
func DeleteFile(filePath string) error {
	op := "internal.pkg.static.DeleteFile"

	// Удаление файла
	err := os.Remove(filePath)
	if err != nil {
		return errors.Wrap(err, op)
	}
	return nil
}

// ReadFile возвращает файл из файловой системы
// В качестве аргумента принимает путь к файлу
func ReadFile(filePath string) ([]byte, error) {
	op := "internal.pkg.static.ReadFile"

	// Чтение содержимого файла
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	return data, nil
}
