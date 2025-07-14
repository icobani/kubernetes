## 📁 Dosya Olarak Enjekte Etme

Bu YAML, bir Secret verisini /etc/secrets klasörüne dosya olarak mount eder. Her key, o dizin altında bir dosya olarak
görünür.

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: my-app-pod
spec:
  containers:
    - name: my-app-container
      image: my-app-image:v1
      volumeMounts:
        - name: secret-volume
          mountPath: /etc/secrets
          readOnly: true
  volumes:
    - name: secret-volume
      secret:
        secretName: db-secret
```