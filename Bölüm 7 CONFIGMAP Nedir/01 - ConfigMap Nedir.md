# ConfigMap Nedir ?

Kubernetes'in önemli nesnelerinden biridir. Örneğin connection stringler.

System Environments içine bu configleri kullanabiliriz.

Bunları configmap yada secret içinde tutuyoruz.

ConfigMap açık bir şekilde environments 'e yazıyor. Secret ise güvenli bir şekilde saklamak için kullanılıyor.

* ConfigMap, Kubernetes'teki yapılandırma verilerini bir veya daha fazla pod'a sağlamak için kullanılan bir kaynak
  türüdür.
* ConfigMap, uygulamanızın çalışma zamanında değişken yapılandırmalarını dışarıdan almasına ve pod'lara enjecte etmesine
  olanak tanır.