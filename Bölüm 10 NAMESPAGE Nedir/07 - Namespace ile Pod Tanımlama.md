# Namespace ile Pod Tanımlama

# 1. Namespace tanımı

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: my-namespace
```

---

# 2. Pod tanımı (namespace içinde)

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: my-app-pod
  namespace: my-namespace
  labels:
    app: my-app
spec:
  containers:
    - name: my-container
      image: nginx:latest
      ports:
        - containerPort: 80
```
