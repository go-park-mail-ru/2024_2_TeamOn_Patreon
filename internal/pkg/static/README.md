
### Сценарий сохранения (обновления) файла

* #### Controller

1. Получить файл и его MIME-тип из тела запроса с помощью **ExtractFileFromMultipart()**
2. Из MIME-типа извлечь расширение файла через **GetFileExtension()**
3. Используя **ConvertMultipartToBytes()** конвертировать multipart в []byte. 
	*Замечание: multipart дальше controller не уходит, иначе будет нарушение чистой архитектуры.*
4. В service отправить сам файл *file* в []byte и его расширение *extension* в формате string: ".jpg", например.

* #### Service

Далее необходимо сохранить файл в файловой системе и указать путь сохранения В БД. 

Поскольку service не должен знать ни про файловую систему, ни про БД, то *file* и *extension* передать сразу в **repository**. 

* #### Repository

1. Сформировать путь сохранения файла *filePath* (URL) через **CreateFilePath()**. На вход передать путь к директории хранения статики, сгенерированный ID файла и расширение файла. 
	`Пример аргументов: ("./static/avatar", ".jpg")`
2. Сохранить файл на диске с помощью **SaveFile()**. На вход передать *file* в []byte и *filePath*.
3. Создать запись в БД.


### Сценарий удаления файла

* #### Service

1. Обратиться в repository для удаления файла

* #### Repository

1. Получить путь до удаляемого файла *filePath*
2. Удалить запись в БД
3. Удалить файл с хранилища, используя **DeleteFile()**. На вход передать *filePath*.


### Сценарий получения файла

* #### Controller

1. Обратиться в service для получения файла
2. После получения файла. **Важно!** Проставить заголовки: указать MIME-тип и название файла. Пример для получения изображения:
```
w.Header().Set("Content-Type", "image/jpeg")
w.Header().Set("Content-Disposition", "attachment; filename=\"avatar.jpg\"")
```
3. Положить файл в response:
```
w.Write(avatar)
```

* #### Service

1. Обратиться в repository для получения искомого файла *fileBytes* 
2. Отправить файл в controller.

* ### Repository

1. Получить *filePath* из БД. 
2. Получить файл в []byte по *filePath* с помощью **ReadFile()**. На вход передать *filePath*.
3. Отправить файл в service.