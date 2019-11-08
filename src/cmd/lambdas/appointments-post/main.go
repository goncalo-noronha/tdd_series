package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"github.com/xeipuuv/gojsonschema"
)

func postHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	var appointment struct {
		ID      string `json:"id,omitempty"`
		Patient struct {
			Name       string `json:"name"`
			DocumentID string `json:"document_id"`
		}
		Specialty string `json:"specialty"`
		Date      string `json:"date"`
	}

	srcPath, err := os.Getwd()

	if err != nil {
		log.Println("Couldn't get root folder", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	schemaLoader := gojsonschema.NewReferenceLoader("file://" + srcPath + "/assets/appointments.json")
	jsonLoader := gojsonschema.NewStringLoader(request.Body)

	result, err := gojsonschema.Validate(schemaLoader, jsonLoader)

	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, nil
	}

	if !result.Valid() {
		unMarErrs, _ := json.Marshal(result.Errors())
		log.Println(unMarErrs)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       string(unMarErrs),
		}, nil
	}

	err = json.Unmarshal([]byte(request.Body), &appointment)

	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, nil
	}

	uuid, err := uuid.NewRandom()

	if err != nil {
		log.Println("Couldn't generate UUI", err.Error())
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	appointment.ID = uuid.String()

	av, err := dynamodbattribute.MarshalMap(appointment)

	if err != nil {
		log.Println("Couldn't unmarshall appointment", err.Error())
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Appointments"),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Println("Couldn't save appointment", err.Error())
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	marshalledJSON, err := json.Marshal(appointment)

	if err != nil {
		log.Println(err.Error())
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusAccepted,
		Body:       string(marshalledJSON),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(postHandler)
}
