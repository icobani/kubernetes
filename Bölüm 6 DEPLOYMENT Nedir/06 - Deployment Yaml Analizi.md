# Deployment Yaml Analizi

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-app-deployment
spec:
  replicas: 3  # Aynı anda çalışacak pod sayısı
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

### Açıklamalar:

* replicas: Kaç adet pod çalıştırılacağını belirtir.
* selector: Bu deployment’ın hangi pod’ları yöneteceğini tanımlar (matchLabels ile).
* template: Pod şablonudur. Burada container’lar tanımlanır.
* containers: İçinde çalışacak uygulama tanımlanır. Bu örnekte nginx container’ı kullanılıyor.