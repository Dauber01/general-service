package table

import "general-service/library/resource"

const UserTableName = "user"

// table user
type User struct {
	ID       uint    `gorm:"column:id"`
	UserName string  `gorm:"column:username"`
	PassWord string  `gorm:"column:password"`
	Email    string  `gorm:"column:email"`
	Age      int     `gorm:"column:age"`
	Sex      string  `gorm:"column:sex"`
	Tel      string  `gorm:"column:tel"`
	Addr     string  `gorm:"column:addr"`
	Card     string  `gorm:"column:card"`
	Married  int     `gorm:"column:married"`
	Salary   float32 `gorm:"column:salary"`
}

// find user info by user name
func SelectUserByName(userName string) (*User, error) {
	var data User
	if err := resource.MysqlClient.Table(UserTableName).
		Where("username = ?", userName).First(&data).Error; err != nil {

		return nil, err
	}
	return &data, nil
}
