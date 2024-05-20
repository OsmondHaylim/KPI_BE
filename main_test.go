package main

import (
	"bytes"
	"goreact/db"
	"goreact/model"
	"testing"

	// _ "embed"
	"encoding/json"
	"fmt"
	"io"

	"net/http"
	"net/http/httptest"
	// "strconv"
	"strings"
	"sync"
	// "testing"
	"time"

	"github.com/gin-gonic/gin"
	// _ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func start(port string) (*gin.Engine, *sync.WaitGroup){
	gin.SetMode(gin.ReleaseMode)
	config := &db.Config{
		Host:     "aws-0-ap-southeast-1.pooler.supabase.com",
		Port:     "5432",
		Password: "Technosport@2024",
		User:     "postgres.mjmfiwrmndhoitdltbud",
		SSLMode:  "disable",
		DBName:   "postgres",
	}
	router := gin.New()
	db := db.NewDB()
	conn, err := db.Connect(config)
	if err != nil {panic(err)}

	conn.Migrator().DropTable(&model.Yearly{})
	conn.Migrator().DropTable(&model.Item{})
	conn.Migrator().DropTable(&model.Result{})
	conn.Migrator().DropTable(&model.Factor{})
	conn.Migrator().DropTable(&model.Attendance{})
	conn.Migrator().DropTable(&model.MiniPAP{}, &model.Analisa{}, &model.Summary{})
	conn.Migrator().DropTable(&model.Monthly{}, &model.Masalah{}, &model.Project{}, &model.UploadFile{}, &model.User{}, &model.Session{})

	conn.AutoMigrate(&model.Monthly{}, &model.Masalah{}, &model.Project{}, &model.UploadFile{}, &model.User{}, &model.Session{}) 
	conn.AutoMigrate(&model.MiniPAP{}, &model.Analisa{}, &model.Summary{})
	conn.AutoMigrate(&model.Attendance{})
	conn.AutoMigrate(&model.Factor{})
	conn.AutoMigrate(&model.Result{})
	conn.AutoMigrate(&model.Item{})
	conn.AutoMigrate(&model.Yearly{})

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		router = RunServer(conn, router)
		err := router.Run(port)
		if err != nil {
			panic(err)
		}
		fmt.Println("Test Server is running on port " + port)
	}()
	return router, &wg
}
func getSuccess(w *bytes.Buffer){
	var temp model.SuccessResponse
	body, err := io.ReadAll(w)
	if err != nil {
		fmt.Println("Error reading response:", err)
	}
	resJson := []byte(body)
	err = json.Unmarshal(resJson, &temp)
	if err != nil {
		var temp2 model.ErrorResponse
		err = json.Unmarshal(resJson, &temp2)
		if err != nil{
			fmt.Println("Error parsing JSON:", err)
		}
		fmt.Println(temp2.Error)
	}
}
func stop(router *gin.Engine, wg *sync.WaitGroup){
	wg.Done()
}

