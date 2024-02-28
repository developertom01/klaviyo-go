# Klaviyo's Unofficial Golang SDK

### Other Klaviyo Resources

- [API Reference](https://developers.klaviyo.com/en/v2024-02-15/reference/)
- [API Guides](https://developers.klaviyo.com/en/v2024-02-15/docs)
- [Postman Workspace](https://www.postman.com/klaviyo/workspace/klaviyo-developers)

## Design & Approach

This SDK is a thin wrapper around our API. See our API Reference for full documentation on API behavior.

This SDK exactly mirrors the organization and naming convention of the above language-agnostic resources, with a few namespace changes to make it fit better with Golang

## Organization

This SDK is organized into the following resources:

- AccountsApi
- CampaignsApi
- FlowsApi

## Installation

```sh
 go get github.com/developertom01/klaviyo-go
```

## Usage Example

- Initialization

 ```go
 var apiKey = "api-key"
    //Initialize options
 opt := options.NewOptionsWithDefaultValues().WithApiKey(apiKey)
    session := common.NewApiKeySession(opt, common.NewRetryOptionsWithDefaultValues())

    klaviyoApi := NewAccountsApi(session,nil)

 ```

## Filter Builder

```go
fb := commons.NewFilerBuilder()

//Simple operations

// equals
fb.Contains("name","test name") 
fb.Build()// "contains(name, test name)"

//Chain operators
fb.Equals("field1","value1").LessThan("field2","value2") 
fb.Build()// "equals(field1,value1),less-than(field2,value2)"

//Boolean Operators

op1 := commons.NewFilerBuilder()
api.Equals("field1","value1")
op2 := commons.NewFilerBuilder()
api.LessThan("field2","value2")

fb.And(op1,op2)

fb.Build() // and(equals(field1,value1),less-than(field2,value2))

```
