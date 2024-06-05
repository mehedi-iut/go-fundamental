//package middlewares
//
//import (
//	"context"
//	"example.com/rest-api/utils"
//	"log"
//	"net/http"
//)
//
//// Middleware type as before
//type Middleware func(http.Handler) http.Handler
//
//// App struct to hold our routes and middleware
//type App struct {
//	mux         *http.ServeMux
//	middlewares []Middleware
//}
//
//// NewApp creates and returns a new App with an initialized ServeMux and middleware slice
//func NewApp() *App {
//	return &App{
//		mux:         http.NewServeMux(),
//		middlewares: []Middleware{},
//	}
//}
//
//// Use adds middleware to the chain
//func (a *App) Use(mw Middleware) {
//	a.middlewares = append(a.middlewares, mw)
//}
//
//// Handle registers a handler for a specific route, applying all middleware
//func (a *App) Handle(pattern string, handler http.Handler) {
//	finalHandler := handler
//	for i := len(a.middlewares) - 1; i >= 0; i-- {
//		finalHandler = a.middlewares[i](finalHandler)
//	}
//	a.mux.Handle(pattern, finalHandler)
//}
//
//// ListenAndServe starts the application server
//func (a *App) ListenAndServe(address string) error {
//	return http.ListenAndServe(address, a.mux)
//}
//
//// Example middleware for demonstration purposes
//func LoggingMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		log.Printf("Request: %s %s", r.Method, r.URL.Path)
//		next.ServeHTTP(w, r)
//	})
//}
//
//func Authenticate(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		token := r.Header.Get("Authorization")
//
//		if token == "" {
//			http.Error(w, "Not authorized", http.StatusUnauthorized)
//			return
//		}
//
//		userId, err := utils.VerifyToken(token)
//		if err != nil {
//			http.Error(w, "Not authorized", http.StatusUnauthorized)
//			return
//		}
//
//		// Set user ID in context (optional, depending on your application logic)
//		ctx := context.WithValue(r.Context(), "userId", userId)
//		r = r.WithContext(ctx)
//		next.ServeHTTP(w, r)
//	})
//}

//func EnsureProtected(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//
//	})
//}

package middlewares

import (
	"context"
	"example.com/rest-api/utils"
	"net/http"
)

//type Middleware func(http.Handler) http.Handler

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		userId, err := utils.VerifyToken(token)
		if err != nil {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		// Set user ID in context (optional, depending on your application logic)
		ctx := context.WithValue(r.Context(), "userId", userId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)

		// Call the next handler in the chain
		//next.ServeHTTP(w, r)
	})
}
