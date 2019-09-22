## http://127.0.0.1/GetFeature



### Method

GET/POST

### Request

|   参数名    | 类型  |              说明              | 是否必须 |
|-------------|-------|--------------------------------|----------|
| lo          | Point |  One corner of the rectangle.  | Y        |
|   latitude  | int32 |                                | Y        |
|   longitude | int32 |                                | Y        |
| hi          | Point |  The other corner of the       | Y        |
|             |       | rectangle.                     |          |
|   latitude  | int32 |                                | Y        |
|   longitude | int32 |                                | Y        |


### Reply

|   参数名    |  类型  |              说明              | 是否必须 |
|-------------|--------|--------------------------------|----------|
| name        | string |  The name of the feature.      |          |
| location    | Point  |  The point where the feature   |          |
|             |        | is detected.                   |          |
|   latitude  | int32  |                                |          |
|   longitude | int32  |                                |          |

## http://127.0.0.1/RecordRoute



### Method

GET/POST

### Request

|   参数名    |  类型  |              说明              | 是否必须 |
|-------------|--------|--------------------------------|----------|
| location    | Point  |  The location from which the   | Y        |
|             |        | message is sent.               |          |
|   latitude  | int32  |                                | Y        |
|   longitude | int32  |                                | Y        |
| message     | string |  The message to be sent.       | N        |


### Reply

|   参数名    |    类型     |              说明              | 是否必须 |
|-------------|-------------|--------------------------------|----------|
| routes      | []RouteNote |                                |          |
|   location  | Point       |  The location from which the   |          |
|             |             | message is sent.               |          |
|   latitude  | int32       |                                |          |
|   longitude | int32       |                                |          |
|   message   | string      |  The message to be sent.       |          |

