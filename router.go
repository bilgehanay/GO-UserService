package main

import "net/http"

func RouterGroup() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/api/v1/user/", userMux())
	mux.Handle("/api/v1/auth/", authMux())

	return mux
}

func authMux() http.Handler {
	authMux := http.NewServeMux()

	authMux.HandleFunc("/signup", MethodMiddleware("POST", SignUp))
	authMux.HandleFunc("/login", MethodMiddleware("POST", Login))

	return http.StripPrefix("/api/v1/auth", authMux)
}

func userMux() http.Handler {
	userMux := http.NewServeMux()

	userMux.HandleFunc("/profile", MethodMiddleware("GET", JWT(GetProfile)))
	userMux.HandleFunc("/update", MethodMiddleware("PUT", JWT(UpdateProfile)))

	return http.StripPrefix("/api/v1/user", userMux)
}
