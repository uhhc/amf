# Another Micro-Framework

## 框架说明

1. 底层框架基于 Go-kit 构建: 
    - https://gokit.io
    - https://github.com/go-kit/kit

2. 整体分为三层：
    - Services
        - 业务逻辑层。该层提供用户需要的业务逻辑
    - Endpoints
        - 端点层。将业务逻辑转化为一个个协议无关的端点，从而可以提供给各类协议使用
    - Transports
        - 传输层。将端点转化为各类具体协议能使用的 handler。每个端点可以有一个或多个 transport

3. 数据库连接层采用 GORM：
    - https://gorm.io/
    - https://github.com/jinzhu/gorm


## 目录结构

目录结构如下：

```
|-assets：放置静态文件。目前主要存放着 swagger-ui 相关的文件
|-bin：放置编译生成的二进制文件，文件在 git 中会忽略，不会提交到 git 仓库
|-cmd：放置需要运行的命令入口文件
|-db：放置数据库 migration 文件
|-devbox：放置本地开发调试所需的一些文件
|-pkg：框架核心目录
|  |-application：整体业务目录
|      |-example：`example`业务相关目录，其他业务采用对应的名称即可
|          |-endpoint：将业务逻辑转化为 gokit endpoint
|          |-repository：业务相关的存储逻辑
|          |-service：业务逻辑核心目录
|          |-transport：将 endpoint 转化为各类协议的 handler
|  |-config：业务逻辑相关配置文件目录
|  |-middleware：中间件目录
|  |-response：程序共用的输出定义
|  |-util：程序共用的工具目录
|-.config.yaml：默认配置文件
|-.gitignore：git 仓库需要忽略的文件配置
|-Dockerfile：Docker 镜像构建文件
|-Makefile：放置常用的 make 命令
|-README.md：项目的说明文件
|-main.go：程序入口文件
```

参考：

- https://github.com/go-kit/kit/tree/master/examples/shipping
- https://github.com/shijuvar/gokit-examples
- https://medium.com/@shijuvar/go-microservices-with-go-kit-introduction-43a757398183


## 本地运行说明

1. 启动 MySQL。在程序根目录下运行：

    ```bash
    make mysql-reload
    ```

2. 启动 MongoDB。在程序根目录下运行：

    ```bash
    make mongo-reload
    ```

3. 安装 Go 1.11 或以上版本，在程序根目录下直接运行：

    ```bash
    make run
    ```

4. 也可以使用 `docker-compose` 来运行，在程序根目录下运行：

    ```bash
    make amf-reload
    ```

    运行完以后，就可以通过地址 http://127.0.0.1 来访问了。


## GRPC

### 安装软件（protoc 和 protoc-gen-go）

```
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
```

### 编译

在 `sdk-common-go` 项目根目录下运行：

```
make grpc
```

也可以直接运行：

```
protoc -I <proto_import_path> --go_out=<plugins={plugin1+plugin2+...}>:<pb_output_path> <proto_file_path>
```

### Optional：使用 gogofaster 来编译

说明：可以去掉官方编译生成的 `XXX_unrecognized` 等信息。但目前生成的 PB 文件有些问题，因此暂未启用。

参考：

- https://github.com/gogo/protobuf

1. 安装步骤：

    ```
    go get github.com/gogo/protobuf/proto
    go get github.com/gogo/protobuf/protoc-gen-gogofaster
    go get github.com/gogo/protobuf/gogoproto
    ```

2. 编译命令：

    ```
    protoc -I <proto_import_path> --gogofaster_out=<plugins={plugin1+plugin2+...}>:<pb_output_path> <proto_file_path>
    ```

### 编写 GRPC 服务步骤

1. 在 `sdk-common-go` 的 grpc/proto 目录中编写 protobuf 文件
2. 在 `sdk-common-go` 根目录下运行 `make grpc` 生成 pb 文件
3. 在 `sdk-common-go` 的 grpc/client 目录中编写对应业务的 client 文件
4. 在各业务的 transport 目录下新增 grpc 目录，参考样例编写 grpc handler
5. 在各业务的 endpoint 目录下新增 grpc 目录，参考样例编写 grpc endpoint

### GRPC Client

#### gRPCurl

1. 官网：https://github.com/fullstorydev/grpcurl

2. 安装（Mac环境）

    ```
    brew install grpcurl
    ```

3. 使用

    ```
    # 列出支持的 GRPC 服务
    grpcurl -plaintext localhost:8080 list
    # 列出某个服务支持的方法
    grpcurl -plaintext localhost:8080 list pb.ExampleService
    # 查看某个服务的详细信息
    grpcurl -plaintext localhost:8080 describe pb.ExampleService
    # 调用某个服务的方法
    grpcurl -plaintext -d '{"id": "c9320ad5-6a97-4b8a-a31f-6bbd2f57b53b"}' localhost:8080 pb.ExampleService/Get
    ```
    
    说明：如果 GRPC 不支持 SSL，则请求时需要加上 `-plaintext` 参数；如果 GRPC 支持 SSL，但请求时不想带证书，则需要加上 `-insecure` 参数。

#### Client 编写

样例：

```
grpcClient := client.NewDatasetGRPCClient(
    viper.GetString("GRPC_HOST"),
    viper.GetString("GRPC_PORT"),
    logger,
)
res, err := grpcClient.Get(datasetID)
```

### 参考文档

- https://www.ru-rocker.com/2017/02/24/micro-services-using-go-kit-grpc-endpoint
- https://github.com/qneyrat/grpc-go-kit-example
- https://github.com/ru-rocker/gokit-playground/tree/master/lorem-grpc


## Migration

### 说明

采用 [Goose](https://github.com/steinbacher/goose) 来进行数据库的 migration。
Goose 支持 SQL 文件和 Go 文件作为 migration 的脚本。

### 配置

1. 默认的 migration 目录结构如下：

    ```
    |-db：migration所在的根目录
    |   |-migrations：放置 SQL 或 Go 类型的 migration 文件
    |   dbconf.yaml：goose 运行时默认使用的配置文件
    ```

2. `dbconf.yaml` 配置说明

    在 `dbconf.yaml` 文件中可以配置各个环境所使用的数据库配置。通过在运行 goose 命令时加上 `-env` 来指定。
    
    如果不加该参数，则默认的环境为 `development`。
    
    可以在配置文件中使用环境变量值。格式为：`$MY_ENV_VAR` 或 `${MY_ENV_VAR}`。

### 常用命令

1. 创建一个 migration 脚本：

    1）创建 SQL 文件脚本

    ```
    goose create <脚本名称标识>
    ```

    2）创建 Go 文件脚本
    
    ```
    goose create -type go <脚本名称标识>
    ```

    **注：** 脚本名称标识必须采用驼峰的写法，比如`InitDb`。

2. 执行 migration 脚本：

    ```
    goose up
    ```

3. 回滚上次执行的 migration：

    ```
    goose down
    ```


## API 文档

### 说明

采用 swagger 作为 API 文档，使用 [swag](https://github.com/swaggo/swag) 来生成 swagger.json。

### 使用

1. 安装 swag

    ```
    go get -u github.com/swaggo/swag/cmd/swag
    ```

2. 生成或更新 swagger.json 文件

    ```
    make swagger
    ```

3. 运行主程序

4. 通过地址 http://127.0.0.1/swagger 查看 API 文档

### 参考

- https://github.com/swaggo/swag
- https://github.com/swaggo/swag/tree/master/example/celler/controller
