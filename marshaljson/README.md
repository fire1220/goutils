# marshaljson


### Getting marshaljson
With Go's module support, go `[build|run|test]` automatically fetches the necessary dependencies when you add the import in your code:
```shell
import "github.com/fire1220/goutils/marshaljson"
```

Alternatively, use go get:

```shell
go get github.com/fire1220/goutils/marshaljson
```

# 使用参数说明
- `datetime`格式化`time.Time`格式为字符串，格式规则和`time.Format`通用
- `default`设置默认值,规则如下
  - 基础类型 : 默认值只能是对应的基础类型值
  - slice : 默认值只能是:`[]`空数组、`{}`空对象
  - map : 默认值只能是:`{}`空对象、任意字符串
  - struct : 默认值只能是:`{}`空对象
  - 指针类型 : 默认值只能是:`[]`空数组、`{}`空对象、任意字符串
- `defaultString` 设置默认值为字符串

### Running marshaljson

```go
package test

import (
	"encoding/json"
	"fmt"
	"github.com/fire1220/marshaljson"
	"testing"
	"time"
)

type GoodInfo struct {
	Title    string    `json:"title" default:"ABC"`
	Like     string    `json:"like"`
	PlayTime time.Time `json:"play_time" datetime:"2006-01-02 15:04:05"`
}

func (t GoodInfo) MarshalJSON() ([]byte, error) {
	return marshaljson.MarshalFormat(t)
}

type Good struct {
  ID      int32       `json:"id" default:"456"`
  Float1  float64     `json:"float1" default:"11.1"`
  Float2  float64     `json:"float2" default:"-11.1"`
  Float3  float64     `json:"float3" defaultString:"hello"`
  Uint    uint        `json:"uint" default:"111"`
  Bool1   bool        `json:"bool1" default:"false"`
  Bool2   bool        `json:"bool2" default:"true"`
  Bool3   bool        `json:"bool3" defaultString:"hello"`
  Slice1  []bool      `json:"slice1" default:"[]"`
  Slice2  []bool      `json:"slice2" default:"{}"`
  Slice3  []bool      `json:"slice3" defaultString:"hello"`
  Slice4  []bool      `json:"slice4" defaultString:""`
  Map1    map[int]int `json:"map1" default:"[]"`
  Map2    map[int]int `json:"map2" default:"hello"`
  Struct1 GoodInfo    `json:"struct1"`
  Struct2 GoodInfo    `json:"struct2" default:"{}"`
  Struct3 GoodInfo    `json:"struct3" default:"hello"`
  Ptr1    *GoodInfo   `json:"ptr1" default:"{}"`
  Ptr2    *GoodInfo   `json:"ptr2" default:"[]"`
  Ptr3    *GoodInfo   `json:"ptr3" default:"hello"`
  String1 string      `json:"string1" defaultString:"hello"`
  String2 string      `json:"string2" default:"hello"`
  Name    string      `json:"name" default:"123"`
  Time1   time.Time   `json:"time1" datetime:"2006-01-02 15:04:05"`
  Time2   time.Time   `json:"time2" datetime:"2006-01-02"`
  Time3   time.Time   `json:"time3" datetime:"15:04:05" default:"-"`
  Time4   time.Time   `json:"time4" datetime:"2006-01-02 15:04:05" default:"0000-00-00"`
  Time5   time.Time   `json:"time5" default:""`
  Time6   time.Time   `json:"time6" default:"-"`
}

func (t Good) MarshalJSON() ([]byte, error) {
  return marshaljson.MarshalFormat(t)
}

func TestMarshal(t *testing.T) {
  good := Good{
    ID:    0,
    Name:  "",
    Time1: time.Now(),
    Time2: time.Now(),
    Time3: time.Now(),
    Time4: time.Now(),
  }
  bytes, err := json.Marshal(good)
  fmt.Printf("%s\n", bytes)
  fmt.Println(err)
  goodList := make([]Good, 0)
  goodList = append(goodList, good)
  bytes, err = json.Marshal(goodList)
  fmt.Printf("%s\n", bytes)
  fmt.Println(err)
}
```
### output:
```json
[
  {
    "id": 456,
    "float1": 11.1,
    "float2": -11.1,
    "float3": "hello",
    "uint": 111,
    "bool1": false,
    "bool2": true,
    "bool3": "hello",
    "slice1": [],
    "slice2": {},
    "slice3": "hello",
    "slice4": "",
    "map1": "[]",
    "map2": "hello",
    "struct1": {
      "title": "ABC",
      "like": "",
      "play_time": "0000-00-00 00:00:00"
    },
    "struct2": {},
    "struct3": "hello",
    "ptr1": {},
    "ptr2": [],
    "ptr3": "hello",
    "string1": "hello",
    "string2": "hello",
    "name": "123",
    "time1": "2025-01-22 18:23:01",
    "time2": "2025-01-22",
    "time3": "18:23:01",
    "time4": "2025-01-22 18:23:01",
    "time5": "",
    "time6": "-"
  }
]
```

