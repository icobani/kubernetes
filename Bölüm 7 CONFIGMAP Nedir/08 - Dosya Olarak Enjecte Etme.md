# Dosya Olarak Enjecte Etme

Bir Pod yaml dosyamız aşağıdaki şekilde

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: configmap-volume-pod
spec:
  containers:
    - name: app-container
      image: alpine  # Basit bir Alpine imajı
      command: [ "sleep", "3600" ]  # Pod'un çalışmaya devam etmesi için
      volumeMounts:
        - name: config-volume
          mountPath: /etc/files  # Bu dizine mount edilecek
          readOnly: true  # Dosyalar salt okunur olacak
  volumes:
    - name: config-volume
      configMap:
        name: my-config-files  # Burada kullanacağımız ConfigMap'in adı
```

config dizinimizin içinde

app.config ve db.config dosyalarımızın olduğunu farzedelim.

Bu komutla, `app.conf` ve `db.conf` dosyalarını `my-config-files` adında bir `ConfigMap` içine koyuyoruz.

```bash
    kubectl create configmap my-config-files \
    --from-file=app.conf=config/app.conf\
    --from-file=db.conf=config/db.conf`
```

Aşağıdaki şekilde iki dosyanın config dizini içinde olduğunu farzedelim.

```
config/
|___ app.conf
|___ db.cong
```
