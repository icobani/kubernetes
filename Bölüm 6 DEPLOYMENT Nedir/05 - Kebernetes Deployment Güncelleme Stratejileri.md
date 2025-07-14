# Kebernetes Deployment Güncelleme Stratejileri

* Rolling Update
    * Kademeli güncelleme
    * Parametre olarak aynı anda kaçar node'u güncelleyebilirsin.MaxResource
    * Bir de Max UnAvailable: aynı anda kaç nod'u devredışı bırakabilirsin.
    * yavaş yavaş devreye alarak trafik almaya devam etmesini sağlıyor.
    * Avantajı sistemin sürekli çalışır durumda kalmasını sağlıyor. Kullanıcılar kesinti yaşamıyorlar.
    * Podlar adım adım güncellenir
    * Eski ve yeni pod'lar aynı anda çalışmaz, böylece uygulamanın kesintisiz çalışması sağlanır.
    * Güncelleme süreci yönetilebilir ve ayarlanabilir
* Recreate (Yeniden oluşturulma)
    * Bir kerede tüm node'lar kapatılıyor. Sonrasında yeni versiyon olarak ayağa kaldırılıyor.
    * Örneğin veritabanında mandatory field'lar var ve yapacak bişi yok bunu kullanıyoruz.
    * Sade ve doğrudan bir strateji
    * Hepsini yayından kaldırıyorsunuz. Sonra yayına alıyorsunuz.
    * Bir miktar kesinti hissedilebilir.
    * Tüm pod'lar bir anda silinir ve ardından yeni pod'lar başlatılır.
    * Bu yöntem genellikle kesinti gerektiren durumlar için kullanılır.
* Blue/Green Deployment
    * İki farklı ortam kuruluyor.
    * Eski ve yeni versiyon gibi
    * Tamamını bir defada yayına almıyoruz kademeli bir yayına alıyoruz.
    * Blue Eski versiyon
    * Green Yeni versiyon
    * Kullanıcıların bir kısmının green versiyonda denenmesi için kullanılabiliyor.
* Canary Deployment
    * Kanarya diye de geçiyor.
    * Küçük bir kullanıcı kısmını örneğin %5'ini yeni versiyona açıp. Güvenli bir geçiş yapmak için kullanılıyor.
* Progressive Delivery (Argo Rollouts, Flagger, vb.)
    * Canary Deployment'in otoamatik olarak yavaş yavaş deployment olmasını sağlıyor.
    * Kademe kademe küçük grupları geçirerek geri bildirimleri ölçmek için kullanılıyor.
    * Örneğin rest apinin versiyon 2'den 3'e geçilmesi gibi.
    * Exception'a düşmüyorsa yavaş yavaş kademeli geçiş yapıyor.