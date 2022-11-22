# Clover

<div align=center>
<svg t="1669107962890" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="712" width="128" height="128">
<path d="M518.46 466.146c-0.25 0.968-25.462 96.968-100.158 208.154-68.574 101.968-194.222 238.208-403.596 307.968v0.032C6.24 985.11 0.118 993.078 0.118 1002.512c0 11.778 9.56 21.336 21.336 21.336a20.76 20.76 0 0 0 6.716-1.124v0.032c97.814-32.614 186.598-81.538 263.982-145.518 61.856-51.14 116.528-111.904 162.512-180.6 78.32-117.058 104.062-215.778 105.124-219.902l-41.328-10.59z" fill="#79AA41" p-id="713" data-spm-anchor-id="a313x.7781069.0.i4" class=""></path><path d="M110.616 325.19C-19.5 451.058 167.754 567.802 167.754 567.802s-116.058 187.692 56.358 242.8c172.416 55.108 339.398-335.334 339.398-335.334S240.732 199.352 110.616 325.19z" fill="#8AC054" p-id="714" data-spm-anchor-id="a313x.7781069.0.i3" class=""></path><path d="M977.882 634.344c130.116-125.834-57.138-242.612-57.138-242.612s116.058-187.662-56.36-242.77c-172.414-55.108-339.426 335.334-339.426 335.334s322.81 275.916 452.924 150.048z" fill="#8AC054" p-id="715" data-spm-anchor-id="a313x.7781069.0.i1" class=""></path><path d="M389.656 913.416c125.868 130.116 242.614-57.14 242.614-57.14s187.694 116.058 242.8-56.356c55.108-172.418-335.334-320.746-335.334-320.746s-275.918 304.124-150.08 434.242z" fill="#9ED36A" p-id="716" data-spm-anchor-id="a313x.7781069.0.i0" class=""></path><path d="M698.81 46.15c-125.836-130.116-242.614 57.138-242.614 57.138s-187.66-116.058-242.768 56.358c-55.14 172.416 326.308 319.528 326.308 319.528S824.68 176.266 698.81 46.15z" fill="#9ED36A" p-id="717" data-spm-anchor-id="a313x.7781069.0.i2" class=""></path>
</svg>

[![GitHub](https://img.shields.io/github/license/zhanjunjie2019/clover)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![GitHub go.mod Go version (branch)](https://img.shields.io/github/go-mod/go-version/zhanjunjie2019/clover/master)](https://github.com/golang/go.git)
[![CodeQL](https://github.com/zhanjunjie2019/clover/actions/workflows/codeql.yml/badge.svg)](https://github.com/zhanjunjie2019/clover/actions/workflows/codeql.yml)
[![Codacy Security Scan](https://github.com/zhanjunjie2019/clover/actions/workflows/codacy.yml/badge.svg)](https://github.com/zhanjunjie2019/clover/actions/workflows/codacy.yml)
</div>

基于领域驱动设计思想，支持微服务的快速开发框架和系统架构。
代码分层参考[alibaba/COLA](https://github.com/alibaba/COLA)的设计。

## 架构图

### 服务分层依赖架构

![服务分层依赖架构](https://assets.processon.com/chart_image/63745a045653bb3a8405069a.png)

### 领域分层领域图

### 代码分层依赖架构

![分层依赖架构](https://assets.processon.com/chart_image/637334050791290b4b9a005e.png)

## 配置优先级

合并成最终配置，0值配置无效

环境变量（可选） > 配置中心服务<u>版本级</u>配置（可选） > 配置中心<u>服务级</u>配置（可选）> 本地配置文件

## 服务环境变量定义

无默认值，若空值，最终读取配置文件的配置内容

| 变量名                     | 定义                     | 示例                             |
|-------------------------|------------------------|--------------------------------|
| SVC_MODE                | 服务模式,1正式2测试            | 2                              |
| SVC_NAME                | 服务部署名                  | clover-example                 |
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
| clover-example  | 8700                      | [http://127.0.0.1:8700/swagger/index.html](http://127.0.0.1:8700/swagger/index.html) |
| clover-auth     | 8900                      | [http://127.0.0.1:8900/swagger/index.html](http://127.0.0.1:8900/swagger/index.html) |

## 开发时本地运行设置

### hosts文件，添加host配置

```text
127.0.0.1       nsqd
```
