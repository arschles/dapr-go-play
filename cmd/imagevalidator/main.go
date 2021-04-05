package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/grpc"
)

// Declare global so don't have to pass it to all of the tasks.
var computerVisionContext context.Context

func main() {
	// 	computerVisionKey := "PASTE_YOUR_COMPUTER_VISION_SUBSCRIPTION_KEY_HERE"
	// 	endpointURL := "PASTE_YOUR_COMPUTER_VISION_ENDPOINT_HERE"

	s, err := daprd.NewService(":5001")
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}

	// add handler to the service
	if err := s.AddServiceInvocationHandler("validate_image", validateHandler); err != nil {
		log.Fatalf("error adding handler: %v", err)
	}

	// start the server to handle incoming events
	if err := s.Start(); err != nil {
		log.Fatalf("server error: %v", err)
	}

}

func validateHandler(ctx context.Context, in *common.InvocationEvent) (*common.Content, error) {
	type body struct {
		URL string `json:"url"`
	}

	log.Printf(
		"Invocation (ContentType:%s, Verb:%s, QueryString:%s, Data:%s)",
		in.ContentType, in.Verb, in.QueryString, string(in.Data),
	)

	req := new(body)
	if err := json.Unmarshal(in.Data); err != nil {
		return nil, err
	}

	// get image URL
	url := req.URL

	// download image

	// meow hash image

	// check dapr storage for hash

	// if hash is in dapr storage, return result

	// otherwise:
	// - send image URL to Azure
	// - get response
	// - put response into dapr cache
	// - return response

	// TODO: implement handling logic here
	out := &common.Content{
		ContentType: in.ContentType,
		Data:        in.Data,
	}

	return out, nil
}

// https://api.thecatapi.com/v1/images/search
// https://dog.ceo/api/breeds/image/random

// stanford dogs dataset: http://vision.stanford.edu/aditya86/ImageNetDogs/
