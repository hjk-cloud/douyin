# simple-demo

## 抖音项目服务端简单示例

导入项目后，在/define/define.go中修改相应参数适配本地信息

```shell
go build && ./simple-demo
```

### 功能说明

抖音接口文档
https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18345145

### 待优化

token应该放在session中，我直接写到数据库的user表里了，后期改进

