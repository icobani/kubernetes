# Neden SERVICE Kullanılır ?

* Pod'ların dinamik IP değişikliklerini yönetir ==> Pod yeniden başlatıldığında IP'si değişebilir.
* Yük dengeleme (LB) sağlar ==> Gelen istekleri çalışan pod'lar arasında dengeler.
* İç ve dış erişimi yönetir ==> Cluster içindeki veya dışındaki uygulamaların pod'lara erişimini sağlar.

Podların önüne bir servis kuruyoruz ve dışardaki podlar bu servis üzerinden giderler. çünkü bir sürü pod olabilir.

Servisler bize statik ip sağlar. dns hizmeti ile de çalışması sağlanır. örneğin api.yakitmatik.com gibi.

En basit'i yük dengelemesini sağlayacaktır.

nodelar farklı şehirlerde de olabilir.

Podlarımız default olarak dışarı kapalıdır. Dışarıya açmak için kullanılır.

Dış dünyaya da kapalı bunu da izin dahilinde verebiliyoruz. 