package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/goncalo-noronha/tdd_series/src/infra/db"
	"log"
	"net/http"
)

func getHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	appointmentId := request.PathParameters["id"]

	repo := db.AppointmentRepoFactory{}.Make(startConnection())

	appointment, err := repo.FindOneByID(appointmentId)

	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, err
	}

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

func startConnection() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return dynamodb.New(sess)
}

func main() {
	lambda.Start(getHandler)
}
