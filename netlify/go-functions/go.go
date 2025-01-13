package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

// const DEFAULT_API_HOST = "api.netlify.com"

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
	Blobs                           string                 `json:"blobs,omitempty"`
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
	fmt.Println("cc.Env.NETLIFY_BLOBS_CONTEXT", cc.Env["NETLIFY_BLOBS_CONTEXT"])

	for key, value := range cc.Custom {
		fmt.Printf("cc.Custom.%s value is %v\n", key, value)
	}

	for key, value := range request.Headers {
		fmt.Printf("request.Headers.%s value is %v\n", key, value)
	}

	fmt.Println(request)
	fmt.Println("")
	fmt.Println(lc)

	// deploy_id := request.Headers["x-nf-deploy-id"]
	// api_url := "https://api.netlify.com"
	// site_id := request.Headers["x-nf-site-id"]
	// const payload = {
	// apiURL: "https://api.netlify.com",
	// deployID: reques,
	// siteID: siteId,
	// token,
	// }
	// const encodedPayload = Buffer.from(JSON.stringify(payload)).toString('base64')

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello, " + cc.Client.AppPackageName,
	}, nil
}

func main() {
	lambda.Start(handler)
}
