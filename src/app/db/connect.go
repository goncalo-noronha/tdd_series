package db

import "github.com/aws/aws-sdk-go/service/dynamodb"

type DAO interface {
	Create(item interface{}) error
	Read(input interface{}) (interface{}, error)
	Update(item interface{}) error
	Delete(input interface{}) error
}

type DynamoDAO struct {
	connection *dynamodb.DynamoDB
}

