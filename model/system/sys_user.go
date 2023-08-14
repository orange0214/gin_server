package system

import (
	"project/global"

	"github.com/gofrs/uuid/v5"
)

type SysUser struct {
	global.GVA_MODEL
	UUID        uuid.UUID      `json:"uuid" gorm:"index;comment:用户UUID"` // 用户UUID
	Username    string         `json:"userName" gorm:"index;comment:用户登录名"`
	Password    string         `json:"-"  gorm:"comment:用户登录密码"`
	AuthorityId uint           `json:"authorityId" gorm:"default:888;comment:用户角色ID"` // 用户角色ID
	Authority   SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
	Enable      int            `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
}
