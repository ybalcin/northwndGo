package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
)

func Client() *firestore.Client {
	ctx := context.Background()
	opt := option.WithCredentialsFile("firebase-credentials.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatal(err.Error())
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	return client
}
