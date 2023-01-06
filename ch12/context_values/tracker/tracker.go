package tracker

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type guidKey int

const gKey guidKey = 1

func contextWithGUID(ctx context.Context, guid string) context.Context {
	return context.WithValue(ctx, gKey, guid)
}

func guidFromContext(ctx context.Context) (string, bool) {
	g, ok := ctx.Value(gKey).(string)
	return g, ok
}

func GuidMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		if guid := req.Header.Get("X-GUID"); guid != "" {
			ctx = contextWithGUID(ctx, guid)
		} else {
			ctx = contextWithGUID(ctx, uuid.New().String())
		}
		req = req.WithContext(ctx)
		h.ServeHTTP(rw, req)
	})
}

type Logger struct{}

func (Logger) Log(ctx context.Context, message string) {
	if guid, ok := guidFromContext(ctx); ok {
		message = fmt.Sprintf("GUID: %s - %s", guid, message)
	}
	// выполняем журналирование
	fmt.Println(message)
}

func Request(req *http.Request) *http.Request {
	ctx := req.Context()
	if guid, ok := guidFromContext(ctx); ok {
		req.Header.Add("X-GUID", guid)
	}
	return req
}
