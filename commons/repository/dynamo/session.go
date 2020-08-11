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

    dynamo := dynamodb.New(awsSession)

    sessionDynamo = &SessionDynamo{
        instance: dynamo,
    }

    logrus.Print("Create connection to dynamo database")
    return sessionDynamo, nil
}

