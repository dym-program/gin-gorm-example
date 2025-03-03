package image

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImageController struct {
	ImageService ImageServiceInterface
}

func NewImageController(service ImageServiceInterface) *ImageController {
	return &ImageController{ImageService: service}
}

func (ic *ImageController) GetImageList(c *gin.Context) {
	images, err := ic.ImageService.GetImageList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve images"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"images": images})
}

func (ic *ImageController) UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer src.Close()

	fileData := make([]byte, file.Size)
	_, err = src.Read(fileData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	image, err := ic.ImageService.UploadImage(file.Filename, fileData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "image": image})
}
