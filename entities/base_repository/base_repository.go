package repository

import (
	"github.com/jinzhu/gorm"

	"uit_payment/database"
	"uit_payment/model"
)

type BaseRepository struct {
	DB *gorm.DB
}

func (base *BaseRepository) Init() {
	base.DB = database.PG.GetInstance()
	base.DB.LogMode(true)
}

func (base *BaseRepository) Find(id int, obj interface{}) error {
	query := model.BaseModel{}
	query.ID = id
	return base.FindBy(query, obj)
}

func (base *BaseRepository) FindBy(condition interface{}, obj interface{}) error {
	return base.DB.Where(condition).Last(obj).Error
}

func (base *BaseRepository) Create(obj interface{}) error {
	tx := base.DB.Begin()

	if err := tx.Create(obj).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (base *BaseRepository) Update(obj interface{}) error {
	return base.DB.Model(obj).Updates(obj).Error
}

func (base *BaseRepository) Delete(obj interface{}) error {
	return base.DB.Delete(obj).Error
}

func (base *BaseRepository) Count(model interface{}, query interface{}) int {
	var count int
	base.DB.Model(model).Where(query).Count(&count)
	return count
}
