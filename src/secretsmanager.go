package main

import (
	"context"
	"encoding/json"
	"log"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type secret struct {
	Region              string
	AccountId           string
	SecretName          string
	SecretKey           string
	SecretArnWithoutKey string
}

func secretsmanagerArnParser(secretArn string) secret {
	// get the secret region and key in ARN
	// AWS secrets manager ARN format is :
	//         arn:aws:secretsmanager:<Region>:<AccountId>:secret:<SecretName>-6RandomCharacters/<OptionalKey>

	awsRegionFromArn := regexp.MustCompile(`arn:aws:secretsmanager:([a-z]{2}-[a-z]+-[0-9]):([0-9]+):secret:(.+)/(.+)`)

	s := secret{
		Region:              awsRegionFromArn.FindStringSubmatch(secretArn)[1],
		AccountId:           awsRegionFromArn.FindStringSubmatch(secretArn)[2],
		SecretName:          awsRegionFromArn.FindStringSubmatch(secretArn)[3],
		SecretKey:           awsRegionFromArn.FindStringSubmatch(secretArn)[4],
		SecretArnWithoutKey: strings.TrimSuffix(secretArn, "/"+awsRegionFromArn.FindStringSubmatch(secretArn)[4]),
	}

	return s
}

func getSecretsmanagerSecret(secretArn string) string {

	s := secretsmanagerArnParser(secretArn)
	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(s.Region))
	if err != nil {
		log.Fatal(err)
	}

	// Create Secrets Manager client
	awsClient := secretsmanager.NewFromConfig(config)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(s.SecretArnWithoutKey), // Use the ARN without the key part
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := awsClient.GetSecretValue(context.TODO(), input)
	if err != nil {
		log.Fatal(err.Error())
	}

	if s.SecretKey != "" {
		var secretMap map[string]string

		err = json.Unmarshal([]byte(*result.SecretString), &secretMap)
		if err != nil {
			log.Fatal(err.Error())
		}

		return secretMap[s.SecretKey]
	}

	// If no secretKey is provided, we return the entire secret string
	return *result.SecretString

}
