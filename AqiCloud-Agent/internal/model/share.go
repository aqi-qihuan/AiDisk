package model

import "time"

type Share struct {
	ID           int64     `gorm:"column:id;primaryKey" json:"id"`
	ShareName    string    `gorm:"column:share_name" json:"shareName"`
	ShareType    string    `gorm:"column:share_type" json:"shareType"`
	ShareDayType int       `gorm:"column:share_day_type" json:"shareDayType"`
	ShareDay     int       `gorm:"column:share_day" json:"shareDay"`
	ShareEndTime time.Time `gorm:"column:share_end_time" json:"shareEndTime"`
	ShareURL     string    `gorm:"column:share_url" json:"shareUrl"`
	ShareCode    string    `gorm:"column:share_code" json:"shareCode"`
	ShareStatus  string    `gorm:"column:share_status" json:"shareStatus"`
	AccountID    int64     `gorm:"column:account_id" json:"accountId"`
	GmtCreate    time.Time `gorm:"column:gmt_create;autoCreateTime" json:"gmtCreate"`
	GmtModified  time.Time `gorm:"column:gmt_modified" json:"gmtModified"`
}

func (Share) TableName() string { return "share" }

type ShareFile struct {
	ID            int64     `gorm:"column:id;primaryKey" json:"id"`
	ShareID       int64     `gorm:"column:share_id" json:"shareId"`
	AccountFileID int64     `gorm:"column:account_file_id" json:"accountFileId"`
	AccountID     int64     `gorm:"column:account_id" json:"accountId"`
	GmtCreate     time.Time `gorm:"column:gmt_create;autoCreateTime" json:"gmtCreate"`
	GmtModified   time.Time `gorm:"column:gmt_modified" json:"gmtModified"`
}

func (ShareFile) TableName() string { return "share_file" }

type ShareDTO struct {
	ID           int64  `json:"id"`
	ShareName    string `json:"shareName"`
	ShareType    string `json:"shareType"`
	ShareDayType int    `json:"shareDayType"`
	ShareDay     int    `json:"shareDay"`
	ShareEndTime string `json:"shareEndTime"`
	ShareURL     string `json:"shareUrl"`
	ShareCode    string `json:"shareCode"`
	ShareStatus  string `json:"shareStatus"`
	AccountID    int64  `json:"accountId"`
	GmtCreate    string `json:"gmtCreate"`
}

type ShareSimpleDTO struct {
	ID              int64       `json:"id"`
	ShareName       string      `json:"shareName"`
	ShareType       string      `json:"shareType"`
	ShareDayType    int         `json:"shareDayType"`
	ShareDay        int         `json:"shareDay"`
	ShareEndTime    string      `json:"shareEndTime"`
	ShareURL        string      `json:"shareUrl"`
	ShareAccountDTO *ShareAccountDTO `json:"shareAccountDTO"`
	ShareToken      string           `json:"shareToken,omitempty"`
	FileCount       int         `json:"fileCount"`
}

type ShareDetailDTO struct {
	ID            int64            `json:"id"`
	ShareName     string           `json:"shareName"`
	ShareType     string           `json:"shareType"`
	ShareDayType  int              `json:"shareDayType"`
	ShareDay      int              `json:"shareDay"`
	ShareEndTime  string           `json:"shareEndTime"`
	ShareURL      string           `json:"shareUrl"`
	ShareAccountDTO *ShareAccountDTO  `json:"shareAccountDTO"`
	FileDTOList   []*AccountFileDTO `json:"fileDTOList"`
}

type ShareAccountDTO struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatarUrl"`
}
