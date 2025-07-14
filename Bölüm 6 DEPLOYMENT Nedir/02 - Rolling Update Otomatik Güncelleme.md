# Otomatik Güncellemeler (Rolling Updates)

Özellikle yayına alma politikaları konusunda çok güçlü. Bu politikaları ele alacağız.

Bunlardan bir tanesi Rolling Updates

* Yeni bir sürüm yayımlandığında, eski sürümün yerine yeni sürümü adım adım günceller.
* Pod'ları kesintisiz bir şekilde güncellenir.

yaml dosyasında örneğin tüm test süreçlerinden sonra
1.1'e geç dediğimiz zaman rolling update devreye giriyor. Rolling update'de 5 replikanın hepsini öldürüp aynı anda
yayına almak yanlış bir turum olur.

Recreate 'de mümkün buna tarif için

sıra ile öldürüp sırasıyla ayağa kaldıracak.