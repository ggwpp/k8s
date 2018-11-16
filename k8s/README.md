## Statefulset

### Mark pod
Edit tomcat homepage at `webapps/ROOT/index.jsp` or your html content. We have to do that for all pod.
```sh
kubectl -n <namespace> exec -it <pod name> /bin/bash
```

After finish, try refresh browser and see the result.

#### But wait where is our edit file

From `Mark pod` we edit some content in the pod. When we delete all pod those file is missing?

Because deployment is stateless application. If pod is removed, all of data will be removed too.

If we want to host somthing that have to use persistent data like `databsae`, `stateful application`. We have to use other controller named `statefulset` :floppy_disk:

Now it time to create `statefulset` for host somthing that need to store persistent data.

You can edit `k8s/statefulset.yaml` for your own deplyoment such as naming, path to mount.

```sh
kubectl -n <namespace> create -f statefulset.yaml
```

### Go make some data for testing.

Get into the pod and create some content in your persistent path.

```sh
kubectl -n <namespace> exec <pod-name> /bin/bash
```
In pod. You can do it in your own way na.
```sh
echo "Is this persistent xso ?" >  yesitis
```

After we create some data, we can delete that pod and see result.
