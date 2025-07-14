# Node'ların Çalışma Mantığı

1. Kullanıcı veya yönetici, Kubernetes API Server'a bir pod çalıştırma talebi gönderir. Bu talep etcd'ye kaydediliyor.
   Burda talebi verirken ram (şu kadar ram), disk (hızlı disk mesela) ve cpu detayı verilebilir.
2. Scheduler, uygun bir worker node seçer ve pod'u burada çalıştırır.
3. Kubelet, pod'un sağlıklı olup olmadığını kontrol eder.
4. Kube Proxy, ağ trafiğini yöneterek pod'ların birbirleriyle iletişim kurmasınız sağlar. verilen ip'ler her zaman
   dinamik ip. Statik erişim için servis üzerinden gideceğiz.
5. Node'un durumu bozulursa, Kubernetes bu pod'un başka bir node'a taşır. Eğer kaybolmaması gereken veri var ise mount
   edilecek olan disk alanı üzerinden bunu yönetmeliyiz. Ölmesi durumunda bu data kabolmasın diye