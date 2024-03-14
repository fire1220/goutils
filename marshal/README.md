# marshal

---

### describe
> 使用json.Marshal把结构体转换成json时，
> 会自动将time.Time类型解析成"2024-03-14 18:10:09"格式

## Usage

```
import "github.com/fire1220/goutils/marshal"
```

### example:
``` go
package marshal

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type Good struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	PlayTime  time.Time `json:"play_time"`
	CreatedAt time.Time `json:"created_at" datetime:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `json:"updated_at" datetime:"15:04:05"`
}

func (g Good) MarshalJSON() ([]byte, error) {
	return Marshal(g)
}

func TestMarshal(t *testing.T) {
	good := Good{ID: 123, Name: "jock", PlayTime: time.Now(), CreatedAt: time.Now()}
	bytes, _ := json.Marshal(good)
	// {"id":123,"name":"jock","play_time":"2024-03-14 19:40:19","created_at":"2024-03-14 19:40:19","updated_at":"00:00:00"}
	fmt.Printf("%s\n", bytes)
}

```
