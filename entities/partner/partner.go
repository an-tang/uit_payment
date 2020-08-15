package repository

import (
	baseRepo "uit_payment/entities/base_repository"
	"uit_payment/model"
)

type PartnerRepository struct {
	baseRepo.BaseRepository
}

func NewPartnerRepository() PartnerRepositoryInterface {
	repo := &PartnerRepository{}
	repo.BaseRepository.Init()
	return repo
}

func (repo *PartnerRepository) FindByKey(key string, obj *model.Partner) error {
	return repo.BaseRepository.DB.
		Where(model.Partner{Key: key}).Last(obj).Error
}
