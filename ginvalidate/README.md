# ginvalidate

---

### describe
> gin包的绑定验证器


## Usage

```
import "github.com/fire1220/goutils/ginvalidate"
```

### example:
``` go
package controller

import (
	"github.com/fire1220/goutils/ginvalidate"
	"github.com/gin-gonic/gin"
	"net/http"
)

type baseController struct {
}

func (b *baseController) Validate(ctx *gin.Context, param any) bool {
	if ok, errs := ginvalidate.SimpleValidate(ctx, &param); !ok {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"code": 405, "msg": errs[0].Error()})
		return false
	}
	return true
}

```
