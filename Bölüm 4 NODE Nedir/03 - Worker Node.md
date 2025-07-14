# Worker Node (Çalışan Düğün)

Kübernetes'de iki farklı node vardır. Master node ve Worker node. Master node 1 tane bulunurken bir cluster içinde
Worker node
5000 adede kadar node olabilir. 110.000 adede kadar da pod çalışabilir

* Pod'ları çalıştıran aslın makineler bunlardır. içinde birden fazla container olabilir
* Podların çalışmasınnı ve yönetimini sağlayan bileşenler içerir:
    * Kubelet ==> Node'un Kubernetes ile iletişimini sağlar.
    * Container Runtime (Docker, containerd, CRI-O, vb.) ==> Kontenerleri çalıştırır
    * Kube Proxy ==> Node içindeki ağ trafiğini yönetir ve yönlendirir.

📌 Worker node'lar, Kubernetes'in ana işlem gücüdür! burada pod'lar çalışır ve ölçeklendirilir.