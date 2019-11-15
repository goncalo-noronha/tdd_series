package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	db2 "github.com/goncalo-noronha/tdd_series/src/app/db"
	"github.com/goncalo-noronha/tdd_series/src/infra/db"
	"log"
	"net/http"
)

func getHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	appointmentId := request.PathParameters["id"]

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)

	dao := db.DynamoDBDAO{TypedDynDBDAO: &db.AppointmentDynDAO{}, Connection: svc}
	repo := db2.AppointmentRepository{Dao: &dao}

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

func main() {
	lambda.Start(getHandler)
}
