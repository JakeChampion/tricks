package main

import (
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

// const DEFAULT_API_HOST = "api.netlify.com"

type EnvironmentContext struct {
	Edge_URL          string `json:"url,omitempty"`
	Primary_Region    string `json:"primary_region,omitempty"`
	Token             string `json:"token,omitempty"`
	Uncached_Edge_URL string `json:"url_uncached,omitempty"`
}

type InvocationMetadata struct {
	AccountTier      string `json:"accountTier,omitempty"`
	BuildbotVersion  string `json:"buildbotVersion,omitempty"`
	BuildVersion     string `json:"buildVersion,omitempty"`
	Branch           string `json:"branch,omitempty"`
	Framework        string `json:"framework,omitempty"`
	FrameworkVersion string `json:"frameworkVersion,omitempty"`
	FunctionName     string `json:"function_name,omitempty"`
	Generator        string `json:"generator,omitempty"`
}

// APIGatewayRequestIdentity contains identity information for the request caller.
type APIGatewayRequestIdentity struct {
	CognitoIdentityPoolID         string `json:"cognitoIdentityPoolId,omitempty"`
	AccountID                     string `json:"accountId,omitempty"`
	CognitoIdentityID             string `json:"cognitoIdentityId,omitempty"`
	Caller                        string `json:"caller,omitempty"`
	APIKey                        string `json:"apiKey,omitempty"`
	APIKeyID                      string `json:"apiKeyId,omitempty"`
	AccessKey                     string `json:"accessKey,omitempty"`
	SourceIP                      string `json:"sourceIp"`
	CognitoAuthenticationType     string `json:"cognitoAuthenticationType,omitempty"`
	CognitoAuthenticationProvider string `json:"cognitoAuthenticationProvider,omitempty"`
	UserArn                       string `json:"userArn,omitempty"` //nolint: stylecheck
	UserAgent                     string `json:"userAgent"`
	User                          string `json:"user,omitempty"`
}

// APIGatewayProxyRequestContext contains the information to identify the AWS account and resources invoking the
// Lambda function. It also includes Cognito identity information for the caller.
type APIGatewayProxyRequestContext struct {
	AccountID         string                    `json:"accountId"`
	ResourceID        string                    `json:"resourceId"`
	OperationName     string                    `json:"operationName,omitempty"`
	Stage             string                    `json:"stage"`
	DomainName        string                    `json:"domainName"`
	DomainPrefix      string                    `json:"domainPrefix"`
	RequestID         string                    `json:"requestId"`
	ExtendedRequestID string                    `json:"extendedRequestId"`
	Protocol          string                    `json:"protocol"`
	Identity          APIGatewayRequestIdentity `json:"identity"`
	ResourcePath      string                    `json:"resourcePath"`
	Path              string                    `json:"path"`
	Authorizer        map[string]interface{}    `json:"authorizer"`
	HTTPMethod        string                    `json:"httpMethod"`
	RequestTime       string                    `json:"requestTime"`
	RequestTimeEpoch  int64                     `json:"requestTimeEpoch"`
	APIID             string                    `json:"apiId"` // The API Gateway rest API Id
}

// APIGatewayProxyRequest contains data coming from the API Gateway proxy
type APIGatewayProxyRequest struct {
	Resource       string                        `json:"resource"` // The resource path defined in API Gateway
	PathParameters map[string]string             `json:"pathParameters"`
	StageVariables map[string]string             `json:"stageVariables"`
	RequestContext APIGatewayProxyRequestContext `json:"requestContext"`

	RawURL                          string                 `json:"rawUrl"`
	RawQuery                        string                 `json:"rawQuery"`
	Path                            string                 `json:"path"`
	HTTPMethod                      string                 `json:"httpMethod"`
	Headers                         map[string]string      `json:"headers"`
	MultiValueHeaders               map[string][]string    `json:"multiValueHeaders"`
	QueryStringParameters           map[string]string      `json:"queryStringParameters"`
	MultiValueQueryStringParameters map[string][]string    `json:"multiValueQueryStringParameters"`
	Body                            string                 `json:"body"`
	IsBase64Encoded                 bool                   `json:"isBase64Encoded"`
	Route                           string                 `json:"route,omitempty"`
	Blobs                           string                 `json:"blobs"`
	Flags                           map[string]interface{} `json:"flags,omitempty"`
	InvocationMetadata              InvocationMetadata     `json:"invocationMetadata,omitempty"`
	LogIngestionToken               string                 `json:"logToken,omitempty"`
}

// func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
func handler(ctx context.Context, request APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	lc, ok := lambdacontext.FromContext(ctx)
	if !ok {
		return &events.APIGatewayProxyResponse{
			StatusCode: 503,
			Body:       "Something went wrong",
		}, nil
	}

	cc := lc.ClientContext

	for key, value := range cc.Env {
		fmt.Printf("cc.Env.%s value is %v\n", key, value)
	}

	for key, value := range cc.Custom {
		fmt.Printf("cc.Custom.%s value is %v\n", key, value)
	}

	for key, value := range request.Headers {
		fmt.Printf("request.Headers.%s value is %v\n", key, value)
	}

	fmt.Println("request.Blobs", request.Blobs)
	fmt.Println("request.InvocationMetadata", request.InvocationMetadata)
	fmt.Printf("request: %+v\n", request)
	fmt.Printf("lc: %+v\n", lc)

	blob, _ := b64.StdEncoding.DecodeString(request.Blobs)

	fmt.Println(string(blob))

	var blobContext EnvironmentContext
	err := json.Unmarshal([]byte(blob), &blobContext)
	if err != nil {
		return nil, err
	}

	fmt.Printf("blobContext: %+v\n", blobContext)

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello, " + cc.Client.AppPackageName,
	}, nil
}

func main() {
	lambda.Start(handler)
}
