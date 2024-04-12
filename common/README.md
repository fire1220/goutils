# common 常用工具

---

### describe
- Round 四舍五入,保留precision位小数(默认保留2位小数)
- RoundStr 四舍五入,保留precision位小数(默认保留2位小数)
- GetAge 根据生日获取年龄
- IsCancel 判断上下文是否关闭
- ContextKeys 获取上下文所有key
- ContextDuplicate 复值上下文的key和val到新的上下文
- SliceColumn 取出slice里元素结构体的key成员，返回slice
- SliceColumnMap 取出slice里元素结构体的key成员，返回map
- NumberConvChinese 数字转汉字；例："100" 转成 "一百"

## Usage

```
import "github.com/fire1220/goutils/common"
```
