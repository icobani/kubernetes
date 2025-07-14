# Secret ve ConfigMap Arasındaki Farklar

| Özellik    | Secret                                         | ConfigMap                                      |
|------------|------------------------------------------------|------------------------------------------------|
| Amaç       | Gizli verileri güvenli şekilde saklamak        | Uygulama yapılandırmalarını dışarıdan sağlamak |
| Veri Şekli | Şifreli veriler (parolalar, API anahtarları)   | Genel yapılandırma verileri (JSON, YAML)       |
| Güvenlik   | Şifreli depolama                               | Şifreli değil, düz metin                       |
| Kullanım   | Parola, API anahtarı, token gibi gizli veriler | Yapılandırma bilgileri, ortam değişkenleri     |

