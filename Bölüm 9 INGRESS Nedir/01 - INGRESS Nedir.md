# INGRESS Nedir ?

Servisi dışarı açmak ve load balancer yapısını devreye giriyor. Bizim servisimizi dışardaki bir adrese yönlendirme işi.

Ingress, Kubernetes'te dış dünyadan gelen HTTP/HTTPS trafiğini pod'lara yönlendiren bir bileşendir. Ingress, daha geniş
anlamda yük dengeleme, SSL sonlandırma, URL yönlendirme gibi işlevleri yerine getirir.

Ingress, genellikle dışa açık bir API veya web uygulaması sağlayan servislerin önünde bir kontrol noktası olarak
çalışır. ve gelen istekleri doğru pod'lara yönlendirir.