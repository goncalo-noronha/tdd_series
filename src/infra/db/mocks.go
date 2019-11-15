package db

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/stretchr/testify/mock"
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
