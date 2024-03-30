package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
	"tugas/auth"
	"tugas/campaign"
	"tugas/handler"
	"tugas/helper"
	"tugas/user"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/tugas?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)

	authService := auth.NewJwtService()
	userService := user.NewService(userRepository)
	campaignService := campaign.NewService(campaignRepository)

	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampagnHandler(campaignService)

	router := gin.Default()
	router.Static("/images/profile", "./images/profile")
	api := router.Group("/api/v1")

	//Router
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/cek_email", userHandler.CekEmail)
	api.POST("/upload_avatar", AuthMiddleware(authService, userService), userHandler.UploadAvatar)

	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.POST("/campaign", AuthMiddleware(authService, userService), campaignHandler.CampaignCreate)
	api.PUT("/campaigns/:id", AuthMiddleware(authService, userService), campaignHandler.UpdateCampaign)

	router.Run()
}

func AuthMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIresponse("Unauthorized", http.StatusUnauthorized, "Error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIresponse("Unauthorized", http.StatusUnauthorized, "Error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIresponse("Unauthorized", http.StatusUnauthorized, "Error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserById(userID)

		if err != nil {
			response := helper.APIresponse("Unauthorized", http.StatusUnauthorized, "Error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)
	}
}
