# Namespace Kullanım Senaryoları

* Çoklu Ortam Yönetimi : Geliştirme (dev), Test (test), Üretim (prod) gibi farklı ortamları izole etmek için Namespace
  kullanabilirsiniz. Her ortam kendi kaynaklarına sahip olacak şekilde yapılandırılabilir.
* Çoklu Takım ve Proje Yönetimi: Birden fazla takımın çalıştığı büyük organizasyonlarda, her takımın kendi Namespace'ini
  kullanarak kaynakları bibirinden ayırması sağlanabilir.
* Kaynak Yönetimi ve Kotası : Büyük uygulamalar için, belirli Namespace'lere kaynak kotaları ekleyerek her takımın adil
  bir kaynak kullanımı sağlaaması sağlanabilir.
* Erişim Denetimi : Farklı gruplara belirli Namespace'lere erişim izni vererek güvenlik sağlanabilir. Örneğin, sadece
  belirli bir takımın test Namespace'ine erişimi olabilir.