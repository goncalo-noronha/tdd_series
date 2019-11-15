package main

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func listHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	panic("HEYYY")

	var appointments []struct {
		Id      string `json:"id,omitempty"`
		Patient struct {
			Name       string `json:"name"`
			DocumentId string `json:"document_id"`
		}
		Specialty string `json:"specialty"`
		Date      string `json:"date"`
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	params := &dynamodb.ScanInput{
		TableName: aws.String("Appointments"),
	}

	result, err := svc.Scan(params)
	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &appointments)

	response, err := json.Marshal(appointments)

	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, err
	}
	return events.APIGatewayProxyResponse{
		Body:       string(response),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(listHandler)
}
