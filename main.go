package main

import (
	"goreact/api"
	"goreact/db"
	"goreact/model"
	"goreact/repository"
	"goreact/service"
	"goreact/middleware"

	_ "embed"
	"fmt"
	"sync"
	// "log"
	// "os"

	// "github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type APIHandler struct{
	KpiAPIHandler		api.KpiAPI
	FileAPIHandler		api.FileAPI
	AnalisaAPIHandler	api.AnalisaAPI
	ProjectAPIHandler 	api.ProjectAPI
	UserAPIHandler		api.UserAPI
}

func RunServer(db *gorm.DB, gin *gin.Engine) *gin.Engine {
	analisaRepo		:= repository.NewAnalisaRepo(db)
	attendanceRepo 	:= repository.NewAttendanceRepo(db)
	factorRepo 		:= repository.NewFactorRepo(db)
	itemRepo 		:= repository.NewItemRepo(db)
	masalahRepo 	:= repository.NewMasalahRepo(db)
	minipapRepo 	:= repository.NewMiniPAPRepo(db)
	monthlyRepo 	:= repository.NewMonthlyRepo(db)
	resultRepo 		:= repository.NewResultRepo(db)
	yearlyRepo 		:= repository.NewYearlyRepo(db)
	fileRepo		:= repository.NewFileRepo(db)
	projectRepo 	:= repository.NewProjectRepo(db)
	summaryRepo		:= repository.NewSummaryRepo(db)
	userRepo 		:= repository.NewUserRepo(db)
	sessionRepo		:= repository.NewSessionRepo(db)

	crudService 	:= service.NewCrudService(
		attendanceRepo,
		analisaRepo,
		factorRepo,
		fileRepo,
		itemRepo,
		masalahRepo,
		minipapRepo,
		monthlyRepo,
		projectRepo,
		resultRepo,
		summaryRepo,
		userRepo,
		yearlyRepo,
	)
	
	parseService	:= service.NewParseService(
		fileRepo,
	)

	userService 	:= service.NewUserService(userRepo, sessionRepo)
	
	kpiAPIHandler := api.NewKpiAPI(crudService)
	fileAPIHandler := api.NewFileAPI(crudService, parseService)
	analisaAPIHandler := api.NewAnalisaAPI(crudService)
	projectAPIHandler := api.NewProjectAPI(crudService)
	userAPIHandler := api.NewUserAPI(crudService, userService)
	apiHandler := APIHandler{
		KpiAPIHandler: kpiAPIHandler,
		FileAPIHandler: fileAPIHandler,
		AnalisaAPIHandler: analisaAPIHandler,
		ProjectAPIHandler: projectAPIHandler,
		UserAPIHandler: userAPIHandler,
	}
	kpi := gin.Group("/kpi")
	{
		analisa := kpi.Group("/analisa")
		{
			analisa.GET("", apiHandler.AnalisaAPIHandler.GetAnalisaList)
			analisa.GET("/:id", apiHandler.AnalisaAPIHandler.GetAnalisaByID)
			analisa.Use(middleware.Auth())
			analisa.POST("", apiHandler.AnalisaAPIHandler.AddAnalisa)
			analisa.POST("/entire", apiHandler.AnalisaAPIHandler.AddEntireAnalisa)
			analisa.PUT("/:id", apiHandler.AnalisaAPIHandler.UpdateAnalisa)
			analisa.DELETE("/:id", apiHandler.AnalisaAPIHandler.DeleteAnalisa)
			analisa.DELETE("/entire/:id", apiHandler.AnalisaAPIHandler.DeleteEntireAnalisa)
		}
		attendance := kpi.Group("/attendance") //inefficient endpoint
		{
			attendance.GET("", apiHandler.KpiAPIHandler.GetAttendanceList)
			attendance.GET("/:id", apiHandler.KpiAPIHandler.GetAttendanceByID)
			attendance.Use(middleware.Auth())
			attendance.POST("", apiHandler.KpiAPIHandler.AddAttendance)
			attendance.POST("/entire", apiHandler.KpiAPIHandler.AddEntireAttendance)
			attendance.PUT("/:id", apiHandler.KpiAPIHandler.UpdateAttendance)
			attendance.DELETE("/:id", apiHandler.KpiAPIHandler.DeleteAttendance)
			attendance.DELETE("/entire/:id", apiHandler.KpiAPIHandler.DeleteEntireAttendance)
		}
		factor := kpi.Group("/factor") //inefficient endpoint
		{
			factor.GET("", apiHandler.KpiAPIHandler.GetFactorList)
			factor.GET("/:id", apiHandler.KpiAPIHandler.GetFactorByID)
			factor.Use(middleware.Auth())
			factor.POST("", apiHandler.KpiAPIHandler.AddFactor)
			factor.POST("/entire", apiHandler.KpiAPIHandler.AddEntireFactor)
			factor.PUT("/:id", apiHandler.KpiAPIHandler.UpdateFactor)
			factor.DELETE("/:id", apiHandler.KpiAPIHandler.DeleteFactor)
			factor.DELETE("/entire/:id", apiHandler.KpiAPIHandler.DeleteEntireFactor)
		}
		item := kpi.Group("/item") //inefficient endpoint
		{
			item.GET("", apiHandler.KpiAPIHandler.GetItemList)
			item.GET("/:id", apiHandler.KpiAPIHandler.GetItemByID)
			item.Use(middleware.Auth())
			item.POST("", apiHandler.KpiAPIHandler.AddItem)
			item.POST("/entire", apiHandler.KpiAPIHandler.AddEntireItem)
			item.PUT("/:id", apiHandler.KpiAPIHandler.UpdateItem)
			item.DELETE("/:id", apiHandler.KpiAPIHandler.DeleteItem)
			item.DELETE("/entire/:id", apiHandler.KpiAPIHandler.DeleteEntireItem)
		}
		masalah := kpi.Group("/masalah")
		{
			masalah.GET("", apiHandler.AnalisaAPIHandler.GetMasalahList)
			masalah.GET("/:id", apiHandler.AnalisaAPIHandler.GetMasalahByID)
			masalah.Use(middleware.Auth())
			masalah.POST("", apiHandler.AnalisaAPIHandler.AddMasalah)
			masalah.PUT("/:id", apiHandler.AnalisaAPIHandler.UpdateMasalah)
			masalah.DELETE("/:id", apiHandler.AnalisaAPIHandler.DeleteMasalah)
		}
		minipap := kpi.Group("/minipap") //inefficient endpoint
		{
			minipap.GET("", apiHandler.KpiAPIHandler.GetMinipapList)
			minipap.GET("/:id", apiHandler.KpiAPIHandler.GetMinipapByID)
			minipap.Use(middleware.Auth())
			minipap.POST("", apiHandler.KpiAPIHandler.AddMinipap)
			minipap.PUT("/:id", apiHandler.KpiAPIHandler.UpdateMinipap)
			minipap.DELETE("/:id", apiHandler.KpiAPIHandler.DeleteMinipap)
		}
		monthly := kpi.Group("/monthly") //inefficient endpoint
		{
			monthly.GET("", apiHandler.KpiAPIHandler.GetMonthlyList)
			monthly.GET("/:id", apiHandler.KpiAPIHandler.GetMonthlyByID)
			monthly.Use(middleware.Auth())
			monthly.POST("", apiHandler.KpiAPIHandler.AddMonthly)
			monthly.PUT("/:id", apiHandler.KpiAPIHandler.UpdateMonthly)
			monthly.DELETE("/:id", apiHandler.KpiAPIHandler.DeleteMonthly)
		}
		project := kpi.Group("/project")
		{
			project.GET("", apiHandler.ProjectAPIHandler.GetProjectList)
			project.GET("/:id", apiHandler.ProjectAPIHandler.GetProjectByID)
			project.Use(middleware.Auth())
			project.POST("", apiHandler.ProjectAPIHandler.AddProject)
			project.PUT("/:id", apiHandler.ProjectAPIHandler.UpdateProject)
			project.DELETE("/:id", apiHandler.ProjectAPIHandler.DeleteProject)
		}
		result := kpi.Group("/result") //inefficient endpoint
		{
			result.GET("", apiHandler.KpiAPIHandler.GetResultList)
			result.GET("/:id", apiHandler.KpiAPIHandler.GetResultByID)
			result.Use(middleware.Auth())
			result.POST("", apiHandler.KpiAPIHandler.AddResult)
			result.POST("/entire", apiHandler.KpiAPIHandler.AddEntireResult)
			result.PUT("/:id", apiHandler.KpiAPIHandler.UpdateResult)
			result.DELETE("/:id", apiHandler.KpiAPIHandler.DeleteResult)
			result.DELETE("/entire/:id", apiHandler.KpiAPIHandler.DeleteEntireResult)
		}
		summary := kpi.Group("/summary")
		{
			summary.GET("", apiHandler.ProjectAPIHandler.GetSummaryList)
			summary.GET("/:id", apiHandler.ProjectAPIHandler.GetSummaryByID)
			summary.Use(middleware.Auth())
			summary.POST("", apiHandler.ProjectAPIHandler.AddSummary)
			summary.POST("/entire", apiHandler.ProjectAPIHandler.AddEntireSummary)
			summary.PUT("/:id", apiHandler.ProjectAPIHandler.UpdateSummary)
			summary.DELETE("/:id", apiHandler.ProjectAPIHandler.DeleteSummary)
			summary.DELETE("/entire/:id", apiHandler.ProjectAPIHandler.DeleteEntireSummary)
		}
		user := kpi.Group("/user")
		{
			user.POST("/login", apiHandler.UserAPIHandler.Login)
			user.POST("/register", apiHandler.UserAPIHandler.Register)
			user.POST("/logout", apiHandler.UserAPIHandler.Logout)
			user.GET("/profile", apiHandler.UserAPIHandler.Profile)
		}
		yearly := kpi.Group("/yearly") 
		{
			yearly.GET("", apiHandler.KpiAPIHandler.GetYearlyList)
			yearly.GET("/:id", apiHandler.KpiAPIHandler.GetYearlyByID)
			yearly.Use(middleware.Auth())
			yearly.POST("", apiHandler.KpiAPIHandler.AddYearly)
			yearly.PUT("/:id", apiHandler.KpiAPIHandler.UpdateYearly)
			yearly.DELETE("/:id", apiHandler.KpiAPIHandler.DeleteYearly)
			yearly.POST("/entire", apiHandler.KpiAPIHandler.AddEntireYearly)
			yearly.DELETE("/entire/:id", apiHandler.KpiAPIHandler.DeleteEntireYearly)
		}
		file := kpi.Group("/file", middleware.Auth())
		{
			file.POST("/kpi", apiHandler.FileAPIHandler.KpiFileUpload)
			file.POST("/analisa", apiHandler.FileAPIHandler.AnalisaFileUpload)
			file.POST("/summary", apiHandler.FileAPIHandler.SummaryFileUpload)
			file.POST("/save", apiHandler.FileAPIHandler.SaveFile)
		}
	}
	return gin
}

