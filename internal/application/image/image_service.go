package image

import (
	"fmt"
	"gin-gorm-example/internal/model"
	"gin-gorm-example/internal/repository"
	"os"
	"path/filepath"
	"time"
)

type ImageServiceInterface interface {
	GetImageList() ([]model.Image, error)
	UploadImage(filename string, fileData []byte) (*model.Image, error)
}

type ImageService struct {
	ImageRepo *repository.ImageRepository
	UploadDir string // 图片存储目录（如：./uploads/）
	ImageURL  string // 前端访问图片的 URL 前缀（如：/static/images/）
}

func NewImageService(repo *repository.ImageRepository, uploadDir, imageURL string) *ImageService {
	return &ImageService{
		ImageRepo: repo,
		UploadDir: uploadDir,
		ImageURL:  imageURL,
	}
}

func (s *ImageService) GetImageList() ([]model.Image, error) {
	return s.ImageRepo.GetAllImages()
}

func (s *ImageService) UploadImage(filename string, fileData []byte) (*model.Image, error) {
	uniqueName := fmt.Sprintf("%d_%s", time.Now().Unix(), filename)
	savePath := filepath.Join(s.UploadDir, uniqueName)

	// 将文件数据保存到磁盘，此处示例省略错误处理，实际使用中请使用 os.WriteFile 或 ioutil.WriteFile
	// 例如：os.WriteFile(savePath, fileData, 0644)
	// 此处假设文件保存成功
	os.WriteFile(savePath, fileData, 0644)
	imgURL := s.ImageURL + uniqueName
	image := &model.Image{
		ImgName: uniqueName,
		ImgURL:  imgURL,
	}
	if err := s.ImageRepo.CreateImage(image); err != nil {
		return nil, err
	}
	return image, nil
}
