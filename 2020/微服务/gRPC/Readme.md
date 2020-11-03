
# gRPC

https://www.bilibili.com/video/BV1Fa4y1i7C6

## grpc-go 
    https://github.com/grpc/grpc-go

### Protobuf

Google Protocol Buffer(简称Protobuf)

轻便高效的序列化数据结构的协议，可以用于网络通信和存储

特点：
    1. 性能高
    2. 传输快
    3. 维护方便
    4. 反正是各种好，各种棒

* github地址 https://github.com/protocolbuffers/protobuf
* golang库所属地址 https://github.com/golang/protobuf

安装下载protobuf编译器 

    https://github.com/protocolbuffers/protobuf/releases


解压后放到喜欢的路径下/Users/smzdm/webroot/golang/bin

然后加到环境变量里


protocol-buffers文档：https://developers.google.com/protocol-buffers/docs/gotutorial


安装插件

    go get github.com/golang/protobuf/protoc-gen-go

    此时会在GOPATH的bin目录下生成可执行文件.protobuf的编译器插件protoc-gen-go
    等下我们执行protoc命令时会自动调用这个插件

创建中间文件

    syntax="proto3";
    package services;

    message ProductRequest {
        int32 prod_id = 1; // 传入的商品ID
    }
    message ProdResponse {
        int32 prod_stock=1;// 商品库存
    }

然后执行 protoc --go_out=../services/Prod.proto


添加service 

    syntax="proto3";
    package services;
    option go_package=".;services";
    message ProductRequest {
        int32 prod_id = 1; // 传入的商品ID
    }
    message ProdResponse {
        int32 prod_stock=1;// 商品库存
    }
    service ProdService {
        rpc GetProdStock(ProductRequest) returns (ProdResponse);
    }


在执行 protoc --go_out=plugins=grpc:../services Product.proto