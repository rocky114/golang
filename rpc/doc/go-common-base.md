## 1. 第三方组件客户端框架搭建

* 1.1 数据库组件（sqlite（单机） + tidb（集群））

TODO: 接口抽象（客户端创建/销毁方法 + 通用CRUD方法） + 配置抽象
    + 需要自适应增量升级（可以使用migration）

Tips: 目前Viper-Lite并未支持sqlite + tidb，寻找优质开源项目，边抄边改

* 1.2 缓存组件（redis（单机） + redis（集群））

TODO: 接口抽象（客户端创建/销毁方法 + 通用CRUD方法） + 配置抽象

Tips: 目前一人一档组件已支持redis（单机 + 集群），直接复用，再加改进（可选）

* 1.3 搜索引擎组件（elasticsearch（单机） + elasticsearch（集群））

TODO: 接口抽象（客户端创建/销毁方法 + 通用CRUD方法） + 配置抽象

Tips: 目前Viper-Lite并未支持elasticsearch（单机 + 集群），寻找优质开源项目，边抄边改

* 1.4 消息队列组件（mosquitto（单机） + kafka（集群））

TODO: 接口抽象（客户端创建/销毁方法 + 通用produce/consume方法） + 配置抽象

Tips: 目前Viper-Lite已支持mosquitto + kafka，直接复用，再加改进（可选）

* 1.5 文件存储组件（localstore + seaweedfs）

TODO: 接口抽象（客户端创建/销毁方法 + 通用CRUD方法） + 配置抽象
    + 需要适配（裸文件存储/查询/下载+osg存储/查询/下载）

Tips: 目前OSG已支持localstore + seaweedfs，直接从Viper-Lite IPS/VPS中找出OSG的最佳实践代码，再加改进（可选）

以上是否都要做SSL支持？

## 2. RPC Server/Client框架搭建

TODO: Server/Client搭建示例（创建/销毁方法 + 通用CRUD方法） + 配置抽象

Tips: 使用grpc还是go-zero zrpc（可以利用goctl生成大坨的skeleton code），需要对比下，哪种更利于加速开发用哪个？

## 3. 监控框架搭建（不用promethus，用DMS, 暂时先不做）

TODO: 监控模型搭建示例（创建/销毁方法 + Counter/Summary等） + 配置抽象

Tips: 直接抄 https://github.com/zeromicro/go-zero/tree/master/core/prometheus + 找出Viper-Lite IPS/VPS中最佳实践代码，再加改进（可选）

## 4. 日志框架搭建

TODO: 日志模型搭建示例（创建/销毁方法 + 标准输出 + 文件输出 + 格式定制化 + rotate等） + 配置抽象

    *  支持日志分割时间设置
    *  支持日志路径设置
    *  支持保存最长时间设置

Tips: 直接抄 https://github.com/zeromicro/go-zero/tree/master/core/logx，再加改进（可选）

## 5. panic/recover框架搭建

TODO: 奔溃/恢复模型搭建示例 + 配置抽象

Tips: 寻找优质开源项目，边抄边改 + 找出Viper-Lite IPS/VPS中最佳实践代码，再加改进（可选）

## 6. Trace框架搭建

TODO: Trace模型搭建示例（创建/销毁方法 + 通用CRUD方法） + 配置抽象

Tips: 直接抄 https://github.com/zeromicro/go-zero/tree/master/core/trace，再加改进（可选）

---

## 7. sql字段的增量升级
TODO: 可以使用migration（支持sqlite、mysql），但是不支持tidb，需要调研下

待补充...
