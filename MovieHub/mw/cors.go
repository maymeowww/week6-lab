package mw

import (
	"net/http"
	"strings"
)

// CORS อนุญาตเฉพาะ origin ที่อยู่ในรายชื่อ หรืออนุญาตทั้งหมดถ้าใส่ "*"
func CORS(allowed []string, next http.Handler) http.Handler {
	ok := map[string]bool{}
	allowAll := false

	for _, o := range allowed {
		trimmed := strings.TrimSpace(o)
		if trimmed == "*" {
			allowAll = true
		}
		ok[trimmed] = true
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		// ถ้าตั้งค่าอนุญาตทั้งหมด หรือ origin ตรงกับรายการที่กำหนด
		if allowAll || ok[origin] {
			h := w.Header()
			if allowAll && origin != "" {
				h.Set("Access-Control-Allow-Origin", origin)
			} else if allowAll {
				h.Set("Access-Control-Allow-Origin", "*")
			} else {
				h.Set("Access-Control-Allow-Origin", origin)
			}
			h.Add("Vary", "Origin")
			h.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			h.Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			h.Set("Access-Control-Max-Age", "600")
		}

		// หากเป็น OPTIONS Preflight request ให้ส่ง 204 กลับทันที
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}