package system

import (
	"errors"
	"fmt"
	"project/global"
	system "project/model/system"

	"project/utils"

	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

type UserService struct{}

func (userService *UserService) Register(u system.SysUser) (userInter system.SysUser, err error) {
	var user system.SysUser
	if !errors.Is(global.GVA_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return userInter, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码hash加密 注册
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.Must(uuid.NewV4())
	err = global.GVA_DB.Create(&u).Error
	return u, err
}

// 用户登录
func (UserService *UserService) Login(u *system.SysUser) (userInter *system.SysUser, err error) {
	// connect to database and judge
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}
	// compare the code and return
	var user system.SysUser
	err = global.GVA_DB.Where("username = ?", u.Username).Preload("Authorities").Preload("Authority").First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
		MenuServiceApp.UserAuthorityDefaultRouter(&user)
	}
	return &user, err

}
