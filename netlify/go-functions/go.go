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

// APIGatewayProxyRequest contains data coming from the API Gateway proxy
type APIGatewayProxyRequest struct {
	// Resource                        string                        `json:"resource"` // The resource path defined in API Gateway
	// Path                            string                        `json:"path"`     // The url path for the caller
	// HTTPMethod                      string                        `json:"httpMethod"`
	// Headers                         map[string]string             `json:"headers"`
	// MultiValueHeaders               map[string][]string           `json:"multiValueHeaders"`
	// QueryStringParameters           map[string]string             `json:"queryStringParameters"`
	// MultiValueQueryStringParameters map[string][]string           `json:"multiValueQueryStringParameters"`
	// PathParameters                  map[string]string             `json:"pathParameters"`
	// StageVariables                  map[string]string             `json:"stageVariables"`
	// RequestContext                  APIGatewayProxyRequestContext `json:"requestContext"`
	// Body                            string                        `json:"body"`
	// IsBase64Encoded                 bool                          `json:"isBase64Encoded,omitempty"`
	// Blobs                           string                 `json:"blobs,omitempty"`

	RawURL                          string                 `json:"rawUrl"`
	RawQuery                        string                 `json:"rawQuery"`
	Path                            string                 `json:"path"`
	Method                          string                 `json:"httpMethod"`
	Headers                         map[string]string      `json:"headers"`
	MultiValueHeaders               map[string][]string    `json:"multiValueHeaders"`
	Params                          map[string]string      `json:"queryStringParameters"`
	MultiValueQueryStringParameters map[string][]string    `json:"multiValueQueryStringParameters"`
	Body                            string                 `json:"body"`
	IsBase64                        bool                   `json:"isBase64Encoded"`
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
