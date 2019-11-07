package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func getHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	appointmentId := request.PathParameters["id"]

	uuid, err := uuid.Parse(appointmentId)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	var appointment struct {
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

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Appointments"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(uuid.String()),
			},
		},
	})

	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	if result.Item == nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
		}, nil
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &appointment)

	response, err := json.Marshal(appointment)

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
	lambda.Start(getHandler)
}
