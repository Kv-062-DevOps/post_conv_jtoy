

```
minikube start

kubectl apply -f kube

kubectl config set-context --current --namespace=demo

```


`minikube service front-srv -n demo --url`
> http://172.17.0.2:30808



`minikube dashboard`


http://172.17.0.2:30808/about
http://172.17.0.2:30808/employee
http://172.17.0.2:30808/home


```
kubectl config set-context --current --namespace=demo
 
kubectl delete deployments back get post db front
kubectl delete services back-srv get-srv post-srv db-srv front-srv 
 
kubectl config set-context --current --namespace=default
 
minikube stop
 
```





---

`minikube service get-srv -n demo --url`
> http://172.17.0.2:30593

http://172.17.0.2:30593


`minikube service back-srv -n demo --url`
> http://172.17.0.2:31486

http://172.17.0.2:31486/list











