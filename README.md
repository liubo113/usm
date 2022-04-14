# USM

## 模型图

![ddd_image](./doc/ddd.png)

### 名词定义

- DTO (Data Transfer Object), 数据传输对象
- BO (Business Object), 业务对象
- PO (Persistent Object)，持久化对象，对应 Entity
- Service API 定义的实现
- Usecase Service 的核心业务逻辑
- Repository 单个资源数据持久化，定义在 biz 层，实现在 data 层

采用依赖倒置原则，service，data 依赖 biz，biz 层定义 BO 和 repo 由 data 层实现

## 相关文档

- [kratos](https://go-kratos.dev)
- [entgo](https://entgo.io)
- [domain driver design (ddd)](https://domain-driven-design.org/)

## 使用说明

```shell
make help
```

## 目录结构

```text
├── api # api 目录，通过 IDL 文件可以生成若干 stub 代码和 openapi 文档
│   └── account # 按照服务/模块作为目录进行划分
│       └── v1 # 服务/模块必须包含版本号
├── cmd # 可执行程序目录
├── configs # 配置文件目录
│   └── config.yaml
├── internal # 业务核心代码，使用 internal 关键字防止被错误引用
│   ├── biz # biz 层
│   │   ├── biz.go # biz 公共定义
│   │   ├── repo # 包含各 model 及接口定义，由 data 层实现，repo 只关心定义，不关心由何种方式实现（db、cache、file、etc...）
│   │   ├── usecase # 调用多个 repo 进行组合，usecase 为业务的核心实现，复杂的事务逻辑应在此实现，且此处应有完备的单元测试
│   │   │   └── account # usecase 由 服务/模块划分
│   │   └── wire.go
│   ├── conf
│   ├── data # data 目录实现 biz 层的 repo 的实现，这里只针对单个资源做粗力度的实现，多个资源的组合应放在 biz/usecase 中
│   │   ├── data.go # data 公共定义
│   │   ├── ent # 实体层，使用 entgo 作为 orm
│   │   │   ├── schema # 在此处定义数据库的 schema，并通过 go generate 生成 stub 代码
│   │   │   │   └── user.go
│   │   ├── user.go # 按照资源作为文件划分
│   │   └── wire.go
│   ├── server
│   └── service # api 层的实现，仅可在此处打印日志，同时服务的外部调用、iam、接口鉴权认证等逻辑都应在此处实现，这里不实现复杂的业务逻辑
│       ├── account # 按照服务/模块作为目录进行划分
├── openapi.yaml
├── pkg
└── third_party
```

## 单元测试

单元测试是非常重要的，对于业务逻辑必须编写单元测试，因为：

1. 单元测试是最好的使用文档
2. 好的代码是易于测试的
3. 单元测试可以极大提高服务可靠性，防止因为一些修修改改但却没有测试导致一些低级问题出现

### 最佳实践

利用 go 官方提供的 subtest 和 gomock 完成单元测试，对于每层代码：

- /api：适合进行集成测试，待设计
- /data：利用 sqlite3 内存模式模拟真实的数据库操作进行测试，可覆盖绝大数情况的数据库 case
- /biz：依赖 repo，利用 gomock 实现 repo 接口来进行单元测试
- /service：service 层为很薄的一层，一般会调用 biz usecase、iam、pubsub、rpc/http 外部调用、认证等逻辑，需在各逻辑的实现处进行单元测试，service 无需测试
