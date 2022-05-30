# 抖音项目使用

导入项目后，在/define/define.go中修改相应参数适配本地信息

# 功能说明

抖音接口文档
https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18345145

# 已完成

## user模块
用户登录、注册、个人信息
## video模块
视频流、投稿、作品列表
## relation模块
关注、取关操作、关注列表、粉丝列表
## favorite模块
点赞、取消操作、喜欢列表

# 待优化

## user模块
个人信息中获赞数、作品数、喜欢数（目前前端未提供该属性接口）

## video模块
video封面url未获取

前端请求参数中latest_time属性未使用

next_time属性未使用
