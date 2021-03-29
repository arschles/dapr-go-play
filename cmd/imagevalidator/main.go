package main

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.0/computervision"
    "github.com/Azure/go-autorest/autorest"
    "io"
    "log"
    "os"
    "strings"
    "time"
)

// Declare global so don't have to pass it to all of the tasks.
var computerVisionContext context.Context

func main() {
    computerVisionKey := "PASTE_YOUR_COMPUTER_VISION_SUBSCRIPTION_KEY_HERE"
    endpointURL := "PASTE_YOUR_COMPUTER_VISION_ENDPOINT_HERE"
