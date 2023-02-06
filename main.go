package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func main() {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithSharedConfigProfile("aphexlog"),
	)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	client := sts.NewFromConfig(cfg)

	identity, err := client.GetCallerIdentity(
		context.TODO(),
		&sts.GetCallerIdentityInput{},
	)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	fmt.Printf(
		"Account: %s\nUserID: %s\nARN: %s\n",
		aws.ToString(identity.Account),
		aws.ToString(identity.UserId),
		aws.ToString(identity.Arn),
	)
}
