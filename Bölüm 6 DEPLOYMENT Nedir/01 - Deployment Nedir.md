# Deployment Nedir ?

* Kubernetes Deployment, pod'ların yaşam döngüsünü yöneten ve uygulamanın istenilen durumunu (desired state) sağlamak
  için kullanılan bir Kubernetes kaynağıdır.
* Deployment, uygulamanızın doğru sürümünü çalıştırmak, ölçeklendirmek, güncellemek ve gerektiğinde geri almak için
  kullanılır.
* Pod'ların dağıtımı, güncellenmesi ve yönetilmesi için genellikle ReplicaSet ile birlikte çalışır.

Kubernetes'in en derin nesnelerinden birisidir.

Önceden replica setler vardı üstüne deployment geldi.

replicaset mevcut olarak kullanılıyor. ama daha gelişmiş olarak deployment devam ediyor.

deployment 'da replicas = 5 dediğimizde bunu oluşturuyor ve ReplicaSet'i buna bekçi olarak koyuyor. Dolayısıyla sonradan
bunların yönetimini replicaSet'e devrediyor

sürekli olarak 5 pod'un