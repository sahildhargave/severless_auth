package handlers

import (
	"net/http"
    "github.com/aws/aws-lambda-go-events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws-sdk-go/servce/dynamodb/dynamodbiface"
    "github.com/sahildhargave/serverless-auth/user"
)


var ErrorMethodNotAllowed = "method Not Allowed"

type ErrorBody struct{
   ErrorMsg *String `json:"error,omitempty"`
}


func GetUser(){

}

func CreateUser(){}

func UpdateUser(){}

func DeleteUser(){}

func UnhandledMethod()(*events.API){}