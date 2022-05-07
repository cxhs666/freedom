## 概述
    本人是个phper，初学go，搭建了一个基础应用框架，直接可以作为项目服务提供者使用，有什么想法可以一起交流学习(564210220@qq.com)。

### 使用的框架
    go web框架 gin，主要是处理路由
    gorm 数据库包
    air 热更新
### 实现了些什么
    路由到控制器的设置，都已经ok
    数据库的使用，只支持了mysql
    认证的中间件已经实现可以直接用，用的jwt
    你可以直接运行，给你的应用提供服务
### 怎么使用呢
    1.安装air的go包
    2.cp .env.example .env 然后配置自己的数据库信息
    3.到项目根目录下运行 air 命令即可
    
