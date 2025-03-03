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

	if err := os.MkdirAll(s.UploadDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create upload directory: %w", err)
	}

	err := os.WriteFile(savePath, fileData, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to write file: %w", err)
	}

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
