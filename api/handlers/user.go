package handler

import (
	"Hackathon_Backend/service/user"
	"Hackathon_Backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp(c *gin.Context) {
	res, err := user.SignUp(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func SignIn(c *gin.Context) {
	res, err := user.SignInUser(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func AddBasicInfo(c *gin.Context) {
	res, err := user.AddBasicInfo(c)
	fmt.Printf("Err %v\n", err)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func GetBasicInfo(c *gin.Context) {
	res, err := user.GetBasicInfoByUserId(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func UpdateBasicInfo(c *gin.Context) {
	res, err := user.UpdateBasicInfo(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func AddMeasurement(c *gin.Context) {
	res, err := user.AddMeasurement(c)
	fmt.Printf("Err %v\n", err)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func GetMeasurementByType(c *gin.Context) {
	res, err := user.GetMeasurement(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func GetAllMeasurementByDate(c *gin.Context) {
	res, err := user.GetAllMeasurement(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func GetAllMeasurementOfDate(c *gin.Context) {
	res, err := user.GetAllMeasurementByDate(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func UploadFile(c *gin.Context) {
	res, err := user.UploadFile(c)
	fmt.Printf("Err %v\n", err)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func DownloadFile(c *gin.Context) {
	res, err := user.DownloadFile(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func AddVisit(c *gin.Context) {
	res, err := user.AddVisit(c)
	fmt.Printf("Err %v\n", err)

	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}

func GetAllVisit(c *gin.Context) {
	res, err := user.GetAllVisits(c)
	statusCode, data := utils.FormatResponseMessage(res, err, http.StatusOK)
	c.JSON(statusCode, data)
}
