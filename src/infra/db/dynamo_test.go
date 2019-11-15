package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/goncalo-noronha/tdd_series/src/app"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"strconv"
	"testing"
)

type MockDynamoConnection struct {
	dynamodbiface.DynamoDBAPI
	mock.Mock
}

func (mc *MockDynamoConnection) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	args := mc.Called(input)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*dynamodb.GetItemOutput), args.Error(1)
}

func TestDynamoDBDAO_Read(t *testing.T) {
	t.Run("Test input field text", func(t *testing.T) {
		mockConnection := new(MockDynamoConnection)

		mockValue := map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String("UUID")},
		}

		stubRet := dynamodb.GetItemOutput{Item: mockValue}

		mockConnection.On(
			"GetItem",
			&dynamodb.GetItemInput{
				TableName: aws.String("Appointments"),
				Key: map[string]*dynamodb.AttributeValue{
					"id": {S: aws.String("UUID")},
				},
			},
		).Return(&stubRet, nil)

		sut := DynamoDBDAO{TypedDynDBDAO: &AppointmentDynDAO{}, Connection: mockConnection}
		result, err := sut.Read(map[string]interface{}{"id": "UUID"})

		assert.Nil(t, err)
		assert.Equal(t, app.Appointment{ID: "UUID"}, result)
	})

	t.Run("Test input field number", func(t *testing.T) {
		mockConnection := new(MockDynamoConnection)

		stubValue := map[string]*dynamodb.AttributeValue{
			"id": {N: aws.String(strconv.Itoa(1))},
		}

		stubRet := dynamodb.GetItemOutput{Item: stubValue}

		mockConnection.On(
			"GetItem",
			&dynamodb.GetItemInput{
				TableName: aws.String("Appointments"),
				Key: map[string]*dynamodb.AttributeValue{
					"id": {N: aws.String(strconv.Itoa(1))},
				},
			},
		).Return(&stubRet, nil)

		sut := DynamoDBDAO{TypedDynDBDAO: &AppointmentDynDAO{}, Connection: mockConnection}
		result, err := sut.Read(map[string]interface{}{"id": 1})

		assert.Nil(t, err)
		assert.Equal(t, app.Appointment{ID: "1"}, result)
	})
}
