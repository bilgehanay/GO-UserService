package main

import (
	"context"
	"net/http"
)

func JWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")
		if token == "" {
			JsonResponse(w, http.StatusUnauthorized, "Invalid Token", "", nil)
			return
		}

		claims, err := VerifyToken(token)
		if err != nil {
			JsonResponse(w, http.StatusUnauthorized, "Invalid Token", "", nil)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "id", claims.ID)))
	}
}

func MethodMiddleware(allowedMethod string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != allowedMethod {
			JsonResponse(w, http.StatusMethodNotAllowed, "Method not allowed", "", nil)
			return
		}
		next.ServeHTTP(w, r)
	}
}
