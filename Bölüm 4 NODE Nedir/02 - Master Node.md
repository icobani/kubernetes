# Master Node (Yönetici Düğüm) 🛠️

* Kubernetes kümesini yönetir ve denetler
* Küme içinde pod'ların dağıtımı, ölçeklenmesi ve yönetimini sağlar.
* Ana bileşenleri
    * API SERVER ==> Kubernetes API'sini yönetir. Bütün varolan kübernetes clusturundaki iletişim noktası
    * Scheduler ==> Pod'ları uygun worker node'lara yerleştirir.Podların hangi worker nodların üzerinde çalışması
      gerektiğine karar verir.
    * Controller Manager ==> Sistem durumunu kontrol eder (ReplicaSet, Node Health, vb.). Sistemin durumunu kontrol
      ediyor. Sağlık durumları, replica setleri, etcd veritabanı na erişen bileşen
    * etcd ==> Küme yapılandırmasını ve durumu saklayan veritabanıdır. Aslında bir veritabanı. Bütün bir kluster
      üzerindeki tüm durumların tutulduğu database'dedir. durumlar bu database üzerindedir. bunu kaybetmememiz
      gerekiyor. Korumamız gerekiyor. Backuplarının alınması gerekiyor.

📌 Master node, doğrurdan uygulamaları çalıştırmaz!