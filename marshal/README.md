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
import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type Good struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (g Good) MarshalJSON() ([]byte, error) {
	return Marshal(g)
}

func TestMarshal(t *testing.T) {
	good := Good{123, "jock", time.Now(), time.Now()}
	bytes, _ := json.Marshal(good)
	// {"id":123,"name":"jock","created_at":"2024-03-14 17:55:03","updated_at":"2024-03-14 17:55:03"}
	fmt.Printf("%s\n", bytes)
}
```
