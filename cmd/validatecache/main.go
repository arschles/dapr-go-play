package main

import (
	"context"
	"encoding/json"
	"fmt"
	"hash"
	"io"
	"log"
	"net/http"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/grpc"
	"github.com/mmcloughlin/meow"
)

func main() {
	meowHasher := meow.New64(8675309)

	s, err := daprd.NewService(":5001")
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}

	daprCl, err := dapr.NewClient()
	if err != nil {
		log.Fatalf("Error creating dapr storage client: %s", err)
	}

	// add handler to the service
	if err := s.AddServiceInvocationHandler(
		"validate_from_cache",
		validateHandler(daprCl, meowHasher),
	); err != nil {
		log.Fatalf("error adding handler: %v", err)
	}

	// start the server to handle incoming events
	if err := s.Start(); err != nil {
		log.Fatalf("server error: %v", err)
	}

}

func validateHandler(cl dapr.Client, hasher hash.Hash64) func(context.Context, *common.InvocationEvent) (*common.Content, error) {
	type body struct {
		URL string `json:"url"`
	}

	return func(ctx context.Context, in *common.InvocationEvent) (*common.Content, error) {
		req := new(body)
		if err := json.Unmarshal(in.Data, req); err != nil {
			return nil, err
		}

		// get image URL
		url := req.URL

		// download image
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		imgData, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		// meow hash image
		hasher.Reset()
		hasher.Write(imgData)
		imgHash := hasher.Sum64()

		// check dapr storage for hash=
		imgMetadata, err := cl.GetState(ctx, "image-metadata", fmt.Sprintf("%d", imgHash))
		if err == nil {
			// if hash is in dapr storage, return result
			return &common.Content{
				ContentType: in.ContentType,
				Data:        imgMetadata.Value,
			}, nil
		}

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
}

// https://api.thecatapi.com/v1/images/search
// https://dog.ceo/api/breeds/image/random

// stanford dogs dataset: http://vision.stanford.edu/aditya86/ImageNetDogs/
