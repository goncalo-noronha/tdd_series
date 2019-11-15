package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/goncalo-noronha/tdd_series/src/app/db"
)

type AppointmentDynDAO struct {
}

func (ad *AppointmentDynDAO) getSource() *string {
	return aws.String("Appointments")
}

func (ad *AppointmentDynDAO) mapToObj(input map[string]*dynamodb.AttributeValue) (interface{}, error) {
	var appointment db.Appointment

	err := dynamodbattribute.UnmarshalMap(input, &appointment)

	if err != nil {
		return db.Appointment{}, err
	}

	return appointment, nil
}
