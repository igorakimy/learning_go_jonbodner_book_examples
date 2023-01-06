package main

import (
	"context"
	"context_values/identity"
	"context_values/tracker"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	bl := BusinessLogic{
		RequestDecorator: tracker.Request,
		Logger:           tracker.Logger{},
		Remote:           "http://www.example.com/query",
	}
	c := Controller{
		Logic: bl,
	}
	router := chi.NewRouter()
	router.Use(tracker.GuidMiddleware, identity.UserMiddleware)
	router.Get("/", c.handleRequest)
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}

type Logic interface {
	businessLogic(ctx context.Context, user string, data string) (string, error)
}

type Logger interface {
	Log(context.Context, string)
}

type RequestDecorator func(*http.Request) *http.Request

type BusinessLogic struct {
	RequestDecorator RequestDecorator
	Logger           Logger
	Remote           string
}

func (bl BusinessLogic) businessLogic(ctx context.Context, user string,
	data string) (string, error) {
	bl.Logger.Log(ctx, "starting businessLogic for "+user+" with "+data)
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		bl.Remote+"?query="+data,
		nil,
	)
	if err != nil {
		bl.Logger.Log(ctx, "error building remote request:"+err.Error())
		return "", err
	}
	req = bl.RequestDecorator(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		bl.Logger.Log(ctx, "failed to get response from client")
		return "", err
	}
	defer resp.Body.Close()
	return "Successful response!", nil
}

type Controller struct {
	Logic Logic
}

func (c Controller) handleRequest(rw http.ResponseWriter, req *http.Request) {
	// получить контекст из экземпляра запроса
	ctx := req.Context()
	// извлечь пользователя из контекста
	user, ok := identity.UserFromContext(ctx)
	if !ok {
		fmt.Println("user not found")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	data := req.URL.Query().Get("data")
	// вызвать бизнес-логику
	result, err := c.Logic.businessLogic(ctx, user, data)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	rw.Write([]byte(result))
}
