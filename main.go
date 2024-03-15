package main

import (
	"goreact/db"
	"goreact/model"
	"goreact/api"
	"goreact/service"

	_ "embed"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type APIHandler struct{
	KpiAPIHandler	api.KpiAPI

}

func RunServer(db *gorm.DB, gin *gin.Engine) *gin.Engine {
	attendanceService := service.NewAttendanceService(db)
	factorService := service.NewFactorService(db)
	monthlyService := service.NewMonthlyService(db)
	minipapService := service.NewMiniPAPService(db)
	resultService := service.NewResultService(db)

	kpiAPIHandler := api.NewKpiAPI(
		attendanceService, 
		factorService, 
		monthlyService, 
		minipapService, 
		resultService)
	apiHandler := APIHandler{
		KpiAPIHandler: kpiAPIHandler,
	}
	kpi := gin.Group("/kpi")
	{
		factor := kpi.Group("/factor")
		{
			factor.GET("", apiHandler.KpiAPIHandler.GetFactorList)
			factor.GET("/:id", apiHandler.KpiAPIHandler.GetFactorByID)
			factor.POST("", apiHandler.KpiAPIHandler.AddFactor)
			factor.PUT("/:id", apiHandler.KpiAPIHandler.UpdateFactor)
			factor.DELETE("/:id", apiHandler.KpiAPIHandler.DeleteFactor)
		}
		monthly := kpi.Group("/monthly")
		{
			monthly.GET("", apiHandler.KpiAPIHandler.GetMonthlyList)
			monthly.GET("/:id", apiHandler.KpiAPIHandler.GetMonthlyByID)
			monthly.POST("", apiHandler.KpiAPIHandler.AddMonthly)
			monthly.PUT("/:id", apiHandler.KpiAPIHandler.UpdateMonthly)
			monthly.DELETE("/:id", apiHandler.KpiAPIHandler.DeleteMonthly)
		}
	}
	return gin
}

func main(){
	gin.SetMode(gin.ReleaseMode)

	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Missing .env file. Probably okay on dockerized environment")
	}
	config := &db.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		router := gin.New()
		db := db.NewDB()
		router.Use(gin.Recovery())
		router.ForwardedByClientIP = true
		router.SetTrustedProxies([]string{"127.0.0.1"})

		conn, err := db.Connect(config)
		if err != nil {
			panic(err)
		}

		conn.AutoMigrate(&model.Monthly{}) 
		conn.AutoMigrate(&model.MiniPAP{})
		conn.AutoMigrate(&model.Attendance{})
		conn.AutoMigrate(&model.PAP{})
		conn.AutoMigrate(&model.Factor{})
		conn.AutoMigrate(&model.Result{})
		conn.AutoMigrate(&model.Item{})
		conn.AutoMigrate(&model.Yearly{})

		router = RunServer(conn, router)

		fmt.Println("Server is running on port 8080")
		err = router.Run(":8080")
		if err != nil {
			panic(err)
		}

	}()

	wg.Wait()
}