# Comandos

```sh
kubectl apply -f <PATH>
```

**Nodes**

```sh
kubectl get nodes
```

**Pods**

```sh
kubectl get pods
kubectl delete pod <NAME>
kubectl describe pod <NAME>
```

**Replicasets**

```sh
kubectl get replicasets
kubectl delete replicasets <NAME>
```

**Deployments**

```sh
kubectl get deployments
```

**Rollout**

```sh
kubectl rollout history deployment <NAME>
kubectl rollout undo deployment <NAME>
kubectl rollout undo deployment <NAME> --to-revision=<NUMBER>
```

**Port forward**

```sh
kubectl port-forward <TYPE>/<NAME> <PORT>:<TARGET_PORT>
```

**Services**

```sh
kubectl get services
kubectl get svc
```
