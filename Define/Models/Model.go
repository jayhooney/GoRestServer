package models

type UserInfo struct {
	ID   int
	Name string
}

type AUser struct {
	UserID int `form:"userId" binding:"required"`
}

type AddUser struct {
	UserName string `json:"userName" binding:"required"`
}

type DeleteUser struct {
	UserID int `form:"userId" binding:"required"`
}

type UpdateUser struct {
	UserName string `json:"userName" binding:"required"`
	UserID   int    `json:"userId" binding:"required"`
}
