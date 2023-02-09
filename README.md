# go-utils

一个简单的Go项目工具包

## 配置
> 基于 [github.com/spf13/viper v1.15.0](https://github.com/spf13/viper)
> 
> 用法
```
	config.Init(ENV)
	stringValue := config.Instance().GetString(KEY)
```
## 数据库
> 基于 [github.com/go-gorm/gorm v1.24.5](https://github.com/go-gorm/gorm)
> 
> 用法
```
	database.Init()
	baseDb := database.Instance("base")
```
## 缓存
> 基于 [github.com/redis/go-redis v9](https://github.com/redis/go-redis)
> 
```
	cache.Init()
	var ctx = context.Background()
	cache.Instance().Set(ctx, "test", "123456", 0).Err()
```
