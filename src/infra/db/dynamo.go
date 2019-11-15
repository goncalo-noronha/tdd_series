package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"log"
	"strconv"
)

type TypedDynDBDAO interface {
	getSource() *string
	mapToObj(map[string]*dynamodb.AttributeValue) (interface{}, error)
}

type DynamoDBDAO struct {
	TypedDynDBDAO

	Connection dynamodbiface.DynamoDBAPI
}

func (d *DynamoDBDAO) Read(input map[string]interface{}) (interface{}, error) {

	rawResult, err := d.Connection.GetItem(&dynamodb.GetItemInput{
		TableName: d.getSource(),
		Key:       d.mapInputs(input),
	})

	if err != nil {
		log.Println(err.Error())
		return rawResult, err
	}

	result, err := d.mapToObj(rawResult.Item)

	if err != nil {
		log.Println(err.Error())
		return result, err

	}

	return result, nil
}

func (d *DynamoDBDAO) Create(item interface{}) error {
	return nil
}

func (d *DynamoDBDAO) Update(item interface{}) error {
	return nil
}

func (d *DynamoDBDAO) Delete(input interface{}) error {
	return nil
}

func (d *DynamoDBDAO) mapInputs(input map[string]interface{}) map[string]*dynamodb.AttributeValue {

	queryMap := make(map[string]*dynamodb.AttributeValue)

	for i, value := range input {
		switch v := value.(type) {
		case string:
			in := dynamodb.AttributeValue{S: aws.String(v)}
			queryMap[i] = &in
			break
		case int:
			in := dynamodb.AttributeValue{N: aws.String(strconv.Itoa(v))}
			queryMap[i] = &in
			break
		}
	}

	return queryMap
}
