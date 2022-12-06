package user

import (
	"Hackathon_Backend/model"
	repository "Hackathon_Backend/repository/db"
	"Hackathon_Backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
	"path/filepath"
)

func UploadFile(c *gin.Context) (interface{}, error) {
	userId := c.GetString("userId")
	if userId == "" {
		return nil, fmt.Errorf("user Id No specified in token")
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return nil, err
	}

	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String()

	if err := c.SaveUploadedFile(file, "/recordsTemp/upload/"+newFileName+extension); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return nil, err
	}

	err = utils.Encrypt("/recordsTemp/upload/"+newFileName+extension, newFileName)
	if err != nil {
		return nil, err
	}
	err = utils.AwsFileUpload("/recordsTemp/upload/"+newFileName+".bin", newFileName)
	if err != nil {
		return nil, err
	}

	record := model.Record{
		UserId:   userId,
		AwsLink:  newFileName,
		RecordId: uuid.New().String(),
	}

	if err := validator.New().Struct(record); err != nil {
		return nil, err
	}
	res, err := repository.CreateOne(&record)
	if err != nil {
		return nil, err
	}

	return res, nil

}