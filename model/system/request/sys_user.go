package request

type Login struct {
	// User login structure
	Username    string         `json:"username"`                                      // 用户名
	Password    string         `json:"password"`                                      // 密码
}
