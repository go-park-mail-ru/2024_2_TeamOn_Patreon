package middlewares

//
//// CreateMiddlewareWithCommonRepository - создает мидлваре, который использует единый объект репозитория, во всех обработчиках
//// При этом структура бизнес-логики создается новая при каждом вызове
//func CreateMiddlewareWithCommonRepository(repository interfaces.Repository, newBehavior interfaces.NewBehavior) func(next http.Handler) http.Handler {
//	// Middleware для разделения одного репозитория и использования отдельной бизнес логики
//	commonRep := func(next http.Handler) http.Handler {
//		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			op := "internal.common.middleware.CreateMiddlewareWithCommonRepository.commonRep"
//
//			// Используем type switch для проверки типа поведения
//			intBehavior := newBehavior(repository)
//			switch tempBehavior := intBehavior.(type) {
//			case interfaces.Behavior:
//				logger.StandardDebug(
//					fmt.Sprintf("Created new behavior %v", tempBehavior),
//					op,
//				)
//
//				// Создаем новый контекст с данными
//				ctx := context.WithValue(r.Context(), global.BehaviorKey, tempBehavior)
//
//				// Передаем новый контекст дальше в цепочке
//				r = r.WithContext(ctx)
//				logger.StandardDebug(
//					fmt.Sprintf("Added new behavior %v in ctx", tempBehavior),
//					op,
//				)
//			default:
//				logger.StandardError("Unknown behavior type", op)
//			}
//
//			// Вызываем следующий обработчик
//			next.ServeHTTP(w, r)
//		})
//	}
//
//	return commonRep
//}
