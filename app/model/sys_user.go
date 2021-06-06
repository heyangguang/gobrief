package model

type UserModel struct {
	UserId int	`gorm:"column:user_id;type:int;primaryKey;autoIncrement" json:"user_id"`
	UserName string	`gorm:"column:user_name;unique;type:varchar(50)" json:"user_name" validate:"required,max=10,min=3,ValidationUserNameFormat"`
	UserPwd string `gorm:"column:user_pwd;type:varchar(50)" json:"user_pwd" `
}

func (*UserModel) TableName() string {
	return "sys_user"
}