func TestMain(t *testing.T){
	t.Run("Running Server", func(t *testing.T){
		router, wg := start(":8080") 
		defer stop(router, wg)
		time.Sleep(1 * time.Second)
		apiKey := ""
		t.Run("Account", func(t *testing.T){
			t.Run("Register", func(t *testing.T){
				registerBody := `{
					"username": "test",
					"email": "test@mail.com",
					"password": "12345678",
					"confirm_password": "12345678"
				}`
				req, _ := http.NewRequest("POST", "/kpi/user/register", strings.NewReader(registerBody))
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusCreated, w.Code)
			})
			t.Run("Login", func(t *testing.T){
				registerBody := `{
					"username": "test",
					"password": "12345678"
				}`
				req, _ := http.NewRequest("POST", "/kpi/user/login", strings.NewReader(registerBody))
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
				temp := model.LoginResponse{}
				_ = json.Unmarshal(w.Body.Bytes(), &temp)
				apiKey = temp.Token
			})
		})
		t.Run("Monthly", func(t *testing.T){
			t.Run("Add Single", func(t *testing.T){
				monthlyBody := `{
					"January": 9,
					"February": 6,
					"March": 6,
					"April": 5,
					"May": 6,
					"June": 6,
					"July": 7,
					"August": 6,
					"September": 5,
					"October": 4,
					"November": 5,
					"December": 4,
					"Remarks": ""
				}`
				req, _ := http.NewRequest("POST", "/kpi/monthly", strings.NewReader(monthlyBody))
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusCreated, w.Code)
			})
			t.Run("Get Single", func(t *testing.T){
				req, _ := http.NewRequest("GET", "/kpi/monthly/1", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Update Single", func(t *testing.T){
				monthlyBody := `{
					"Monthly_ID": 1,
					"January": 1,
					"February": 1,
					"March": 1,
					"April": 5,
					"May": 6,
					"June": 6,
					"July": 7,
					"August": 6,
					"September": 5,
					"October": 4,
					"November": 5,
					"December": 4,
					"Remarks": "test"
				}`
				req, _ := http.NewRequest("PUT", "/kpi/monthly/1", strings.NewReader(monthlyBody))
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)

				req, _ = http.NewRequest("GET", "/kpi/monthly/1", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w = httptest.NewRecorder()
				router.ServeHTTP(w, req)
			})
			
			t.Run("Delete Single", func(t *testing.T){
				req, _ := http.NewRequest("DELETE", "/kpi/monthly/1", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Get All", func(t *testing.T){
				req, _ := http.NewRequest("GET", "/kpi/monthly", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
		})
		t.Run("Attendance", func(t *testing.T){
			t.Run("Add Single", func(t *testing.T){
				attendanceBody := `{
					"Year": 2000
				}`
				req, _ := http.NewRequest("POST", "/kpi/attendance", strings.NewReader(attendanceBody))
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusCreated, w.Code)
			})
			t.Run("Add Entire", func(t *testing.T){
				attendanceBody := `{
					"Year": 2023,
					"Planned": {
						"MiniPAP": null,
						"January": 0.94,
						"February": 0.94,
						"March": 0.94,
						"April": 0.94,
						"May": 0.94,
						"June": 0,
						"July": 0.94,
						"August": 0.94,
						"September": 0,
						"October": 0.94,
						"November": 0.94,
						"December": 0.94,
						"YTD": 0,
						"Remarks": "",
						"MinipapID": null
					},
					"Actual": {
						"MiniPAP": null,
						"January": 0.9469,
						"February": 0.9412,
						"March": 0.9412,
						"April": 0.9765,
						"May": 0.9144,
						"June": 0,
						"July": 0.9441,
						"August": 0.9251,
						"September": 0,
						"October": 0.9517,
						"November": 0.94,
						"December": 0.9564,
						"YTD": 0,
						"Remarks": "",
						"MinipapID": null
					},
					"Cuti": {
						"MiniPAP": null,
						"January": 0.0088,
						"February": 0.0082,
						"March": 0.0117,
						"April": 0.0023,
						"May": 0.017,
						"June": 0,
						"July": 0.01,
						"August": 0.014,
						"September": 0,
						"October": 0.008,
						"November": 0.0075,
						"December": 0.006,
						"YTD": 0,
						"Remarks": "",
						"MinipapID": null
					},
					"Izin": {
						"MiniPAP": null,
						"January": 0.0005,
						"February": 0.0017,
						"March": 0.0012,
						"April": 0,
						"May": 0.0006,
						"June": 0,
						"July": 0.0005,
						"August": 0.0011,
						"September": 0,
						"October": 0.0018,
						"November": 0.0056,
						"December": 0.001,
						"YTD": 0,
						"Remarks": "",
						"MinipapID": null
					},
					"Lain": {
						"MiniPAP": null,
						"January": 0.0005,
						"February": 0.0012,
						"March": 0.0012,
						"April": 0,
						"May": 0.0006,
						"June": 0,
						"July": 0.0005,
						"August": 0.0005,
						"September": 0,
						"October": 0.0006,
						"November": 0.0006,
						"December": 0.0006,
						"YTD": 0,
						"Remarks": "",
						"MinipapID": null
					}
				}`
				req, _ := http.NewRequest("POST", "/kpi/attendance/entire", strings.NewReader(attendanceBody))
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				getSuccess(w.Body)
				assert.Equal(t, http.StatusCreated, w.Code)
			})
			t.Run("Get Single", func(t *testing.T){
				req, _ := http.NewRequest("GET", "/kpi/attendance/2000", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			// t.Run("Update Single", func(t *testing.T){
			// 	attendanceBody := `{
			// 		"Title": "a. LKK",
			// 		"Unit": "case",
			// 		"Target": "4 /Month"
			// 	}`
			// 	req, _ := http.NewRequest("PUT", "/kpi/attendance/1", strings.NewReader(attendanceBody))
				// req.Header.Set("Authorization", "Bearer "+ apiKey)
			// 	w := httptest.NewRecorder()
			// 	router.ServeHTTP(w, req)
			// 	// getSuccess(w.Body)
			// 	assert.Equal(t, http.StatusOK, w.Code)
			// })
			t.Run("Delete Single", func(t *testing.T){
				req, _ := http.NewRequest("DELETE", "/kpi/attendance/2000", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Delete Entire", func(t *testing.T){
				req, _ := http.NewRequest("DELETE", "/kpi/attendance/entire/2023", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Get All", func(t *testing.T){
				req, _ := http.NewRequest("GET", "/kpi/attendance", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
		})
		t.Run("Minipap", func(t *testing.T){
			t.Run("Add Single", func(t *testing.T){
				minipapBody := `{}`
				req, _ := http.NewRequest("POST", "/kpi/minipap", strings.NewReader(minipapBody))
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusCreated, w.Code)
			})
			t.Run("Get Single", func(t *testing.T){
				req, _ := http.NewRequest("GET", "/kpi/minipap/1", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			// t.Run("Update Single", func(t *testing.T){
			// 	minipapBody := `{
			// 		"Minipap_ID": 2
			// 	}`
			// 	req, _ := http.NewRequest("PUT", "/kpi/minipap/1", strings.NewReader(minipapBody))
				// req.Header.Set("Authorization", "Bearer "+ apiKey)
			// 	w := httptest.NewRecorder()
			// 	router.ServeHTTP(w, req)
			// 	// getSuccess(w.Body)
			// 	assert.Equal(t, http.StatusOK, w.Code)
			// })
			t.Run("Delete Single", func(t *testing.T){
				req, _ := http.NewRequest("DELETE", "/kpi/minipap/1", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Get All", func(t *testing.T){
				req, _ := http.NewRequest("GET", "/kpi/minipap", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
		})
		t.Run("Factor", func(t *testing.T){
			t.Run("Add Single", func(t *testing.T){
				factorBody := `{
					"Title": "Zero Fire, Work & Traffic accident",
					"Unit": "case",
					"Target": "Zero"
				}`
				req, _ := http.NewRequest("POST", "/kpi/factor", strings.NewReader(factorBody))
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusCreated, w.Code)
			})
			t.Run("Add Entire", func(t *testing.T){
				factorBody := `{
					"Title": "Zero Fire, Work & Traffic accident",
					"Unit": "case",
					"Target": "Zero",
					"Planned": {
						"Monthly": [
							{
								"Monthly_ID": 0,
								"MiniPAP": null,
								"January": 0,
								"February": 0,
								"March": 0,
								"April": 0,
								"May": 0,
								"June": 0,
								"July": 0,
								"August": 0,
								"September": 0,
								"October": 0,
								"November": 0,
								"December": 0,
								"YTD": 0,
								"Remarks": "",
								"MinipapID": null
							}
						]
					},
					"Actual": {
						"Monthly": [
							{
								"Monthly_ID": 0,
								"MiniPAP": null,
								"January": 0,
								"February": 0,
								"March": 0,
								"April": 0,
								"May": 0,
								"June": 0,
								"July": 0,
								"August": 0,
								"September": 0,
								"October": 0,
								"November": 0,
								"December": 0,
								"YTD": 0,
								"Remarks": "",
								"MinipapID": null
							}
						]
					}
				}`
				req, _ := http.NewRequest("POST", "/kpi/factor/entire", strings.NewReader(factorBody))
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				getSuccess(w.Body)
				assert.Equal(t, http.StatusCreated, w.Code)
			})
			t.Run("Get Single", func(t *testing.T){
				req, _ := http.NewRequest("GET", "/kpi/factor/1", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Update Single", func(t *testing.T){
				factorBody := `{
					"Title": "a. LKK",
					"Unit": "case",
					"Target": "4 /Month"
				}`
				req, _ := http.NewRequest("PUT", "/kpi/factor/1", strings.NewReader(factorBody))
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Delete Single", func(t *testing.T){
				req, _ := http.NewRequest("DELETE", "/kpi/factor/1", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Delete Entire", func(t *testing.T){
				req, _ := http.NewRequest("DELETE", "/kpi/factor/entire/2", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Get All", func(t *testing.T){
				req, _ := http.NewRequest("GET", "/kpi/factor", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
		})
		t.Run("Result", func(t *testing.T){
			t.Run("Add Single", func(t *testing.T){
				resultBody := `{
					"Name": "Test"
				}`
				req, _ := http.NewRequest("POST", "/kpi/result", strings.NewReader(resultBody))
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusCreated, w.Code)
			})
			t.Run("Add Entire", func(t *testing.T){
				resultBody := `{
                    "Result_ID": 0,
                    "Name": "2. Safety minded/awareness ",
                    "Factors": [
                        {
                            "Factor_ID": 0,
                            "Title": "a. LKK",
                            "Unit": "case",
                            "Target": "4 /Month",
                            "Planned": {
                                "Minipap_ID": 0,
                                "Monthly": [
                                    {
                                        "Monthly_ID": 0,
                                        "MiniPAP": null,
                                        "January": 4,
                                        "February": 4,
                                        "March": 4,
                                        "April": 4,
                                        "May": 4,
                                        "June": 4,
                                        "July": 4,
                                        "August": 4,
                                        "September": 4,
                                        "October": 4,
                                        "November": 4,
                                        "December": 4,
                                        "YTD": 0,
                                        "Remarks": "",
                                        "MinipapID": null
                                    }
                                ]
                            },
                            "Actual": {
                                "Minipap_ID": 0,
                                "Monthly": [
                                    {
                                        "Monthly_ID": 0,
                                        "MiniPAP": null,
                                        "January": 9,
                                        "February": 6,
                                        "March": 6,
                                        "April": 5,
                                        "May": 6,
                                        "June": 6,
                                        "July": 7,
                                        "August": 6,
                                        "September": 5,
                                        "October": 4,
                                        "November": 5,
                                        "December": 4,
                                        "YTD": 0,
                                        "Remarks": "",
                                        "MinipapID": null
                                    }
                                ]
                            },
                            "Percentage": null
                        },
                        {
                            "Factor_ID": 0,
                            "Title": "b. LKBK",
                            "Unit": "case",
                            "Target": "14/Month",
                            "Planned": {
                                "Minipap_ID": 0,
                                "Monthly": [
                                    {
                                        "Monthly_ID": 0,
                                        "MiniPAP": null,
                                        "January": 14,
                                        "February": 14,
                                        "March": 14,
                                        "April": 14,
                                        "May": 14,
                                        "June": 14,
                                        "July": 14,
                                        "August": 14,
                                        "September": 14,
                                        "October": 14,
                                        "November": 14,
                                        "December": 14,
                                        "YTD": 0,
                                        "Remarks": "",
                                        "MinipapID": null
                                    }
                                ]
                            },
                            "Actual": {
                                "Minipap_ID": 0,
                                "Monthly": [
                                    {
                                        "Monthly_ID": 0,
                                        "MiniPAP": null,
                                        "January": 15,
                                        "February": 16,
                                        "March": 16,
                                        "April": 12,
                                        "May": 13,
                                        "June": 15,
                                        "July": 17,
                                        "August": 14,
                                        "September": 14,
                                        "October": 14,
                                        "November": 14,
                                        "December": 14,
                                        "YTD": 0,
                                        "Remarks": "",
                                        "MinipapID": null
                                    }
                                ]
                            },
                            "Percentage": null
                        }
                    ]
                },`
				req, _ := http.NewRequest("POST", "/kpi/result/entire", strings.NewReader(resultBody))
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				getSuccess(w.Body)
				assert.Equal(t, http.StatusCreated, w.Code)
			})
			t.Run("Get Single", func(t *testing.T){
				req, _ := http.NewRequest("GET", "/kpi/result/1", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Update Single", func(t *testing.T){
				resultBody := `{
					"Name": "a. LKK"
				}`
				req, _ := http.NewRequest("PUT", "/kpi/result/1", strings.NewReader(resultBody))
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Delete Single", func(t *testing.T){
				req, _ := http.NewRequest("DELETE", "/kpi/result/1", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Delete Entire", func(t *testing.T){
				req, _ := http.NewRequest("DELETE", "/kpi/result/entire/2", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Get All", func(t *testing.T){
				req, _ := http.NewRequest("GET", "/kpi/result", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
		})
		t.Run("Item", func(t *testing.T){
			t.Run("Add Single", func(t *testing.T){
				itemBody := `{
					"Name": "Test"
				}`
				req, _ := http.NewRequest("POST", "/kpi/item", strings.NewReader(itemBody))
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusCreated, w.Code)
			})
			t.Run("Add Entire", func(t *testing.T){
				itemBody := `{
					"Item_ID": 0,
					"Name": "S",
					"Results": [
						{
							"Result_ID": 0,
							"Name": "1. Fire, Work & Traffic \naccident ",
							"Factors": [
								{
									"Factor_ID": 0,
									"Title": "Zero Fire, Work & Traffic accident",
									"Unit": "case",
									"Target": "Zero",
									"Planned": {
										"Minipap_ID": 0,
										"Monthly": [
											{
												"Monthly_ID": 0,
												"MiniPAP": null,
												"January": 0,
												"February": 0,
												"March": 0,
												"April": 0,
												"May": 0,
												"June": 0,
												"July": 0,
												"August": 0,
												"September": 0,
												"October": 0,
												"November": 0,
												"December": 0,
												"YTD": 0,
												"Remarks": "",
												"MinipapID": null
											}
										]
									},
									"Actual": {
										"Minipap_ID": 0,
										"Monthly": [
											{
												"Monthly_ID": 0,
												"MiniPAP": null,
												"January": 0,
												"February": 0,
												"March": 0,
												"April": 0,
												"May": 0,
												"June": 0,
												"July": 0,
												"August": 0,
												"September": 0,
												"October": 0,
												"November": 0,
												"December": 0,
												"YTD": 0,
												"Remarks": "",
												"MinipapID": null
											}
										]
									},
									"Percentage": null
								}
							]
						},
						{
							"Result_ID": 0,
							"Name": "2. Safety minded/awareness ",
							"Factors": [
								{
									"Factor_ID": 0,
									"Title": "a. LKK",
									"Unit": "case",
									"Target": "4 /Month",
									"Planned": {
										"Minipap_ID": 0,
										"Monthly": [
											{
												"Monthly_ID": 0,
												"MiniPAP": null,
												"January": 4,
												"February": 4,
												"March": 4,
												"April": 4,
												"May": 4,
												"June": 4,
												"July": 4,
												"August": 4,
												"September": 4,
												"October": 4,
												"November": 4,
												"December": 4,
												"YTD": 0,
												"Remarks": "",
												"MinipapID": null
											}
										]
									},
									"Actual": {
										"Minipap_ID": 0,
										"Monthly": [
											{
												"Monthly_ID": 0,
												"MiniPAP": null,
												"January": 9,
												"February": 6,
												"March": 6,
												"April": 5,
												"May": 6,
												"June": 6,
												"July": 7,
												"August": 6,
												"September": 5,
												"October": 4,
												"November": 5,
												"December": 4,
												"YTD": 0,
												"Remarks": "",
												"MinipapID": null
											}
										]
									},
									"Percentage": null
								},
								{
									"Factor_ID": 0,
									"Title": "b. LKBK",
									"Unit": "case",
									"Target": "14/Month",
									"Planned": {
										"Minipap_ID": 0,
										"Monthly": [
											{
												"Monthly_ID": 0,
												"MiniPAP": null,
												"January": 14,
												"February": 14,
												"March": 14,
												"April": 14,
												"May": 14,
												"June": 14,
												"July": 14,
												"August": 14,
												"September": 14,
												"October": 14,
												"November": 14,
												"December": 14,
												"YTD": 0,
												"Remarks": "",
												"MinipapID": null
											}
										]
									},
									"Actual": {
										"Minipap_ID": 0,
										"Monthly": [
											{
												"Monthly_ID": 0,
												"MiniPAP": null,
												"January": 15,
												"February": 16,
												"March": 16,
												"April": 12,
												"May": 13,
												"June": 15,
												"July": 17,
												"August": 14,
												"September": 14,
												"October": 14,
												"November": 14,
												"December": 14,
												"YTD": 0,
												"Remarks": "",
												"MinipapID": null
											}
										]
									},
									"Percentage": null
								}
							]
						},
						{
							"Result_ID": 0,
							"Name": "3. BSFA Activity ",
							"Factors": [
								{
									"Factor_ID": 0,
									"Title": "a. 3S Activity & Assessment",
									"Unit": "item",
									"Target": "As schedule",
									"Planned": {
										"Minipap_ID": 0,
										"Monthly": [
											{
												"Monthly_ID": 0,
												"MiniPAP": null,
												"January": 1,
												"February": 1,
												"March": 1,
												"April": 1,
												"May": 1,
												"June": 1,
												"July": 1,
												"August": 1,
												"September": 1,
												"October": 1,
												"November": 1,
												"December": 1,
												"YTD": 0,
												"Remarks": "",
												"MinipapID": null
											}
										]
									},
									"Actual": {
										"Minipap_ID": 0,
										"Monthly": [
											{
												"Monthly_ID": 0,
												"MiniPAP": null,
												"January": 1,
												"February": 1,
												"March": 1,
												"April": 1,
												"May": 1,
												"June": 1,
												"July": 1,
												"August": 1,
												"September": 1,
												"October": 1,
												"November": 1,
												"December": 1,
												"YTD": 0,
												"Remarks": "",
												"MinipapID": null
											}
										]
									},
									"Percentage": null
								},
								{
									"Factor_ID": 0,
									"Title": "b. KY Small group (4 group)",
									"Unit": "case",
									"Target": "4/Month",
									"Planned": {
										"Minipap_ID": 0,
										"Monthly": [
											{
												"Monthly_ID": 0,
												"MiniPAP": null,
												"January": 4,
												"February": 4,
												"March": 4,
												"April": 4,
												"May": 4,
												"June": 4,
												"July": 4,
												"August": 4,
												"September": 4,
												"October": 4,
												"November": 4,
												"December": 4,
												"YTD": 0,
												"Remarks": "",
												"MinipapID": null
											}
										]
									},
									"Actual": {
										"Minipap_ID": 0,
										"Monthly": [
											{
												"Monthly_ID": 0,
												"MiniPAP": null,
												"January": 4,
												"February": 4,
												"March": 4,
												"April": 4,
												"May": 4,
												"June": 4,
												"July": 4,
												"August": 4,
												"September": 4,
												"October": 4,
												"November": 4,
												"December": 4,
												"YTD": 0,
												"Remarks": "",
												"MinipapID": null
											}
										]
									},
									"Percentage": null
								},
								{
									"Factor_ID": 0,
									"Title": "c. Follow Up Hopper",
									"Unit": "item",
									"Target": "As schedule",
									"Planned": {
										"Minipap_ID": 0,
										"Monthly": [
											{
												"Monthly_ID": 0,
												"MiniPAP": null,
												"January": 20,
												"February": 15,
												"March": 0,
												"April": 18,
												"May": 2,
												"June": 0,
												"July": 0,
												"August": 0,
												"September": 0,
												"October": 0,
												"November": 0,
												"December": 0,
												"YTD": 0,
												"Remarks": "Finish",
												"MinipapID": null
											}
										]
									},
									"Actual": {
										"Minipap_ID": 0,
										"Monthly": [
											{
												"Monthly_ID": 0,
												"MiniPAP": null,
												"January": 27,
												"February": 7,
												"March": 0,
												"April": 6,
												"May": 0,
												"June": 15,
												"July": 0,
												"August": 0,
												"September": 0,
												"October": 0,
												"November": 0,
												"December": 0,
												"YTD": 0,
												"Remarks": "",
												"MinipapID": null
											}
										]
									},
									"Percentage": null
								},
								{
									"Factor_ID": 0,
									"Title": "d. Follow Up RA Cat III ",
									"Unit": "item",
									"Target": "As schedule",
									"Planned": {
										"Minipap_ID": 0,
										"Monthly": [
											{
												"Monthly_ID": 0,
												"MiniPAP": null,
												"January": 0,
												"February": 0,
												"March": 0,
												"April": 0,
												"May": 0,
												"June": 10,
												"July": 18,
												"August": 35,
												"September": 33,
												"October": 95,
												"November": 132,
												"December": 121,
												"YTD": 0,
												"Remarks": "",
												"MinipapID": null
											}
										]
									},
									"Actual": {
										"Minipap_ID": 0,
										"Monthly": [
											{
												"Monthly_ID": 0,
												"MiniPAP": null,
												"January": 0,
												"February": 0,
												"March": 0,
												"April": 0,
												"May": 0,
												"June": 41,
												"July": 12,
												"August": 30,
												"September": 51,
												"October": 136,
												"November": 45,
												"December": 75,
												"YTD": 0,
												"Remarks": "",
												"MinipapID": null
											}
										]
									},
									"Percentage": null
								},
								{
									"Factor_ID": 0,
									"Title": "e.  Re-inforcement Building \nStructure",
									"Unit": "item",
									"Target": "As schedule",
									"Planned": {
										"Minipap_ID": 0,
										"Monthly": [
											{
												"Monthly_ID": 0,
												"MiniPAP": null,
												"January": 0,
												"February": 0,
												"March": 0,
												"April": 1,
												"May": 1,
												"June": 1,
												"July": 1,
												"August": 1,
												"September": 1,
												"October": 1,
												"November": 1,
												"December": 2,
												"YTD": 0,
												"Remarks": "Finish",
												"MinipapID": null
											}
										]
									},
									"Actual": {
										"Minipap_ID": 0,
										"Monthly": [
											{
												"Monthly_ID": 0,
												"MiniPAP": null,
												"January": 0,
												"February": 0,
												"March": 0,
												"April": 0,
												"May": 0,
												"June": 2,
												"July": 1,
												"August": 2,
												"September": 2,
												"October": 1,
												"November": 1,
												"December": 1,
												"YTD": 0,
												"Remarks": "",
												"MinipapID": null
											}
										]
									},
									"Percentage": null
								}
							]
						}
					]
				}`
				req, _ := http.NewRequest("POST", "/kpi/item/entire", strings.NewReader(itemBody))
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				getSuccess(w.Body)
				assert.Equal(t, http.StatusCreated, w.Code)
			})
			t.Run("Get Single", func(t *testing.T){
				req, _ := http.NewRequest("GET", "/kpi/item/1", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Update Single", func(t *testing.T){
				itemBody := `{
					"Name": "a. LKK"
				}`
				req, _ := http.NewRequest("PUT", "/kpi/item/1", strings.NewReader(itemBody))
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Delete Single", func(t *testing.T){
				req, _ := http.NewRequest("DELETE", "/kpi/item/1", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Delete Entire", func(t *testing.T){
				req, _ := http.NewRequest("DELETE", "/kpi/item/entire/2", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Get All", func(t *testing.T){
				req, _ := http.NewRequest("GET", "/kpi/item", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
		})
		t.Run("Year", func(t *testing.T){
			t.Run("Add Single", func(t *testing.T){
				yearlyBody := `{
					"Year": 2025
				}`
				req, _ := http.NewRequest("POST", "/kpi/yearly", strings.NewReader(yearlyBody))
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusCreated, w.Code)
			})
			t.Run("Add Entire", func(t *testing.T){
				yearlyBody := `{
					"Year": 2023,
					"Items": [
						{
							"Item_ID": 0,
							"Name": "S",
							"Results": [
								{
									"Result_ID": 0,
									"Name": "1. Fire, Work & Traffic \naccident ",
									"Factors": [
										{
											"Factor_ID": 0,
											"Title": "Zero Fire, Work & Traffic accident",
											"Unit": "case",
											"Target": "Zero",
											"Planned": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 0,
														"May": 0,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 0,
														"October": 0,
														"November": 0,
														"December": 0,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Actual": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 0,
														"May": 0,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 0,
														"October": 0,
														"November": 0,
														"December": 0,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Percentage": null
										}
									]
								},
								{
									"Result_ID": 0,
									"Name": "2. Safety minded/awareness ",
									"Factors": [
										{
											"Factor_ID": 0,
											"Title": "a. LKK",
											"Unit": "case",
											"Target": "4 /Month",
											"Planned": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 4,
														"February": 4,
														"March": 4,
														"April": 4,
														"May": 4,
														"June": 4,
														"July": 4,
														"August": 4,
														"September": 4,
														"October": 4,
														"November": 4,
														"December": 4,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Actual": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 9,
														"February": 6,
														"March": 6,
														"April": 5,
														"May": 6,
														"June": 6,
														"July": 7,
														"August": 6,
														"September": 5,
														"October": 4,
														"November": 5,
														"December": 4,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Percentage": null
										},
										{
											"Factor_ID": 0,
											"Title": "b. LKBK",
											"Unit": "case",
											"Target": "14/Month",
											"Planned": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 14,
														"February": 14,
														"March": 14,
														"April": 14,
														"May": 14,
														"June": 14,
														"July": 14,
														"August": 14,
														"September": 14,
														"October": 14,
														"November": 14,
														"December": 14,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Actual": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 15,
														"February": 16,
														"March": 16,
														"April": 12,
														"May": 13,
														"June": 15,
														"July": 17,
														"August": 14,
														"September": 14,
														"October": 14,
														"November": 14,
														"December": 14,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Percentage": null
										}
									]
								},
								{
									"Result_ID": 0,
									"Name": "3. BSFA Activity ",
									"Factors": [
										{
											"Factor_ID": 0,
											"Title": "a. 3S Activity & Assessment",
											"Unit": "item",
											"Target": "As schedule",
											"Planned": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 1,
														"February": 1,
														"March": 1,
														"April": 1,
														"May": 1,
														"June": 1,
														"July": 1,
														"August": 1,
														"September": 1,
														"October": 1,
														"November": 1,
														"December": 1,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Actual": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 1,
														"February": 1,
														"March": 1,
														"April": 1,
														"May": 1,
														"June": 1,
														"July": 1,
														"August": 1,
														"September": 1,
														"October": 1,
														"November": 1,
														"December": 1,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Percentage": null
										},
										{
											"Factor_ID": 0,
											"Title": "b. KY Small group (4 group)",
											"Unit": "case",
											"Target": "4/Month",
											"Planned": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 4,
														"February": 4,
														"March": 4,
														"April": 4,
														"May": 4,
														"June": 4,
														"July": 4,
														"August": 4,
														"September": 4,
														"October": 4,
														"November": 4,
														"December": 4,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Actual": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 4,
														"February": 4,
														"March": 4,
														"April": 4,
														"May": 4,
														"June": 4,
														"July": 4,
														"August": 4,
														"September": 4,
														"October": 4,
														"November": 4,
														"December": 4,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Percentage": null
										},
										{
											"Factor_ID": 0,
											"Title": "c. Follow Up Hopper",
											"Unit": "item",
											"Target": "As schedule",
											"Planned": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 20,
														"February": 15,
														"March": 0,
														"April": 18,
														"May": 2,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 0,
														"October": 0,
														"November": 0,
														"December": 0,
														"YTD": 0,
														"Remarks": "Finish",
														"MinipapID": null
													}
												]
											},
											"Actual": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 27,
														"February": 7,
														"March": 0,
														"April": 6,
														"May": 0,
														"June": 15,
														"July": 0,
														"August": 0,
														"September": 0,
														"October": 0,
														"November": 0,
														"December": 0,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Percentage": null
										},
										{
											"Factor_ID": 0,
											"Title": "d. Follow Up RA Cat III ",
											"Unit": "item",
											"Target": "As schedule",
											"Planned": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 0,
														"May": 0,
														"June": 10,
														"July": 18,
														"August": 35,
														"September": 33,
														"October": 95,
														"November": 132,
														"December": 121,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Actual": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 0,
														"May": 0,
														"June": 41,
														"July": 12,
														"August": 30,
														"September": 51,
														"October": 136,
														"November": 45,
														"December": 75,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Percentage": null
										},
										{
											"Factor_ID": 0,
											"Title": "e.  Re-inforcement Building \nStructure",
											"Unit": "item",
											"Target": "As schedule",
											"Planned": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 1,
														"May": 1,
														"June": 1,
														"July": 1,
														"August": 1,
														"September": 1,
														"October": 1,
														"November": 1,
														"December": 2,
														"YTD": 0,
														"Remarks": "Finish",
														"MinipapID": null
													}
												]
											},
											"Actual": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 0,
														"May": 0,
														"June": 2,
														"July": 1,
														"August": 2,
														"September": 2,
														"October": 1,
														"November": 1,
														"December": 1,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Percentage": null
										}
									]
								}
							]
						},
						{
							"Item_ID": 0,
							"Name": "E",
							"Results": [
								{
									"Result_ID": 0,
									"Name": "Environmental  ",
									"Factors": [
										{
											"Factor_ID": 0,
											"Title": "No Environment Trouble by Project",
											"Unit": "case",
											"Target": "Zero",
											"Planned": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 0,
														"May": 0,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 0,
														"October": 0,
														"November": 0,
														"December": 0,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Actual": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 0,
														"May": 0,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 0,
														"October": 0,
														"November": 0,
														"December": 0,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Percentage": null
										}
									]
								}
							]
						},
						{
							"Item_ID": 0,
							"Name": "Q",
							"Results": [
								{
									"Result_ID": 0,
									"Name": "Quality Trouble ",
									"Factors": [
										{
											"Factor_ID": 0,
											"Title": "No Quality trouble causes by Project",
											"Unit": "theme",
											"Target": "Zero",
											"Planned": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 0,
														"May": 0,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 0,
														"October": 0,
														"November": 0,
														"December": 0,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Actual": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 0,
														"May": 0,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 0,
														"October": 0,
														"November": 0,
														"December": 0,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Percentage": null
										}
									]
								}
							]
						},
						{
							"Item_ID": 0,
							"Name": "C",
							"Results": [
								{
									"Result_ID": 0,
									"Name": "1. Overtime Control ",
									"Factors": [
										{
											"Factor_ID": 0,
											"Title": "As Budget Equipment ",
											"Unit": "MRp.",
											"Target": "As budget",
											"Planned": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 45.2,
														"February": 25.01,
														"March": 25.01,
														"April": 131.31,
														"May": 36.33,
														"June": 36.33,
														"July": 25.01,
														"August": 25.01,
														"September": 24.3,
														"October": 19.72,
														"November": 19.72,
														"December": 37.9,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Actual": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 46,
														"February": 29.9,
														"March": 22.3,
														"April": 129.5,
														"May": 23,
														"June": 26.53,
														"July": 20.08,
														"August": 22.21,
														"September": 18.56,
														"October": 20.47,
														"November": 5.59,
														"December": 37.9,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 1.0176991150442478,
														"February": 1.1955217912834866,
														"March": 0.8916433426629348,
														"April": 0.9862158251465997,
														"May": 0.63308560418387,
														"June": 0.7302504816955685,
														"July": 0.8028788484606156,
														"August": 0.8880447820871651,
														"September": 0.7637860082304526,
														"October": 1.0380324543610548,
														"November": 0.28346855983772823,
														"December": 1,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Percentage": null
										}
									]
								},
								{
									"Result_ID": 0,
									"Name": "2. Investment Budget ",
									"Factors": [
										{
											"Factor_ID": 0,
											"Title": "As Budget Equipment + PC Tooling",
											"Unit": "MRp.",
											"Target": "As budget",
											"Planned": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 11300,
														"February": 12950,
														"March": 18100,
														"April": 52250,
														"May": 29371.18,
														"June": 27625,
														"July": 25083.375,
														"August": 33742.65,
														"September": 45451,
														"October": 45006.281,
														"November": 62445.82,
														"December": 85846.6,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 3849.631278,
														"February": 8155.767809,
														"March": 11515.0009,
														"April": 10557.33218,
														"May": 26539.48068,
														"June": 52771.5155,
														"July": 14631.618,
														"August": 18788.25,
														"September": 12561.321,
														"October": 14490.673,
														"November": 58926.747873,
														"December": 117950.967441,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 3849.631278,
														"February": 8155.767809,
														"March": 24845.466424,
														"April": 3817.737918,
														"May": 22891.532109,
														"June": 16148.242189,
														"July": 10880.990547,
														"August": 45438.802769,
														"September": 16159.341525,
														"October": 23180.58799,
														"November": 50975.157609,
														"December": 98395.047223,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 3849.631278,
														"February": 8155.767809,
														"March": 24845.466424,
														"April": 3817.737918,
														"May": 22891.532109,
														"June": 16148.242189,
														"July": 17310.771096,
														"August": 38423.333329,
														"September": 16812.631137,
														"October": 13653.69846,
														"November": 45284.288994,
														"December": 89036.610755,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Actual": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 3849.631278,
														"February": 8155.767809,
														"March": 24797.767424,
														"April": 3817.737918,
														"May": 22497.132109,
														"June": 16148.242089,
														"July": 17280.443642,
														"August": 38423.333329,
														"September": 18100.369706,
														"October": 15963.568724,
														"November": 52245.900008,
														"December": 75900.934,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0.34067533433628316,
														"February": 0.62978902,
														"March": 1.3700423991160222,
														"April": 0.07306675441148325,
														"May": 0.7659594237957071,
														"June": 0.5845517498280542,
														"July": 0.6889201968235933,
														"August": 1.1387171229586295,
														"September": 0.3982391961893028,
														"October": 0.35469646390022763,
														"November": 0.8366596836745838,
														"December": 0.8841460698501745,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 1,
														"February": 1,
														"March": 2.1535184963815333,
														"April": 0.36161956949999086,
														"May": 0.8476854683126376,
														"June": 0.30600300059603175,
														"July": 1.181034362843535,
														"August": 2.045072496320839,
														"September": 1.4409606844694123,
														"October": 1.101644397330614,
														"November": 0.8866245278052899,
														"December": 0.6434956460867202,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 1,
														"February": 1,
														"March": 0.998080172890056,
														"April": 1,
														"May": 0.9827709216612487,
														"June": 0.9999999938073755,
														"July": 1.5881314819048706,
														"August": 0.845606199712942,
														"September": 1.1201180244873872,
														"October": 0.6886610784371221,
														"November": 1.0249286605202304,
														"December": 0.7713897817232617,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 1,
														"February": 1,
														"March": 0.998080172890056,
														"April": 1,
														"May": 0.9827709216612487,
														"June": 0.9999999938073755,
														"July": 0.9982480587472495,
														"August": 1,
														"September": 1.0765935182010888,
														"October": 1.1691754267729741,
														"November": 1.1537312646097277,
														"December": 0.8524688143044309,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Percentage": null
										}
									]
								}
							]
						},
						{
							"Item_ID": 0,
							"Name": "D",
							"Results": [
								{
									"Result_ID": 0,
									"Name": "1 Project 2023 & Additional\n Project",
									"Factors": [
										{
											"Factor_ID": 0,
											"Title": "On Time On Schedule",
											"Unit": "item",
											"Target": "As schedule",
											"Planned": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 2,
														"April": 5,
														"May": 3,
														"June": 2,
														"July": 2,
														"August": 1,
														"September": 8,
														"October": 4,
														"November": 5,
														"December": 23,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 2,
														"April": 5,
														"May": 4,
														"June": 2,
														"July": 2,
														"August": 2,
														"September": 3,
														"October": 4,
														"November": 5,
														"December": 47,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Actual": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 1,
														"April": 8,
														"May": 3,
														"June": 0,
														"July": 3,
														"August": 0,
														"September": 3,
														"October": 4,
														"November": 4,
														"December": 10,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 1,
														"April": 8,
														"May": 3,
														"June": 0,
														"July": 3,
														"August": 1,
														"September": 4,
														"October": 6,
														"November": 5,
														"December": 12,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 1,
														"February": 1,
														"March": 0.5,
														"April": 1.5,
														"May": 1,
														"June": 0,
														"July": 1.3333333333333333,
														"August": 0,
														"September": 0.4444444444444444,
														"October": 1,
														"November": 0.8333333333333334,
														"December": 0.4583333333333333,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 1,
														"February": 1,
														"March": 0.5,
														"April": 1.5,
														"May": 0.8,
														"June": 0,
														"July": 1.3333333333333333,
														"August": 0.5,
														"September": 1.25,
														"October": 1.4,
														"November": 1,
														"December": 0.2708333333333333,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Percentage": null
										}
									]
								},
								{
									"Result_ID": 0,
									"Name": "a. Capacity",
									"Factors": [
										{
											"Factor_ID": 0,
											"Title": "On Time On Schedule",
											"Unit": "item",
											"Target": "As schedule",
											"Planned": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 2,
														"April": 0,
														"May": 0,
														"June": 0,
														"July": 0,
														"August": 1,
														"September": 0,
														"October": 0,
														"November": 1,
														"December": 3,
														"YTD": 0,
														"Remarks": "Continue Y24 : 3 project",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 2,
														"April": 0,
														"May": 0,
														"June": 0,
														"July": 0,
														"August": 1,
														"September": 0,
														"October": 0,
														"November": 1,
														"December": 3,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Actual": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 1,
														"April": 1,
														"May": 0,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 0,
														"October": 0,
														"November": 1,
														"December": 1,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 1,
														"February": 1,
														"March": 0.6666666666666666,
														"April": 2,
														"May": 1,
														"June": 1,
														"July": 1,
														"August": 0,
														"September": 1,
														"October": 1,
														"November": 1,
														"December": 0.3333333333333333,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 1,
														"February": 1,
														"March": 0.6666666666666666,
														"April": 2,
														"May": 1,
														"June": 1,
														"July": 1,
														"August": 0,
														"September": 1,
														"October": 1,
														"November": 1,
														"December": 0.3333333333333333,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Percentage": null
										}
									]
								},
								{
									"Result_ID": 0,
									"Name": "b. Product Mix Change",
									"Factors": [
										{
											"Factor_ID": 0,
											"Title": "On Time On Schedule",
											"Unit": "item",
											"Target": "As schedule",
											"Planned": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 0,
														"May": 0,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 0,
														"October": 0,
														"November": 0,
														"December": 0,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 0,
														"May": 0,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 0,
														"October": 0,
														"November": 0,
														"December": 1,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Actual": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 0,
														"May": 0,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 0,
														"October": 0,
														"November": 0,
														"December": 0,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 1,
														"February": 1,
														"March": 1,
														"April": 1,
														"May": 1,
														"June": 1,
														"July": 1,
														"August": 1,
														"September": 1,
														"October": 1,
														"November": 1,
														"December": 1,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 1,
														"February": 1,
														"March": 1,
														"April": 1,
														"May": 1,
														"June": 1,
														"July": 1,
														"August": 1,
														"September": 1,
														"October": 1,
														"November": 1,
														"December": 0,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Percentage": null
										}
									]
								},
								{
									"Result_ID": 0,
									"Name": "c. Enliten",
									"Factors": [
										{
											"Factor_ID": 0,
											"Title": "On Time On Schedule",
											"Unit": "item",
											"Target": "As schedule",
											"Planned": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 0,
														"May": 1,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 1,
														"October": 1,
														"November": 2,
														"December": 1,
														"YTD": 0,
														"Remarks": "Continue Y24 : 1 project",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 0,
														"May": 1,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 0,
														"October": 1,
														"November": 2,
														"December": 2,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Actual": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 0,
														"May": 1,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 1,
														"October": 1,
														"November": 1,
														"December": 1,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 1,
														"February": 1,
														"March": 1,
														"April": 1,
														"May": 1,
														"June": 1,
														"July": 1,
														"August": 1,
														"September": 1,
														"October": 1,
														"November": 0.5,
														"December": 1,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 1,
														"February": 1,
														"March": 1,
														"April": 1,
														"May": 1,
														"June": 1,
														"July": 1,
														"August": 1,
														"September": 2,
														"October": 1,
														"November": 0.5,
														"December": 0.5,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Percentage": null
										}
									]
								},
								{
									"Result_ID": 0,
									"Name": "d. Smart ",
									"Factors": [
										{
											"Factor_ID": 0,
											"Title": "On Time On Schedule",
											"Unit": "item",
											"Target": "As schedule",
											"Planned": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 0,
														"May": 0,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 1,
														"October": 0,
														"November": 0,
														"December": 0,
														"YTD": 0,
														"Remarks": "Continue Y24 : 1 project",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 0,
														"May": 0,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 0,
														"October": 0,
														"November": 0,
														"December": 1,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Actual": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 0,
														"May": 0,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 0,
														"October": 0,
														"November": 0,
														"December": 0,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 1,
														"February": 1,
														"March": 1,
														"April": 1,
														"May": 1,
														"June": 1,
														"July": 1,
														"August": 1,
														"September": 0,
														"October": 1,
														"November": 1,
														"December": 1,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 1,
														"February": 1,
														"March": 1,
														"April": 1,
														"May": 1,
														"June": 1,
														"July": 1,
														"August": 1,
														"September": 1,
														"October": 1,
														"November": 1,
														"December": 0,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Percentage": null
										}
									]
								},
								{
									"Result_ID": 0,
									"Name": "e. Quality Improvement",
									"Factors": [
										{
											"Factor_ID": 0,
											"Title": "On Time On Schedule",
											"Unit": "item",
											"Target": "As schedule",
											"Planned": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 2,
														"May": 0,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 3,
														"October": 1,
														"November": 0,
														"December": 5,
														"YTD": 0,
														"Remarks": "Continue Y24 : 1 project",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 2,
														"May": 1,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 3,
														"October": 1,
														"November": 0,
														"December": 16,
														"YTD": 0,
														"Remarks": "Cancel : 1 project",
														"MinipapID": null
													}
												]
											},
											"Actual": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 2,
														"May": 0,
														"June": 0,
														"July": 1,
														"August": 0,
														"September": 1,
														"October": 0,
														"November": 0,
														"December": 4,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 2,
														"May": 0,
														"June": 0,
														"July": 1,
														"August": 0,
														"September": 2,
														"October": 0,
														"November": 1,
														"December": 5,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 1,
														"February": 1,
														"March": 1,
														"April": 1,
														"May": 1,
														"June": 1,
														"July": 2,
														"August": 1,
														"September": 0.3333333333333333,
														"October": 0,
														"November": 1,
														"December": 0.8,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 1,
														"February": 1,
														"March": 1,
														"April": 1,
														"May": 0,
														"June": 1,
														"July": 2,
														"August": 1,
														"September": 0.75,
														"October": 0,
														"November": 2,
														"December": 0.35294117647058826,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Percentage": null
										}
									]
								},
								{
									"Result_ID": 0,
									"Name": "f. Cost Improvement",
									"Factors": [
										{
											"Factor_ID": 0,
											"Title": "On Time On Schedule",
											"Unit": "item",
											"Target": "As schedule",
											"Planned": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 0,
														"May": 0,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 0,
														"October": 0,
														"November": 1,
														"December": 0,
														"YTD": 0,
														"Remarks": "Finish",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 0,
														"May": 0,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 0,
														"October": 0,
														"November": 1,
														"December": 1,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Actual": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 0,
														"May": 0,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 0,
														"October": 1,
														"November": 0,
														"December": 0,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 0,
														"May": 0,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 0,
														"October": 2,
														"November": 0,
														"December": 0,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 1,
														"February": 1,
														"March": 1,
														"April": 1,
														"May": 1,
														"June": 1,
														"July": 1,
														"August": 1,
														"September": 1,
														"October": 2,
														"November": 0,
														"December": 1,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 1,
														"February": 1,
														"March": 1,
														"April": 1,
														"May": 1,
														"June": 1,
														"July": 1,
														"August": 1,
														"September": 1,
														"October": 3,
														"November": 0,
														"December": 0,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Percentage": null
										}
									]
								},
								{
									"Result_ID": 0,
									"Name": "g. QA Equipment",
									"Factors": [
										{
											"Factor_ID": 0,
											"Title": "On Time On Schedule",
											"Unit": "item",
											"Target": "As schedule",
											"Planned": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 1,
														"May": 0,
														"June": 2,
														"July": 1,
														"August": 0,
														"September": 0,
														"October": 1,
														"November": 0,
														"December": 6,
														"YTD": 0,
														"Remarks": "Continue Y24 : 2 project",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 1,
														"May": 0,
														"June": 2,
														"July": 1,
														"August": 0,
														"September": 0,
														"October": 1,
														"November": 0,
														"December": 7,
														"YTD": 0,
														"Remarks": "Cancel : 1 project",
														"MinipapID": null
													}
												]
											},
											"Actual": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 1,
														"May": 1,
														"June": 0,
														"July": 2,
														"August": 0,
														"September": 1,
														"October": 1,
														"November": 1,
														"December": 1,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 1,
														"May": 1,
														"June": 0,
														"July": 2,
														"August": 0,
														"September": 1,
														"October": 2,
														"November": 1,
														"December": 1,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 1,
														"February": 1,
														"March": 1,
														"April": 1,
														"May": 2,
														"June": 0,
														"July": 1.5,
														"August": 1,
														"September": 2,
														"October": 1,
														"November": 2,
														"December": 0.16666666666666666,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 1,
														"February": 1,
														"March": 1,
														"April": 1,
														"May": 2,
														"June": 0,
														"July": 1.5,
														"August": 1,
														"September": 2,
														"October": 2,
														"November": 2,
														"December": 0.14285714285714285,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Percentage": null
										}
									]
								},
								{
									"Result_ID": 0,
									"Name": "h. Environment - Safety",
									"Factors": [
										{
											"Factor_ID": 0,
											"Title": "On Time On Schedule",
											"Unit": "item",
											"Target": "As schedule",
											"Planned": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 0,
														"May": 1,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 2,
														"October": 1,
														"November": 1,
														"December": 6,
														"YTD": 0,
														"Remarks": "Continue Y24 : 7 project",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 0,
														"May": 1,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 0,
														"October": 1,
														"November": 0,
														"December": 10,
														"YTD": 0,
														"Remarks": "Cancel : 1 project",
														"MinipapID": null
													}
												]
											},
											"Actual": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 1,
														"May": 0,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 0,
														"October": 1,
														"November": 0,
														"December": 2,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 1,
														"May": 0,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 0,
														"October": 1,
														"November": 0,
														"December": 2,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 1,
														"February": 1,
														"March": 1,
														"April": 2,
														"May": 0,
														"June": 1,
														"July": 1,
														"August": 1,
														"September": 0,
														"October": 1,
														"November": 0,
														"December": 0.3333333333333333,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 1,
														"February": 1,
														"March": 1,
														"April": 2,
														"May": 0,
														"June": 1,
														"July": 1,
														"August": 1,
														"September": 1,
														"October": 1,
														"November": 1,
														"December": 0.2,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Percentage": null
										}
									]
								},
								{
									"Result_ID": 0,
									"Name": "i. Maintenance of Bussines",
									"Factors": [
										{
											"Factor_ID": 0,
											"Title": "On Time On Schedule",
											"Unit": "item",
											"Target": "As schedule",
											"Planned": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 2,
														"May": 1,
														"June": 0,
														"July": 1,
														"August": 0,
														"September": 1,
														"October": 0,
														"November": 0,
														"December": 2,
														"YTD": 0,
														"Remarks": "Continue Y24 : 1 project",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 2,
														"May": 1,
														"June": 0,
														"July": 1,
														"August": 0,
														"September": 0,
														"October": 0,
														"November": 1,
														"December": 6,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Actual": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 3,
														"May": 1,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 0,
														"October": 0,
														"November": 1,
														"December": 1,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 3,
														"May": 1,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 0,
														"October": 0,
														"November": 1,
														"December": 2,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 1,
														"February": 1,
														"March": 1,
														"April": 1.3333333333333333,
														"May": 1,
														"June": 1,
														"July": 0,
														"August": 1,
														"September": 0,
														"October": 1,
														"November": 2,
														"December": 0.5,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 1,
														"February": 1,
														"March": 1,
														"April": 1.3333333333333333,
														"May": 1,
														"June": 1,
														"July": 0,
														"August": 1,
														"September": 1,
														"October": 1,
														"November": 1,
														"December": 0.3333333333333333,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Percentage": null
										}
									]
								},
								{
									"Result_ID": 0,
									"Name": "j. General Improvement",
									"Factors": [
										{
											"Factor_ID": 0,
											"Title": "On Time On Schedule",
											"Unit": "item",
											"Target": "As schedule",
											"Planned": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 0,
														"May": 0,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 0,
														"October": 0,
														"November": 0,
														"December": 0,
														"YTD": 0,
														"Remarks": "Finish",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 0,
														"May": 0,
														"June": 0,
														"July": 0,
														"August": 1,
														"September": 0,
														"October": 0,
														"November": 0,
														"December": 0,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Actual": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 0,
														"April": 0,
														"May": 0,
														"June": 0,
														"July": 0,
														"August": 1,
														"September": 0,
														"October": 0,
														"November": 0,
														"December": 0,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 1,
														"February": 1,
														"March": 1,
														"April": 1,
														"May": 1,
														"June": 1,
														"July": 1,
														"August": 2,
														"September": 1,
														"October": 1,
														"November": 1,
														"December": 1,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													},
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 1,
														"February": 1,
														"March": 1,
														"April": 1,
														"May": 1,
														"June": 1,
														"July": 1,
														"August": 1,
														"September": 1,
														"October": 1,
														"November": 1,
														"December": 1,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Percentage": null
										}
									]
								},
								{
									"Result_ID": 0,
									"Name": "k. Additional Project",
									"Factors": [
										{
											"Factor_ID": 0,
											"Title": "As Requested",
											"Unit": "item",
											"Target": "As Request",
											"Planned": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 26,
														"February": 7,
														"March": 0,
														"April": 6,
														"May": 0,
														"June": 10,
														"July": 4,
														"August": 3,
														"September": 4,
														"October": 4,
														"November": 1,
														"December": 2,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Actual": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 26,
														"February": 7,
														"March": 0,
														"April": 6,
														"May": 0,
														"June": 10,
														"July": 4,
														"August": 3,
														"September": 4,
														"October": 4,
														"November": 1,
														"December": 2,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Percentage": null
										}
									]
								},
								{
									"Result_ID": 0,
									"Name": "2 Maintenance PC & Server M/C",
									"Factors": [
										{
											"Factor_ID": 0,
											"Title": "On Schedule\n(PM Schedule)",
											"Unit": "unit",
											"Target": "As schedule",
											"Planned": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 36,
														"February": 36,
														"March": 37,
														"April": 36,
														"May": 36,
														"June": 37,
														"July": 36,
														"August": 36,
														"September": 37,
														"October": 37,
														"November": 36,
														"December": 37,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Actual": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 36,
														"February": 36,
														"March": 37,
														"April": 36,
														"May": 36,
														"June": 37,
														"July": 36,
														"August": 36,
														"September": 37,
														"October": 37,
														"November": 36,
														"December": 37,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Percentage": null
										}
									]
								},
								{
									"Result_ID": 0,
									"Name": "3 Back Up Program Server",
									"Factors": [
										{
											"Factor_ID": 0,
											"Title": "On Schedule\n(PM Schedule)",
											"Unit": "Unit",
											"Target": "As schedule",
											"Planned": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 109,
														"February": 0,
														"March": 0,
														"April": 0,
														"May": 0,
														"June": 109,
														"July": 0,
														"August": 0,
														"September": 0,
														"October": 0,
														"November": 0,
														"December": 0,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Actual": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 109,
														"February": 0,
														"March": 0,
														"April": 109,
														"May": 0,
														"June": 0,
														"July": 0,
														"August": 0,
														"September": 0,
														"October": 0,
														"November": 0,
														"December": 0,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Percentage": null
										}
									]
								},
								{
									"Result_ID": 0,
									"Name": "4 Trouble Eng IT",
									"Factors": [
										{
											"Factor_ID": 0,
											"Title": "Max. 8 case / month",
											"Unit": "Unit",
											"Target": "As schedule",
											"Planned": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 8,
														"February": 8,
														"March": 8,
														"April": 8,
														"May": 8,
														"June": 8,
														"July": 8,
														"August": 8,
														"September": 8,
														"October": 8,
														"November": 8,
														"December": 8,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Actual": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 4,
														"February": 4,
														"March": 13,
														"April": 8,
														"May": 11,
														"June": 5,
														"July": 8,
														"August": 6,
														"September": 3,
														"October": 5,
														"November": 7,
														"December": 4,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Percentage": null
										}
									]
								}
							]
						},
						{
							"Item_ID": 0,
							"Name": "PEOPLE",
							"Results": [
								{
									"Result_ID": 0,
									"Name": "Ownership and Level up ",
									"Factors": [
										{
											"Factor_ID": 0,
											"Title": "Kaizen Activity (SGK)",
											"Unit": "Theme",
											"Target": "By schedule",
											"Planned": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 1,
														"April": 0,
														"May": 0,
														"June": 1,
														"July": 0,
														"August": 0,
														"September": 1,
														"October": 0,
														"November": 0,
														"December": 1,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Actual": {
												"Minipap_ID": 0,
												"Monthly": [
													{
														"Monthly_ID": 0,
														"MiniPAP": null,
														"January": 0,
														"February": 0,
														"March": 1,
														"April": 0,
														"May": 0,
														"June": 1,
														"July": 0,
														"August": 0,
														"September": 1,
														"October": 0,
														"November": 0,
														"December": 1,
														"YTD": 0,
														"Remarks": "",
														"MinipapID": null
													}
												]
											},
											"Percentage": null
										}
									]
								}
							]
						}
					],
					"Attendance": {
						"Year": 2023,
						"Planned": {
							"Monthly_ID": 0,
							"MiniPAP": null,
							"January": 0.94,
							"February": 0.94,
							"March": 0.94,
							"April": 0.94,
							"May": 0.94,
							"June": 0,
							"July": 0.94,
							"August": 0.94,
							"September": 0,
							"October": 0.94,
							"November": 0.94,
							"December": 0.94,
							"YTD": 0,
							"Remarks": "",
							"MinipapID": null
						},
						"Actual": {
							"Monthly_ID": 0,
							"MiniPAP": null,
							"January": 0.9469,
							"February": 0.9412,
							"March": 0.9412,
							"April": 0.9765,
							"May": 0.9144,
							"June": 0,
							"July": 0.9441,
							"August": 0.9251,
							"September": 0,
							"October": 0.9517,
							"November": 0.94,
							"December": 0.9564,
							"YTD": 0,
							"Remarks": "",
							"MinipapID": null
						},
						"Cuti": {
							"Monthly_ID": 0,
							"MiniPAP": null,
							"January": 0.0088,
							"February": 0.0082,
							"March": 0.0117,
							"April": 0.0023,
							"May": 0.017,
							"June": 0,
							"July": 0.01,
							"August": 0.014,
							"September": 0,
							"October": 0.008,
							"November": 0.0075,
							"December": 0.006,
							"YTD": 0,
							"Remarks": "",
							"MinipapID": null
						},
						"Izin": {
							"Monthly_ID": 0,
							"MiniPAP": null,
							"January": 0.0005,
							"February": 0.0017,
							"March": 0.0012,
							"April": 0,
							"May": 0.0006,
							"June": 0,
							"July": 0.0005,
							"August": 0.0011,
							"September": 0,
							"October": 0.0018,
							"November": 0.0056,
							"December": 0.001,
							"YTD": 0,
							"Remarks": "",
							"MinipapID": null
						},
						"Lain": {
							"Monthly_ID": 0,
							"MiniPAP": null,
							"January": 0.0005,
							"February": 0.0012,
							"March": 0.0012,
							"April": 0,
							"May": 0.0006,
							"June": 0,
							"July": 0.0005,
							"August": 0.0005,
							"September": 0,
							"October": 0.0006,
							"November": 0.0006,
							"December": 0.0006,
							"YTD": 0,
							"Remarks": "",
							"MinipapID": null
						}
					}
				}`
				req, _ := http.NewRequest("POST", "/kpi/yearly/entire", strings.NewReader(yearlyBody))
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				getSuccess(w.Body)
				assert.Equal(t, http.StatusCreated, w.Code)
			})
			t.Run("Get Single", func(t *testing.T){
				req, _ := http.NewRequest("GET", "/kpi/yearly/2023", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			// t.Run("Update Single", func(t *testing.T){
			// 	yearlyBody := `{
			// 		"Name": "a. LKK"
			// 	}`
			// 	req, _ := http.NewRequest("PUT", "/kpi/yearly/1", strings.NewReader(yearlyBody))
				// req.Header.Set("Authorization", "Bearer "+ apiKey)
			// 	w := httptest.NewRecorder()
			// 	router.ServeHTTP(w, req)
			// 	// getSuccess(w.Body)
			// 	assert.Equal(t, http.StatusOK, w.Code)
			// })
			t.Run("Delete Single", func(t *testing.T){
				req, _ := http.NewRequest("DELETE", "/kpi/yearly/2025", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Delete Entire", func(t *testing.T){
				req, _ := http.NewRequest("DELETE", "/kpi/yearly/entire/2023", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Get All", func(t *testing.T){
				req, _ := http.NewRequest("GET", "/kpi/yearly", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
		})
		t.Run("Project", func(t *testing.T){
			t.Run("Add Single", func(t *testing.T){
				projectBody := `{
					"Name": "uhh",
					"INYS": 0,
					"QNYS": 0,
					"IDR": 0,
					"QDR": 0,
					"IPR": 0,
					"QPR": 0,
					"II": 0,
					"QI": 0,
					"IF": 0,
					"QF": 0,
					"IC": 0,
					"QC": 0
				}`
				req, _ := http.NewRequest("POST", "/kpi/project", strings.NewReader(projectBody))
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusCreated, w.Code)
			})
			t.Run("Get Single", func(t *testing.T){
				req, _ := http.NewRequest("GET", "/kpi/project/1", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Update Single", func(t *testing.T){
				projectBody := `{
					"Name": "test",
					"INYS": 1,
					"QNYS": 1,
					"IDR": 1,
					"QDR": 1,
					"IPR": 1,
					"QPR": 1,
					"II": 1,
					"QI": 1,
					"IF": 1,
					"QF": 1,
					"IC": 1,
					"QC": 1
				}`
				req, _ := http.NewRequest("PUT", "/kpi/project/1", strings.NewReader(projectBody))
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Delete Single", func(t *testing.T){
				req, _ := http.NewRequest("DELETE", "/kpi/project/1", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Get All", func(t *testing.T){
				req, _ := http.NewRequest("GET", "/kpi/project", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
		})
		t.Run("Summary", func(t *testing.T){
			t.Run("Add Single", func(t *testing.T){
				summaryBody := `{
					"IssuedDate": "2024-01-24T00:00:00Z"
				}`
				req, _ := http.NewRequest("POST", "/kpi/summary", strings.NewReader(summaryBody))
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				// fmt.Printf(w.Body.String())
				assert.Equal(t, http.StatusCreated, w.Code)
			})
			t.Run("Add Entire", func(t *testing.T){
				summaryBody := `{
					"Projects": [{
						"Name": "Test"
					},
					{
						"Name": "Test2"
					}],
					"IssuedDate": "2024-01-24T00:00:00Z"
				}`
				req, _ := http.NewRequest("POST", "/kpi/summary/entire", strings.NewReader(summaryBody))
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				// fmt.Printf(w.Body.String())
				assert.Equal(t, http.StatusCreated, w.Code)
			})
			t.Run("Get Single", func(t *testing.T){
				req, _ := http.NewRequest("GET", "/kpi/summary/1", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Update Single", func(t *testing.T){
				summaryBody := `{
					"IssuedDate": "2024-02-24T00:00:00Z"
				}`
				req, _ := http.NewRequest("PUT", "/kpi/summary/1", strings.NewReader(summaryBody))
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Delete Single", func(t *testing.T){
				req, _ := http.NewRequest("DELETE", "/kpi/summary/1", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Delete Entire", func(t *testing.T){
				req, _ := http.NewRequest("DELETE", "/kpi/summary/entire/2", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// fmt.Printf(w.Body.String())
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Get All", func(t *testing.T){
				req, _ := http.NewRequest("GET", "/kpi/summary", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
		})
		t.Run("Masalah", func(t *testing.T){
			t.Run("Add Single", func(t *testing.T){
				masalahBody := `{
					"Masalah": "Test",
					"Why":["cape","pegel"],
					"Tindakan":"tiada"
				}`
				req, _ := http.NewRequest("POST", "/kpi/masalah", strings.NewReader(masalahBody))
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				// fmt.Printf(w.Body.String())
				assert.Equal(t, http.StatusCreated, w.Code)
			})
			t.Run("Get Single", func(t *testing.T){
				req, _ := http.NewRequest("GET", "/kpi/masalah/1", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				fmt.Printf(w.Body.String()) 
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Update Single", func(t *testing.T){
				masalahBody := `{
					"Masalah": "Tests",
					"Why":["cape","pegel"],
					"Tindakan":"tiada"
				}`
				req, _ := http.NewRequest("PUT", "/kpi/masalah/1", strings.NewReader(masalahBody))
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Delete Single", func(t *testing.T){
				req, _ := http.NewRequest("DELETE", "/kpi/masalah/1", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Get All", func(t *testing.T){
				req, _ := http.NewRequest("GET", "/kpi/masalah", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
		})
		t.Run("Analisa", func(t *testing.T){
			t.Run("Add Single", func(t *testing.T){
				analisaBody := `{
					"Year": 2023
				}`
				req, _ := http.NewRequest("POST", "/kpi/analisa", strings.NewReader(analisaBody))
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				// fmt.Printf(w.Body.String())
				assert.Equal(t, http.StatusCreated, w.Code)
			})
			t.Run("Add Entire", func(t *testing.T){
				analisaBody := `{
					"Masalah": [{
						"Masalah": "Test",
						"Why":["cape","pegel"],
						"Tindakan":"tiada"
					}],
					"Year": 2024
				}`
				req, _ := http.NewRequest("POST", "/kpi/analisa/entire", strings.NewReader(analisaBody))
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				// fmt.Printf(w.Body.String())
				assert.Equal(t, http.StatusCreated, w.Code)
			})
			t.Run("Get Single", func(t *testing.T){
				req, _ := http.NewRequest("GET", "/kpi/analisa/2023", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			// t.Run("Update Single", func(t *testing.T){
			// 	analisaBody := `{
			// 		"IssuedDate": "2024-02-24T00:00:00Z"
			// 	}`
			// 	req, _ := http.NewRequest("PUT", "/kpi/analisa/1", strings.NewReader(analisaBody))
				// req.Header.Set("Authorization", "Bearer "+ apiKey)
			// 	w := httptest.NewRecorder()
			// 	router.ServeHTTP(w, req)
			// 	// getSuccess(w.Body)
			// 	assert.Equal(t, http.StatusOK, w.Code)
			// })
			t.Run("Delete Single", func(t *testing.T){
				req, _ := http.NewRequest("DELETE", "/kpi/analisa/2023", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Delete Entire", func(t *testing.T){
				req, _ := http.NewRequest("DELETE", "/kpi/analisa/entire/2024", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// fmt.Printf(w.Body.String())
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
			t.Run("Get All", func(t *testing.T){
				req, _ := http.NewRequest("GET", "/kpi/analisa", nil)
				req.Header.Set("Authorization", "Bearer "+ apiKey)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				// getSuccess(w.Body)
				assert.Equal(t, http.StatusOK, w.Code)
			})
		})	
		t.Run("Logout", func(t *testing.T){
			req, _ := http.NewRequest("POST", "/kpi/user/logout", nil)
			req.Header.Set("Authorization", "Bearer "+ apiKey)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			// getSuccess(w.Body)
			assert.Equal(t, http.StatusOK, w.Code)
		})
	})
}