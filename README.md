# 抖音项目使用

导入项目后，在/define/define.go中修改相应参数适配本地信息

# 功能说明

抖音接口文档
https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18345145

# 已完成

## user模块
用户登录、注册、个人信息
## video模块
视频流、投稿、个人发布列表
## relation模块
关注、取关操作、关注列表、粉丝列表

# 待优化

## user模块
个人信息中获赞数显示（目前前端未提供该属性接口）

token应该放在session中，我直接写到数据库的user表里了，肯定不能这么做
## video模块
video封面url未获取
## relation模块
关注列表中，应显示取消关注按钮，和user的is_follow属性有关

