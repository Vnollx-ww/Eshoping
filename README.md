# 一、项目介绍
 基于"Hertz HTTP框架+Kitex微服务框架"完成的购物系统
 # 二、项目实现
采用微服务架构，分为用户服务，商品服务，订单服务,架构分为HTTP,RPC和DAL三层

其中 HTTP 层使用 Hertz 框架接收客户端发来的 HTTP 请求；
RPC 层使用 Kitex 框架，并用 ETCD 做服务注册与发现；
DAL 层为数据访问层，包含 Redis 和 MySQL 两部份。
当客户端发来 HTTP 请求时，HTTP 层会调用 RPC层 的 RPC Client，然后 RPC Client 去 ETCD 中心寻找已经注册的对应的微服务，交给 RPC Server 处理。RPC Server 会去调用 DAL 层的数据库，数据库处理完毕后把结果返回给 RPC Server，RPC Server 将其返回给 RPC Client，最后 RPC Client 返回给 Hertz 并由 Hertz 返回 HTTP 响应结果。

技术栈

Hertz：提供 HTTP 服务；

Kitex：提供 RPC 服务；

ETCD：服务注册与发现；

JWT：token 的生成与校验；

Gorm：对 MySQL 进行 ORM 操作，Go-Redis：操作 Redis 对频繁访问的数据进行缓存；

Redis：对频繁访问的数据进行缓存，按照一定策略使键过期，并定时同步数据到数据库，并使用分布式锁；

Viper：读取配置文件；

Kafka：消息队列，用于发送异步请求和流量控制；

Gocron：定时任务，用于定时同步Redis与MYSQL之间的数据；

Nginx：反向代理与负载均衡；

Zap：日志打印；

# 三、项目部署
采用docker部署微服务
