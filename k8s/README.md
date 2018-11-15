# Workshop3
It's time to play on k8s

We will deploy tomcat application and make it accessible form internet. 



## Create our first app.
Create deployment for our app first. 
Just 

### Create namespace.
```sh
$ kubectl create namespace <yourName>
```

Rememeber to add `-n <yourName>` to specify which namespace you will be working on.

### Create deployment.
```sh
$ kubectl -n <namespace> create -f deployment.yaml
```

To check deployment and pod status
```sh
$ kubectl get deployment
$ kubectl get pod

### status will be pending, running 
### if it not something like that we can run this to check what is happen
$ kubectl describe deployment <deployment name>
$ kubectl describe pod <pod name> 
$ kubectl describe logs <pod name>
```


### Create service.
```sh
$ kubectl -n <namespace> create -f service.yaml
```
To check service status
```sh
$ kubectl get service # can add -w to watch changes and ctrl+c to exit
```
When you see the `external IP` you can use it to access you app.:tada:

## Let get some grean

### Scalability and Load balancing  
We can scale our pod to serve workload with 1 single command.

```sh
$ kubectl -n <namespace> scale deployment <deployment name> --replicas=<Number of pod | just 2 or 3 is enough nah we have limited budget :cry:>
``` 

To check that is loadbalancing we have to do some work. 

### Mark pod
Edit tomcat homepage at `webapps/ROOT/index.jsp` or your html content. We have to do that for all pod. 

```sh
$ kubectl -n <namespace> exec -it <pod name> /bin/bash
```

After finish, try refresh browser and see the result.

### Pod are mmortal
Pod can be self-healing, if pod crash or node down. k8s will create pod to desire state.

#Be careful pod has got too many aegis.:trollface: 

To delete pod 
```sh
$ kubectl -n <namespace> delete pod <podName>
```
Try to delete all pod.
Check the result by kubectl or refresh browser.
 

### But wait where is our edit file

From `Mark pod` we edit some content in the pod. When we delete all pod those file is missing? 

Because deployment is stateless application. If pod is removed, all of data will be removed too.

If we want to host somthing that have to use persistent data like `databsae`, `stateful application`. We have to use other controller named `statefulset` :floppy_disk:

## Create statefulset 
Now it time to create `statefulset` for host somthing that need to store persistent data.


You can edit `./statefulset.yaml` for your own deplyoment such as naming, path to mount.

```sh
$ kubectl -n <namespace> create -f statefulset.yaml
```

### Go make some data for testing.

Get into the pod and create some content in your persistent path.

```sh
$ kubectl -n <namespace> exec <pod-name> /bin/bash
```
In pod. You can do it in your own way na.
```sh
echo "is this persistent xso ?" >  yesitis
```

After we create some data, we can delete that pod and see result.





