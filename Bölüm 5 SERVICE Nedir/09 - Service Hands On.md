# Service Hands On

Öncelikle docker var dimi

```bash
    docker --version
```

Bakalım kubernetes çalışıyor mu

```bash
    kubectl get pods
```

nano ile bir namespace.yaml dosyası oluşturalım.

```yaml
# Namespace tanımı
apiVersion: v1
kind: Namespace
metadata:
  name: demo-app
```

şimdi bu namespace.yaml'ı uygulayalım.

```bash
    kubectl apply -f namespace.yaml
```

Oluşturulan namespace'lerin listesini aşağıdaki şekilde kontrol edebiliriz.

```bash
    kubectl get ns
    
    # çıktısı aşağıdaki şekilde olacaktır.
    
    NAME              STATUS   AGE
    default           Active   5h44m
    demo-app          Active   65s
    kube-node-lease   Active   5h44m
    kube-public       Active   5h44m
    kube-system       Active   5h44m
```

nano ile bir deployment.yaml dosyası oluşturalım.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: backend-deployment
  namespace: demo-app
  labels:
    app: nginx
spec:
  replicas: 3  # 3 adet nginx pod'u çalışacak
  selector:
    matchLabels:
      app: backend
  template:
    metadata:
      labels:
        app: backend
    spec:
      containers:
        - name: backend-container
          image: nginx:1.25
          ports:
            - containerPort: 80
```

şimdi önce bir podslarımızın listesini alalım.

```bash
    kubectl get pods
    
    ## çıktısı aşağıdaki gibi olmalı
    
    NAME        READY   STATUS    RESTARTS   AGE
    nginx-pod   1/1     Running   0          3h39m

```

şimdi de bu deployement.yaml dosyasını apply edelim.

```bash
    kubectl apply -f deployment.yaml
    
    ## aşağıdaki şekilde cevap almalıyız.
    deployment.apps/backend-deployment created
```

şimdi deploymentslarımızı görelim

```bash
  ## -n ile özellikle namespace belirtiyoruz.
  kubectl get deployments -n demo-app
  
  ## aşağıdaki şekilde çıktı almalıyız.
  
  NAME                 READY   UP-TO-DATE   AVAILABLE   AGE
  backend-deployment   3/3     3            3           3m4s
  
```

şimdi demo-app name space'indeki podları bana ver dediğimizde.

```bash
    kubectl get pods -n demo-app
    
    ## çıktının aşağıdaki gibi olmasını bekleriz.
    
    NAME                                  READY   STATUS    RESTARTS   AGE
    backend-deployment-6d8bc7b5d4-8rgbm   1/1     Running   0          4m38s
    backend-deployment-6d8bc7b5d4-j6jk5   1/1     Running   0          4m38s
    backend-deployment-6d8bc7b5d4-nzg2m   1/1     Running   0          4m38s

```

şimdi demo-app name space'i için bir servis oluşturulmuş mu ona bakalım.

```bash
    kubectl get services -n demo-app
    
    ## henüz bir servis oluşturmadığımız için aşağıdaki gibi bir çıktı alırız.
    No resources found in demo-app namespace.
```

şimdi bir service.yaml dosyası oluşturalım.

```yaml
# service.yaml
# Statik IP atanmış = ClusterIP tipi Service
apiVersion: v1
kind: Service
metadata:
  name: backend-service
  namespace: demo-app
spec:
  type: ClusterIP     # varsayılan, içeriden erişim sağlar
  clusterIP: 10.96.0.250
  selector:
    app: backend  # deployment'taki label ile eşleşmeli
  ports:
    - protocol: TCP
      port: 80        # Servisin dışa açtığı port
      targetPort: 80  # Pod içindeki container port
```

şimdi bunu apply edelim.

```bash
    kubectl apply -f service.yaml
    
    # çıktı olarak aşağıdaki şekilde bir sonuç görmeliyiz.
    service/backend-service created
```

şimdi bu servis oluşmuş mu bakalım

```bash
    kubectl get services -n demo-app
    
    # Bunun çıktısını aşağıdaki şekilde görmeliyiz.
    
    NAME              TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)   AGE
    backend-service   ClusterIP   10.96.0.250   <none>        80/TCP    101s
 
```

Describe'ına bakalım.

```bash
    kubectl describe services -n demo-app
    
    
    # Çıktısı aşağıdaki şekilde olmalı.
    Name:                     backend-service
Namespace:                demo-app
Labels:                   <none>
Annotations:              <none>
Selector:                 app=backend
Type:                     ClusterIP
IP Family Policy:         SingleStack
IP Families:              IPv4
IP:                       10.96.0.250
IPs:                      10.96.0.250
Port:                     <unset>  80/TCP
TargetPort:               80/TCP
Endpoints:                10.1.0.8:80,10.1.0.7:80,10.1.0.9:80
Session Affinity:         None
Internal Traffic Policy:  Cluster
Events:                   <none>

```