package repository

import (
	"gin-gorm-example/internal/model"

	"gorm.io/gorm"
)

type ImageRepository struct {
	DB *gorm.DB
}

func NewImageRepository(db *gorm.DB) *ImageRepository {
	return &ImageRepository{DB: db}
}

func (r *ImageRepository) GetAllImages() ([]model.Image, error) {
	var images []model.Image
	if err := r.DB.Find(&images).Error; err != nil {
		return nil, err
	}
	return images, nil
}

func (r *ImageRepository) CreateImage(image *model.Image) error {
	return r.DB.Create(image).Error
}
