# starter-example

示例服务

## 目录结构

- bc --限界上下文(允许一个服务有多个限界上下文，但不建议。根据限界领域职责命名)
    - adapter --适配器（表现层）
        - consumer --消息队列监听
        - scheduler --定时任务调度
        - controller --http请求监听
        - grpc --grpc请求监听
    - app --应用层（主要实现）
    - domain --领域层
        - biserrs --业务异常定义
        - gateway --下游适配器网关（防腐层接口）
        - model --领域模型
    - infr --基础设施层
        - bcconsts --常量集合
        - gatewayimpl --下游网关实现（防腐层实现）
            - convs --跨域对象转换器
        - repo --仓储层
            - po --仓储层对象
- configs --本地配置
- docs --swagger文件
- main.go --启动程序
- zz_generated.ioc.go --ioc文件
