# Node Nedir ?

* Node, Kubernetes kümesindeki çalışan makineleri ifade eder.
* Her node, pod'ları barındıran fiziksel veya sanal bir sunucudur.
* Kubernetes kümesi (cluster) birden fazla node'dan oluşur.
* İki tür node vardır:
    * Master Node (Yönetim Düğümü) Bizim cluster'ımızı yöneten beyin.
    * Worker Node (Çalışan Düğüm) Fiziksel bir sunucu da olabilir sanal sunucu da olabilir.

### Kubelet

Podların yaşam döngülerini yöneten bileşen. Sahlık durumlarını yöneten bileşen. Worker nodelar üzerindeki podları
yöneten bileşen. Çalıştırılması veya öldürülmesi.

### KubeProxy

Bir node ayağa kalktığı zaman kendisini master node'a register ediyor bana node gönderebilirsin diye. Master node ile
worker node arasındaki
İletişim kubeproxy üzerinden sağlanıyor olacak.

### Container Runtime

Bunlar ilgili container'ın container.d olabilir yada docker container olabilir docker dosyalarını çalıştıran
bileşenlerdir.

Bir Cluster'da en az bir tane master node bir tane worker node olmak zorunda. Worker node çok sayıda olabilir. bir
master node'a
5000 adet worker node olabilir. Yani 5000 tane fiziksel makina gibi düşünebilirsiniz.

Toplam pos sayısı 110.000 pod çalışabilir.

slave master node la olabilir.

etcd veritabenı bunların hepsini tutacaktır.

nodelar ölçeklendirilebiliyor.
horizantel auto sclae olabildiği gibi cluster autosclaer'da olabilir. nodelar bellek kullanımlarını optimize etmee
çalışacaktır.

1 node sanal sanal makinayı ifade edebilir.

