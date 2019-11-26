docker build -t yoyogodemo:v1 .

docker run --rm -p 8080:8080  --name yoyodemov1 yoyogodemo:v1

kubectl create -f k8s-deploy.yaml

kubectl create -f k8s-service.yaml

kubectl expose deployment yoyogodemo --type="LoadBalancer"
