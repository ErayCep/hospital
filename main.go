package main

import (
	"fmt"
	"hospital/db"
	"hospital/handlers"
	"hospital/middleware"
	"hospital/models"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("[ERROR] Failed to load .env file\n")
		return
	}
	var db_user = os.Getenv("POSTGRES_USER")
	var db_name = os.Getenv("POSTGRES_DB")
	var db_password = os.Getenv("POSTGRES_PASSWORD")
	var db_port = os.Getenv("POSTGRES_PORT")
	var db_host = os.Getenv("POSTGRES_HOST")

	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", db_host, db_user, db_password, db_name, db_port)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("[ERROR] Failed to connect to PostgreSQL database\n")
		return
	}

	database.AutoMigrate(&models.Hospital{})
	database.AutoMigrate(&models.Staff{})
	database.AutoMigrate(&models.Polyclinic{})
	database.AutoMigrate(&models.Title{})
	database.AutoMigrate(&models.Skill{})

	var redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "redisPass",
		DB:       0,
	})

	db.HospitalStorage = db.Storage{DB: database, Redis: redisClient}
	handler := handlers.NewHandler(db.HospitalStorage)

	r := gin.Default()

	r.POST("/signup", handler.Signup)
	r.POST("/login", handler.Login)
	r.GET("/validate", middleware.RequireAuth, handler.Validate)

	r.GET("/hospital", handler.GetHospitalsHandler)
	r.GET("/hospital/:id", handler.GetHospitalHandler)
	r.POST("/hospital", middleware.RequireAuth, handler.PostHospitalHandler)
	r.DELETE("/hospital/:id", middleware.RequirePrivileged, handler.DeleteHospitalHandler)

	r.GET("/hospital/:id/polyclinic", handler.GetPolyclinicsHandler)
	r.GET("/hospital/:id/polyclinic/:p_id", handler.GetPolyclinicHandler)
	r.POST("/hospital/:id/polyclinic", handler.PostPolyclinicHandler)

	r.GET("/polyclinic", handler.GetPolyclinicsHandler)
	r.POST("/polyclinic", middleware.RequirePrivileged, handler.PostPolyclinicHandler2)
	r.GET("/polyclinic/:id", handler.GetPolyclinicHandler2)
	r.DELETE("/polyclinic/:id", middleware.RequirePrivileged, handler.DeletePolyclinicHandler)

	r.GET("/polyclinic/:id/staff", handler.GetPolyclinicStaffHandler)
	r.POST("/polyclinic/:id/staff", middleware.RequirePrivileged, handler.PostPolyclinicStaffHandler)
	r.DELETE("/polyclinic/:id/staff/:s_id", middleware.RequirePrivileged, handler.DeletePolyclinicStaffHandler)

	r.Run("0.0.0.0:8080")
}
