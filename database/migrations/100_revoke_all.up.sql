--— Запретить доступ всем пользователям к схеме public
REVOKE ALL ON SCHEMA public FROM public;

-- Запретить доступ ко всем таблицам в схеме public
REVOKE ALL ON ALL TABLES IN SCHEMA public FROM public;
