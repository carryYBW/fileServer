package handler

import (
	"net/http"
)

// 拦截器    传入一个 handlerfunc   让后放回一个加工的handlerfunc（前面加了一点东西）
func HTTPInterceptor(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		username := r.Form.Get("username")
		token := r.Form.Get("token")

		if len(username) < 3 || !isTokenValid(token) {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		h(w, r)
	}
}
