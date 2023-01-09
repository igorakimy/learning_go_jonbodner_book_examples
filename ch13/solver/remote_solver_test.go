package solver

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRemoteSolver_Resolve(t *testing.T) {
	type info struct {
		expression string
		code       int
		body       string
	}
	var io info

	// Настройка своей имитации удаленного сервера
	// Функция httptest.NewServer создает и запускает HTTP-сервер на случайно
	// выбранном неиспользуемом порте
	server := httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			expression := req.URL.Query().Get("expression")
			if expression != io.expression {
				rw.WriteHeader(http.StatusBadRequest)
				rw.Write([]byte("invalid expression: " + io.expression))
				return
			}
			rw.WriteHeader(io.code)
			rw.Write([]byte(io.body))
		}))
	// Закрываем сервер после выполнения запроса
	defer server.Close()
	rs := RemoteSolver{
		MathServerURL: server.URL,
		Client:        server.Client(),
	}
	data := []struct {
		name   string
		io     info
		result float64
		errMsg string
	}{
		{
			"case1",
			info{"2 + 2 * 10", http.StatusOK, "22"},
			22,
			"",
		},
		// остальные случаи
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			io = d.io
			result, err := rs.Resolve(context.Background(), d.io.expression)
			if result != d.result {
				t.Errorf("io `%f`, got `%f`", d.result, result)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("io error `%s`, got `%s`", d.errMsg, errMsg)
			}
		})
	}
}
