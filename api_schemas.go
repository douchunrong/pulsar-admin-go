/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"fmt"
	"strings"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type SchemasApiService service

/*
SchemasApiService Delete the schema of a topic
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param tenant
 * @param namespace
 * @param topic
 * @param optional nil or *DeleteSchemaOpts - Optional Parameters:
 * @param "Authoritative" (optional.Bool) - 
@return DeleteSchemaResponse
*/

type DeleteSchemaOpts struct {
	Authoritative optional.Bool
}

func (a *SchemasApiService) DeleteSchema(ctx context.Context, tenant string, namespace string, topic string, localVarOptionals *DeleteSchemaOpts) (DeleteSchemaResponse, *http.Response, error) {
	var (
		localVarHttpMethod   = http.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DeleteSchemaResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/schemas/{tenant}/{namespace}/{topic}/schema"
	localVarPath = strings.Replace(localVarPath, "{"+"tenant"+"}", fmt.Sprintf("%v", tenant), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", fmt.Sprintf("%v", namespace), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topic"+"}", fmt.Sprintf("%v", topic), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Authoritative.IsSet() {
		localVarQueryParams.Add("authoritative", parameterToString(localVarOptionals.Authoritative.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v DeleteSchemaResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
SchemasApiService Get the schema of a topic
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param tenant
 * @param namespace
 * @param topic
 * @param optional nil or *GetSchemaOpts - Optional Parameters:
 * @param "Authoritative" (optional.Bool) - 
@return GetSchemaResponse
*/

type GetSchemaOpts struct {
	Authoritative optional.Bool
}

func (a *SchemasApiService) GetSchema(ctx context.Context, tenant string, namespace string, topic string, localVarOptionals *GetSchemaOpts) (GetSchemaResponse, *http.Response, error) {
	var (
		localVarHttpMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetSchemaResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/schemas/{tenant}/{namespace}/{topic}/schema"
	localVarPath = strings.Replace(localVarPath, "{"+"tenant"+"}", fmt.Sprintf("%v", tenant), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", fmt.Sprintf("%v", namespace), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topic"+"}", fmt.Sprintf("%v", topic), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Authoritative.IsSet() {
		localVarQueryParams.Add("authoritative", parameterToString(localVarOptionals.Authoritative.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v GetSchemaResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
SchemasApiService Get the schema of a topic at a given version
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param tenant
 * @param namespace
 * @param topic
 * @param version
 * @param optional nil or *GetSchema_1Opts - Optional Parameters:
 * @param "Authoritative" (optional.Bool) - 
@return GetSchemaResponse
*/

type GetSchema_1Opts struct {
	Authoritative optional.Bool
}

func (a *SchemasApiService) GetSchema_1(ctx context.Context, tenant string, namespace string, topic string, version string, localVarOptionals *GetSchema_1Opts) (GetSchemaResponse, *http.Response, error) {
	var (
		localVarHttpMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetSchemaResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/schemas/{tenant}/{namespace}/{topic}/schema/{version}"
	localVarPath = strings.Replace(localVarPath, "{"+"tenant"+"}", fmt.Sprintf("%v", tenant), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", fmt.Sprintf("%v", namespace), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topic"+"}", fmt.Sprintf("%v", topic), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", fmt.Sprintf("%v", version), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Authoritative.IsSet() {
		localVarQueryParams.Add("authoritative", parameterToString(localVarOptionals.Authoritative.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v GetSchemaResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
SchemasApiService Update the schema of a topic
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param tenant
 * @param namespace
 * @param topic
 * @param optional nil or *PostSchemaOpts - Optional Parameters:
 * @param "Authoritative" (optional.Bool) - 
 * @param "Body" (optional.Interface of PostSchemaPayload) -  A JSON value presenting a schema playload. An example of the expected schema can be found down here.
@return PostSchemaResponse
*/

type PostSchemaOpts struct {
	Authoritative optional.Bool
	Body optional.Interface
}

func (a *SchemasApiService) PostSchema(ctx context.Context, tenant string, namespace string, topic string, localVarOptionals *PostSchemaOpts) (PostSchemaResponse, *http.Response, error) {
	var (
		localVarHttpMethod   = http.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PostSchemaResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/schemas/{tenant}/{namespace}/{topic}/schema"
	localVarPath = strings.Replace(localVarPath, "{"+"tenant"+"}", fmt.Sprintf("%v", tenant), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", fmt.Sprintf("%v", namespace), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topic"+"}", fmt.Sprintf("%v", topic), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Authoritative.IsSet() {
		localVarQueryParams.Add("authoritative", parameterToString(localVarOptionals.Authoritative.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {
		localVarOptionalBody, localVarOptionalBodyok := localVarOptionals.Body.Value().(PostSchemaPayload)
		if !localVarOptionalBodyok {
			return localVarReturnValue, nil, reportError("body should be PostSchemaPayload")
		}
		localVarPostBody = &localVarOptionalBody
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v PostSchemaResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
