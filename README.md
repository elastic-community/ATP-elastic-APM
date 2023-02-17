

# English Version - Version Ingles
## The best way to generate observability and monitoring for our cloud applications.
 [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome)
This application was presented at Nerdearla Argentina in 2019, so you can watch an explanatory video at the following link: https://www.youtube.com/watch?v=e-JhnuYfoyw.

The application consists of a frontend developed using React technology and a backend built with Go (also known as Golang). The combination of these two technologies was used to create a demo of the ATP tennis ranking, with the addition of the APM agent for observability from Elastic Search.

Here's how to run both projects using containers.

```bash
docker compose -up
```

### Frontend 
In order to run the frontend locally, we need to navigate to the frontend folder and execute the following commands:

```bash
npm install 
npm run dev
```

To configure the frontend with the URL of Elastic Search, we need to modify the file src/apm.js#L9.

```js
const apm = initApm({
  // Set required service name (allowed characters: a-z, A-Z, 0-9, -, _, and space)
  serviceName: 'frontend',
  // Set custom APM Server URL (default: http://localhost:8200)
  serverUrl: '<Replace here by endpoint>,
  // Set service version (required for sourcemap feature)
  serviceVersion: '0.1'
})
```

```js
const apm = initApm({
  // Set required service name (allowed characters: a-z, A-Z, 0-9, -, _, and space)
  serviceName: 'frontend',
  // Set custom APM Server URL (default: http://localhost:8200)
  serverUrl: '<Replace here by endpoint>,
  // Set service version (required for sourcemap feature)
  serviceVersion: '0.1'
})
```

### Backend 

Run the backend locally, we need to navigate to the backend folder and execute the following commands:
```bash
go mod tidy
go mod download
go run cmd/main.go
```

To configure the backend with the APM of Elastic Search, we need to modify the file cmd/main.go.
```go
import (
	"atp"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
)

func main() {
	r := gin.Default()
	r.Use(apmgin.Middleware(r))
	r.GET("/api/ranking", atp.Ranking)
	r.Run("0.0.0.0:8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```
Before running the program, you must configure the environment variables with the Elastic endpoint, token, and service name.
```bash
export ELASTIC_APM_SERVICE_NAME="backend-k8s"
export ELASTIC_APM_SERVER_URL="<Replace by apm url server>"
export ELASTIC_APM_SECRET_TOKEN="<Replace By Token>"

go run cmd/main.go
```

###Deploy

To run the code on your Kubernetes server, you can use the support manifest and instantiate it with the following command.

```console
kubectl apply -f k8s.yaml
```


Happy codding and enjoy elastic search in your project

### Update

The code was updated in February of 2023.


# Version Español - Spanish Version
## La mejor forma de generar observabilidad y monitoreo a nuestras aplicaciones en la nube [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome)

Esta aplicación fue presentada en nerdearla argentina el año 2019, por lo cual puedes ver un video explicativo en el siguiente link https://www.youtube.com/watch?v=e-JhnuYfoyw 

Esta aplicación se compone de un frontend desarrollado con la tecnología React y un backend construido con Go (también conocido como Golang). La combinación de estas dos tecnologías para mostrar un demo del ranking del tennis ATP agregando el agente apm de  observabilidad de elastic search.

Como levantar ambos proyectos con contendores

```bash
docker compose -up
```

### Frontend 
Para poder ejecutar el frontend de manera local, debemos situarnos en la carpeta correspondiente al frontend y ejecutar los siguientes comandos:
```bash
npm install 
npm run dev
```

Para configurar el frontend con el rum de elastic search debemos modificar el archivo src/apm.js#L9

```js
const apm = initApm({
  // Set required service name (allowed characters: a-z, A-Z, 0-9, -, _, and space)
  serviceName: 'frontend',
  // Set custom APM Server URL (default: http://localhost:8200)
  serverUrl: '<Replace here by endpoint>,
  // Set service version (required for sourcemap feature)
  serviceVersion: '0.1'
})
```

### Backend 

Para poder ejecutar el backend de manera local, debemos situarnos en la carpeta correspondiente al backend y ejecutar los siguientes comandos:
```bash
go mod tidy
go mod download
go run cmd/main.go
```

Para configurar el backend con el apm de elastic search en el archivo cmd/main.go

```go
import (
	"atp"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
)

func main() {
	r := gin.Default()
	r.Use(apmgin.Middleware(r))
	r.GET("/api/ranking", atp.Ranking)
	r.Run("0.0.0.0:8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```
antes de ejecutar el programa debes configurar las variables de ambientes con el endpoint token y service name de elastic
```bash
export ELASTIC_APM_SERVICE_NAME="backend-k8s"
export ELASTIC_APM_SERVER_URL="<Replace by apm url server>"
export ELASTIC_APM_SECRET_TOKEN="<Replace By Token>"

go run cmd/main.go
```

Espero que esto te sirva para incorporar elastic search en tu propio proyecto

###Deploy

Para ejecutar el codigo en tu servidor de kubernetes puedes usar el manifiesto de apoyo y instanciarlo con el siguiente comando 
```console
kubectl apply -f k8s.yaml
```


### Update

El codigo actualizado en febrero del año 2023

## License

[![CC0](http://mirrors.creativecommons.org/presskit/buttons/88x31/svg/cc-zero.svg)](https://creativecommons.org/publicdomain/zero/1.0/)

Our projects are built with the mindset of open-source applications, using the MIT license.
