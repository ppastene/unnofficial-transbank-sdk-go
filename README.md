# Unnofficial SDK Transbank Go
SDK no oficial de Transbank Webpay hecho en el lenguage de programación Go

## Requisitos
- Go 1.19 o superior

## Instalación
Crea el archivo go.mod si no lo haz creado
``` go mod init <nombre-proyecto> ```

Descarga el modulo con go get
``` go get github.com/ppastene/unnofficial-transbank-sdk-go ```

Importa el modulo en los archivos donde se usará
```
import(
	"github.com/ppastene/unnofficial-transbank-sdk-go/src/common"
	"github.com/ppastene/unnofficial-transbank-sdk-go/src/webpayplus"
)

func main() {
	var options common.Options
	transaction := webpayplus.NewTransaction(options.ForIntegration(webpayplus.WEBPAY_PLUS_COMMERCE_CODE, webpayplus.WEBPAY_API_KEY))
}
```

## Uso
Crea una variable de tipo Options donde se debe setear el entorno para la comunicación con Transbank. Se provee dos metodos llamados ForIntegration() y ForProduction() que permite setear el commerceCode y la apiKey para el entorno de integración o producción (se recomienda que por el estado actual del modulo solo setear el entorno para integración)
``` var options common.Options = common.Options.ForIntegration(webpayplus.WEBPAY_PLUS_COMMERCE_CODE, webpayplus.WEBPAY_API_KEY) ```

Crea una variable de tipo Transaction usando el constructor NewTransaction() pasando como parametro las opciones antes definidas. Todas las consultas hacia el servicio de Transbank se harán a traves de la variable de tipo Transaction
``` transaction := webpayplus.NewTransaction(options) ```

Los metodos para consultar el servicio de Transbank cumplen con lo establecido en la documentación oficial. Por ejemplo si se desea crear una transaccion de Webpay Plus escriba lo siguiente y se imprimirá por consola el token y url dados por Transbank (se debe importar las librerias fmt y time):
```
t := time.Now().Local()
order := fmt.Sprintf("order%d%d%d%d%d%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
session := fmt.Sprintf("session%d%d%d%d%d%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
amount := 50000
returnUrl := "http://www.micomercio.cl/boleta"
response := transaction.Create(order, session, amount, "http://localhost:3000/resumen")
fmt.Println("token:", response.token, "url:", response.url)
```
## Metodos Webpay Plus
### Create
Se crea una transaccion
```
transaction.create(buy_order string, session_id string, amount float64, return_url string)
```
### Commit
Se confirma la transaccion
```
transaction.commit(token string)
```
### Refund
Se solicita devolucion parcial o total de la transaccion.
```
transaction.refund(token string, amount float64)
```
### Status
Muestra información de la transacción, disponible por 7 dias desde su creación
```
transaction.status(token string)
``` 
### Capture
Se realiza una captura diferida de la transaccion.
```
transaction.capture(token string, buy_order string, authorization_code string, capture_amount float64)
```

## TODO
- [x] Flujos Webpay Plus
- [ ] Flujos Webpay Plus Mall
- [ ] Flujos Webpay Modal
- [ ] Flujos OnePay
- [ ] Excepciones
- [ ] HTTPS y uso de certificados
- [ ] Pruebas Unitarias

## Documentación
Si bien este es un SDK no oficial toda la informacion relevante a la implementación del servicio de Webpay en este y otros lenguajes lo puedes encontrar en el sitio https://www.transbankdevelopers.cl.

La documentación relevante para usar este y otros SDK es:
- Primeros pasos con [Webpay](https://www.transbankdevelopers.cl/documentacion/webpay)
- Documentación sobre [ambientes, deberes del comercio, puesta en producción,
  etc](https://www.transbankdevelopers.cl/documentacion/como_empezar#ambientes).
  
También puedes encontrar: 
- Documentación general sobre los productos y sus diferencias:
  [Webpay](https://www.transbankdevelopers.cl/producto/webpay) y
  [Onepay](https://www.transbankdevelopers.cl/producto/onepay).
- Referencia detallada sobre [Webpay](https://www.transbankdevelopers.cl/referencia/webpay) y [Onepay](https://www.transbankdevelopers.cl/referencia/onepay).

Personalmente recomiendo una mirada a los siguientes repositorios
- [Documentación](https://github.com/TransbankDevelopers/transbank-developers-docs)
- [Credenciales SSL](https://github.com/TransbankDevelopers/transbank-webpay-credenciales)

## Disclaimer
A pesar de ser un SDK no oficial se hace todo lo posible para cumplir con la documentación oficial para un correcto uso e integración. Aun así este SDK no está terminado y falta por implementar varias funcionalidades para considerar un modulo bien hecho hacia el desarrollador, así que se recomienda NO USAR este modulo en entornos de producción.

Este proyecto está abierto a recibir issues y Pull Requests, como también estudiar el codigo y realizar sus propias integraciones
