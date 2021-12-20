### 命令

- docker build -t rocky114/server:v1 . //打包镜像
- docker tag local-image:tagname new-repo:tagname //打标签
- docker push new-repo:tagname //镜像上传
- docker run -p 5001:5001 -d grpc-server:v1 //启动docker镜像
