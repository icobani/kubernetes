- Pod, Kubernetes'in en küçük ve en temel dağıtım birimidir.
- Bir veya daha fazla konteyneri (genellikle Docker kontainerleri) içinde barındırır ve bu kontenerler ortak bir ağ ve depolama alanı paylaşır.
- Yani şöyle düşün bir pod açtın ve içine birden çok container'ı yapılandırdın bu durumda bu pod içindeki container'lar bir birlerine localhost şeklinde ulaşabilirler. sanki laptop'umda çalışıyormuş gibi.
- Kubernetes, uygulamaları pod'lar halinde çalıştırır ve yönetir.


## POD'un Temel Özellikleri

- Tek veya Çoklu Kontainer içerebilir
    - Bir pod, tek bir konteyner çalıştırılabilir (yaygın kullanım)
    - Alternatif olarak, birden fazla konteyneri de içerebilir.
- Ortak Ağ Paylaşımı
    - Pod içindeki tüm konteynerler aynı IP adresini paylaşır.
    - Konteynerler localhost üzerinden birbirleriyle iletişim kurabilir
- Ortak Depolama Kullanımı
    - Pod içindeki konteynerler, aynı volume (depoloma alanını) paylaşabilir.
    - Örneğin, bir konteyner veriyi yazarken, diğeri bu veriyi okuyabilir.
- Bağımsız Olmayan Birimdir.
    - Tek başına çalıştırılmaz, Kubernetes tarafından yönetilir.
    - Pod'lar geçici (ephemeral) birimlerdir, yani sistem tarafından bozulduğunda otomatik olarak yeniden oluşturulabilir.
    - Bu yüzden kalıcı bir disk alanını mutlaka dışarı volume olarak mount edin. Kalıcı dataları buradan yönetin.
    - Pod değersiz küt diye siliniyor.

## POD Kullanım Seneryoları

1. Tek konteynerli Pod (En yaygın kullanım)
    1. Her pod tek bir uygulama konteyner'i çalıştırır
    2. Örneğin, bir Nginx web sunucusunu çalıştıran pod.
2. Yan Araç (Sidecar) Konteynerli Pod
    1. Bir pod'un içinde ana uygulama + yardımcı bir konteyner bulunur.
    2. Örneğin:
        1. Ana konteyner: Uygulama (örn: Node.js)
        2. Yan konteyner: Log toplama veya veri senkranizasyonu
3. Bağımlı Konteynerler Çalıştıran Pod
    1. Bir pod'un içinde birden fazla uygulama çalışabilir ve birbirleriyle etkileşime girebilir.
    2. Örneğin:
        1. Ana konteyner: PostgreSQL veritabanı
        2. Yan konteyner: Veritabanı izleme (Monitoring uygulaması)