
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
    
使用第三方库进行字段验证

https://github.com/envoyproxy/protoc-gen-validate

安装 go get -u github.com/envoyproxy/protoc-gen-validate


proto 文件

    // 引人 protoc-gen-validate 文件
    import "validate.proto";
    
    
    // 主订单信息
    message OrderMain {
      int32 order_id = 1; // 订单id
      //int32 pro_id = 2; // 商品id
      string order_no = 3; // 订单号
      int32 user_id = 4[(validate.rules).int32.gt = 1000]; // 用户id,加个校验必须大于1000
      google.protobuf.Timestamp order_time = 5;
    
      // 嵌套 模型
      repeated OrderDetail order_details = 6;
    }
    
增加编译选项

    --validate_out=lang=go:
    
    #!/usr/bin/env bash
    
    current_path=$PWD
    
    cd $current_path/protofiles
    
    
    protoc --go_out=plugins=grpc:../server/services --validate_out=lang=go:../server/services *.proto
    protoc --grpc-gateway_out=logtostderr=true:../server/services --validate_out=lang=go:../server/services *.proto
    
    protoc --go_out=plugins=grpc:../client/services --validate_out=lang=go:../client/services *.proto
    protoc --grpc-gateway_out=logtostderr=true:../client/services --validate_out=lang=go:../client/services *.proto
    
 
 golang 代码引人校验逻辑
     .....
     err := request.OrderMain.Validate()
        if err != nil {
            resp := &OrderResponse{
                ErrCode: 1,
                ErrMsg:  err.Error(),
            }
            return resp,nil
        }
     ..........
     
流模式入门

为何要用流模式

基本模式是客户端请求--->服务端响应

如果是传输较大数据呢？会带来

1. 数据包过大导致压力徒增
2. 需要等待客户端包全部发完，才能处理及响应

举例：

假设我们要从库里取一批(x万到几十万)批量查询用户的积分

先创建用户模型

    message UserInfo {
        int32 user_id = 1;
        int32 user_score = 2;
    }
    
服务端流模式 (服务端分批返回)

场景:客户端批量查询用户积分

1. 客户端一次性把用户列表发生过来，单查询的用户列表很小时，很快返回，要是查询比较多时，由于服务端响应较慢
2. 服务端查询积分比较耗时，因此需要分批返回查到一批返回一批，二不是完全查完在全部返回

修改proto 文件

    service UserService {
      rpc GetUserScore(UserScoreRequest) returns (UserScoreResponse);
      // 服务端流模式
      rpc GetUserScoreByServerStream(UserScoreRequest) returns (stream UserScoreResponse);
    }
    
生成protoc 生成go 文件

修改服务端service 方法

    // GetUserScoreByServerStream 服务端流模式方法.
    func (u *UserService) GetUserScoreByServerStream(request *UserScoreRequest, stream UserService_GetUserScoreByServerStreamServer) error {
        var score int32 = 101
        users := make([]*UserInfo,0)
        for index,user := range request.Users{
            user.UserScore = score
            score++
            users = append(users,user)
            // 分批每两次发一条
            if (index+1) %2 == 0 && index >0 {
                err := stream.Send(&UserScoreResponse{Users: users})
                if err != nil {
                    return err
                }
                // 每次清空切片
                users = (users)[0:0]
            }
        }
        // users 里面还有值在发一次
        if len(users) > 0 {
            err := stream.Send(&UserScoreResponse{Users: users})
            if err != nil {
                return err
            }
        }
        return nil
    }
    
修改客户端调用方法

    // 客户端请求服务端流模式方法
        serverStream, err := serviceClient.GetUserScoreByServerStream(ctx, &request)
        if err != nil {
            log.Fatal(err)
        }
        // 循环读
        for {
            userScoreResponse, err := serverStream.Recv()
            // 结束了
            if err == io.EOF {
                break
            }
            if err != nil {
                log.Fatal(err)
            }
            fmt.Println(userScoreResponse)
        }


