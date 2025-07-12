package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func getSecretsmanagerSecret(secretArn string) string {

	// get the secret region and key in ARN
	// AWS secrets manager ARN format is :
	//         arn:aws:secretsmanager:<Region>:<AccountId>:secret:<SecretName>-6RandomCharacters/<OptionalKey>

	awsRegionFromArn := regexp.MustCompile(`arn:aws:secretsmanager:([a-z]{2}-[a-z]+-[0-9]):[0-9]+:secret:(.+)/(.+)`)
	region := awsRegionFromArn.FindStringSubmatch(secretArn)[1]
	secretName := awsRegionFromArn.FindStringSubmatch(secretArn)[2]
	secretKey := awsRegionFromArn.FindStringSubmatch(secretArn)[3]
	secretArnWithoutKey := strings.TrimSuffix(secretArn, "/"+secretKey)

	fmt.Println("\tRegion: ", region)
	fmt.Println("\tSecret Name: ", secretName)
	fmt.Println("\tSecret Key: ", secretKey)

	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatal(err)
	}

	// Create Secrets Manager client
	awsClient := secretsmanager.NewFromConfig(config)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretArnWithoutKey), // Use the ARN without the key part
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := awsClient.GetSecretValue(context.TODO(), input)
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
