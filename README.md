<div align=center>

![Icon](https://gitee.com/zhanjunjie/clover/raw/master/docs/imgs/icon.svg)

# Clover

[![GitHub](https://img.shields.io/github/license/zhanjunjie2019/clover)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![GitHub go.mod Go version (branch)](https://img.shields.io/github/go-mod/go-version/zhanjunjie2019/clover/master)](https://github.com/golang/go.git)
[![CodeQL](https://github.com/zhanjunjie2019/clover/actions/workflows/codeql.yml/badge.svg)](https://github.com/zhanjunjie2019/clover/actions/workflows/codeql.yml)
[![Codacy Security Scan](https://github.com/zhanjunjie2019/clover/actions/workflows/codacy.yml/badge.svg)](https://github.com/zhanjunjie2019/clover/actions/workflows/codacy.yml)
</div>

基于领域驱动设计思想，支持微服务的快速开发框架和系统架构。
代码分层参考[alibaba/COLA](https://github.com/alibaba/COLA)的设计。

本项目的期望是通过提供给开源项目的开发者，统一的代码框架和开发规范，贡献更多的可任意插拔的通用领域服务和基础设施服务。
最终让业务开发者在这套微服务架构的帮助下，在做核心领域业务开发过程中，可以更快的集成实现各类业务和扩展。

项目先从Saas化的直播电商平台开始做起。

## 架构图

### 服务分层架构

![服务分层依赖架构](https://gitee.com/zhanjunjie/clover/raw/master/docs/imgs/services.svg?r=v3)

### 领域分层架构图

### 代码分层架构

![分层依赖架构](https://gitee.com/zhanjunjie/clover/raw/master/docs/imgs/depend.svg?r=v1)

## 配置优先级

合并成最终配置，0值配置无效

环境变量（可选，不支持热更） > 配置中心节点，优先级越往后越高（必要，支持热更） > 本地配置文件（必要，不支持热更）

## 服务环境变量定义

无默认值，若空值，最终读取配置文件的配置内容

| 变量名                     | 定义                     | 示例                             |
|-------------------------|------------------------|--------------------------------|
| SVC_MODE                | 服务模式,1正式2测试            | 2                              |
| SVC_NAME                | 服务部署名                  | clover-example                 |
| SVC_NUM                 | 服务实例序号                 | 1                              |
| SVC_PORT                | 服务端口号                  | 8800                           |
| SVC_VERSION             | 服务版本号                  | v1.0.0                         |
| CONSUL_ADDR             | consul服务配置中心路径         | consul:8500                    |
| CONSUL_REGISTER_TTL     | consul服务注册时限           | 2                              |
| LOG_LEVEL               | 日志级别                   | debug                          |
| LOG_DIRECTOR            | 日志输出文件夹                | ../deploy/docker/filebeat/logs |
| LOG_MAX_AGE             | 日志最大天数                 | 7                              |
| LOG_IN_CONSOLE          | 是否输出到控制台，1是2否          | 1                              |
| JWT_SIGNING_KEY         | jwt签名密钥                | clover                         |
| JWT_EXPIRES_TIME        | jwt续约时限，单位秒            | 14400                          |
| REDIS_ENABLED           | 是否启用redis，1是2否         | 1                              |
| REDIS_ADDR              | redis域名地址              | redis:6379                     |
| REDIS_POSSWORD          | redis密码                | clover                         |
| REDIS_DB                | redis子库                | 1                              |
| OTEL_ENABLED            | 是否启用遥测，1是2否            | 1                              |
| OTEL_COLLECTOR_ENDPOINT | 遥测collector的grpc切入点    | otel-collector:4317            |
| NSQ_ENABLED             | 是否开启NSQ消息队列，1是2否       | 1                              |
| NSQ_NSQD_ADDR           | 用于生产者使用的nsqd服务地址       | nsqd:4150                      |
| NSQ_LOOKUPD_ADDR        | 用于消费者使用的nsqlookupd服务地址 | nsqlookupd:4161                |

## 端口分配

| 服务名             | 端口                        | 说明          | 可视化界面路径                                                                                   |
|-----------------|---------------------------|-------------|-------------------------------------------------------------------------------------------|
| consul          | 8300/8301/8500/8600       | 注册中心/配置中心   | [http://127.0.0.1:8500/](http://127.0.0.1:8500/)                                          |
| elasticsearch   | 9200                      | 搜索引擎        | 无                                                                                         |
| kibana          | 5601                      | 数据视图中心      | [http://127.0.0.1:5601/](http://127.0.0.1:5601/)                                          |
| jaeger          | 14250/16686               | 链路追踪器       | [http://127.0.0.1:16686/](http://127.0.0.1:16686/)                                        |
| otel-collector  | 1888/4317/8888/8889/13133 | 遥测系统数据收集器   | 无                                                                                         |
| postgres-master | 9432                      | 主关系数据库      | 无                                                                                         |
| postgres-slave  | 9433                      | 从关系数据库      | 无                                                                                         |
| prometheus      | 9090                      | 指标分析服务      | [http://127.0.0.1:9090/](http://127.0.0.1:9090/)                                          |
| redis           | 6379                      | 分布式缓存       | 无                                                                                         |
| nsqlookupd      | 4160/4161                 | 消息路由分配服务    | 无                                                                                         |
| nsqd            | 4150/4151                 | 消息队列引擎      | 无                                                                                         |
| nsqadmin        | 4171                      | 消息队列管理中心    | [http://127.0.0.1:4171/](http://127.0.0.1:4171/)                                          |
| livego          | 1935/8090/7001/7002       | 直播视频推拉流中间服务 | 无                                                                                         |
| clover-example  | 8800（http）/8801（gRPC）     | 示例项目        | 接口文档：[http://127.0.0.1:8800/swagger/index.html](http://127.0.0.1:8800/swagger/index.html) |
| clover-auth     | 8810（http）/8811（gRPC）     | 权限中心        | 接口文档：[http://127.0.0.1:8810/swagger/index.html](http://127.0.0.1:8810/swagger/index.html) |

### redis数据库子库序号分配

| 序号  | 场景                          |
|-----|-----------------------------|
| 0   | livego：直播通道凭证               |
| 1   | starter-example：示例服务        |
| 2   | starter-auth：权限中心分布式锁+用户授权码 |

## 开发时本地运行设置

### hosts文件，添加host配置

```text
127.0.0.1       nsqd
```
