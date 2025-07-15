# Varsayılan Namespace'ler

Büyük kubernetes kurulumları tamamlandığı zaman 4 tane default namespacelerin olduğunu görüyoruz.

Kubernetes'te, belirli bir Namaspace belirtilmediğinde, kaynaklar varsayılan olarak default Namaspace'ine yerleştirilir.
Kubernetes, aşağıdaki varsayılan Namespace'lere sahiptir:

1. *default*: Herhangi bir Namespace belirtilmediğinde kullanılan varsayılan Namespace'tir.
2. *kube-system*: Kubernetes kontrol düzlemi(kontrol plane) bileşenlerinin çalıştığı Namespace'dir.
3. *kube-public*: Tüm kullanıcılara açık olan, genellikle sistemin genel yapılandırmalarını içeren Namespace'dir.
4. *kube-node-lease*: Kubernetes node'larının durumlarını izlemek için kullanılan bir Namespace'tir.