# Workshop3
It's time to play on k8s

We will deploy our application and make it accessible form internet.



## Deploy our first app.
We will create `deployment`, controller of pod in k8s.
In k8s there is a component named `namespaces`. We can use it for manage our resourse. Every resources will be inside 1 namespace only. Can not span across namespace.

### Create namespace.
```sh
kubectl create namespace patty
```

Rememeber to add `-n <yourNameSpace>` to specify which namespace you will be working on.

### Create deployment.
```sh
kubectl -n <namespace> create -f k8s/deployment.yaml
```

To check deployment and pod status
```sh
kubectl  -n <namespace> get deployment
kubectl -n <namespace> get pod

### status will be pending, running
### if it not something like that we can run this to check what is happen
kubectl -n <namespace> describe deployment <deployment name>
kubectl -n <namespace> describe pod <pod name>
```
From this we should see our pod is in Running status.
Then we will make our app access from internet.

### Create service.
```sh
kubectl -n <namespace> create -f k8s/service.yaml
```
To check service status
```sh
kubectl get -n <namespace> service # can add -w to watch changes and ctrl+c to exit
```
When you see the `external IP` you can use it to access you app.:tada:

`/` Welcome page
`/version` Display information

## Let's play more

### Scalability and Load balancing
We can scale our pod to serve workload with 1 single command.

```sh
kubectl -n <namespace> scale deployment workshop3 --replicas=2
```
After finish, try refresh browser and see the result.

### Logging
We can get logs from pod by run
```sh
kubectl -n <namespace> logs <pod name>
```
We can also add paramater `-f` to follow new comming logs.
Note that this logs will come from stdout/stderror only.

### Pod are mmortal
Pod can be self-healing, if pod crash or node down. k8s will create pod to desire state.

To delete pod
```sh
kubectl -n <namespace> delete pod <podName>
```
Try to delete all pod.
Check the result by kubectl or refresh browser.


### Secrets
Some of application use some credential for connecting to other service. Such as database url, api key.  It is not good to hardcode it inside the code. This might lead to some bad situation.

In kubernets, there is component name secret for store our credentials and provide it for application. So we no longer need to hardcode it anymore.

#### Create secrets
This secret is in form of key-values. You can edit key and vault if you want to.
But for the value needs to encode with base64.

```sh
kubectl -n <namespace> create -f k8s/secrets.yaml
```

#### Use secrets as environment variable
Now it's time to let applcation use secret. In k8s, we can retrive secret by environment variable. So let edit our app to use env.

In `src/handles.go` line 24,25 You can uncomment it and comment the line 26,27.
And `k8s/deployment.yaml` uncomment 20-30.
Please make sure env and key of secrets are match.

When we finish coding then how can we deploy to k8s?
### Update deployment
Let deployment know to use env.
```sh
kubectl -n <namespace> apply -f k8s/deployment.yaml
```
But application still not update yet.

### Update new image
In k8s we can use `kubectl set image` to update image of container.

But before do that we have to prepare new image. Can you remember how to do that ?

```sh
docker build -t asia.gcr.io/workshop-mfec/workshop3:v2 .
```

Then push image to GCR
```sh
docker push asia.gcr.io/workshop-mfec/workshop3:v2
```

Now our image is ready. Next we'll update `deployment`
```sh
kubectl -n <namespace> set image deployment/workshop3 workshop3=asia.gcr.io/workshop-mfec/workshop3:v1
```
We can check status by
```sh
kubectl -n <namespace> rollout status deployment workshop3
```

When k8s've updated already, go to your service and check result.
