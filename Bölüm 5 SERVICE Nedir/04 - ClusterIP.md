# Cluster IP (Varsayılan Servis Türü)

* Cluster içindeki bileşenlerin birbirine erişmesini sağlar.
* Varsayılan olarak atanır ve küme dışında erişim yoktur.

Örnek :

<service-name>.<namespace>.svc.cluster.local

backend-service.demo-app.svc.cluster.local

```bash
    kubectl get services
    kubectl port-forward svc/backend-service 8080:80
```

```yaml
apiVersion: v1
kind: Service
metadata:
  name: backend-service
  namespace: demo-app
spec:
  type: ClusterIP
  clusterIP: 10.96.0.250 # Statik IP (CIDR içinde olmalı)
  selector:
    app: backend
  ports:
    - protocol: TCP
      port: 80        # Servis portu (Cluster içinden erişim bu portla olur)
      targetPort: 80  # Pod içindeki container portu
```