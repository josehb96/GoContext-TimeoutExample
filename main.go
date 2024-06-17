package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

// callThirdPartyAPI calls a third-party API and returns a boolean indicating whether the user is subscribed
func callThirdPartyAPI(ctx context.Context, userID int) (bool, error) {

	// Simulate a delay to mimic the time it takes to call the third-party API
	time.Sleep(400 * time.Millisecond)

	// Check if the context has been cancelled or timed out
	if ctx.Err() == context.DeadlineExceeded {
		// If the context has timed out, return an error
		return false, errors.New("context timeout exceeded")
	}

	// If the context has not timed out, return true indicating the user is subscribed
	return true, nil

}

// main is the entry point for the program
func main() {

	// Create a context with a timeout of 500 milliseconds
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel() // Cancel the context when the function returns

	// Define the user ID
	userId := 43253

	// Call the third-party API using the context
	isUserSubbed, err := callThirdPartyAPI(ctx, userId)
	if err != nil {
		// If an error occurs, log a fatal error message
		log.Fatalf("error fetching user status for :%d error: %s", userId, err)
	}

	// If the user is subscribed, print a message
	if isUserSubbed {
		fmt.Printf("this user is subbed %d\n", userId)
	}

}
