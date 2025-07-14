# Ortam Değişkeni Olarak Enjecte Etme

Aşağıda bir pod yaml dosyasında ilgili configMap tanımını enjecte etme örneğini görüyorsunuz.

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: my-app-pod
spec:
  containers:
    - name: my-app-container
      image: my-app-image:v1
      envFrom:
        - configMapRef:
            name: app-config
```

ConfigMap Yaml Tanımı

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-config
data:
  APP_ENV: "production"
  APP_PORT: "8080"
  DATABASE_URL: "postgres://user:password@db.example.com:5432/mydb"
```