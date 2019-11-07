package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func deleteHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	appointmentId := request.PathParameters["id"]

	uuid, err := uuid.Parse(appointmentId)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	allOld := dynamodb.ReturnValueAllOld

	result, err := svc.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String("Appointments"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(uuid.String()),
			},
		},
		ReturnValues: &allOld,
	})

	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	if len(result.Attributes) == 0 {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
		}, nil

	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(deleteHandler)
}
