# FurballCommunity_backend
“Furball Community” is a community-based pet sharing and boarding platform app, which is a course design project.



### 项目架构

- config： 项目配置模块，将集成 mysql、token、redis 等配置部分。
- controller： 负责请求转发，接受页面过来的参数，传给 Model 处理，接到返回值，再传给页面。
- models： 对应数据表的增删查改。
- routers：处理路由。
- utils：定义项目工具组件，包括错误代码，返回类型等。

