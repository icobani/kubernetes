## Load Balancer (Bulut Servisleri için)

* Bulut servis sağlayucuları (AWS, GCP [Google Cloud Platform] Azure) için kullanılır.
* Otomatik olarak dış dünyaya bir yük dengeleyici oluşturur.

Örnek :

Dış IP adresi atandıktan sonra:

```bash
    curl http://<External-IP>
```

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-loadbalancer-service
spec:
  type: LoadBalancer
  selector:
    app: my-app
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8080  
```