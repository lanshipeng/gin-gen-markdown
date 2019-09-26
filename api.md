## http://127.0.0.1/GetFeature



### Method

GET/POST

### Request

| 参数名 | 类型  |                说明                 | 是否必须 |
|--------|-------|-------------------------------------|----------|
| lo     | Point |  One corner of the rectangle.       | Y        |
| hi     | Point |  The other corner of the rectangle. | Y        |


### Point
|  参数名   | 类型  | 说明 | 是否必须 |
|-----------|-------|------|----------|
| latitude  | int32 |      | Y        |
| longitude | int32 |      | Y        |



### Reply

|  参数名  |  类型  |                   说明                    | 是否必须 |
|----------|--------|-------------------------------------------|----------|
| name     | string |  The name of the feature.                 |          |
| location | Point  |  The point where the feature is detected. |          |

## http://127.0.0.1/RecordRoute



### Method

GET/POST

### Request

|  参数名  |  类型  |                     说明                      | 是否必须 |
|----------|--------|-----------------------------------------------|----------|
| location | Point  |  The location from which the message is sent. | Y        |
| message  | string |  The message to be sent.                      | N        |


### routeSummary
|    参数名     | 类型  |                               说明                               | 是否必须 |
|---------------|-------|------------------------------------------------------------------|----------|
| point_count   | int32 |  The number of points received.                                  | N        |
| feature_count | int32 |  The number of known features passed while traversing the route. | N        |
| distance      | int32 |  The distance covered in metres.                                 | N        |
| elapsed_time  | int32 |  The duration of the traversal in seconds.                       | N        |



### Reply

| 参数名 |      类型      | 说明 | 是否必须 |
|--------|----------------|------|----------|
| routes | []routeSummary |      |          |

