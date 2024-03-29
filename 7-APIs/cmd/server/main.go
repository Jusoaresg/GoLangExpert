package main

import (
	"log"
	"net/http"

	"github.com/Jusoaresg/GoLangExpert/7-APIs/configs"
	_ "github.com/Jusoaresg/GoLangExpert/7-APIs/docs"
	"github.com/Jusoaresg/GoLangExpert/7-APIs/internal/entity"
	"github.com/Jusoaresg/GoLangExpert/7-APIs/internal/infra/database"
	"github.com/Jusoaresg/GoLangExpert/7-APIs/internal/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @tittle Go Expert API Example
// @version 1.0
// @description Product API with authentication
// @termsOfService http://swagger.io/terms/

// @contact.name Juliano Gregorio
// @contact.url https://www.linkedin.com/in/juliano-gregorio/
// @contact.email julianosgreg@gmail.com

// @license.name Full Cycle License
// @license.url http://www.fullcycle.com.br

// @host localhost:8000
// @BasePath /
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("JwtExpiresIn", configs.JwtExperesIn))

	r.Route("/products", func(r chi.Router) {

		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)

	})

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_token", userHandler.GetJWT)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
