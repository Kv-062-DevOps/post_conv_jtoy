## How to start the Demo project with all services in the Minikube

Download repository:
`git clone https://github.com/Kv-062-DevOps/post_conv_jtoy.git`

YAML configuration files for Kubernetes contain in the `kube` directory.  

Open the folder **`post_conv_jtoy`** in commandline console and execute:
```
minikube start
kubectl apply -f kube
kubectl config set-context --current --namespace=demo  
```
To initialize Database you can try two ways. Using three commands step-by-step:
```
kubectl apply -f kube-init-separate\create-kube.yaml
kubectl apply -f kube-init-separate\load-kube.yaml
kubectl apply -f kube-init-separate\back-only-kube.yaml
```
or next one:
    `kubectl apply -f kube-back-allinone`

**Important notice**. In last case there are no separate files to create DB and load data. 
All initiations applied in one Backend config (`back-kube.yaml`), section _**initContainers**_.

After that open in your web browser link from command
`minikube service front-srv -n demo --url`
(for example, <http://172.17.0.2:30808>).  

---
To visit a Kubernetes WEB dashboard, use command **`minikube dashboard`**

If you want to see a raw data in the GET service, use:
`minikube service get-srv -n demo --url`
for example, link will be <http://172.17.0.2:30593>

If you want to see a raw data in the Backend service, get the address using command:
`minikube service back-srv -n demo --url`
and open this link in browser, adding a path `/list`, for example: <http://172.17.0.2:31486/list>

### Clear resources after work:
```
kubectl config set-context --current --namespace=demo

kubectl delete deployments back get post db front
kubectl delete services back-srv get-srv post-srv db-srv front-srv 
 
kubectl config set-context --current --namespace=default
 
minikube stop

```
___