客户端流模式

场景：客户端批量查询用户积分

1.客户端一次性把用户列表发送过去（客户端获取列表较慢）
2. 服务查询比较快
因此需要使用客户端流模式


修改proto 文件

    service UserService {
      rpc GetUserScore(UserScoreRequest) returns (UserScoreResponse);
      // 服务端流模式
      rpc GetUserScoreByServerStream(UserScoreRequest) returns (stream UserScoreResponse);
      // 客户端流模式
      rpc GetUserScoreByClientStream(stream UserScoreRequest) returns (UserScoreResponse);
    }

生成protoc 生成go 文件

修改服务端service 方法

    // GetUserScoreByClientStream 客户端流模式方法.
    func (u *UserService)GetUserScoreByClientStream(stream UserService_GetUserScoreByClientStreamServer) error {
        var score int32 = 101
        users := make([]*UserInfo,0)
        for {
            request, err := stream.Recv()
    
            // 接收完毕,发送出去
            if err == io.EOF {
                return stream.SendAndClose(&UserScoreResponse{Users: users})
            }
            if err != nil {
                return err
            }
    
            // 服务端的业务逻辑
            for _,user := range request.Users{
                user.UserScore = score
                score++
                users = append(users,user)
            }
        }
    
        return nil
    }
    
修改客户端调用方法

    // 客户端流模式调用
	clientStreamClient, err := serviceClient.GetUserScoreByClientStream(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// 循环5次每次发5条 批量
	for j:=1;j<=5;j++{
		request2 := &services.UserScoreRequest{
			Users: users,
		}
		users2 := make([]*services.UserInfo,0)
		for i:=0;i<5;i++ {
			user := services.UserInfo{
				UserId: int32(i),
			}
			users2 = append(users2,&user)
		}
		err := clientStreamClient.Send(request2)
		if err != nil {
			log.Println(err)
		}
	}
	// 最后获取响应
	recv, err := clientStreamClient.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(recv)
    



双向流的基本套路

修改proto 文件

    service UserService {
      rpc GetUserScore(UserScoreRequest) returns (UserScoreResponse);
      // 服务端流模式
      rpc GetUserScoreByServerStream(UserScoreRequest) returns (stream UserScoreResponse);
      // 客户端流模式
      rpc GetUserScoreByClientStream(stream UserScoreRequest) returns (UserScoreResponse);
      // 双向流模式
      rpc GetUserScoreByStream(stream UserScoreRequest) returns (stream UserScoreResponse);
    }
    
protoc 生成go文件

修改服务端service 

    // GetUserScoreByStream 双向流
    func (u *UserService)GetUserScoreByStream(stream UserService_GetUserScoreByStreamServer) error {
    
        var score int32 = 101
        users := make([]*UserInfo,0)
        for {
            request, err := stream.Recv()
    
            // 接收完毕,发送出去
            if err == io.EOF {
                return nil
            }
            if err != nil {
                return err
            }
    
            // 服务端的业务逻辑
            for _,user := range request.Users{
                user.UserScore = score
                score++
                users = append(users,user)
            }
            err = stream.Send(&UserScoreResponse{Users: users})
            if err != nil {
                return err
            }
            users = (users)[0:0]
        }
    }
    
修改客户端调用

    // 双向流模式
    scoreByStream, err := serviceClient.GetUserScoreByStream(ctx)
    if err != nil {
        log.Fatal(err)
    }
    // 循环5次每次发5条 批量
    for j:=1;j<=5;j++{
        request2 := &services.UserScoreRequest{
            Users: users,
        }
        users2 := make([]*services.UserInfo,0)
        for i:=0;i<5;i++ {
            user := services.UserInfo{
                UserId: int32(i),
            }
            users2 = append(users2,&user)
        }
        err := scoreByStream.Send(request2)
        if err != nil {
            log.Println(err)
        }
        scoreResponse, err := scoreByStream.Recv()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(scoreResponse)
    }