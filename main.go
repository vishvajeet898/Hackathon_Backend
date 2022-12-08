package main

import (
	"Hackathon_Backend/api"
	repository "Hackathon_Backend/repository/db"
	"fmt"
	_ "github.com/aws/aws-sdk-go-v2/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	_ "log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		return
	}
	dbUrl := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	DB_URL := "postgres://" + dbName + ":" + dbPassword + "@" + dbUrl

	dialector := postgres.Open(DB_URL)
	db := repository.NewDatabse(dialector)
	db.Connect()

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	routersInit := api.InitRouter()
	gin.SetMode(gin.DebugMode)
	routersInit.Run(port)

	/*res := model.Aws_File{
		AwsLink: "41e6779d-a856-4ec7-a5c2-cdaa7984f6c2",
		Extension: "pdf",
	}

	err := godotenv.Load(".env")
	if err != nil {
		return
	}

	err = utils.S3FileDownloader(res.AwsLink)
	if err != nil {
		fmt.Printf("%v", err)
		return 	}

	err = utils.Decrypt("./recordsTemp/download/"+res.AwsLink, res.AwsLink+res.Extension)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}*/
	//c.FileAttachment("./recordsTemp/download/"+res.AwsLink+res.Extension,res.AwsLink+"."+res.Extension)

	/*	utils.Encrypt("./test.pdf", "test")
		utils.AwsFileUpload("./recordsTemp/upload/test.bin", "test.bin")

		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.Printf("error: %v", err)
			return
		}

		utils.S3FileDownloader(cfg, "healthmate", "ciphertext")
	*/
}
