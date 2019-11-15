package db

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/goncalo-noronha/tdd_series/src/app"
)

type AppointmentRepoFactory struct {}

func (f AppointmentRepoFactory) Make(conn dynamodbiface.DynamoDBAPI) app.AppointmentRepository {

	dao := DynamoDBDAO{TypedDynDBDAO: &AppointmentDynDAO{}, Connection: conn}
	return app.AppointmentRepository{Gateway: &dao}

}