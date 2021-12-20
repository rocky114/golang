#### 命令
- minikube start --image-mirror-country='cn'

- minikube dashboard --url

- kubectl create deployment hello --image=nginx:1.7.9 --port=80

- kubectl get deployments

- kubectl get pods -o wide

- kubectl expose deployment hello --type=NodePort

- kubectl get services

- minikube service hello --url

- kubectl scale deployments/kubernetes-bootcamp --replicas=4 

- kubectl describe deployments/kubernetes-bootcamp

- kubectl get rs

#### 搭建集群
- kubectl create deployment grpc-server --image=rocky114/server:v1 --port=5001


#### k8s授权
- kubectl create role pod-reader-role --verb=get --verb=watch --resource=endpoints,services
- kubectl create rolebinding pod-reader-rb --role=pod-reader-role --serviceaccount=default:default