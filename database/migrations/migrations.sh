# ----- Скрипт не используется -----
# Поднимает миграции фалй Docker-migrate, уже всключенный в docker-compose.yml
# Данный файл необходимо удалить

for file in /docker-entrypoint-initdb.d/*.sql; do
 if [[ -f "$file" ]]; then
   psql -U admin testdb < "$file"
 else
   echo "Файл не найден: $file"
 fi
done