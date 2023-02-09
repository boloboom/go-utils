# go-utils

一个简单的Go项目工具包

## 配置
> 用法
```
	config.Init(ENV)
	stringValue := config.Instance().GetString(KEY)
```
## 数据库
> 用法
```
	database.Init()
	baseDb := database.Instance("base")
```
## 缓存
```
	cache.Init()
	var ctx = context.Background()
	cache.Instance().Set(ctx, "test", "123456", 0).Err()
```
