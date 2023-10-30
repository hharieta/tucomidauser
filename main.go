// go get github.com/aws/aws-lambda-go/lambda

package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hharieta/tucomidauser/awsgo"
	"github.com/hharieta/tucomidauser/bd"
	"github.com/hharieta/tucomidauser/models"
)

func main() {
	lambda.Start(LambdaExec)
}

func LambdaExec(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.AwsInit()

	if !ValidateParams() {
		fmt.Println("Error: 'SecretName' has been not send")
		err := errors.New("Error: must be sent SecretName")
		return event, err
	}

	var cognitoData models.SignUp //cognito

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			cognitoData.UserEmail = att
			fmt.Println("Email = " + cognitoData.UserEmail)
		case "sub":
			cognitoData.UserUUID = att
			fmt.Println("Sub = " + cognitoData.UserUUID)
		}

	}

	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("Error: Read Secret" + err.Error())

		return event, err
	}

	err = bd.SignUp(cognitoData)

	return event, err

}

func ValidateParams() bool {
	var retriveParam bool
	_, retriveParam = os.LookupEnv("SecretName")

	return retriveParam
}
