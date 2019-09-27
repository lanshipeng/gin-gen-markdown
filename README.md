## goals
- 基于gin的http接口,自动化构建markdown接口文档

## features

- 基于ast实现，通过ast解析源代码结构。将参数名、类型、注释等字段提取生成到md文档。

- 构建文档前需要添加一些必要注释@request、 @response和对外暴露提供的接口名。具体可参考api.go。

- 使用时可以将代码嵌入到工程项目。当作脚本运行，可以遍历整个目录树。

- 通过mark_down命令行参数可以指定目录树路径，prefix命令行参数指定需要遍历的文件前缀或后缀。


## quick start
```bash
go run main.go doc --mark_down=. --prefix=api.go
```
