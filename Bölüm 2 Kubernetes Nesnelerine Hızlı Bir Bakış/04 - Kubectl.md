- Docker Desktop'un çalıştığından emin olun.


kubectl cluster-info

Aşağıdaki sonucu almanız durumunda control plane'e 6443 port'undan erişebildiğinizi gösterir.
```
❯ kubectl cluster-info
Kubernetes control plane is running at https://127.0.0.1:6443
CoreDNS is running at https://127.0.0.1:6443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy

To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.
```




kubectl get pods -A

Sistemde yüklü olan kubernetes bileşenlerini gösterir. Name space'i kube-system olanlar kubernetes bileşenleri bunları -A paremetresi ile görebiliyoruz. -A parametresini vermezseniz kendi yapılandırdıklarınızı görürsünüz.
```bash
❯ kubectl get pods -A
NAMESPACE     NAME                                     READY   STATUS    RESTARTS   AGE
kube-system   coredns-668d6bf9bc-fqf5b                 1/1     Running   0          8m41s
kube-system   coredns-668d6bf9bc-zw5sp                 1/1     Running   0          8m41s
kube-system   etcd-docker-desktop                      1/1     Running   0          8m45s
kube-system   kube-apiserver-docker-desktop            1/1     Running   0          8m41s
kube-system   kube-controller-manager-docker-desktop   1/1     Running   0          8m40s
kube-system   kube-proxy-dqc9b                         1/1     Running   0          8m41s
kube-system   kube-scheduler-docker-desktop            1/1     Running   0          8m47s
kube-system   storage-provisioner                      1/1     Running   0          8m39s
kube-system   vpnkit-controller                        1/1     Running   0          8m39s
```

kubectl get nodes

kubernetes nodelarının listesini gösterir.
```bash
❯ kubectl get nodes
NAME             STATUS   ROLES           AGE   VERSION
docker-desktop   Ready    control-plane   14m   v1.32.2
```
