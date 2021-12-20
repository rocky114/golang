### 安装
- go get -u google.golang.org/protobuf/cmd/protoc-gen-go
- go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
- go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway

---

### 生成pb文件
- protoc --proto_path=./pb --go_out . --go-grpc_out . --grpc-gateway_out . ./pb/*.proto

### 打包docker镜像
- docker build -t rocky114/client:0.1.1 -f build/client/Dockerfile .
- docker build -t rocky114/server:0.1.2 -f build/server/Dockerfile .
- docker push new-repo:tagname //镜像上传


### k8s部署
- kubectl create deployment gprc-server --image=rocky114/server:0.1.2 --port=5001 //拉取镜像部署server实例
- kubectl expose deployment gprc-server --type=NodePort //暴露端口服务
- kubectl scale deployments/gprc-server --replicas=4 //扩容pods
- kubectl create deployment grpc-client --image=rocky114/client:0.1.1 --port=9999 //拉取镜像部署client实例
- kubectl expose deployment/gprc-client --type=NodePort //暴露端口服务