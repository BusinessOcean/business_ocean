package infrastructure

import (
	"becore/belogger"
	"context"
	"path/filepath"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

// NewFirebaseApp creates new firebase app instance
func NewFirebaseApp(logger *belogger.BeLogger) *firebase.App {

	ctx := context.Background()

	serviceAccountKeyFilePath, err := filepath.Abs("./serviceAccountKey.json")
	if err != nil {
		logger.Errorf("Unable to load serviceAccountKey.json file")
		panic(err)
	}

	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		logger.Errorf("Firebase NewApp: %v", err)
	}
	logger.Info("✅ Firebase app initialized.")
	return app
}

// NewFirebaseAuth creates new firebase auth client
func NewFirebaseAuth(logger *belogger.BeLogger, app *firebase.App) *auth.Client {

	ctx := context.Background()

	firebaseAuth, err := app.Auth(ctx)
	if err != nil {
		logger.Errorf("Firebase Authentication: %v", err)
	}

	return firebaseAuth
}

// NewFirestoreClient creates new firestore client
func NewFirestoreClient(logger *belogger.BeLogger, app *firebase.App) *firestore.Client {
	ctx := context.Background()

	firestoreClient, err := app.Firestore(ctx)
	if err != nil {
		logger.Fatalf("Firestore client: %v", err)
	}

	return firestoreClient
}

// NewFirebaseMessagingClient creates new firebase cloud messaging client
func NewFirebaseMessagingClient(logger *belogger.BeLogger, app *firebase.App) *messaging.Client {
	ctx := context.Background()
	messagingClient, err := app.Messaging(ctx)
	if err != nil {
		logger.Fatalf("Firebase messaing: %v", err)
	}
	return messagingClient
}
