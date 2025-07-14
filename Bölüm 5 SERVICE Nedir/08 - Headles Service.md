# Headless Service

* Kubernetes dışındaki bir servise DNS üzerinden yönlendirme yapar.
* Harici bir servisi cluster içinde kullanmak için kullanılır.

Veritabanı gibi sabit bir ip'ye sahip bir yapı. Cluster ip aynı.

Örnek :

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-headless-service
spec:
  clusterIP: None
  selector:
    app: my-app
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8080  
```

```bash
    nslookup my-headless-service.default.svc.cluster.local
```