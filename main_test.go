package main

import (
	// "bytes"
	"goreact/db"
	"goreact/model"

	// _ "embed"
	// "encoding/json"
	"fmt"
	// "io/ioutil"
	// "net/http"
	// "net/http/httptest"
	// "strconv"
	// "strings"
	"sync"
	// "testing"
	// "time"

	"github.com/gin-gonic/gin"
	// _ "github.com/lib/pq"
	// "github.com/stretchr/testify/assert"
)

func start() (*gin.Engine, *sync.WaitGroup){
	gin.SetMode(gin.ReleaseMode)
	config := &db.Config{
		Host:     "db.kkmkegheitaqvhygnixh.supabase.co",
		Port:     "5432",
		Password: "kalomaugenerikaja",
		User:     "postgres",
		SSLMode:  "disable",
		DBName:   "postgres",
	}
	router := gin.New()
	db := db.NewDB()
	conn, err := db.Connect(config)
	if err != nil {panic(err)}

	conn.AutoMigrate(&model.Monthly{}) 
	conn.AutoMigrate(&model.MiniPAP{})
	conn.AutoMigrate(&model.Attendance{})
	// conn.AutoMigrate(&model.PAP{})
	conn.AutoMigrate(&model.Factor{})
	conn.AutoMigrate(&model.Result{})
	conn.AutoMigrate(&model.Item{})
	conn.AutoMigrate(&model.Yearly{})

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		router = RunServer(conn, router)
		fmt.Println("Test Server is running on port 8080")
		err := router.Run(":8080")
		if err != nil {
			panic(err)
		}
	}()

	return router, &wg
}
