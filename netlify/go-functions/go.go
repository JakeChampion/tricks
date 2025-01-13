package main

import (
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

// Metadata type is a map representing arbitrary key-value pairs.
type Metadata map[string]interface{}

// ListResponseBlob represents a blob's metadata from a list response.
type ListResponseBlob struct {
	ETag         string `json:"etag"`
	LastModified string `json:"last_modified"`
	Size         int64  `json:"size"`
	Key          string `json:"key"`
}

// EnvironmentContext represents configuration context for API and deployment URLs.
// type EnvironmentContext struct {
// 	APIURL          string `json:"apiURL,omitempty"`
// 	DeployID        string `json:"deployID,omitempty"`
// 	EdgeURL         string `json:"edgeURL,omitempty"`
// 	PrimaryRegion   string `json:"primaryRegion,omitempty"`
// 	SiteID          string `json:"siteID,omitempty"`
// 	Token           string `json:"token,omitempty"`
// 	UncachedEdgeURL string `json:"uncachedEdgeURL,omitempty"`
// }

// BlobInput represents a possible input for a Blob, which can be a string, ArrayBuffer, or a Blob.
type BlobInput interface{}

// Fetcher type represents the Fetch function.
type Fetcher func(url string, options *http.Request) (*http.Response, error)

// HTTPMethod type represents HTTP request methods.
type HTTPMethod string

const (
	HTTPMethodDelete HTTPMethod = "delete"
	HTTPMethodGet    HTTPMethod = "get"
	HTTPMethodHead   HTTPMethod = "head"
	HTTPMethodPut    HTTPMethod = "put"
)

// SIGNED_URL_ACCEPT_HEADER is the constant for signed URL content type.
const SIGNED_URL_ACCEPT_HEADER = "application/json;type=signed-url"

// ConsistencyMode represents the consistency modes available.
type ConsistencyMode string

const (
	ConsistencyModeEventual ConsistencyMode = "eventual"
	ConsistencyModeStrong   ConsistencyMode = "strong"
)

// MakeStoreRequestOptions represents options for making a request to store.
type MakeStoreRequestOptions struct {
	Body        BlobInput         `json:"body,omitempty"`
	Consistency *ConsistencyMode  `json:"consistency,omitempty"`
	Headers     map[string]string `json:"headers,omitempty"`
	Key         string            `json:"key,omitempty"`
	Metadata    Metadata          `json:"metadata,omitempty"`
	Method      HTTPMethod        `json:"method"`
	Parameters  map[string]string `json:"parameters,omitempty"`
	StoreName   string            `json:"storeName,omitempty"`
}

// ClientOptions represents configuration options for the client.
type ClientOptions struct {
	APIURL          string          `json:"apiURL,omitempty"`
	Consistency     ConsistencyMode `json:"consistency,omitempty"`
	EdgeURL         string          `json:"edgeURL,omitempty"`
	Fetch           Fetcher         `json:"fetch,omitempty"`
	SiteID          string          `json:"siteID"`
	Token           string          `json:"token"`
	UncachedEdgeURL string          `json:"uncachedEdgeURL,omitempty"`
}

// InternalClientOptions extends ClientOptions with region.
type InternalClientOptions struct {
	ClientOptions
	Region string `json:"region,omitempty"`
}

// GetFinalRequestOptions represents the final options for a request.
type GetFinalRequestOptions struct {
	Consistency *ConsistencyMode  `json:"consistency,omitempty"`
	Key         string            `json:"key,omitempty"`
	Metadata    Metadata          `json:"metadata,omitempty"`
	Method      HTTPMethod        `json:"method"`
	Parameters  map[string]string `json:"parameters,omitempty"`
	StoreName   string            `json:"storeName,omitempty"`
}

// Client represents the client to interact with the API.
type Client struct {
	APIURL          string
	Consistency     ConsistencyMode
	EdgeURL         string
	Fetch           Fetcher
	Region          string
	SiteID          string
	Token           string
	UncachedEdgeURL string
}

// GetFinalRequest prepares the final request options.
func (c *Client) GetFinalRequest(options GetFinalRequestOptions) (map[int]interface{}, string, error) {
	// Simulate building the final request
	return nil, "", nil
}

// MakeRequest performs a request to the store.
func (c *Client) MakeRequest(options MakeStoreRequestOptions) (*http.Response, error) {
	// Simulate a store request
	return nil, nil
}

// BlobsConsistencyError represents an error related to blob consistency.
type BlobsConsistencyError struct {
	Message string
}

// NewBlobsConsistencyError creates a new BlobsConsistencyError.
func NewBlobsConsistencyError() *BlobsConsistencyError {
	return &BlobsConsistencyError{
		Message: "Blob consistency error",
	}
}

// Constants for store prefixes.
const (
	DEPLOY_STORE_PREFIX          = "deploy:"
	LEGACY_STORE_INTERNAL_PREFIX = "netlify-internal/legacy-namespace/"
	SITE_STORE_PREFIX            = "site:"
)

// BaseStoreOptions represents common options for store operations.
type BaseStoreOptions struct {
	Client      *Client
	Consistency *ConsistencyMode
}

// NamedStoreOptions represents options for a named store.
type NamedStoreOptions struct {
	BaseStoreOptions
	Name string `json:"name"`
}

// Store represents a store object in the system.
type Store struct {
	Client *Client
	Name   string
}

func validateStoreName(name string) error {
	if strings.Contains(name, "/") || strings.Contains(name, "%2F") {
		return fmt.Errorf("Store name must not contain forward slashes (/)")
	}

	if len(name) > 64 {
		return fmt.Errorf(
			"Store name must be a sequence of Unicode characters whose UTF-8 encoding is at most 64 bytes long",
		)
	}
	return nil
}

// NewStore creates a new store instance.
func NewStore(storeName string, client Client) (*Store, error) {
	err := validateStoreName(storeName)
	if err != nil {
		return nil, err
	}
	return &Store{
		Client: &client,
		Name:   storeName,
	}, nil
}

// Delete removes a key from the store.
func (s *Store) Delete(key string) error {
	// Simulate deleting the key
	return nil
}

// Get retrieves a value from the store.
func (s *Store) Get(key string) (string, error) {
	// Simulate getting the value
	return "", nil
}

// ListOptions represents options for listing store items.
type ListOptions struct {
	Directories bool   `json:"directories,omitempty"`
	Paginate    bool   `json:"paginate,omitempty"`
	Prefix      string `json:"prefix,omitempty"`
}

// ListResult represents the result of a list operation.
type ListResult struct {
	Blobs       []ListResultBlob `json:"blobs"`
	Directories []string         `json:"directories"`
}

// List lists store items based on the options.
func (s *Store) List(options *ListOptions) (*ListResult, error) {
	// Simulate listing the result
	return &ListResult{}, nil
}

// Set stores data in the store.
func (s *Store) Set(key string, data BlobInput, options *SetOptions) error {
	// Simulate storing the data
	return nil
}

// SetOptions represents options when setting data in the store.
type SetOptions struct {
	Metadata Metadata `json:"metadata,omitempty"`
}

// SetJSON stores JSON data in the store.
func (s *Store) SetJSON(key string, data interface{}, options *SetOptions) error {
	// Simulate storing JSON data
	return nil
}

// ListResultBlob represents a blob in the list result.
type ListResultBlob struct {
	ETag string `json:"etag"`
	Key  string `json:"key"`
}

// ///////////////////////////////////////////////////////////////////////////
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

// getStore: {
// 	(name: string): Store
// 	(options: GetStoreOptions): Store
//   } = (input) => {
// 	if (typeof input === 'string') {
// 	  const clientOptions = getClientOptions({})
// 	  const client = new Client(clientOptions)

// 	  return new Store({ client, name: input })
// 	}

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

	deploy_id := request.Headers["x-nf-deploy-id"]
	api_url := "https://api.netlify.com"
	site_id := request.Headers["x-nf-site-id"]
	region := "auto"
	println(deploy_id, api_url, site_id, region)

	store, err := NewStore("construction", Client{
		SiteID:      site_id,
		Token:       blobContext.Token,
		Consistency: ConsistencyModeEventual,
	})
	if err != nil {
		return nil, err
	}
	fmt.Printf("store: %+v\n", store)

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello, " + cc.Client.AppPackageName,
	}, nil
}

func main() {
	lambda.Start(handler)
}
