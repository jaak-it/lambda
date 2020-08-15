package dynamo

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type SessionDynamo struct {
    instance *dynamodb.DynamoDB
}

func (cd *SessionDynamo) SaveItem(item interface{}, tableName string) error {
    elementMap, err := dynamodbattribute.MarshalMap(item)
    if err != nil {
        return err
    }

    intentSave := dynamodb.PutItemInput{
        Item:      elementMap,
        TableName: aws.String(tableName),
    }

    _, err = cd.instance.PutItem(&intentSave)
    if err != nil {
        return err
    }

    return nil
}

func (cd *SessionDynamo) DeleteItem(id string, tableName string) error {
    intentDelete := &dynamodb.DeleteItemInput{
        Key: map[string]*dynamodb.AttributeValue{
            "id": {
                S: aws.String(id),
            },
        },
        TableName: aws.String(tableName),
    }

    _, err := cd.instance.DeleteItem(intentDelete)
    if err != nil {
        return err
    }

    return nil
}

func (cd *SessionDynamo) GetItem(id string, tableName string) (*dynamodb.GetItemOutput, error) {
    query := dynamodb.GetItemInput{
        TableName: aws.String(tableName),
        Key: map[string]*dynamodb.AttributeValue{
            "id": {
                S: aws.String(id),
            },
        },
    }

    result, err := cd.instance.GetItem(&query)
    if err != nil {
        return nil, err
    }

    //if result.Item == nil {
    //    msg := "Could not find '" + id + "'"
    //    return nil, errors.New(msg)
    //}

    return result, nil
}

func (cd *SessionDynamo) FindUserId(userId string, tableName string) (*dynamodb.GetItemOutput, error) {
    query := dynamodb.GetItemInput{
        TableName: aws.String(tableName),
        Key: map[string]*dynamodb.AttributeValue{
            "userId": {
                S: aws.String(userId),
            },
        },
    }

    result, err := cd.instance.GetItem(&query)
    if err != nil {
        return nil, err
    }

    //if result.Item == nil {
    //    msg := "Could not find '" + id + "'"
    //    return nil, errors.New(msg)
    //}

    return result, nil

}
