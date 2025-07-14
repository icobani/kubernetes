# ConfigMap'in Kullanım Seneryoları

1. Yapılandırma Dosyaları
    * ConfigMap, yapılandırma dosyalarını dışarıdan almak için kullanılır. Örneğin, bir JSON veya YAML dosyasını pod'un
      içine yerleştirebilirsiniz.
2. Ortamsal Değişkenler
    * Uygulama yapılandırmaları, ortam değişkenleri olarak ConfigMap üzerinden pod'lara aktarılabilir.
3. Komut Satırı Argümanları
    * Uygulamanızın başlatılması için gerekli olan parametreler veya komut satırı argümanları ConfigMap aracılığıyla
      sağlanabilir.