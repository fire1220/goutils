# goutils 包介绍

### parallel
> 启动多协成，统一处理并返回各个协成结果,
> 方法名和返回值都是slice，会额外返回一个error

### marshal
> 使用json.Marshal把结构体转换成json时，
> 会自动将time.Time类型解析成"2024-03-14 18:10:09"格式

### common 常用工具
- Round 四舍五入,保留precision位小数(默认保留2位小数)
- RoundStr 四舍五入,保留precision位小数(默认保留2位小数)
- GetAge 根据生日获取年龄
- IsCancel 判断上下文是否关闭
- ContextKeys 获取上下文所有key
- ContextDuplicate 复值上下文的key和val到新的上下文
- SliceColumn 取出slice里元素结构体的key成员，返回slice
- SliceColumnMap 取出slice里元素结构体的key成员，返回map
- NumberConvChinese 数字转汉字；例："100" 转成 "一百"

### ginvalidate
- SimpleValidate 是gin包的绑定验证器