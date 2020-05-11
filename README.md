# code_playground_server
大创主服务器后端工程

## 使用前注意:
- 命令行输入 go mod tidy 下载项目所有依赖
- 修改config文件夹中yaml配置文件
- 创建codep mysql数据库 导入sql中文件 

## 目录:
- config:配置文件 使用yaml
- controller:业务控制层
- dao:数据库配置
- middleware:路由中间件
- model:模型及dao
- router:路由(请编写v1路由组 勿新建)
- session:会话模块
- sql:数据库脚本
- utils:自定义工具类 若需要使用不明第三方插件 请新建vendor文件 
- main.go:服务器入口函数

