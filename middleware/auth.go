package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"
)

// AuthMiddleware is a middleware for authenticating requests with Firebase
func (f *FirebaseApp) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "Bearer" { // 空の時
			log.Println("Authorization header is missing")
			ctx := context.WithValue(r.Context(), "Authorization", "Authorization header is missing")
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		idToken := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))
		token, err := f.authClient.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			if strings.Contains(err.Error(), "ID token has expired") {
				log.Println("ID token has expired")
				ctx := context.WithValue(r.Context(), "Authorization", "ID token has expired")
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			} else if strings.Contains(err.Error(), "ID token must be a non-empty string") {
				log.Println("ID token is not an ID token")
				ctx := context.WithValue(r.Context(), "Authorization", "ID token is not an ID token")
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
			log.Println("Error verifying ID token:", err)
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Token is valid, set the token in the context and call the next handler
		key := "uid"
		ctx := context.WithValue(r.Context(), key, token.UID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
