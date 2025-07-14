## Bir tane pod.yaml dosyası oluşturuyoruz
içeriği aşağıdaki şekildedir.
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx-pod
spec:
  containers:
    - name: nginx
      image: nginx:latest
      ports:
        - containerPort: 80

```

Aşağıdaki şekilde bunu çalıştırıyoruz.

```bash
  kubectl apply -f pod.yaml
```


## Get pods
Şimdi bunu kontrol etmek için aşağıdaki komutu çalıştırdığımızda bunun kurulduğunu görüyoruz.

```bash
  kubectl get pods
```

## Describe komutu
```bash
    kubectl describe pod nginx-pod
```
çıktı olarak aşağıdaki bir çıktı görürüz.
```bash
Name:             nginx-pod
Namespace:        default
Priority:         0
Service Account:  default
Node:             docker-desktop/192.168.65.3
Start Time:       Mon, 14 Jul 2025 13:48:34 +0300
Labels:           <none>
Annotations:      <none>
Status:           Running
IP:               10.1.0.6
IPs:
  IP:  10.1.0.6
Containers:
  nginx:
    Container ID:   docker://173222ffe22151d8a92056dd14069148a647d7a8e10b27f47bbc6dabbe49863d
    Image:          nginx:latest
    Image ID:       docker-pullable://nginx@sha256:93230cd54060f497430c7a120e2347894846a81b6a5dd2110f7362c5423b4abc
    Port:           80/TCP
    Host Port:      0/TCP
    State:          Running
      Started:      Mon, 14 Jul 2025 13:48:44 +0300
    Ready:          True
    Restart Count:  0
    Environment:    <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from kube-api-access-kttx9 (ro)
Conditions:
  Type                        Status
  PodReadyToStartContainers   True 
  Initialized                 True 
  Ready                       True 
  ContainersReady             True 
  PodScheduled                True 
Volumes:
  kube-api-access-kttx9:
    Type:                    Projected (a volume that contains injected data from multiple sources)
    TokenExpirationSeconds:  3607
    ConfigMapName:           kube-root-ca.crt
    Optional:                false
    DownwardAPI:             true
QoS Class:                   BestEffort
Node-Selectors:              <none>
Tolerations:                 node.kubernetes.io/not-ready:NoExecute op=Exists for 300s
                             node.kubernetes.io/unreachable:NoExecute op=Exists for 300s
Events:
  Type    Reason     Age    From               Message
  ----    ------     ----   ----               -------
  Normal  Scheduled  7m53s  default-scheduler  Successfully assigned default/nginx-pod to docker-desktop
  Normal  Pulling    7m52s  kubelet            Pulling image "nginx:latest"
  Normal  Pulled     7m43s  kubelet            Successfully pulled image "nginx:latest" in 8.715s (8.715s including waiting). Image size: 68741916 bytes.
  Normal  Created    7m43s  kubelet            Created container: nginx
  Normal  Started    7m43s  kubelet            Started container nginx

```

## log komutuna baktığımızda
Ne durumdayız diye bakmak için.

```bash
  kubectl logs nginx-pod
  
  /docker-entrypoint.sh: /docker-entrypoint.d/ is not empty, will attempt to perform configuration
/docker-entrypoint.sh: Looking for shell scripts in /docker-entrypoint.d/
/docker-entrypoint.sh: Launching /docker-entrypoint.d/10-listen-on-ipv6-by-default.sh
10-listen-on-ipv6-by-default.sh: info: Getting the checksum of /etc/nginx/conf.d/default.conf
10-listen-on-ipv6-by-default.sh: info: Enabled listen on IPv6 in /etc/nginx/conf.d/default.conf
/docker-entrypoint.sh: Sourcing /docker-entrypoint.d/15-local-resolvers.envsh
/docker-entrypoint.sh: Launching /docker-entrypoint.d/20-envsubst-on-templates.sh
/docker-entrypoint.sh: Launching /docker-entrypoint.d/30-tune-worker-processes.sh
/docker-entrypoint.sh: Configuration complete; ready for start up
2025/07/14 10:48:44 [notice] 1#1: using the "epoll" event method
2025/07/14 10:48:44 [notice] 1#1: nginx/1.29.0
2025/07/14 10:48:44 [notice] 1#1: built by gcc 12.2.0 (Debian 12.2.0-14+deb12u1) 
2025/07/14 10:48:44 [notice] 1#1: OS: Linux 6.10.14-linuxkit
2025/07/14 10:48:44 [notice] 1#1: getrlimit(RLIMIT_NOFILE): 1048576:1048576
2025/07/14 10:48:44 [notice] 1#1: start worker processes
2025/07/14 10:48:44 [notice] 1#1: start worker process 29
2025/07/14 10:48:44 [notice] 1#1: start worker process 30
2025/07/14 10:48:44 [notice] 1#1: start worker process 31
2025/07/14 10:48:44 [notice] 1#1: start worker process 32
2025/07/14 10:48:44 [notice] 1#1: start worker process 33
2025/07/14 10:48:44 [notice] 1#1: start worker process 34
2025/07/14 10:48:44 [notice] 1#1: start worker process 35
2025/07/14 10:48:44 [notice] 1#1: start worker process 36
2025/07/14 10:48:44 [notice] 1#1: start worker process 37
2025/07/14 10:48:44 [notice] 1#1: start worker process 38
2025/07/14 10:48:44 [notice] 1#1: start worker process 39
2025/07/14 10:48:44 [notice] 1#1: start worker process 40
2025/07/14 10:48:44 [notice] 1#1: start worker process 41
2025/07/14 10:48:44 [notice] 1#1: start worker process 42
2025/07/14 10:48:44 [notice] 1#1: start worker process 43
2025/07/14 10:48:44 [notice] 1#1: start worker process 44

```

### İlgili Pod'un terminaline erişmek için
Aşağıdaki şekilde ilgili pod'a terminal seviyesinde erişilebiliyor.
- it : Interactive Terminal
- bash yada sh denilebilir
```bash
  kubectl exec -it nginx-pod -- bash
```