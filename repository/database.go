package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var singleton *dynamodb.DynamoDB
var err error
var once sync.Once

func GetDynamo() (*dynamodb.DynamoDB, error) {
	once.Do(func() {
		region := os.Getenv("AWS_REGION")

		awsSession, er := session.NewSession(&aws.Config{
			Region: aws.String(region),
		})

		if er != nil {
			err = er
			return
		}

		singleton = dynamodb.New(awsSession)
		logrus.Info("Created instance to dynamo")

	})
	return singleton, err
}

/*func AddItem() (*dynamodb.DynamoDB, error) {

    av, err := dynamodbattribute.MarshalMap()
    input := &dynamodb.PutItemInput{
        Item:      av,
        TableName: aws.String("user"),
    }

    _, err = svc.PutItem(input)

}*/
