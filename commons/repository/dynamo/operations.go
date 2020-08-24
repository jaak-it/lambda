package dynamo

import (
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
    "github.com/aws/aws-sdk-go/service/dynamodb/expression"
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

func (cd *SessionDynamo) FindUserId(userId string, tableName string) (*dynamodb.ScanOutput, error) {
    filter := expression.Name("userId").Equal(expression.Value(userId))
    proj := expression.NamesList(expression.Name("cards"))

    expr, err := expression.NewBuilder().WithFilter(filter).WithProjection(proj).Build()
    if err != nil {
        fmt.Println("Got error building expression:")
        fmt.Println(err.Error())
        //os.Exit(1)
    }

    query := &dynamodb.ScanInput{
        ExpressionAttributeNames:  expr.Names(),
        ExpressionAttributeValues: expr.Values(),
        //FilterExpression:          expr.Filter(),
        ProjectionExpression: expr.Projection(),
        TableName:            aws.String(tableName),
    }

    result, err := cd.instance.Scan(query)
    if err != nil {
        fmt.Println("Query API call failed:")
        fmt.Println(err.Error())
        //os.Exit(1)
    }

    return result, err
}

func (cd *SessionDynamo) FindUserList(item interface{}, tableName string) (*dynamodb.ScanOutput, error) {
    proj := expression.NamesList(expression.Name("firstname"), expression.Name("lastname"),
        expression.Name("isActive"))

    expr, err := expression.NewBuilder().WithProjection(proj).Build()
    if err != nil {
        fmt.Println("Got error building expression:")
        fmt.Println(err.Error())
    }

    query := &dynamodb.ScanInput{
        ExpressionAttributeNames:  expr.Names(),
        ExpressionAttributeValues: expr.Values(),
        FilterExpression:          expr.Filter(),
        ProjectionExpression:      expr.Projection(),
        TableName:                 aws.String(tableName),
    }

    result, err := cd.instance.Scan(query)
    if err != nil {
        fmt.Println("Query API call failed:")
        fmt.Println(err.Error())
    }

    return result, err
}

/*func (cd *SessionDynamo) FindUserId(userId string, tableName string) (*dynamodb.ScanOutput, error) {
    query := dynamodb.ScanInput{
        TableName:            aws.String(tableName),
        ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
            ":a": {
                S: aws.String(userId),
            },
        },
        FilterExpression: aws.String("userId = :a"),
    }
    result, err := cd.instance.Scan(&query)
    if err != nil {
        return nil, err
    }

    return result, nil
}*/
