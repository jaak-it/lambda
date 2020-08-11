package repository

import (
    "errors"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/sirupsen/logrus"
)

var dynamoDB *dynamodb.DynamoDB
var err error

func GetSessionDynamo() (*dynamodb.DynamoDB, error) {
    if dynamoDB == nil {
        return nil, errors.New("Session to dynamo is empty")
    }

    return dynamoDB, err
}

func CreateSessionDynamo(region string) (*dynamodb.DynamoDB, error) {
    awsSession, er := session.NewSession(&aws.Config{
        Region: aws.String(region),
    })

    if er != nil {
        return nil, er
    }

    dynamoDB = dynamodb.New(awsSession)
    logrus.Info("New session to dynamo created")

    return dynamoDB, nil
}

/*func AddItem() (*dynamodb.DynamoDB, error) {

    av, err := dynamodbattribute.MarshalMap()
    input := &dynamodb.PutItemInput{
        Item:      av,
        TableName: aws.String("user"),
    }

    _, err = svc.PutItem(input)

}*/
