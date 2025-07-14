# ReplicaSet Temel Özellikleri

1. Pod Kopyalarını Yöneten
    * ReplicaSet, belirli sayıda pod'u çalıştırarak, pod'lar arasında yük dengelemesi sağlar.
2. Otomatik Pod Yeniden Başlatma
    * Eğer bir pod çöker veya arızalanırsa, ReplicaSet yeni bir pod başlatır.
3. Sürekli Ölçeklendirme
    * Pod sayısı istenilen seviyeye geldiğinde ReplicaSet, belirli sayıda pod'un çalışmasını sürekli olarak sağlar.
4. İstenmeyen Pod'ları Kaldırma
    * Gereksiz veya fazla pod'ları otomatik olarak siler, böylece istenen pod sayısı korunur.