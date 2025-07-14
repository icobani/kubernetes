# ConfigMap Yaml Analizi

Örnek ConfigMap Yaml Dosyası

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