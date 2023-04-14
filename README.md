

# English Version - Version Ingles
## The best way to generate observability and monitoring for our cloud applications.

<img width="1169" alt="image" src="https://user-images.githubusercontent.com/5631542/232169532-035cfb25-59dd-46f6-99e4-d0ec8bfb5f55.png">


 [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome)
This application was presented at Nerdearla Argentina in 2019, so you can watch an explanatory video at the following link: https://www.youtube.com/watch?v=e-JhnuYfoyw.

The application consists of a frontend developed using React technology and a backend built with Go (also known as Golang). The combination of these two technologies was used to create a demo of the ATP tennis ranking, with the addition of the APM agent for observability from Elastic Search.

This project has been update in abril 2023.

Remeber this same project is available in [spanish](https://github.com/elastic-community/ATP-example-elastic-APM/blob/main/README-ES.md)

### Frontend 
In order to run the frontend locally, we need to navigate to the frontend folder and execute the following commands:

```bash
cd frontend
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

### Backend 

Run the backend locally, we need to navigate to the backend folder and execute the following commands:
```bash
go mod tidy
go mod download
go run cmd/main.go
```

To configure the backend with the APM of Elastic Search in our webservice with framework go gin , we need to modify the file cmd/main.go.
```go
import (
	"atp"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin/v2"
)

func main() {
	r := gin.Default()
	r.Use(apmgin.Middleware(r))
	r.GET("/api/ranking", atp.Ranking)
	r.Run("0.0.0.0:8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```

And for configurate observability on our gorm with our database we need change our driver in db/db.go.
```go
import (
	"context"
	"log"

	sqlite "go.elastic.co/apm/module/apmgormv2/v2/driver/sqlite"

	//"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)
func GetDB(ctx context.Context) *gorm.DB {
	db := DB
	db = db.WithContext(ctx)
	return db
}
```

Before running the program, you must configure the environment variables with the Elastic endpoint, token, and service name.
```bash
export ELASTIC_APM_SERVICE_NAME="backend"
export ELASTIC_APM_SERVER_URL="<Replace by apm url server>"
export ELASTIC_APM_SECRET_TOKEN="<Replace By Token>"
export ELASTIC_APM_EXIT_SPAN_MIN_DURATION=0ms
go run cmd/main.go
```

### Deploy

To run the code on your Kubernetes server, you can use the support manifest and instantiate it with the following command.

```console
kubectl apply -f k8s.yaml
```

Happy codding and enjoy elastic search in your project


## License

[![CC0](http://mirrors.creativecommons.org/presskit/buttons/88x31/svg/cc-zero.svg)](https://creativecommons.org/publicdomain/zero/1.0/)

Our projects are built with the mindset of open-source applications, using the MIT license.
