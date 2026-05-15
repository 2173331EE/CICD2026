package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoClassStore struct {
	client    *dynamodb.Client
	tableName string
}

func NewDynamoClassStore(client *dynamodb.Client, tableName string) *DynamoClassStore {
	return &DynamoClassStore{client: client, tableName: tableName}
}

func (s *DynamoClassStore) List(ctx context.Context) ([]BoatClass, error) {
	out, err := s.client.Scan(ctx, &dynamodb.ScanInput{TableName: aws.String(s.tableName)})
	if err != nil {
		return nil, err
	}

	res := make([]BoatClass, 0, len(out.Items))
	for _, item := range out.Items {
		bc, err := boatClassFromItem(item)
		if err != nil {
			return nil, err
		}
		res = append(res, bc)
	}
	return res, nil
}

func (s *DynamoClassStore) Add(ctx context.Context, bc BoatClass) error {
	// ConditionExpression empêche l'écrasement (nom = clé primaire)
	_, err := s.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(s.tableName),
		Item: boatClassToItem(bc),
		ConditionExpression: aws.String("attribute_not_exists(#n)"),
		ExpressionAttributeNames: map[string]string{"#n": "name"},
	})
	if err != nil {
		var ccf *types.ConditionalCheckFailedException
		if strings.Contains(err.Error(), "ConditionalCheckFailed") || strings.Contains(err.Error(), "conditional request failed") || strings.Contains(err.Error(), "ConditionalCheckFailedException") || (fmt.Sprintf("%T", err) == fmt.Sprintf("%T", ccf)) {
			return ErrClassAlreadyExists
		}
		return err
	}
	return nil
}

func (s *DynamoClassStore) DeleteByName(ctx context.Context, name string) (BoatClass, bool, error) {
	out, err := s.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(s.tableName),
		Key: map[string]types.AttributeValue{"name": &types.AttributeValueMemberS{Value: name}},
		ReturnValues: types.ReturnValueAllOld,
	})
	if err != nil {
		return BoatClass{}, false, err
	}
	if len(out.Attributes) == 0 {
		return BoatClass{}, false, nil
	}
	bc, err := boatClassFromItem(out.Attributes)
	if err != nil {
		return BoatClass{}, false, err
	}
	return bc, true, nil
}

func boatClassToItem(bc BoatClass) map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"name":         &types.AttributeValueMemberS{Value: bc.Name},
		"handicapType": &types.AttributeValueMemberS{Value: bc.HandicapType},
		"handicapValue": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", bc.HandicapValue)},
	}
}

func boatClassFromItem(item map[string]types.AttributeValue) (BoatClass, error) {
	name, _ := item["name"].(*types.AttributeValueMemberS)
	ht, _ := item["handicapType"].(*types.AttributeValueMemberS)
	hv, _ := item["handicapValue"].(*types.AttributeValueMemberN)
	if name == nil || ht == nil || hv == nil {
		return BoatClass{}, fmt.Errorf("invalid item")
	}
	ival, err := strconv.Atoi(hv.Value)
	if err != nil {
		return BoatClass{}, err
	}
	return BoatClass{Name: name.Value, HandicapType: ht.Value, HandicapValue: ival}, nil
}
