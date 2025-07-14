## Node HandsOn, Eller Klavyeye

Elimizdeki nodelar neler miş bakalım.
docker desktop kullandığımız için aşağıdaki şekilde çıkacaktır.

```bash
    kubectl get nodes
    
NAME             STATUS   ROLES           AGE     VERSION
docker-desktop   Ready    control-plane   3h38m   v1.32.2

```

Şimdi bunu bir describe edelim. Çıktısı aşağıdaki şekilde görülecek.

```bash
    kubectl describe node docker-desktop
    
Name:               docker-desktop
Roles:              control-plane
Labels:             beta.kubernetes.io/arch=arm64
                    beta.kubernetes.io/os=linux
                    kubernetes.io/arch=arm64
                    kubernetes.io/hostname=docker-desktop
                    kubernetes.io/os=linux
                    node-role.kubernetes.io/control-plane=
                    node.kubernetes.io/exclude-from-external-load-balancers=
Annotations:        kubeadm.alpha.kubernetes.io/cri-socket: unix:///var/run/cri-dockerd.sock
                    node.alpha.kubernetes.io/ttl: 0
                    volumes.kubernetes.io/controller-managed-attach-detach: true
CreationTimestamp:  Mon, 14 Jul 2025 11:37:43 +0300
Taints:             <none>
Unschedulable:      false
Lease:
  HolderIdentity:  docker-desktop
  AcquireTime:     <unset>
  RenewTime:       Mon, 14 Jul 2025 15:19:43 +0300
Conditions:
  Type             Status  LastHeartbeatTime                 LastTransitionTime                Reason                       Message
  ----             ------  -----------------                 ------------------                ------                       -------
  MemoryPressure   False   Mon, 14 Jul 2025 15:19:23 +0300   Mon, 14 Jul 2025 11:37:42 +0300   KubeletHasSufficientMemory   kubelet has sufficient memory available
  DiskPressure     False   Mon, 14 Jul 2025 15:19:23 +0300   Mon, 14 Jul 2025 11:37:42 +0300   KubeletHasNoDiskPressure     kubelet has no disk pressure
  PIDPressure      False   Mon, 14 Jul 2025 15:19:23 +0300   Mon, 14 Jul 2025 11:37:42 +0300   KubeletHasSufficientPID      kubelet has sufficient PID available
  Ready            True    Mon, 14 Jul 2025 15:19:23 +0300   Mon, 14 Jul 2025 11:37:43 +0300   KubeletReady                 kubelet is posting ready status
Addresses:
  InternalIP:  192.168.65.3
  Hostname:    docker-desktop
Capacity:
  cpu:                16
  ephemeral-storage:  1055761844Ki
  hugepages-1Gi:      0
  hugepages-2Mi:      0
  hugepages-32Mi:     0
  hugepages-64Ki:     0
  memory:             8024600Ki
  pods:               110
Allocatable:
  cpu:                16
  ephemeral-storage:  972990113820
  hugepages-1Gi:      0
  hugepages-2Mi:      0
  hugepages-32Mi:     0
  hugepages-64Ki:     0
  memory:             7922200Ki
  pods:               110
System Info:
  Machine ID:                 1f2f0e750cfc4b6eafba5eaef3060baa
  System UUID:                1f2f0e750cfc4b6eafba5eaef3060baa
  Boot ID:                    32cfca3b-42fe-4a25-b315-63823180331c
  Kernel Version:             6.10.14-linuxkit
  OS Image:                   Docker Desktop
  Operating System:           linux
  Architecture:               arm64
  Container Runtime Version:  docker://28.3.0
  Kubelet Version:            v1.32.2
  Kube-Proxy Version:         v1.32.2
Non-terminated Pods:          (10 in total)
  Namespace                   Name                                      CPU Requests  CPU Limits  Memory Requests  Memory Limits  Age
  ---------                   ----                                      ------------  ----------  ---------------  -------------  ---
  default                     nginx-pod                                 0 (0%)        0 (0%)      0 (0%)           0 (0%)         91m
  kube-system                 coredns-668d6bf9bc-fqf5b                  100m (0%)     0 (0%)      70Mi (0%)        170Mi (2%)     3h41m
  kube-system                 coredns-668d6bf9bc-zw5sp                  100m (0%)     0 (0%)      70Mi (0%)        170Mi (2%)     3h41m
  kube-system                 etcd-docker-desktop                       100m (0%)     0 (0%)      100Mi (1%)       0 (0%)         3h41m
  kube-system                 kube-apiserver-docker-desktop             250m (1%)     0 (0%)      0 (0%)           0 (0%)         3h41m
  kube-system                 kube-controller-manager-docker-desktop    200m (1%)     0 (0%)      0 (0%)           0 (0%)         3h41m
  kube-system                 kube-proxy-dqc9b                          0 (0%)        0 (0%)      0 (0%)           0 (0%)         3h41m
  kube-system                 kube-scheduler-docker-desktop             100m (0%)     0 (0%)      0 (0%)           0 (0%)         3h42m
  kube-system                 storage-provisioner                       0 (0%)        0 (0%)      0 (0%)           0 (0%)         3h41m
  kube-system                 vpnkit-controller                         0 (0%)        0 (0%)      0 (0%)           0 (0%)         3h41m
Allocated resources:
  (Total limits may be over 100 percent, i.e., overcommitted.)
  Resource           Requests    Limits
  --------           --------    ------
  cpu                850m (5%)   0 (0%)
  memory             240Mi (3%)  340Mi (4%)
  ephemeral-storage  0 (0%)      0 (0%)
  hugepages-1Gi      0 (0%)      0 (0%)
  hugepages-2Mi      0 (0%)      0 (0%)
  hugepages-32Mi     0 (0%)      0 (0%)
  hugepages-64Ki     0 (0%)      0 (0%)
Events:              <none>
❯ 
❯ 

```

Aşağıdaki şekilde bu node'u istersek production olarak işaretleyebiliriz. Şimdilik çalıştırmayın.

```bash
    kubectl label node docker-desktop env=production
```

Örneğin aşağıdaki şekilde environment'i test ortamı

```bash
    kubectl label node docker-desktop env=test
```

Örneğin Development ortamı

```bash
    kubectl label node docker-desktop env=dev
```

Örneğin Qa ortamı gibi ayrıştırabilirsiniz.

```bash
    kubectl label node docker-desktop env=qa
```

Node'un geçici olarak askıya alınmasını artık trafik yönlendirilmemesini istiyorsanız aşağıdaki *cordon* komutunu
kullanabilirsiniz.

```bash
    kubectl cordon docker-desktop
```

Tekrardan aktif olarak trafik yönlendirmesi istendiğinde

```bash
    kubectl uncordon docker-desktop
```

Bu node'daki tüm podları boşalt komutu aşağıdaki şekilde. Bir sorun durumunda

```bash
  kubectl drain docker-desktop --ignore-deamonsets --delete-emptydir-data
```

Pod dosyası üzerinde bu

```bash
    nano pod-production.yaml
```

oluşan yaml dosyası aşağıdaki şekilde. nodeSelector'e dikkat edin bunun env production olarak set edildi. böylece
yukardaki işlemlerde selector olarak bunlar kullanılacak.

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
  nodeSelector:
    env: production
```