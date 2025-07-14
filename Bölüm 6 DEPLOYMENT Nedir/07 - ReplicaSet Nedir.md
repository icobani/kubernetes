# ReplicaSet Nedir

Deploymentların önceki versiyonlarında ReplicaSetler varmış. Manuel olarak set edilmesi ile ayağa kaldırılması ile
ilgili çalışıyor.

deployment ile ReplicaSet arasındaki fark yeni versiyona geçerken manuel olarak hepsini silip tekrar ayağa kaldırmak
yerine

bunu seçilen politikaya göre kendisi yapar. Deployment mutlaka otomatik bir ReplicaSet başına koymak için oluşturup
bırakır.

* ReplicaSet, Kubernetes'te bir uygulamanın belili sayıda kopyasının her zaman çalışmasını sağlamaya yönelik kullanılan
  bir kaynaktır.
* Pod'ları yönetir ve belirli bir sayıda pod'u çalıştırmak için tasarlanmıştır.
* ReplicaSet, uygulamanın yüksek erişilebilirliğini garanti ederve pod'lar arızalandığında otomatik olarak yeni pod'ları
  başlatır.