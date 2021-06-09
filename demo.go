package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
)

func main() {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "fake-project-id")
	if err != nil {
		panic(err)
	}

	result, err := client.Collection("mycollection").Doc("myid").Create(context.TODO(), map[string]interface{}{
		"mykey": "my data",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("done", result)
}
