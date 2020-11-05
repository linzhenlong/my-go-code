
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

加入自鉴证书

执行命令

    1. openssl
    2. 执行 `genrsa -des3 -out server.key 2048` 会生成(server.key)的私钥文件
    3. 创建请求证书 `req -new -key server.key -out server.csr` 会生成server.csr 文件，其中Common Name 也就是域名
    4. 删除密码执行:`rsa -in server.key -out server_no_password.key`
    5. 执行 `x509 -req -days 365 -in server.csr -signkey server_no_password.key -out server.crt`  


使用双向证书

使用CA证书

生成CA证书

1. 根证书(root certificate) 是属于根证书办法机构（CA）的公钥证书。用于验证他签发的证书(客户端,服务端)
2. 执行openssl
3. genrsa -out ca.key 2048
4. req -new -x509 -days 3650 -key ca.key -out ca.pem ，Common Name 填localhost

生成服务端证书

0. 执行openssl
1. genrsa -out server.key 2048
2. req -new -key server.key -out server.csr ，注意Common Name 必须填写localhost
3. x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in server.csr -out server.pem 

生成客户端证书

1. ecparam -genkey -name secp384r1 -out client.key
2. req -new -key client.key -out client.csr
3. x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in client.csr -out client.pem

程序中重新覆盖server.pem,ca.pem 和server.key


双向认证下rpc-gateway使用(同时提供rpc和http服务)

grpc-geteway(https://github.com/grpc-ecosystem/grpc-gateway)

相当于在grpc之上加一层代理并转发，转变成protobuf格式来访问grpc服务


    go install \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
        google.golang.org/protobuf/cmd/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc
        
 .proto 文件引入
     import "google/api/annotations.proto";
     
     service SearchService {
       rpc GetArticles(SearchRequest) returns (SearchResponse) {
         option (google.api.http) = {  // 添加 google http api
           get:"/v1/search/{query}"
         };
       }
     }
     
生成两个文件

cd 进入.proto文件夹

1.生成Search.pb.go 文件

    protoc --go_out=plugins=grpc:../services Search.proto
    
2. 在生成gateway 文件 search.pb.gw.go

    protoc --grpc-gateway_out=logtostderr=true:../services Search.proto
    