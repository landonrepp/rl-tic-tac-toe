package api

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"strconv"
)

type Item struct {
	Year   int
	Title  string
	Plot   string
	Rating float64
}

func Start() {
	awsSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	dynamoDbSvc := dynamodb.New(awsSession)

	item := Item{
		Year:   2015,
		Title:  "The Big New Movie",
		Plot:   "Nothing happens at all.",
		Rating: 0.0,
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
	}

	// Create item in table Movies
	tableName := "Test"

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = dynamoDbSvc.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}

	year := strconv.Itoa(item.Year)

	fmt.Println("Successfully added '" + item.Title + "' (" + year + ") to table " + tableName)

}
