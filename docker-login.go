package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func main() {
	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load AWS config:", err)
		os.Exit(1)
	}

	client := ecr.NewFromConfig(cfg)

	resp, err := client.GetAuthorizationToken(ctx, &ecr.GetAuthorizationTokenInput{})
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to get authorization token:", err)
		os.Exit(1)
	}

	if len(resp.AuthorizationData) == 0 {
		fmt.Fprintln(os.Stderr, "no authorization data returned")
		os.Exit(1)
	}

	authToken := aws.ToString(resp.AuthorizationData[0].AuthorizationToken)

	password, err := decodeAuthToken(authToken)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to decode auth token:", err)
		os.Exit(1)
	}

	fmt.Println(password)
}

func decodeAuthToken(token string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", err
	}
	parts := strings.SplitN(string(decoded), ":", 2)
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid auth token format")
	}
	return parts[1], nil // the password
}
