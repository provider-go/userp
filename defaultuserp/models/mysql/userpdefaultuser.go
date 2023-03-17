package mysql

type UserPDefaultUser struct {
	Id         string `json:"id"`
	BelongId   string `json:"belong_id"`   // 所属
	Username   string `json:"username"`    // 用户名
	Password   string `json:"password"`    // 密码
	Nickname   string `json:"nickname"`    // 昵称
	Sex        int    `json:"sex"`         // 性别
	Avatar     string `json:"avatar"`      // 头像
	Phone      string `json:"phone"`       // 电话
	Email      string `json:"email"`       // 邮箱
	Status     int    `json:"status"`      // 账号状态:0(正常)1(禁用)-1(软删除)
	CreateTime string `json:"create_time"` // 创建时间
	UpdateTime string `json:"update_time"` // 更新时间
}
