package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adityarizkyramadhan/template-go-mvc/controller"
	"github.com/adityarizkyramadhan/template-go-mvc/database"
	"github.com/adityarizkyramadhan/template-go-mvc/docs"
	"github.com/adityarizkyramadhan/template-go-mvc/middleware"
	"github.com/adityarizkyramadhan/template-go-mvc/model"
	"github.com/adityarizkyramadhan/template-go-mvc/repositories"
	"github.com/adityarizkyramadhan/template-go-mvc/routes"
	"github.com/adityarizkyramadhan/template-go-mvc/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/natefinch/lumberjack"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Perkemahan Bedengan API
// @version         1.0
// @description     This is API documentation for Islamind. You can use the API with the following hosts:
// @description     - Production: api.perkemahanbedengan.com
// @contact.name    Aditya Rizky Ramadhan
// @contact.email   adityarizky1020@gmail.com
// @host            api.perkemahanbedengan.com
// @BasePath        /api/v1
// @Server 					api.perkemahanbedengan.com Production server
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Environment loaded")
	}

	db, err := database.NewDB()
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Database connected")
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.Perlengkapan{},
		&model.Ground{},
		&model.SubGround{},
		&model.Kavling{},
		&model.Reservasi{},
		&model.InvoiceReservasi{},
	)

	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Migration success")
	}
	redis := database.NewRedis()

	// check redis connection
	_, err = redis.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Redis connected")
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(&lumberjack.Logger{
		Filename:   "logs/request.log",
		MaxSize:    100,  // maksimal ukuran file dalam MB sebelum rotasi
		MaxAge:     4,    // maksimal hari untuk menyimpan file log lama
		MaxBackups: 7,    // maksimal jumlah file log backup
		Compress:   true, // mengkompres file log lama
	})

	router := gin.New(gin.OptionFunc(func(engine *gin.Engine) {
		engine.Use(gin.Recovery())
		engine.Use(middleware.ErrorHandler())
		engine.Use(middleware.LoggerRequest(logger))
		engine.Use(middleware.CORS())
		engine.Use(middleware.CheckToken(redis))
	}))

	router.Use(func(c *gin.Context) {
		log.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})

	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group("/api/v1")

	repoUser := repositories.NewUserRepository(db, redis)
	userController := controller.NewUserController(repoUser)
	userRoutes := routes.NewUserRoutes(userController)
	userRoutes.SetupRoutes(v1)

	repoPerlengkapan := repositories.NewPerlengkapanRepository(db)
	perlengkapanController := controller.NewPerlengkapanController(repoPerlengkapan)
	perlengkapanRoutes := routes.NewPerlengkapanRoutes(perlengkapanController)
	perlengkapanRoutes.SetupRoutes(v1)

	repoGround := repositories.NewGroundRepository(db)
	groundController := controller.NewGroundController(repoGround)
	groundRoutes := routes.NewGroundRoutes(groundController)
	groundRoutes.SetupRoutes(v1)

	repoSubGround := repositories.NewSubGroundRepository(db)
	subGroundController := controller.NewSubGroundController(repoSubGround)
	subGroundRoutes := routes.NewSubGroundRoutes(subGroundController)
	subGroundRoutes.SetupRoutes(v1)

	repoKavling := repositories.NewKavlingRepository(db)
	kavlingController := controller.NewKavlingController(repoKavling)
	kavlingRoutes := routes.NewKavlingRoutes(kavlingController)
	kavlingRoutes.SetupRoutes(v1)

	repoReservasi := repositories.NewInvoiceReservasiRepository(db)
	reservasiController := controller.NewInvoiceReservasiController(repoReservasi)
	reservasiRoutes := routes.NewInvoiceReservasiRoutes(reservasiController)
	reservasiRoutes.SetupRoutes(v1)

	// Serve static files from the "storage/public" directory
	router.Static("/api/v1/storage/public", "./storage/public")

	// Setup server with context for graceful shutdown
	srv := &http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: router,
	}

	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Println(err)
	}
	time.Local = location

	c := cron.New()

	c.AddFunc("0 12 * * *", func() {
		if err := utils.DailyCheckKavlingRawSQL(db); err != nil {
			log.Println(err)
		}
	})

	c.Start()

	// Start server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Printf("Server started on port %s\n", os.Getenv("PORT"))

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Use select to block the main goroutine and keep cron running
	select {
	case <-quit:
		log.Println("Shutting down server...")
	}

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")

}
