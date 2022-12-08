package user

import (
	"Hackathon_Backend/model"
	repository "Hackathon_Backend/repository/db"
	"Hackathon_Backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func DownloadFile(c *gin.Context) (interface{}, error) {
	userId := c.GetString("userId")

	if userId == "" {
		return nil, fmt.Errorf("user Id No specified in token")
	}

	var file model.Aws_File

	err := c.BindJSON(&file)
	if err != nil {
		return nil, err
	}
	file.UserId = userId

	res, err := repository.GetOne(file)

	if err != nil {
		return nil, err
	}

	err = utils.S3FileDownloader(res.AwsLink)
	if err != nil {
		return nil, err
	}

	err = utils.Decrypt("./recordsTemp/download/"+res.AwsLink, res.AwsLink+res.Extension)
	if err != nil {
		return nil, err
	}
	c.FileAttachment("./recordsTemp/download/"+res.AwsLink+res.Extension, res.AwsLink+res.Extension)

	return "Download", nil
}
