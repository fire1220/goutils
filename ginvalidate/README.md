# ginvalidate

---

### describe
> gin包的绑定验证器


## Usage

```shell
go get github.com/fire1220/goutils/ginvalidate
```
或

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

type BaseController struct {
}

func (b *BaseController) Validate(ctx *gin.Context, param any) bool {
	if ok, errs := ginvalidate.SimpleValidate(ctx, param); !ok {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"code": 405, "msg": errs[0].Error()})
		return false
	}
	return true
}
```
