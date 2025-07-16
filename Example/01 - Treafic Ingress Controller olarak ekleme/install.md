# ✅ Adım 1: Traefik Helm repo’sunu ekle ve update et

```bash
    helm repo add traefik https://traefik.github.io/charts
    helm repo update
```

# ✅ Adım 2: Traefik Ingress Controller’ı kur

```bash
    kubectl create namespace traefik
```

Sonra Traefik’i Helm ile kur:

```bash
    helm install traefik traefik/traefik \
      --namespace traefik \
      --set service.type=NodePort \
      --set ports.web.nodePort=32080 \
      --set ports.traefik.nodePort=32090 \
      --set dashboards.enabled=true \
      --set api.dashboard=true \
      --set dashboards.domain=traefik.local
```

```--set service.type=NodePort diyerek dış dünyadan da erişilebilir hale getiriyoruz.```

[dashboard-ingressroute.yaml](dashboard-ingressroute.yaml) dosyasını oluştur:

ve uygula

```bash
  kubectl apply -f dashboard-ingressroute.yaml
```

## Port-Forward

```bash
    kubectl port-forward -n traefik svc/traefik 8080:80
```