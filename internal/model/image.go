package model

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	ImgName string `json:"img_name"`
	ImgURL  string `json:"img_url"`
}
