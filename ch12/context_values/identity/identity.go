package identity

import (
	"context"
	"net/http"
)

type userKey int

const key userKey = 1

// Имя функции, создающей контекст со значением, должно начинаться с ContextWith

func ContextWithUser(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, key, user)
}

// Имя функции, возвращающей значение из контекста, должно оканчиваться на FromContext

func UserFromContext(ctx context.Context) (string, bool) {
	user, ok := ctx.Value(key).(string)
	return user, ok
}

// В реальной реализации следует использовать подпись, чтобы исключить возможность
// подделки идентификатора пользователя

func extractUser(req *http.Request) (string, error) {
	userCookie, err := req.Cookie("user")
	if err != nil {
		return "", err
	}
	return userCookie.Value, nil
}

func UserMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// получить значение пользователя
		user, err := extractUser(req)
		if err != nil {
			// если пользователь не найден, будет возвращен код ошибки 401
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}
		// извлекаем контекст из запроса
		ctx := req.Context()
		// создаем новый контекст со значением пользователя с помощью функции ContextWithUser
		ctx = ContextWithUser(ctx, user)
		// создаем новый запрос на основе старого запроса и нового контекста,
		// используя метод WithContext
		req = req.WithContext(ctx)
		// вызываем следующую функцию в цепи обработчиков, передавая ей новый запрос
		// и полученный в качестве параметра экземпляр типа http.ResponseWriter
		h.ServeHTTP(rw, req)
	})
}
