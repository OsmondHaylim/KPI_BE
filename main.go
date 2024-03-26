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
	KpiAPIHandler		api.KpiAPI
	FileAPIHandler		api.FileAPI
	AnalisaAPIHandler	api.AnalisaAPI
}

func RunServer(db *gorm.DB, gin *gin.Engine) *gin.Engine {
	analisaService		:= service.NewAnalisaService(db)
	attendanceService 	:= service.NewAttendanceService(db)
	factorService 		:= service.NewFactorService(db)
	itemService 		:= service.NewItemService(db)
	masalahService 		:= service.NewMasalahService(db)
	minipapService 		:= service.NewMiniPAPService(db)
	monthlyService 		:= service.NewMonthlyService(db)
	resultService 		:= service.NewResultService(db)
	yearlyService 		:= service.NewYearlyService(db)
	fileService			:= service.NewFileService(db)

	kpiAPIHandler := api.NewKpiAPI(
		attendanceService, 
		factorService, 
		itemService,
		minipapService, 
		monthlyService, 
		resultService,
		yearlyService)
	fileAPIHandler := api.NewFileAPI(fileService)
	analisaAPIHandler := api.NewAnalisaAPI(
		analisaService,
		masalahService)
	apiHandler := APIHandler{
		KpiAPIHandler: kpiAPIHandler,
		FileAPIHandler: fileAPIHandler,
		AnalisaAPIHandler: analisaAPIHandler,
	}
	kpi := gin.Group("/kpi")
	{
		analisa := kpi.Group("/analisa")
		{
			analisa.GET("", apiHandler.AnalisaAPIHandler.GetAnalisaList)
			analisa.GET("/:id", apiHandler.AnalisaAPIHandler.GetAnalisaByID)
			analisa.POST("", apiHandler.AnalisaAPIHandler.AddAnalisa)
			analisa.PUT("/:id", apiHandler.AnalisaAPIHandler.UpdateAnalisa)
			analisa.DELETE("/:id", apiHandler.AnalisaAPIHandler.DeleteAnalisa)
		}
		attendance := kpi.Group("/attendance")
		{
			attendance.GET("", apiHandler.KpiAPIHandler.GetAttendanceList)
			attendance.GET("/:id", apiHandler.KpiAPIHandler.GetAttendanceByID)
			attendance.POST("", apiHandler.KpiAPIHandler.AddAttendance)
			attendance.PUT("/:id", apiHandler.KpiAPIHandler.UpdateAttendance)
			attendance.DELETE("/:id", apiHandler.KpiAPIHandler.DeleteAttendance)
		}
		factor := kpi.Group("/factor")
		{
			factor.GET("", apiHandler.KpiAPIHandler.GetFactorList)
			factor.GET("/:id", apiHandler.KpiAPIHandler.GetFactorByID)
			factor.POST("", apiHandler.KpiAPIHandler.AddFactor)
			factor.PUT("/:id", apiHandler.KpiAPIHandler.UpdateFactor)
			factor.DELETE("/:id", apiHandler.KpiAPIHandler.DeleteFactor)
		}
		item := kpi.Group("/item")
		{
			item.GET("", apiHandler.KpiAPIHandler.GetItemList)
			item.GET("/:id", apiHandler.KpiAPIHandler.GetItemByID)
			item.POST("", apiHandler.KpiAPIHandler.AddItem)
			item.PUT("/:id", apiHandler.KpiAPIHandler.UpdateItem)
			item.DELETE("/:id", apiHandler.KpiAPIHandler.DeleteItem)
		}
		masalah := kpi.Group("/masalah")
		{
			masalah.GET("", apiHandler.AnalisaAPIHandler.GetMasalahList)
			masalah.GET("/:id", apiHandler.AnalisaAPIHandler.GetMasalahByID)
			masalah.POST("", apiHandler.AnalisaAPIHandler.AddMasalah)
			masalah.PUT("/:id", apiHandler.AnalisaAPIHandler.UpdateMasalah)
			masalah.DELETE("/:id", apiHandler.AnalisaAPIHandler.DeleteMasalah)
		}
		minipap := kpi.Group("/minipap")
		{
			minipap.GET("", apiHandler.KpiAPIHandler.GetMinipapList)
			minipap.GET("/:id", apiHandler.KpiAPIHandler.GetMinipapByID)
			minipap.POST("", apiHandler.KpiAPIHandler.AddMinipap)
			minipap.PUT("/:id", apiHandler.KpiAPIHandler.UpdateMinipap)
			minipap.DELETE("/:id", apiHandler.KpiAPIHandler.DeleteMinipap)
		}
		monthly := kpi.Group("/monthly")
		{
			monthly.GET("", apiHandler.KpiAPIHandler.GetMonthlyList)
			monthly.GET("/:id", apiHandler.KpiAPIHandler.GetMonthlyByID)
			monthly.POST("", apiHandler.KpiAPIHandler.AddMonthly)
			monthly.PUT("/:id", apiHandler.KpiAPIHandler.UpdateMonthly)
			monthly.DELETE("/:id", apiHandler.KpiAPIHandler.DeleteMonthly)
		}
		
		result := kpi.Group("/result")
		{
			result.GET("", apiHandler.KpiAPIHandler.GetResultList)
			result.GET("/:id", apiHandler.KpiAPIHandler.GetResultByID)
			result.POST("", apiHandler.KpiAPIHandler.AddResult)
			result.PUT("/:id", apiHandler.KpiAPIHandler.UpdateResult)
			result.DELETE("/:id", apiHandler.KpiAPIHandler.DeleteResult)
		}
		yearly := kpi.Group("/yearly")
		{
			yearly.GET("", apiHandler.KpiAPIHandler.GetYearlyList)
			yearly.GET("/:id", apiHandler.KpiAPIHandler.GetYearlyByID)
			yearly.POST("", apiHandler.KpiAPIHandler.AddYearly)
			yearly.PUT("/:id", apiHandler.KpiAPIHandler.UpdateYearly)
			yearly.DELETE("/:id", apiHandler.KpiAPIHandler.DeleteYearly)
		}
		file := kpi.Group("/file")
		{
			file.POST("/file", apiHandler.FileAPIHandler.FileUpload)
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

		conn.AutoMigrate(&model.Monthly{}, &model.Masalah{}, &model.Project{}) 
		conn.AutoMigrate(&model.MiniPAP{}, &model.Analisa{}, &model.Summary{})
		conn.AutoMigrate(&model.Attendance{})
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