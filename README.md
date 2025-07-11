# Go API client for eskizapi

Отправляйте СМС по всему миру, в любом количестве!

В тестовом статусе для отправки тестовых смс сообщений, Вы можете использовать только нижеуказанные тексты:

- Это тест от Eskiz
    
- Bu Eskiz dan test
    
- This is test from Eskiz

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build date: 2025-07-11T22:39:32.696575900+05:00[Asia/Tashkent]
- Generator version: 7.9.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.iota.uz](https://www.iota.uz)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import eskizapi "github.com/iota-uz/eskiz/eskizapi"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `eskizapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), eskizapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `eskizapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), eskizapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `eskizapi.ContextOperationServerIndices` and `eskizapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), eskizapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), eskizapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://notify.eskiz.uz*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**ExportMessages**](docs/DefaultApi.md#exportmessages) | **Post** /api/message/export | Экспортировать в формате CSV
*DefaultApi* | [**GetDispatchStatus**](docs/DefaultApi.md#getdispatchstatus) | **Post** /api/message/sms/get-dispatch-status | Статус рассылки
*DefaultApi* | [**GetNicknames**](docs/DefaultApi.md#getnicknames) | **Get** /api/nick/me | Получить список ник
*DefaultApi* | [**GetSmsLogs**](docs/DefaultApi.md#getsmslogs) | **Get** /api/logs/sms/{id} | Системные логи
*DefaultApi* | [**GetSmsStatusById**](docs/DefaultApi.md#getsmsstatusbyid) | **Get** /api/message/sms/status_by_id/{id} | Получить статус по ID
*DefaultApi* | [**GetTotalByDispatch**](docs/DefaultApi.md#gettotalbydispatch) | **Post** /api/report/total-by-dispatch | Расходы по рассылке
*DefaultApi* | [**GetTotalByMonth**](docs/DefaultApi.md#gettotalbymonth) | **Get** /api/report/total-by-month | Итого по месяцам
*DefaultApi* | [**GetTotalByRange**](docs/DefaultApi.md#gettotalbyrange) | **Post** /api/report/total-by-range | Расходы по датам
*DefaultApi* | [**GetTotalBySmsc**](docs/DefaultApi.md#gettotalbysmsc) | **Post** /api/report/total-by-smsc | Итого по компаниям
*DefaultApi* | [**GetUserInfo**](docs/DefaultApi.md#getuserinfo) | **Get** /api/auth/user | Данные ползователья
*DefaultApi* | [**GetUserLimit**](docs/DefaultApi.md#getuserlimit) | **Get** /api/user/get-limit | Получить баланс
*DefaultApi* | [**GetUserMessages**](docs/DefaultApi.md#getusermessages) | **Post** /api/message/sms/get-user-messages | Детализация
*DefaultApi* | [**GetUserMessagesByDispatch**](docs/DefaultApi.md#getusermessagesbydispatch) | **Post** /api/message/sms/get-user-messages-by-dispatch | Получить СМС по рассылке
*DefaultApi* | [**GetUserTemplates**](docs/DefaultApi.md#getusertemplates) | **Get** /api/user/templates | Получить список шаблонов
*DefaultApi* | [**GetUserTotals**](docs/DefaultApi.md#getusertotals) | **Post** /api/user/totals | Итог отправленных СМС
*DefaultApi* | [**Login**](docs/DefaultApi.md#login) | **Post** /api/auth/login | Получить токен
*DefaultApi* | [**NormalizeSms**](docs/DefaultApi.md#normalizesms) | **Post** /api/message/sms/normalizer | Нормализации SMS
*DefaultApi* | [**RefreshToken**](docs/DefaultApi.md#refreshtoken) | **Patch** /api/auth/refresh | Обновить токен
*DefaultApi* | [**SendSms**](docs/DefaultApi.md#sendsms) | **Post** /api/message/sms/send | Отправить СМС
*DefaultApi* | [**SendSmsBatch**](docs/DefaultApi.md#sendsmsbatch) | **Post** /api/message/sms/send-batch | Отправить СМС рассылка
*DefaultApi* | [**SendSmsGlobal**](docs/DefaultApi.md#sendsmsglobal) | **Post** /api/message/sms/send-global | Отправить Международный СМС
*DefaultApi* | [**SendTemplate**](docs/DefaultApi.md#sendtemplate) | **Post** /api/user/template | Отправить шаблон


## Documentation For Models

 - [DispatchStatusResponse](docs/DispatchStatusResponse.md)
 - [LinkItem](docs/LinkItem.md)
 - [LogMessage](docs/LogMessage.md)
 - [LoginData](docs/LoginData.md)
 - [LoginResponse](docs/LoginResponse.md)
 - [MessageItem](docs/MessageItem.md)
 - [MessagePart](docs/MessagePart.md)
 - [MonthlyReportItem](docs/MonthlyReportItem.md)
 - [NormalizeSmsResponse](docs/NormalizeSmsResponse.md)
 - [PaginationData](docs/PaginationData.md)
 - [Recommendation](docs/Recommendation.md)
 - [RefreshTokenResponse](docs/RefreshTokenResponse.md)
 - [SendSmsBatchRequest](docs/SendSmsBatchRequest.md)
 - [SendSmsBatchRequestMessagesInner](docs/SendSmsBatchRequestMessagesInner.md)
 - [SendSmsBatchResponse](docs/SendSmsBatchResponse.md)
 - [SendSmsResponse](docs/SendSmsResponse.md)
 - [SendTemplateResponse](docs/SendTemplateResponse.md)
 - [SmsCallbackWebhook](docs/SmsCallbackWebhook.md)
 - [SmsLogsResponse](docs/SmsLogsResponse.md)
 - [SmsStatusData](docs/SmsStatusData.md)
 - [SmsStatusPart](docs/SmsStatusPart.md)
 - [SmsStatusResponse](docs/SmsStatusResponse.md)
 - [SmscReportItem](docs/SmscReportItem.md)
 - [SpecialCharacter](docs/SpecialCharacter.md)
 - [StatusCount](docs/StatusCount.md)
 - [TemplateItem](docs/TemplateItem.md)
 - [TemplatesListResponse](docs/TemplatesListResponse.md)
 - [TotalByMonthResponse](docs/TotalByMonthResponse.md)
 - [TotalBySmscResponse](docs/TotalBySmscResponse.md)
 - [UserBalance](docs/UserBalance.md)
 - [UserInfo](docs/UserInfo.md)
 - [UserInfoResponse](docs/UserInfoResponse.md)
 - [UserLimitResponse](docs/UserLimitResponse.md)
 - [UserMessagesResponse](docs/UserMessagesResponse.md)
 - [UserTotalsItem](docs/UserTotalsItem.md)
 - [UserTotalsResponse](docs/UserTotalsResponse.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### bearerAuth

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), eskizapi.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

danil@iota.uz

