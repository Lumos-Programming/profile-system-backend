# DefaultApi

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**apiProfileBasicInfoPut**](#apiprofilebasicinfoput) | **PUT** /api/profile/basic-info | 基本情報を更新する|

# **apiProfileBasicInfoPut**
> ApiProfileBasicInfoPut200Response apiProfileBasicInfoPut(apiProfileBasicInfoPutRequest)

学籍番号や名前、自己紹介などの基本情報を編集します。

### Example

```typescript
import {
    DefaultApi,
    Configuration,
    ApiProfileBasicInfoPutRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let apiProfileBasicInfoPutRequest: ApiProfileBasicInfoPutRequest; //

const { status, data } = await apiInstance.apiProfileBasicInfoPut(
    apiProfileBasicInfoPutRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **apiProfileBasicInfoPutRequest** | **ApiProfileBasicInfoPutRequest**|  | |


### Return type

**ApiProfileBasicInfoPut200Response**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | 更新成功 |  -  |
|**400** | バリデーションエラー |  -  |
|**500** | サーバーエラー |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

