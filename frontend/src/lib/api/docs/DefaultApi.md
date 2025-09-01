# DefaultApi

All URIs are relative to *http://localhost:8080*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**apiProfileBasicInfoGet**](#apiprofilebasicinfoget) | **GET** /api/profile/basic-info | 基本情報を取得する|
|[**apiProfileBasicInfoPut**](#apiprofilebasicinfoput) | **PUT** /api/profile/basic-info | 基本情報を更新する|

# **apiProfileBasicInfoGet**
> BasicInfo apiProfileBasicInfoGet()

現在登録されているユーザーの基本情報を返します。

### Example

```typescript
import {
    DefaultApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

const { status, data } = await apiInstance.apiProfileBasicInfoGet();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**BasicInfo**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | 取得成功 |  -  |
|**500** | サーバーエラー |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiProfileBasicInfoPut**
> UpdateResponse apiProfileBasicInfoPut(basicInfo)

学籍番号や名前、自己紹介などの基本情報を編集します。

### Example

```typescript
import {
    DefaultApi,
    Configuration,
    BasicInfo
} from './api';

const configuration = new Configuration();
const apiInstance = new DefaultApi(configuration);

let basicInfo: BasicInfo; //

const { status, data } = await apiInstance.apiProfileBasicInfoPut(
    basicInfo
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **basicInfo** | **BasicInfo**|  | |


### Return type

**UpdateResponse**

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

