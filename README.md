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

Kafka：消息队列，用于发送异步请求，减少服务调用之间的耦合，在高并发的时刻对数据库进行压力缓冲；

Minio：对象存储，用于存储图片，视频，日志等非格式化数据；

WebSocket：实现服务器与客户端的双向实时通信，完成好友在线聊天功能；

Nginx：反向代理与负载均衡；

Sentinel：进行熔断限流，控制流量和保证系统的稳定性；

Jaeger： 进行链路追踪与事务监控，提高系统的可观测性；

Zap：日志打印；


# 三、项目部署
采用docker将项目部署至腾讯云云服务器。

在本地docker构建镜像上传至DockerHub ，在服务器docker中拉取镜像，并运行docker-compose文件，开放服务器安全组端口，访问服务器公网ip+端口8889就可以成功访问。

# 四、项目不足(目前所发现的问题)
1.项目部署在同一台机器上，后续期望采用分布式部署并保证分布式数据一致性。

2.采用k8s编排容器

3.考虑消息队列和数据库使用导致数据不一致的情况，当先把请求发送到消息队列时，由于是异步请求，如果未能在下次同步请求前处理完成，这时如果两个请求都会调用同一个属性，就会导致属性数据未能按正确顺序修改，会有潜在风险。

4.数据库性能与事务安全性还需优化，将数据库的读写分离，使用主数据库负责写操作，多个从数据库负责读操作。

5.项目代码冗余，不够简洁,优雅。

6.还需加强系统网络安全和流量控制。

7.加强系统的可观测性和自动化管理，云原生方向。

8:框架只是简单应用，还需加强对框架的深入学习理解和应用。