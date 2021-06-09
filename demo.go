package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"cloud.google.com/go/firestore"
)

func main() {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "fake-project-id")
	if err != nil {
		panic(err)
	}

	max, _ := strconv.Atoi(os.Getenv("NUM_RECORDS"))
	if max < 1 {
		max = 1
	}

	if max > 1 {
		batchWrite(client, max)
		return
	}

	result, err := client.Collection("users").Doc("alovelace").Set(context.TODO(), map[string]interface{}{
		"mykey": "my data",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("done writing", result)
}

func batchWrite(client *firestore.Client, max int) {
	batchSize := 500
	batch := client.Batch()

	for i := 1; i <= max; i++ {
		batch.Set(client.Collection("users").Doc(fmt.Sprint(i)), map[string]interface{}{
			"mykey": "my data",
			"myid":  i,
		})

		if i%batchSize == 0 {
			start := time.Now()
			results, err := batch.Commit(context.TODO())
			if err != nil {
				panic(err)
			}
			timeElapsed := float64(time.Since(start).Seconds())
			rate := float64(batchSize) / timeElapsed
			fmt.Printf("done writing [%d-%d] %d result(s) [rate: %f/s]\n", i-batchSize, i, len(results), rate)
			batch = client.Batch()
		}
	}
}
