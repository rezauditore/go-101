package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	result, err := thirdpartyHTTPCall()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("the response took %v -> %+v\n", time.Since(start), result)
}

func fetchUserID() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100) // If the function takes longer that 100, We're going to cancel it.
	defer cancel()

	type result struct {
		userId string
		err    error
	}
	resultCh := make(chan result, 1)

	go func() {
		res, err := thirdpartyHTTPCall()
		resultCh <- result{
			userId: res,
			err:    err,
		}
	}()

	select {
	//Done ()
	//1. -> the context timeout is exceeded
	//2. -> the context has been manually canceled -> cancel()
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-resultCh:
		return res.userId, res.err
	}
}

func thirdpartyHTTPCall() (string, error) {
	time.Sleep(time.Millisecond * 500)
	return "user id 1", nil
}
