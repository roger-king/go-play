package handlers

import (
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// token := r.Header.Get("Authorization")

		// isAuthed := models.Validate(models.JwtToken{Token: token})

		// logrus.Infoln("PLease work", isAuthed)
		next.ServeHTTP(w,r)
	})
}
