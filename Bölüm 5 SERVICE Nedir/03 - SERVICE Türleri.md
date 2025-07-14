# Kubernetes Servis Türleri

Kubernetes'te servisler, pod'lar arasındaki iletişimi ve dış dünya ile bağlantıyı yöneten temel bileşenlerdir. Dört
farklı servis türü bulunmaktadır.

## Temel Servis Türleri

| Servis Türü            | Açıklama                                                                                                                                             |
|------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------|
| ClusterIP (Varsayılan) | Sadece küme içinde erişim sağlar                                                                                                                     |
| NodePort               | Node'un belirli bir portu üzerinden dış dünyaya açar 30.000 ile 32767 arasında bir port atanabiliyor. Geliştirme için hızlı bir test yapabilmek için |
| LoadBalancer           | Dış dünyaya erişim için bulut sağlayıcısının yük dengeleyicisini kullanır. Amazon, Azure, Google gibi                                                |
| ExternalName           | Harici bir DNS adına yönlendirme yapar. Herhangi bir noktaya erişmek için örneğin plaka sorgulama gibi. bunlara izin vermek lazım                    |
| Headler                | Veritabanı gibi durumlarda kullanılıyor. Statik erişim noktaları sağlıyor.                                                                           |

## Servis Türleri Karşılaştırması

| Özellik         | ClusterIP   | NodePort | LoadBalancer       | External Name  |
|-----------------|-------------|----------|--------------------|----------------|
| Dış Erişim      | ❌           | ✅        | ✅                  | ✅ (DNS ile)    |
| Load Balancing  | ✅           | ✅        | ✅                  | ❌              |
| DNS çözümleme   | ✅           | ✅        | ✅                  | ✅              |
| En yaygın ortam | Cluster içi | Test     | Production (Cloud) | Hybrid / Proxy |
