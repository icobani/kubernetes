# Master Node (Kontrol Düzlemi)
Kubernetes kümesini yöneten kontrol merkezidir. İçerisinde şu bileşenler bulunur.


## API Server

- Kubernetes ile iletişim kurulan merkezi bileşendir.
- Kullanıcılar, komutlar (kubectl, API) aracılığıyla API Server üzerinden küme yönetimi yapar.

## Scheduler (Zamanlayıcı)

- Yeni oluşturulan pod'ları uygun worker node'a yerleştirir.

## Controller Manage

- Pod'ların istenen durumda olup olmadığını denetler ve gerektiğinde otomatik aksiyonlar alır.

## etcd (Veri Depolama)

- Kubernetes'in tüm konfigurasyonlarını ve durum bilgilerini saklayan dağıtılmış anahtar-değer veritabanıdır.
