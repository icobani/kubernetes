# NodePort (Dış Erişim için)

* Her worker node üzerinde belirlenen bir poru dış dünyaya açar.
* Dışarıdan erişmek için: NodePort şeklinde bağlantı yapılır.
* ClusterIP içerir, yani küme içinden de erişilebilir.

Örnek yaml

<service-name>.<namespace>.svc.cluster.local

my-nodeport-service.demo-app.svc.cluster.local

```bash
    curl http://<NodeIP>:30007
```

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-nodeport-service
  namespace: demo-app
spec:
  type: NodePort
  selector:
    app: my-app
  ports:
    - protocol: TCP
      port: 80        # Servis portu (Cluster içinden erişim bu portla olur)
      targetPort: 80  # Pod içindeki container portu
      nodePort: 30007 # 30000-32767 aralığında olmalıdır.
```