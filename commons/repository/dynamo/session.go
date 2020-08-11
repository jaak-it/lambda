package dynamo

import (
    "errors"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/sirupsen/logrus"
)

var sessionDynamo *SessionDynamo
var err error

func GetSessionDynamo() (*SessionDynamo, error) {
    if sessionDynamo == nil {
        return nil, errors.New("Session to dynamo is empty")
    }

    return sessionDynamo, err
}

func NewSessionDynamo(region string) (*SessionDynamo, error) {
    awsSession, erro := session.NewSession(&aws.Config{
        Region: aws.String(region),
    })
    if erro != nil {
        return nil, erro
    }

    instance := dynamodb.New(awsSession)

    logrus.Info("New sessionDynamo to dynamo created")
    return &SessionDynamo{
        instance: instance,
    }, nil
}