func main(){
	gin.SetMode(gin.ReleaseMode)

	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Print("Missing .env file. Probably okay on dockerized environment")
	// }
	// config := &db.Config{
	// 	Host:     os.Getenv("DB_HOST"),
	// 	Port:     os.Getenv("DB_PORT"),
	// 	Password: os.Getenv("DB_PASS"),
	// 	User:     os.Getenv("DB_USER"),
	// 	SSLMode:  os.Getenv("DB_SSLMODE"),
	// 	DBName:   os.Getenv("DB_NAME"),
	// }

	config := &db.Config{
		Host:     "aws-0-ap-southeast-1.pooler.supabase.com",
		Port:     "5432",
		Password: "Technosport@2024",
		User:     "postgres.kfwmmnkrcdvyysxgbame",
		SSLMode:  "disable",
		DBName:   "postgres",
	}



	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		router := gin.New()
		db := db.NewDB()
		router.Use(gin.Recovery())
		router.Use(CORSMiddleware())
		router.ForwardedByClientIP = true
		router.SetTrustedProxies([]string{"127.0.0.1"})

		conn, err := db.Connect(config)
		if err != nil {
			panic(err)
		}

		conn.AutoMigrate(&model.Monthly{}, &model.Masalah{}, &model.Project{}, &model.UploadFile{}, &model.User{}, &model.Session{}) 
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
			fmt.Println("Port 8080 taken, switching port 8081")
			err = router.Run(":8081")
			if err != nil {
				panic(err)
			}
		}

	}()

	wg.Wait()
}
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Content-Length, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}