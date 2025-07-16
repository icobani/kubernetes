# Senaryo

Kendi local Kubernetes ortamımızda 3 replikalı (1 primary, 2 read-only replica) bir PostgreSQL kurulumu yapacağız.  
Dataların kalıcı olması için StatefulSet ve volumeClaimTemplates kullanarak PVC'ler otomatik oluşturulacak.

---

## 🔧 Adım 1: Namespace Oluştur

Öncelikle `postgres-cluster` isminde bir namespace oluşturuyoruz:

```bash
  kubectl create namespace postgres-cluster
```

---

## 🐘 Adım 2: PostgreSQL StatefulSet ile Kurulumu

Bitnami PostgreSQL imajı ile StatefulSet yapısını kullanarak 1 primary ve 2 read-only replica’dan oluşan bir küme
kuracağız.  
PVC'ler `volumeClaimTemplates` ile her pod için otomatik olarak oluşturulacaktır.

[pg-statefulset.yaml](pg-statefulset.yaml):

```bash
  kubectl apply -f pg-statefulset.yaml
```

---

## 🌐 Adım 3: Headless Service Tanımı

StatefulSet pod'larının birbirini bulabilmesi için headless service tanımı yapılır:

[pg-service.yaml](pg-service.yaml)

```bash
  kubectl apply -f pg-service.yaml
```

---

## Adım 4: Kendi makinamdan bu makinaya bağlanmam için nodeport ile dışarı açtım

---

## ✅ Sonuç

- `postgres-0`: Primary (master)
- `postgres-1` ve `postgres-2`: Replica (read-only)
- `data-postgres-0`, `data-postgres-1`, `data-postgres-2`: Otomatik oluşturulmuş kalıcı PVC'ler

```bash
    kubectl get pods -n postgres-cluster
    kubectl get pvc -n postgres-cluster
```

Kurulum tamamlandıktan sonra servis üzerinden veritabanınıza erişebilir ve replikasyon durumunu doğrulayabilirsiniz.