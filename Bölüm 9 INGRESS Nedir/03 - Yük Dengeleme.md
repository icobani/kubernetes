# Yük Dengeleme

Ingress dışardan gelen istekleri içerdeki servislere yönlendirebiliyor. Servislerin arkasında çalışan podlar gelen
trafiğe göre yük paylaşımı görevini de yapar.

Ingress, gelen trafiği çalışan pod'lar arasında dengeler. Bu sayede daha fazla trafik alan pod'lar yerine diğer
pod'lar da trafiği alır ve uygulamanın ölçeklenebilirliği arttırılır.