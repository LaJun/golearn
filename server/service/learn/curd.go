package learn

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model/learn"
	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateApi
//@description: 新增基础api
//@param: api model.SysApi
//@return: err error

type CurdService struct {
}

func (curdService *CurdService) Add(api learn.Curd) (err error) {
	if !errors.Is(global.GVA_DB.Where("operator = ? AND uuid = ?", api.Operator, api.UUID).First(&learn.Curd{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同数据")
	}
	return global.GVA_DB.Create(&api).Error
}