package main

import (
	"log"
	"os"

	_ "bsd/docs"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Barafiri Backend API documentation
// @version 1.0
// @description This is the official data model for BARAFIRI API documentation.
// @description Barafiri documentation port: 8112
// @description apicore port: 7002
// @description mailer port: 7006
// @description sales service: 7003
// @description public endpoints: 7005 (This includes PUBLIC API, PUBLIC SEARCH)
// @description cart service: 7008
// @description products service: 7007
// @termsOfService http://my-gpi.io/terms/

// @contact.name Endy Apinageri
// @contact.url https://www.endyapina.com
// @contact.email endy.apinageri@my-gpi.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8112
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}
func main() {
	r := gin.Default()
	err := godotenv.Load("conf.env")
	if err != nil {
		panic("Error loading .env file")
	}

	appPort := os.Getenv("port")
	// hostAddress := os.Getenv("host_address")
	// c := controller.NewController()

	// v1 := r.Group("/api/v1")
	// {
	// 	// accounts := v1.Group("/authentication")
	// 	{
	// 		accounts.POST("", c.Registration)
	// 		// accounts.GET("", c.ListAccounts)
	// 		// accounts.POST("", c.AddAccount)
	// 		// accounts.DELETE(":id", c.DeleteAccount)
	// 		// accounts.PATCH(":id", c.UpdateAccount)
	// 		// accounts.POST(":id/images", c.UploadAccountImage)
	// 	}
	// 	//...
	// }
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(appPort)
	log.Println("App running on ports" + appPort)
}

// func auth() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		if len(c.GetHeader("Authorization")) == 0 {
// 			httputil.NewError(c, http.StatusUnauthorized, errors.New("Authorization is required Header"))
// 			c.Abort()
// 		}
// 		c.Next()
// 	}
// }
