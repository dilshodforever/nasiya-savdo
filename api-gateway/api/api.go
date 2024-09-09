package api

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/dilshodforever/nasiya-savdo/api/handler"
	//"github.com/dilshodforever/nasiya-savdo/api/middleware"
	_ "github.com/dilshodforever/nasiya-savdo/docs"

	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API Gateway
// @version 1.0
// @description Dilshod's API Gateway
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	// Middleware setup
	ca, err := casbin.NewEnforcer("config/model.conf", "config/policy.csv")
	if err != nil {
		panic(err)
	}

	err = ca.LoadPolicy()
	if err != nil {
		panic(err)
	}
	router := r.Group("/")
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	//router.Use(middleware.NewAuth(ca))

	// Swagger documentation
	url := ginSwagger.URL("/swagger/doc.json") // Adjusted path for Swagger docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler, url))

	// Contract endpoints
	contract := router.Group("/contract")
	{
		contract.POST("/create", h.CreateContract)
		contract.GET("/get/:id", h.GetContract)
		contract.PUT("/update", h.UpdateContract)
		contract.DELETE("/delete/:id", h.DeleteContract)
		contract.GET("/list", h.ListContracts)
		contract.GET("/getpdf/:id", h.GetContractPdf)
	}

	// Exchange endpoints
	exchange := router.Group("/exchange")
	{
		exchange.POST("/create", h.CreateExchange)
		exchange.GET("/get/:id", h.GetExchange)
		exchange.PUT("/update", h.UpdateExchange)
		exchange.DELETE("/delete/:id", h.DeleteExchange)
		exchange.GET("/list", h.ListExchanges)
	}

	// Product endpoints
	product := router.Group("/product")
	{
		product.POST("/create", h.CreateProduct)
		product.GET("/get/:id", h.GetProduct)
		product.PUT("/update", h.UpdateProduct)
		product.DELETE("/delete/:id", h.DeleteProduct)
		product.GET("/list", h.ListProducts)
	}

	// Storage endpoints
	storage := router.Group("/storage")
	{
		storage.POST("/create", h.CreateStorage)
		storage.GET("/get/:id", h.GetStorage)
		storage.PUT("/update", h.UpdateStorage)
		storage.DELETE("/delete/:id", h.DeleteStorage)
		storage.GET("/list", h.ListStorages)
	}

	// Transaction endpoints
	transaction := router.Group("/transaction")
	{
		transaction.POST("/create", h.CreateTransaction)
		transaction.GET("/get/:id", h.GetTransaction)
		transaction.PUT("/update", h.UpdateTransaction)
		transaction.DELETE("/delete/:id", h.DeleteTransaction)
		transaction.GET("/list", h.ListTransactions)
		transaction.POST("/check", h.CheckTransactions)
		transaction.POST("/test", h.TestNotification)
	}

	minIO := router.Group("/minio")
	{
		minIO.POST("/upload", h.UploadFile)
	}

	return r
}
