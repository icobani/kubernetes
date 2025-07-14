# Dry Run
İlgili POD'u oluşturmadan yaml dosyasını oluşturmak için kullanılan fonksiyondur. Bu parametreleri ezberlemek yerine otomatik pod dosyası generate etmek için kullanılan bir yöntemdir.


Aşağıdaki komutu çalıştırdığınızda hiç pod nasıl yazarım diye uğraşmadan bir pod dosyasını sizin için oluşturur

```bash
    kubectl run nginx2-pod --image=nginx:latest --dry-run=client -o yaml > pod2.yaml
```

Oluşan pod2.yaml dosyası aşağıdaki şekilde

```yaml
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    run: nginx2-pod
  name: nginx2-pod
spec:
  containers:
  - image: nginx:latest
    name: nginx2-pod
    resources: {}
  dnsPolicy: ClusterFirst
  restartPolicy: Always
status: {}
```

yani komutu yazıp horizatal mod autoscale'i nasıl yazarım, deployment'i nasıl yazarım bunları bulmak yerine otomatik
oluşturmasını sağlayabiliyorsunuz.