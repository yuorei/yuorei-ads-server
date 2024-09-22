package middleware

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

// FirebaseApp holds a reference to the Firebase app and auth client
type FirebaseApp struct {
	authClient *auth.Client
}

// NewFirebaseApp initializes a new Firebase app
func NewFirebaseApp() (*FirebaseApp, error) {
	opt := option.WithCredentialsFile(os.Getenv("GOOGLE_FIREBASE_CREDENTIALS"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting Auth client: %v", err)
	}

	return &FirebaseApp{authClient: authClient}, nil
}
