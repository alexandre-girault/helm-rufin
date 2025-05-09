package main

import (
	"context"
	"encoding/json"
	"log"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func getSecretsmanagerSecret(secretArn string, secretKey string) string {

	// search for the secret in the region of the ARN
	awsRegionFromArn := regexp.MustCompile(`arn:aws:secretsmanager:([a-z]{2}-[a-z]+-[0-9]):[0-9]+:[a-z]+.+`)
	region := awsRegionFromArn.FindStringSubmatch(secretArn)[1]

	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatal(err)
	}

	// Create Secrets Manager client
	svc := secretsmanager.NewFromConfig(config)
	//fmt.Println("secretArn: ", secretArn)
	//fmt.Println("secretKey: ", secretKey)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretArn),
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		log.Fatal(err.Error())
	}

	if secretKey != "" {
		var secretMap map[string]string

		err = json.Unmarshal([]byte(*result.SecretString), &secretMap)
		if err != nil {
			log.Fatal(err.Error())
		}

		return secretMap[secretKey]
	}
	// If no secretKey is provided, we return the entire secret string
	return *result.SecretString

}
