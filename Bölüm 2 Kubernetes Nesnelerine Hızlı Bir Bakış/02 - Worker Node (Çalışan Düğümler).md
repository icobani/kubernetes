# Worker Node (Çalışan Düğümler)
Uygulamalarınızı çalıştırılan sunuculardır. İçersinde şu bileşenler bulunur. Güvenlik için default olarak dışarı kapalıdır.

## Kubelet

- Her worker node'da çalışan, pod'ları yöneten ve API Server ile iletişim kuran aracı servistir.

## Container Runtime

- Konteynerları çalıştıran bileşenlerdir. Örnek olarak Docker, containerd, CRI-O kullanılır.

## Kube Proxy

- Worker node'lar arasındaki ağ iletişimini yöneten bileşenlerdir.