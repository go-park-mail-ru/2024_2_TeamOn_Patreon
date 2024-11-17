## MOCK тесты
### Установка зависимостей

Устанавливаем библиотеку gomock и утилиту mockgen:

``` bash
go get github.com/golang/mock/gomock
go install github.com/golang/mock/mockgen
```

[//]: # (Директория $GOPATH/bin должна в PATH, чтобы использовать mockgen.)

### Определяем интерфейс


Должен быть интерфейс внешнего для тестируемого модуля сервиса

``` go
package package

type DataStore interface {
    GetData(id int) (string, error)
}
```

### Генерируем мок

С помощью mockgen можно сгенерировать код мока для интерфейса DataStore. Вызовите следующую команду:
``` bash
mockgen -source=path/to/interface.go -destination=path/to/res_mock.go -package=mock_interfaces
```
Эта команда создаст файл mock_datastore.go с автоматически сгенерированным мок-объектом для интерфейса DataStore.

### Пишем тест с использованием gomock

Пример теста с использованием библиотеки gomock может выглядеть так:

``` go
package mypackage_test

import (
"testing"

    "github.com/golang/mock/gomock"
    "github.com/stretchr/testify/assert"
    "mypackage"
    "mypackage/mocks" // Путь к сгенерированным мокам
)

func TestGetDataWithMock(t *testing.T) {
// Создаем контроллер для управления жизненным циклом мока
ctrl := gomock.NewController(t)
defer ctrl.Finish()

    // Создаем мок объекта DataStore
    mockDataStore := mocks.NewMockDataStore(ctrl)

    // Определяем поведение мока
    mockDataStore.
        EXPECT().
        GetData(1).
        Return("mock data", nil)

    // Тестируем функцию, которая использует DataStore
    result, err := mockDataStore.GetData(1)
    assert.NoError(t, err)
    assert.Equal(t, "mock data", result)
}
```
Объяснение кода

    gomock.NewController(t) — создает контроллер для управления жизненным циклом мока. Он отслеживает вызовы методов и проверяет, что все ожидаемые вызовы были выполнены.
    defer ctrl.Finish() — завершает работу контроллера и проверяет, что все ожидания были удовлетворены.
    mockDataStore.EXPECT() — определяет ожидание вызова метода. В этом случае мы ожидаем, что метод GetData будет вызван с аргументом 1 и вернет "mock data" и nil в качестве ошибки.

