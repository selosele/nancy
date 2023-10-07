package middlewares

import "net/http"

/* 인증 미들웨어 */
func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("Authorization")

		// TODO: 인증키 검증
		println("middleware test ::: ", apiKey)

		h.ServeHTTP(w, r)
	})
}
