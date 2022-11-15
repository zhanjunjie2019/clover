# Clover

[![License](https://img.shields.io/badge/license-Apache%202-4EB1BA.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)

基于领域驱动设计思想设计，支持微服务的快速开发框架和系统架构。
代码分层设计参考[alibaba/COLA](https://github.com/alibaba/COLA)的分层设计理念。

## 代码分层架构

![分层依赖架构](http://assets.processon.com/chart_image/637334050791290b4b9a005e.png)

## 配置优先级

合并成最终配置，0值配置无效

环境变量（可选） > 配置中心服务<u>版本级</u>配置（可选） > 配置中心<u>服务级</u>配置（可选）> 本地配置文件

## 服务环境变量定义

无默认值，若空值，最终读取配置文件的配置内容

| 变量名                     | 定义                     | 示例                             |
|-------------------------|------------------------|--------------------------------|
| SVC_MODE                | 服务模式,1正式2测试            | 2                              |
| SVC_NAME                | 服务部署名                  | clover-tenant                  |
| SVC_NUM                 | 服务实例序号                 | 1                              |
| SVC_PORT                | 服务端口号                  | 8900                           |
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

## 接口文档生成

```shell
swag init --pd
```

## 端口分配

| 服务名             | 端口                        | 可视化界面路径                                                                              |
|-----------------|---------------------------|--------------------------------------------------------------------------------------|
| consul          | 8300/8301/8500/8600       | [http://127.0.0.1:8500/](http://127.0.0.1:8500/)                                     |
| elasticsearch   | 9200                      | 无                                                                                    |
| kibana          | 5601                      | [http://127.0.0.1:5601/](http://127.0.0.1:5601/)                                     |
| jaeger          | 14250/16686               | [http://127.0.0.1:16686/](http://127.0.0.1:16686/)                                   |
| otel-collector  | 1888/4317/8888/8889/13133 | 无                                                                                    |
| postgres-master | 9432                      | 无                                                                                    |
| postgres-slave  | 9433                      | 无                                                                                    |
| prometheus      | 9090                      | [http://127.0.0.1:9090/](http://127.0.0.1:9090/)                                     |
| redis           | 6379                      | 无                                                                                    |
| nsqlookupd      | 4160/4161                 | 无                                                                                    |
| nsqd            | 4150/4151                 | 无                                                                                    |
| nsqadmin        | 4171                      | [http://127.0.0.1:4171/](http://127.0.0.1:4171/)                                     |
| clover-auth     | 8900                      | [http://127.0.0.1:8900/swagger/index.html](http://127.0.0.1:8900/swagger/index.html) |

## 开发时本地运行设置

### hosts文件，添加host配置

```text
127.0.0.1       nsqd
```
