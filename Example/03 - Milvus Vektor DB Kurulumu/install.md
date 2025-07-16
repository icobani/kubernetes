# 📦 Milvus Kurulumu (Kubernetes Helm ile)

Bu doküman, Milvus vektör veritabanını on-premise Kubernetes ortamına Helm chart kullanarak hızlıca kurmak için
hazırlanmıştır.

⸻

## 🔧 Gereksinimler

* Kubernetes 1.18+
* Helm 3+
* 2 CPU / 4 GB RAM minimum (test için)

⸻

## 🚀 Adım 1: Helm Repo’yu Ekle

```bash
    helm repo add milvus https://milvus-io.github.io/milvus-helm/
    helm repo update
```

⸻

## 📂 Adım 2: Namespace Oluştur

```bash
  kubectl create namespace milvus
```

⸻

## ⚙️ Adım 3: Standalone (Tek Node) Milvus Kurulumu

Bu yapı; tek bir Milvus sunucusu, etcd, minio ve pulsar içeren hızlı bir kurulum sağlar:

```bash
    helm install my-milvus milvus/milvus \
      --namespace milvus \
      --set cluster.enabled=false \
      --set authentication.enabled=true \
      --set authentication.username=ici \
      --set authentication.password=Zp8mR2xL9vBt4qNh
    
    
```

Not: Bu kurulum “production” yerine test veya POC (Proof of Concept) için uygundur.

⸻

## ✅ Adım 4: Kurulumu Kontrol Et

```bash
  kubectl get pods -n milvus
```

açık olduğu portu da bu şekilde kontrol edebiliriz.

```bash
    kubectl get svc -n milvus my-milvus
```

Node_IP bu şekilde alabiliriz

```bash
  kubectl get nodes -o wide
```

Tüm pod’lar Running durumunda olmalı:
• my-milvus-milvus
• my-milvus-etcd
• my-milvus-minio
• my-milvus-pulsar

⸻

## 🌐 Adım 5 (Opsiyonel): Dış Erişim Açmak

Servisi dış dünyaya açmak için NodePort tipi kullanılabilir:

```bash
    helm upgrade my-milvus milvus/milvus \
      --namespace milvus \
      --set cluster.enabled=false \
      --set service.type=NodePort \
      --set authentication.enabled=true \
      --set authentication.username=ici \
      --set authentication.password=Zp8mR2xL9vBt4qNh
```

IP ve port bilgileri için:

```bash
    kubectl get svc -n milvus
```

## 🔌 Bağlantı (Go SDK ile)

Milvus resmi Go SDK:
[https://github.com/milvus-io/milvus-sdk-go](https://github.com/milvus-io/milvus-sdk-go)

```bash
  go get github.com/milvus-io/milvus-sdk-go/v2
```

## 📦 Dosya Konumu Önerisi

Bu install.md dosyasını projenizde infra/, docs/ veya k8s/ klasörlerine koyabilirsiniz.

## 📌 Ek Notlar

Daha büyük sistemler için --set cluster.enabled=true parametresi ile çok node’lu cluster kurulabilir.

Milvus, vektör indeksleme ve arama için IVF_FLAT, HNSW vb. destekler.

MinIO, etcd, pulsar varsayılan olarak içeriğe dahildir, harici bağlantı mümkündür.

## ✅ Adım 3: Milvus için Traefik IngressRoute Tanımı

[milvus-ingressroute.yaml](milvus-ingressroute.yaml) dosyası oluştur:

```Traefik gRPC’yi desteklemek için scheme: h2c kullanır.```

## 🔧 Ekstra: /etc/hosts Güncellemesi

IP adresi

```bash
    kubectl get nodes -o wide
``` 

çıktısındaki INTERNAL-IP olacak.

Erişim için local makinede /etc/hosts dosyasına şu satırı ekle:

```
  192.168.65.3 milvus.local
```

# 📦 Yaml’ı uygula

```bash
  kubectl apply -f milvus-ingressroute.yaml
```