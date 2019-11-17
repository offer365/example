package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	r := NewRouter()
	r.Use(logger) // 最外层
	// r.Use(timeout)
	r.Use(timeMiddleware)
	r.Add("/", http.HandlerFunc(helloHandler))
}

func helloHandler(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("hello"))
}

type middleware func(http.Handler) http.Handler

type Router struct {
	middlewareChain []middleware
	mux             map[string]http.Handler
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Use(m middleware) {
	r.middlewareChain = append(r.middlewareChain, m)
}

func (r *Router) Add(route string, h http.Handler) {
	var mergedHandler = h

	for i := len(r.middlewareChain) - 1; i >= 0; i-- {
		mergedHandler = r.middlewareChain[i](mergedHandler)
	}

	r.mux[route] = mergedHandler
}

func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()

		// next handler
		next.ServeHTTP(wr, r)

		timeElapsed := time.Since(timeStart)
		fmt.Println(timeElapsed)
	})
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RemoteAddr)
		// next handler
		next.ServeHTTP(wr, r)
	})
}
