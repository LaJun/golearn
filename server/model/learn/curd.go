package learn

import (
"gin-vue-admin/global"
"github.com/satori/go.uuid"
)

type Curd struct {
	global.GVA_MODEL
	UUID        uuid.UUID      `json:"uuid" gorm:"comment:用户UUID"`                                                    // 用户UUID
	Operator    string         `json:"operator" gorm:"comment:用户登录名"`                                                 // 用户登录名
}
