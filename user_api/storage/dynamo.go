package storage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/dachanh/daita-serverless/user_api/model"
)

const tableName = "User"

func CrateUser(user model.User) error {
	userMap, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return err
	}
	sessionDynamo := createDynamoSession()
	input := &dynamodb.PutItemInput{
		Item:      userMap,
		TableName: aws.String(tableName),
	}
	_, err = sessionDynamo.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}

func createDynamoSession() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(
		session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))
	return dynamodb.New(sess)
}
