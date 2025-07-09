# \DefaultApi

All URIs are relative to *https://notify.eskiz.uz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportMessages**](DefaultApi.md#ExportMessages) | **Post** /api/message/export | Экспортировать в формате CSV
[**GetDispatchStatus**](DefaultApi.md#GetDispatchStatus) | **Post** /api/message/sms/get-dispatch-status | Статус рассылки
[**GetNicknames**](DefaultApi.md#GetNicknames) | **Get** /api/nick/me | Получить список ник
[**GetSmsLogs**](DefaultApi.md#GetSmsLogs) | **Get** /api/logs/sms/{id} | Системные логи
[**GetSmsStatusById**](DefaultApi.md#GetSmsStatusById) | **Get** /api/message/sms/status_by_id/{id} | Получить статус по ID
[**GetTotalByDispatch**](DefaultApi.md#GetTotalByDispatch) | **Post** /api/report/total-by-dispatch | Расходы по рассылке
[**GetTotalByMonth**](DefaultApi.md#GetTotalByMonth) | **Get** /api/report/total-by-month | Итого по месяцам
[**GetTotalByRange**](DefaultApi.md#GetTotalByRange) | **Post** /api/report/total-by-range | Расходы по датам
[**GetTotalBySmsc**](DefaultApi.md#GetTotalBySmsc) | **Post** /api/report/total-by-smsc | Итого по компаниям
[**GetUserInfo**](DefaultApi.md#GetUserInfo) | **Get** /api/auth/user | Данные ползователья
[**GetUserLimit**](DefaultApi.md#GetUserLimit) | **Get** /api/user/get-limit | Получить баланс
[**GetUserMessages**](DefaultApi.md#GetUserMessages) | **Post** /api/message/sms/get-user-messages | Детализация
[**GetUserMessagesByDispatch**](DefaultApi.md#GetUserMessagesByDispatch) | **Post** /api/message/sms/get-user-messages-by-dispatch | Получить СМС по рассылке
[**GetUserTemplates**](DefaultApi.md#GetUserTemplates) | **Get** /api/user/templates | Получить список шаблонов
[**GetUserTotals**](DefaultApi.md#GetUserTotals) | **Post** /api/user/totals | Итог отправленных СМС
[**Login**](DefaultApi.md#Login) | **Post** /api/auth/login | Получить токен
[**NormalizeSms**](DefaultApi.md#NormalizeSms) | **Post** /api/message/sms/normalizer | Нормализации SMS
[**RefreshToken**](DefaultApi.md#RefreshToken) | **Patch** /api/auth/refresh | Обновить токен
[**SendSms**](DefaultApi.md#SendSms) | **Post** /api/message/sms/send | Отправить СМС
[**SendSmsBatch**](DefaultApi.md#SendSmsBatch) | **Post** /api/message/sms/send-batch | Отправить СМС рассылка
[**SendSmsGlobal**](DefaultApi.md#SendSmsGlobal) | **Post** /api/message/sms/send-global | Отправить Международный СМС
[**SendTemplate**](DefaultApi.md#SendTemplate) | **Post** /api/user/template | Отправить шаблон



## ExportMessages

> *os.File ExportMessages(ctx).Status(status).End(end).Month(month).Start(start).Year(year).Execute()

Экспортировать в формате CSV



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/eskiz/eskizapi"
)

func main() {
	status := "status_example" // string | `all` или `delivered` или `rejected` (optional)
	end := "end_example" // string | YYYY-MM-DD HH:mm:ss (optional)
	month := "month_example" // string | Месяц (optional)
	start := "start_example" // string | YYYY-MM-DD HH:mm:ss (optional)
	year := "year_example" // string | Год (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ExportMessages(context.Background()).Status(status).End(end).Month(month).Start(start).Year(year).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ExportMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportMessages`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ExportMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | &#x60;all&#x60; или &#x60;delivered&#x60; или &#x60;rejected&#x60; | 
 **end** | **string** | YYYY-MM-DD HH:mm:ss | 
 **month** | **string** | Месяц | 
 **start** | **string** | YYYY-MM-DD HH:mm:ss | 
 **year** | **string** | Год | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDispatchStatus

> DispatchStatusResponse GetDispatchStatus(ctx).DispatchId(dispatchId).UserId(userId).Execute()

Статус рассылки



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/eskiz/eskizapi"
)

func main() {
	dispatchId := "dispatchId_example" // string | ID рассылки (optional)
	userId := "userId_example" // string | ID пользователя (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetDispatchStatus(context.Background()).DispatchId(dispatchId).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetDispatchStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDispatchStatus`: DispatchStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetDispatchStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDispatchStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dispatchId** | **string** | ID рассылки | 
 **userId** | **string** | ID пользователя | 

### Return type

[**DispatchStatusResponse**](DispatchStatusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNicknames

> []string GetNicknames(ctx).Execute()

Получить список ник



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/eskiz/eskizapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetNicknames(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetNicknames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNicknames`: []string
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetNicknames`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNicknamesRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSmsLogs

> SmsLogsResponse GetSmsLogs(ctx, id).Execute()

Системные логи



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/eskiz/eskizapi"
)

func main() {
	id := "id_example" // string | ID СМС

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetSmsLogs(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSmsLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSmsLogs`: SmsLogsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSmsLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID СМС | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSmsLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SmsLogsResponse**](SmsLogsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSmsStatusById

> SmsStatusResponse GetSmsStatusById(ctx, id).Execute()

Получить статус по ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/eskiz/eskizapi"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetSmsStatusById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSmsStatusById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSmsStatusById`: SmsStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSmsStatusById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSmsStatusByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SmsStatusResponse**](SmsStatusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTotalByDispatch

> UserTotalsResponse GetTotalByDispatch(ctx).Status(status).DispatchId(dispatchId).IsAd(isAd).Execute()

Расходы по рассылке



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/eskiz/eskizapi"
)

func main() {
	status := "status_example" // string | [пустой] - все \"delivered\" - только частично и полностью доставленное \"rejected\" - только не доставленное  (optional)
	dispatchId := "dispatchId_example" // string | ID рассылки (optional)
	isAd := "isAd_example" // string | [пустой] - все 1 - только рекламное 0 - только сервисное (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetTotalByDispatch(context.Background()).Status(status).DispatchId(dispatchId).IsAd(isAd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTotalByDispatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTotalByDispatch`: UserTotalsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTotalByDispatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTotalByDispatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | [пустой] - все \&quot;delivered\&quot; - только частично и полностью доставленное \&quot;rejected\&quot; - только не доставленное  | 
 **dispatchId** | **string** | ID рассылки | 
 **isAd** | **string** | [пустой] - все 1 - только рекламное 0 - только сервисное | 

### Return type

[**UserTotalsResponse**](UserTotalsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTotalByMonth

> TotalByMonthResponse GetTotalByMonth(ctx).Year(year).Execute()

Итого по месяцам



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/eskiz/eskizapi"
)

func main() {
	year := "year_example" // string | Год (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetTotalByMonth(context.Background()).Year(year).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTotalByMonth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTotalByMonth`: TotalByMonthResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTotalByMonth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTotalByMonthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **string** | Год | 

### Return type

[**TotalByMonthResponse**](TotalByMonthResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTotalByRange

> UserTotalsResponse GetTotalByRange(ctx).Status(status).IsAd(isAd).StartDate(startDate).ToDate(toDate).Execute()

Расходы по датам



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/eskiz/eskizapi"
)

func main() {
	status := "status_example" // string | [пустой] - все \"delivered\" - только частично и полностью доставленное \"rejected\" - только не доставленное  (optional)
	isAd := "isAd_example" // string | [пустой] - все 1 - только рекламное 0 - только сервисное (optional)
	startDate := "startDate_example" // string | С `%Y-%m-%d %H:%M` (optional)
	toDate := "toDate_example" // string | По `%Y-%m-%d %H:%M` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetTotalByRange(context.Background()).Status(status).IsAd(isAd).StartDate(startDate).ToDate(toDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTotalByRange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTotalByRange`: UserTotalsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTotalByRange`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTotalByRangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | [пустой] - все \&quot;delivered\&quot; - только частично и полностью доставленное \&quot;rejected\&quot; - только не доставленное  | 
 **isAd** | **string** | [пустой] - все 1 - только рекламное 0 - только сервисное | 
 **startDate** | **string** | С &#x60;%Y-%m-%d %H:%M&#x60; | 
 **toDate** | **string** | По &#x60;%Y-%m-%d %H:%M&#x60; | 

### Return type

[**UserTotalsResponse**](UserTotalsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTotalBySmsc

> TotalBySmscResponse GetTotalBySmsc(ctx).Month(month).Year(year).Execute()

Итого по компаниям



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/eskiz/eskizapi"
)

func main() {
	month := "month_example" // string | Месяц (optional)
	year := "year_example" // string | Год (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetTotalBySmsc(context.Background()).Month(month).Year(year).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTotalBySmsc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTotalBySmsc`: TotalBySmscResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTotalBySmsc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTotalBySmscRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **month** | **string** | Месяц | 
 **year** | **string** | Год | 

### Return type

[**TotalBySmscResponse**](TotalBySmscResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserInfo

> UserInfoResponse GetUserInfo(ctx).Execute()

Данные ползователья



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/eskiz/eskizapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetUserInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetUserInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserInfo`: UserInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetUserInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserInfoRequest struct via the builder pattern


### Return type

[**UserInfoResponse**](UserInfoResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserLimit

> UserLimitResponse GetUserLimit(ctx).Execute()

Получить баланс



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/eskiz/eskizapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetUserLimit(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetUserLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserLimit`: UserLimitResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetUserLimit`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserLimitRequest struct via the builder pattern


### Return type

[**UserLimitResponse**](UserLimitResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserMessages

> UserMessagesResponse GetUserMessages(ctx).Status(status).Count(count).IsAd(isAd).PageSize(pageSize).StartDate(startDate).ToDate(toDate).Execute()

Детализация



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/eskiz/eskizapi"
)

func main() {
	status := "status_example" // string | [пустой] - все \"delivered\" - только частично и полностью доставленное \"rejected\" - только не доставленное  (optional)
	count := "count_example" // string | 1 - Если необходимо получить итог по статусу (optional)
	isAd := "isAd_example" // string | [пустой] - все 1 - только рекламное 0 - только сервисное (optional)
	pageSize := "pageSize_example" // string | Количество SMS сообщений (с 20 до 200) (optional)
	startDate := "startDate_example" // string | С `%Y-%m-%d %H:%M` (optional)
	toDate := "toDate_example" // string | По `%Y-%m-%d %H:%M` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetUserMessages(context.Background()).Status(status).Count(count).IsAd(isAd).PageSize(pageSize).StartDate(startDate).ToDate(toDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetUserMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserMessages`: UserMessagesResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetUserMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | [пустой] - все \&quot;delivered\&quot; - только частично и полностью доставленное \&quot;rejected\&quot; - только не доставленное  | 
 **count** | **string** | 1 - Если необходимо получить итог по статусу | 
 **isAd** | **string** | [пустой] - все 1 - только рекламное 0 - только сервисное | 
 **pageSize** | **string** | Количество SMS сообщений (с 20 до 200) | 
 **startDate** | **string** | С &#x60;%Y-%m-%d %H:%M&#x60; | 
 **toDate** | **string** | По &#x60;%Y-%m-%d %H:%M&#x60; | 

### Return type

[**UserMessagesResponse**](UserMessagesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserMessagesByDispatch

> UserMessagesResponse GetUserMessagesByDispatch(ctx).Status(status).Count(count).DispatchId(dispatchId).IsAd(isAd).Execute()

Получить СМС по рассылке



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/eskiz/eskizapi"
)

func main() {
	status := "status_example" // string | [пустой] - все \"delivered\" - только частично и полностью доставленное \"rejected\" - только не доставленное  (optional)
	count := "count_example" // string | 1 - Если необходимо получить итог по статусу (optional)
	dispatchId := "dispatchId_example" // string | ID рассылки (optional)
	isAd := "isAd_example" // string | [пустой] - все 1 - только рекламное 0 - только сервисное (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetUserMessagesByDispatch(context.Background()).Status(status).Count(count).DispatchId(dispatchId).IsAd(isAd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetUserMessagesByDispatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserMessagesByDispatch`: UserMessagesResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetUserMessagesByDispatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserMessagesByDispatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | [пустой] - все \&quot;delivered\&quot; - только частично и полностью доставленное \&quot;rejected\&quot; - только не доставленное  | 
 **count** | **string** | 1 - Если необходимо получить итог по статусу | 
 **dispatchId** | **string** | ID рассылки | 
 **isAd** | **string** | [пустой] - все 1 - только рекламное 0 - только сервисное | 

### Return type

[**UserMessagesResponse**](UserMessagesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserTemplates

> TemplatesListResponse GetUserTemplates(ctx).Execute()

Получить список шаблонов



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/eskiz/eskizapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetUserTemplates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetUserTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserTemplates`: TemplatesListResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetUserTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserTemplatesRequest struct via the builder pattern


### Return type

[**TemplatesListResponse**](TemplatesListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserTotals

> UserTotalsResponse GetUserTotals(ctx).IsGlobal(isGlobal).Month(month).Year(year).Execute()

Итог отправленных СМС



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/eskiz/eskizapi"
)

func main() {
	isGlobal := "isGlobal_example" // string |  (optional)
	month := "month_example" // string |  (optional)
	year := "year_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetUserTotals(context.Background()).IsGlobal(isGlobal).Month(month).Year(year).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetUserTotals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserTotals`: UserTotalsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetUserTotals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserTotalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isGlobal** | **string** |  | 
 **month** | **string** |  | 
 **year** | **string** |  | 

### Return type

[**UserTotalsResponse**](UserTotalsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Login

> LoginResponse Login(ctx).Email(email).Password(password).Execute()

Получить токен



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/eskiz/eskizapi"
)

func main() {
	email := "email_example" // string |  (optional)
	password := "password_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.Login(context.Background()).Email(email).Password(password).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Login``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Login`: LoginResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Login`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | 
 **password** | **string** |  | 

### Return type

[**LoginResponse**](LoginResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NormalizeSms

> NormalizeSmsResponse NormalizeSms(ctx).Message(message).Execute()

Нормализации SMS



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/eskiz/eskizapi"
)

func main() {
	message := "message_example" // string | Сообщение (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.NormalizeSms(context.Background()).Message(message).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.NormalizeSms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NormalizeSms`: NormalizeSmsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.NormalizeSms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNormalizeSmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **message** | **string** | Сообщение | 

### Return type

[**NormalizeSmsResponse**](NormalizeSmsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshToken

> RefreshTokenResponse RefreshToken(ctx).Execute()

Обновить токен



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/eskiz/eskizapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.RefreshToken(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RefreshToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshToken`: RefreshTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RefreshToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshTokenRequest struct via the builder pattern


### Return type

[**RefreshTokenResponse**](RefreshTokenResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendSms

> SendSmsResponse SendSms(ctx).CallbackUrl(callbackUrl).From(from).Message(message).MobilePhone(mobilePhone).Execute()

Отправить СМС



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/eskiz/eskizapi"
)

func main() {
	callbackUrl := "callbackUrl_example" // string | Это необязательное поле, которое используется для автоматического получения смс-статуса с сервера. Укажите URL-адрес обратного вызова, по которому вы будете получать данные POST в следующем формате: `{\\\"request_id\\\": \\\"UUID\\\", \\\"message_id\\\": \\\"4385062\\\", \\\"user_sms_id\\\": \\\"vash_ID_zdes\\\", \\\"country\\\": \\\"UZ\\\", \\\"phone_number\\\": \\\"998991234567\\\", \\\"sms_count\\\": \\\"1\\\", \\\"status\\\" : \\\"DELIVRD\\\", \\\"status_date\\\": \\\"2021-04-02 00:39:36\\\"}`.   request_id — это значение, возвращаемое в поле \\\"id\\\" в результате выполнения API запросов send, send-batch и send-global. (optional)
	from := "from_example" // string | Для того, чтобы использовать никнейм, вам нужно поменять 4546 на свои. (optional)
	message := "message_example" // string | Сообщение (optional)
	mobilePhone := "mobilePhone_example" // string | Номер телефона (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.SendSms(context.Background()).CallbackUrl(callbackUrl).From(from).Message(message).MobilePhone(mobilePhone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SendSms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendSms`: SendSmsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SendSms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendSmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **callbackUrl** | **string** | Это необязательное поле, которое используется для автоматического получения смс-статуса с сервера. Укажите URL-адрес обратного вызова, по которому вы будете получать данные POST в следующем формате: &#x60;{\\\&quot;request_id\\\&quot;: \\\&quot;UUID\\\&quot;, \\\&quot;message_id\\\&quot;: \\\&quot;4385062\\\&quot;, \\\&quot;user_sms_id\\\&quot;: \\\&quot;vash_ID_zdes\\\&quot;, \\\&quot;country\\\&quot;: \\\&quot;UZ\\\&quot;, \\\&quot;phone_number\\\&quot;: \\\&quot;998991234567\\\&quot;, \\\&quot;sms_count\\\&quot;: \\\&quot;1\\\&quot;, \\\&quot;status\\\&quot; : \\\&quot;DELIVRD\\\&quot;, \\\&quot;status_date\\\&quot;: \\\&quot;2021-04-02 00:39:36\\\&quot;}&#x60;.   request_id — это значение, возвращаемое в поле \\\&quot;id\\\&quot; в результате выполнения API запросов send, send-batch и send-global. | 
 **from** | **string** | Для того, чтобы использовать никнейм, вам нужно поменять 4546 на свои. | 
 **message** | **string** | Сообщение | 
 **mobilePhone** | **string** | Номер телефона | 

### Return type

[**SendSmsResponse**](SendSmsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendSmsBatch

> SendSmsBatchResponse SendSmsBatch(ctx).SendSmsBatchRequest(sendSmsBatchRequest).Execute()

Отправить СМС рассылка



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/eskiz/eskizapi"
)

func main() {
	sendSmsBatchRequest := *openapiclient.NewSendSmsBatchRequest() // SendSmsBatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.SendSmsBatch(context.Background()).SendSmsBatchRequest(sendSmsBatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SendSmsBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendSmsBatch`: SendSmsBatchResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SendSmsBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendSmsBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendSmsBatchRequest** | [**SendSmsBatchRequest**](SendSmsBatchRequest.md) |  | 

### Return type

[**SendSmsBatchResponse**](SendSmsBatchResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendSmsGlobal

> SendSmsResponse SendSmsGlobal(ctx).CallbackUrl(callbackUrl).CountryCode(countryCode).Message(message).MobilePhone(mobilePhone).Unicode(unicode).Execute()

Отправить Международный СМС



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/eskiz/eskizapi"
)

func main() {
	callbackUrl := "callbackUrl_example" // string | Это необязательное поле, которое используется для автоматического получения смс-статуса с сервера. Укажите URL-адрес обратного вызова, по которому вы будете получать данные POST в следующем формате: `{\\\"request_id\\\": \\\"UUID\\\", \\\"message_id\\\": \\\"4385062\\\", \\\"user_sms_id\\\": \\\"vash_ID_zdes\\\", \\\"country\\\": \\\"UZ\\\", \\\"phone_number\\\": \\\"998991234567\\\", \\\"sms_count\\\": \\\"1\\\", \\\"status\\\" : \\\"DELIVRD\\\", \\\"status_date\\\": \\\"2021-04-02 00:39:36\\\"}`  request_id — это значение, возвращаемое в поле \\\"id\\\" в результате выполнения API запросов send, send-batch и send-global. (optional)
	countryCode := "countryCode_example" // string | Страна (optional)
	message := "message_example" // string | Сообщение (optional)
	mobilePhone := "mobilePhone_example" // string | Номер телефона (optional)
	unicode := "unicode_example" // string | 1 - если хотите отправить на кирилице (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.SendSmsGlobal(context.Background()).CallbackUrl(callbackUrl).CountryCode(countryCode).Message(message).MobilePhone(mobilePhone).Unicode(unicode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SendSmsGlobal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendSmsGlobal`: SendSmsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SendSmsGlobal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendSmsGlobalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **callbackUrl** | **string** | Это необязательное поле, которое используется для автоматического получения смс-статуса с сервера. Укажите URL-адрес обратного вызова, по которому вы будете получать данные POST в следующем формате: &#x60;{\\\&quot;request_id\\\&quot;: \\\&quot;UUID\\\&quot;, \\\&quot;message_id\\\&quot;: \\\&quot;4385062\\\&quot;, \\\&quot;user_sms_id\\\&quot;: \\\&quot;vash_ID_zdes\\\&quot;, \\\&quot;country\\\&quot;: \\\&quot;UZ\\\&quot;, \\\&quot;phone_number\\\&quot;: \\\&quot;998991234567\\\&quot;, \\\&quot;sms_count\\\&quot;: \\\&quot;1\\\&quot;, \\\&quot;status\\\&quot; : \\\&quot;DELIVRD\\\&quot;, \\\&quot;status_date\\\&quot;: \\\&quot;2021-04-02 00:39:36\\\&quot;}&#x60;  request_id — это значение, возвращаемое в поле \\\&quot;id\\\&quot; в результате выполнения API запросов send, send-batch и send-global. | 
 **countryCode** | **string** | Страна | 
 **message** | **string** | Сообщение | 
 **mobilePhone** | **string** | Номер телефона | 
 **unicode** | **string** | 1 - если хотите отправить на кирилице | 

### Return type

[**SendSmsResponse**](SendSmsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendTemplate

> SendTemplateResponse SendTemplate(ctx).Template(template).Execute()

Отправить шаблон



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/iota-uz/eskiz/eskizapi"
)

func main() {
	template := "template_example" // string | шаблон или текс сообщение (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.SendTemplate(context.Background()).Template(template).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SendTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendTemplate`: SendTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SendTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **template** | **string** | шаблон или текс сообщение | 

### Return type

[**SendTemplateResponse**](SendTemplateResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

