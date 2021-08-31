package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {
	config := &firebase.Config{
		StorageBucket: "<your-bucket-name>",
	}
	opt := option.WithCredentialsFile("<path-to-firebase-key>")
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}

	keyFile := "<path-to-firebase-key>"
	key, err := ioutil.ReadFile(keyFile)
	if err != nil {
		log.Fatalln(err)
	}

	cfg, err := google.JWTConfigFromJSON(key)
	if err != nil {
		log.Fatalln(err)
	}

	attrs, err := bucket.Attrs(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	bucketName := attrs.Name
	method := "GET"
	expires := time.Now().Add(time.Second * 60)

	var names []string
	var downloadLinks []string
	it := bucket.Objects(context.Background(), nil)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		names = append(names, attrs.Name)
		downloadLinks = append(downloadLinks, attrs.MediaLink)

		url, err := storage.SignedURL(bucketName, attrs.Name, &storage.SignedURLOptions{
			GoogleAccessID: cfg.Email,
			PrivateKey:     cfg.PrivateKey,
			Method:         method,
			Expires:        expires,
		})

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(url)
	}
}
