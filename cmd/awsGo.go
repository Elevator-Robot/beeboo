package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func awsGo() {
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

	// create environment variables
	os.Setenv("AWS_ACCOUNT_ID", aws.ToString(identity.Account))
	os.Setenv("AWS_USER_ID", aws.ToString(identity.UserId))
	os.Setenv("AWS_ARN", aws.ToString(identity.Arn))

	fmt.Println("AWS_ACCOUNT_ID:", os.Getenv("AWS_ACCOUNT_ID"))
	fmt.Println("AWS_USER_ID:", os.Getenv("AWS_USER_ID"))
	fmt.Println("AWS_ARN:", os.Getenv("AWS_ARN"))

	// get the credentials
	creds, err := cfg.Credentials.Retrieve(context.TODO())
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	// create environment variables
	os.Setenv("AWS_ACCESS_KEY_ID", creds.AccessKeyID)
	os.Setenv("AWS_SECRET_ACCESS_KEY", creds.SecretAccessKey)
	os.Setenv("AWS_SESSION_TOKEN", creds.SessionToken)

	// print the credentials
	fmt.Println("AWS_ACCESS_KEY_ID:", os.Getenv("AWS_ACCESS_KEY_ID"))
	fmt.Print("AWS_SECRET_ACCESS_KEY:", os.Getenv("AWS_SECRET_ACCESS_KEY"))
	fmt.Print("AWS_SESSION_TOKEN:", os.Getenv("AWS_SESSION_TOKEN"))

}
