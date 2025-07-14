# ExternalName

* Kubernetes dışındaki bir servise DNS üzerinden yönlendirme yapar.
* Harici bir servisi cluster içinde kullanmak için kullanılır.

Örnek :

kubernetes içindeki servislerin internet üzerindeki dışardaki bir servise erişim ihtiyacı varsa bu şekilde tanıtmak
gerekiyor.

Alias tanımlama gibi. Yeni bir dns tabanlı yönlendirme yapılıyor.

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-external-service
spec:
  type: ExternalName
  externalName: example.com
```