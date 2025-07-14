# ReplicaSet Yaml Analizi

name : güncelleme için lazım olabilir.

```bash
apiVersion: apps/v1
kind: ReplicaSet
metadata:
  name: my-replicaset
spec:
  replicas: 3
  selector:
    matchLabels:
      app: my-app
  template:
    metadata:
      labels:
        app: my-app
    spec:
      containers:
        - name: my-app-container
          image: my-app-image:v1 # Kullanılacak container image
          ports:
            - containerPort: 8080
```