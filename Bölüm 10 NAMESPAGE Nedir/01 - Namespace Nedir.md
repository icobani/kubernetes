# Namespace Nedir ?

Namespace, Kubernetes'teki kaynakları gruplamak ve yönetmek için kullanılan bir mantıksal ayrımdır.
Kubernetes üzerinde birden fazla iş yükü (workload) veya ortam (örneğin, geliştirme, test, üretim) çalıştırılabilir ve
bu
kaynaklar Namespace aracılığıyla izole edilebilir. Böylece, aynı kaynak adı altında farklı projeler veya ortamlar
sorunsuz
bir şekilde çalışabilir.

Namespace, özellikle büyük ölçekli Kubernetes kümelerinde kaynak yönetimi ve erişim denetimi gibi işlemleri daha verimli
hale getirir. Aynı kaynak adı altında birden fazla ortam veya uygulama çalıştırmak gerektiğinde, Namespace her birini
izole eder ve karışıklığı engeller.