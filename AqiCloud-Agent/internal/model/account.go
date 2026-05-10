package model

import "time"

type Account struct {
	ID          int64     `gorm:"column:id;primaryKey" json:"id"`
	Username    string    `gorm:"column:username" json:"username"`
	Password    string    `gorm:"column:password" json:"-"`
	AvatarURL   string    `gorm:"column:avatar_url" json:"avatarUrl"`
	Phone       string    `gorm:"column:phone" json:"phone"`
	Role        string    `gorm:"column:role" json:"role"`
	Del         bool      `gorm:"column:del" json:"del"`
	GmtCreate   time.Time `gorm:"column:gmt_create;autoCreateTime" json:"gmtCreate"`
	GmtModified time.Time `gorm:"column:gmt_modified;autoUpdateTime" json:"gmtModified"`
}

func (Account) TableName() string { return "account" }

type Storage struct {
	ID          int64     `gorm:"column:id;primaryKey" json:"id"`
	AccountID   int64     `gorm:"column:account_id" json:"accountId"`
	UsedSize    int64     `gorm:"column:used_size" json:"usedSize"`
	TotalSize   int64     `gorm:"column:total_size" json:"totalSize"`
	GmtCreate   time.Time `gorm:"column:gmt_create;autoCreateTime" json:"gmtCreate"`
	GmtModified time.Time `gorm:"column:gmt_modified;autoUpdateTime" json:"gmtModified"`
}

func (Storage) TableName() string { return "storage" }

type StorageDTO struct {
	ID          int64  `json:"id"`
	AccountID   int64  `json:"accountId"`
	UsedSize    int64  `json:"usedSize"`
	TotalSize   int64  `json:"totalSize"`
	GmtCreate   string `json:"gmtCreate"`
	GmtModified string `json:"gmtModified"`
}

type AccountDTO struct {
	ID          int64       `json:"id"`
	Username    string      `json:"username"`
	AvatarURL   string      `json:"avatarUrl"`
	Phone       string      `json:"phone"`
	Role        string      `json:"role"`
	Del         bool        `json:"del"`
	GmtCreate   string      `json:"gmtCreate"`
	GmtModified string      `json:"gmtModified"`
	RootFileID  int64       `json:"rootFileId"`
	RootFileName string     `json:"rootFileName"`
	Storage     *StorageDTO `json:"storageDTO"`
}
