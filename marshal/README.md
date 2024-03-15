# marshal

---

### describe
> 使用json.Marshal把结构体转换成json时  
> 会自动将time.Time类型解析成"2024-03-14 18:10:09"格式  
> tag中datetime表示自定义结构，后面添加,omitempty表示如果空值，格式化为空字符串   


## Usage

```
import "github.com/fire1220/goutils/marshal"
```

### example:
``` go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/fire1220/goutils/marshal"
	"time"
)

type Good struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	PlayTime    time.Time `json:"play_time" datetime:"omitempty"`
	ExecuteTime time.Time `json:"execute_time" datetime:"2006-01-02"`
	CreatedAt   time.Time `json:"created_at" datetime:"omitempty"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (t Good) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(t)
}

func main() {
	good := Good{ID: 123, Name: "jock", PlayTime: time.Now(), ExecuteTime: time.Now()}
	bytes, _ := json.Marshal(good)
	// {"id":123,"name":"jock","play_time":"2024-03-15 18:23:43","execute_time":"2024-03-15","created_at":"","updated_at":"0000-00-00 00:00:00"}
	fmt.Printf("%s\n", bytes)
}
```
