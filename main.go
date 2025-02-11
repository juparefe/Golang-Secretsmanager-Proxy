package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// Request estructura para recibir el nombre del secreto
type Request struct {
	SecretName string `json:"secretName"`
}

// SecretData estructura de respuesta
type SecretData struct {
	Username             string `json:"username"`
	Password             string `json:"password"`
	Engine               string `json:"engine"`
	Host                 string `json:"host"`
	Port                 int    `json:"port"`
	Ddbname              string `json:"dbname"`
	DbInstanceIdentifier string `json:"dbInstanceIdentifier"`
}

// HandleRequest maneja la invocación de Lambda
func HandleRequest(ctx context.Context, request Request) (SecretData, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return SecretData{}, fmt.Errorf("error loading AWS config: %v", err)
	}

	svc := secretsmanager.NewFromConfig(cfg)
	secretValue, err := svc.GetSecretValue(ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(request.SecretName),
	})
	if err != nil {
		return SecretData{}, fmt.Errorf("error retrieving secret: %v", err)
	}

	var secretData SecretData
	err = json.Unmarshal([]byte(*secretValue.SecretString), &secretData)
	if err != nil {
		return SecretData{}, fmt.Errorf("error unmarshalling secret: %v", err)
	}

	return secretData, nil
}

func main() {
	lambda.Start(HandleRequest)
}
