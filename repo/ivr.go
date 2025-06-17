package repo

import (
	"gorm.io/gorm"

	"github.com/cuongpiger/golang/models"
)

type IVRRepo struct {
	db *gorm.DB
}

func NewIVRRepo(db *gorm.DB) *IVRRepo {
	return &IVRRepo{db: db}
}

func (r *IVRRepo) GetIVRByPhoneNumber(phoneNumber string) (*models.IVR, error) {
	var ivr models.IVR
	if err := r.db.Where("phone_number = ?", phoneNumber).First(&ivr).Error; err != nil {
		return nil, err
	}
	
	return &ivr, nil
}

func (r *IVRRepo) CreateIVR(ivr *models.IVR) error {
	if err := r.db.Create(ivr).Error; err != nil {
		return err
	}
	return nil
}