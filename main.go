package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

func callThirdPartyAPI(ctx context.Context, userID int) (bool, error) {

	time.Sleep(400 * time.Millisecond)

	if ctx.Err() == context.DeadlineExceeded {
		return false, errors.New("context timeout exceeded")
	}

	return true, nil

}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	userId := 43253

	// Calling third party API

	isUserSubbed, err := callThirdPartyAPI(ctx, userId)
	if err != nil {
		log.Fatalf("error fetching user status for :%d error: %s", userId, err)
	}

	if isUserSubbed {
		fmt.Printf("this user is subbed %d\n", userId)
	}

}
