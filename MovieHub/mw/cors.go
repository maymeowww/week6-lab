package mw

import (
	"net/http"
)

func CORS(allowed []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := w.Header()
		
		// อนุญาตทุก Origin
		origin := r.Header.Get("Origin")
		if origin != "" {
			h.Set("Access-Control-Allow-Origin", origin)
		} else {
			h.Set("Access-Control-Allow-Origin", "*")
		}

		h.Add("Vary", "Origin")
		h.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		h.Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		h.Set("Access-Control-Max-Age", "600")

		// ตอบกลับ Preflight Request (OPTIONS) ทันที
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}