package main

import (
	"log"
	"os"

	"github.com/andutruel/go_transacciones/cmd/server/handler"
	"github.com/andutruel/go_transacciones/docs"
	"github.com/andutruel/go_transacciones/internal/products"
	"github.com/andutruel/go_transacciones/internal/transactions"
	"github.com/andutruel/go_transacciones/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	//implementación de godotenv
	//el método Load carga el contenido del archivo .env en la variable de entorno
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error al intentar cargar el archivo .env")
	}

	db := store.New(store.FileType, "transacciones.json")

	repoTransaccion := transactions.NewRepository(db)
	serviceTransaccion := transactions.NewService(repoTransaccion)
	t := handler.NewTransaccion(serviceTransaccion)

	repoProducto := products.NewRepository()
	serviceProducto := products.NewService(repoProducto)
	p := handler.NewProduct(serviceProducto)

	r := gin.Default()

	//localhost:8080/docs/index.html
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pr := r.Group("/products")
	pr.GET("/", p.GetAll())
	pr.POST("/", p.Store())

	tr := r.Group("/transactions")
	tr.GET("/", t.GetAll())
	tr.POST("/", t.Store())
	tr.PUT("/:id", t.Update())

	r.Run()
}
