package middlewares

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/pkg/logger"
)

// Security - мидлваре, которая проставляет заголовки безопасности
func Security(handler http.Handler) http.Handler {
	op := "pkg.middlewares.Security"
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Strict-Transport-Security - принудительно переводит на https
		//						 защита от атак типа "downgrade attacks", это где понижают протокол
		// Мы пока не можем :)
		// nginx: add_header Strict-Transport-Security "max-age=63072000; includeSubDomains" always;
		// includeSubDomains - включает ли поддомены
		// max-age - в секундах сколько правило будет выполняться
		// w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

		// CSP - надо бы, но мне тяжело думать

		// X-XSS-Protection - защита от XSS-атак
		// nginx: add_header X-XSS-Protection "1; mode=block";
		// X-XSS-Protection: 1; mode=block включает фильтр и блокирует отображение страницы,
		// 									если обнаружена потенциальная атака XSS.
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		logger.StandardInfoF(r.Context(), op, "Successful set `X-XSS-Protection: 1; mode=block` header request %s %s %s", r.Method, r.URL.Path)

		// X-Frame-Options - защита от угона кликов
		// Работает за счет размещения  нашего сайта во фрейме
		// nginx: add_header X-Frame-Options "DENY";
		// SAMEORIGIN - позволяет загрузку контента в frame/iframe только если фрейм и страница,
		//					его загружающая, расположены на одном домене. т.е. если мы сами себя загружаем
		// 					можно, в целом, и deny, если фронт ничего такого не планирует
		w.Header().Set("X-Frame-Options", "SAMEORIGIN")
		logger.StandardInfoF(r.Context(), op, "Successful set `X-Frame-Options: SAMEORIGIN` header request %s %s %s", r.Method, r.URL.Path)

		// X-Content-Type-Options - защищает от атак подменой MIME типов
		// nginx: add_header X-Content-Type-Options "nosniff";
		// 		Заголовок содержит инструкции по определению типа файла и не допускает сниффинг контента.
		//		При конфигурации потребуется добавить только один параметр: “nosniff”.
		w.Header().Set("X-Content-Type-Options", "nosniff")
		logger.StandardInfoF(r.Context(), op, "Successful set `X-Content-Type-Options: nosniff` header request %s %s %s", r.Method, r.URL.Path)

		handler.ServeHTTP(w, r)
	})
}
