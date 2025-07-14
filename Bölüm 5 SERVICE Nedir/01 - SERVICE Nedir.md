# Kubernetes Service (Hizmet) Nedir ?

* Kubernetes'te Service (Hizmet), pod'lara ağ üzerinden erişim sağlayan mantıksal bir bileşendir. podlar ayağa
  kalktığında hangi ip ile ayağa kalktığını bilemeyiz. bunu statik bir erişim noktası sağlayan bir bileşendir.
* Pod'lar dinamik IP adreslerine sahip olduğundan, Service sabit bir erişim noktası sağlar.
* Podlar ölçeklendikçe (artıp azaldıkça) Service, otomatik olarak yeni pod'lara trafik yönlendirmeye devam eder.

Horizantal Pod Auto Scaler (HPA) ile pod sayıları artıp azalabilir.

