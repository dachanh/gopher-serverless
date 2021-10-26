package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dachanh/daita-serverless/user_api/model"
	"github.com/dachanh/daita-serverless/user_api/storage"
	"net/http"
)

func main() {
	lambda.Start(Handler)
}

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var user model.User
	err := json.Unmarshal([]byte(req.Body), &user)

	if err != nil {
		return response("Couldn't unmarshal ", http.StatusBadRequest), nil
	}
	err = storage.CrateUser(user)
	if err != nil {
		return response("Could't Create User", http.StatusBadRequest), nil
	}
	return response(user.UserName, http.StatusOK), nil
}

func response(body string, statusCode int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(body),
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
	}
}
