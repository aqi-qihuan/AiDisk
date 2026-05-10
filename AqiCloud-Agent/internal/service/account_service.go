package service

import (
	"context"
	"errors"
	"mime/multipart"
	"time"

	"github.com/aqi/AqiCloud-AgentPan-Go/internal/config"
	"github.com/aqi/AqiCloud-AgentPan-Go/internal/model"
	"github.com/aqi/AqiCloud-AgentPan-Go/internal/util"
	"gorm.io/gorm"
)

type AccountService struct{}

func NewAccountService() *AccountService {
	return &AccountService{}
}

func (s *AccountService) Register(ctx context.Context, req *model.AccountRegisterReq) error {
	db := config.GetDB()
	cfg := config.GetConfig()

	// 校验用户名和密码必填
	if req.Phone == "" && req.Username == "" {
		return errors.New("用户名不能为空")
	}
	if req.Password == "" {
		return errors.New("密码不能为空")
	}

	var exists int64
	if req.Phone != "" {
		db.Model(&model.Account{}).Where("phone = ? AND del = 0", req.Phone).Count(&exists)
	}
	if exists > 0 {
		return errors.New("账号已存在")
	}
	if req.Username != "" {
		db.Model(&model.Account{}).Where("username = ? AND del = 0", req.Username).Count(&exists)
	}
	if exists > 0 {
		return errors.New("账号已存在")
	}

	id := util.NextID()
	account := model.Account{
		ID:        id,
		Username:  req.Username,
		Password:  util.MD5Password(req.Password),
		Phone:     req.Phone,
		AvatarURL: req.AvatarURL,
		Role:      "COMMON",
		Del:       false,
	}
	now := time.Now()
	account.GmtCreate = now
	account.GmtModified = now

	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&account).Error; err != nil {
			return err
		}

		// 创建默认存储空间
		storage := model.Storage{
			ID:        util.NextID(),
			AccountID: id,
			UsedSize:  0,
			TotalSize: cfg.DefaultStorageSize,
		}
		storage.GmtCreate = now
		storage.GmtModified = now
		if err := tx.Create(&storage).Error; err != nil {
			return err
		}

		// 创建根文件夹
		rootFile := model.AccountFile{
			ID:        util.NextID(),
			AccountID: id,
			IsDir:     1,
			ParentID:  cfg.RootParentID,
			FileName:  cfg.RootFolderName,
			Del:       false,
		}
		rootFile.GmtCreate = now
		rootFile.GmtModified = now
		return tx.Create(&rootFile).Error
	})
}

func (s *AccountService) Login(ctx context.Context, req *model.AccountLoginReq) (string, error) {
	db := config.GetDB()

	if req.Phone == "" && req.Username == "" {
		return "", errors.New("请输入用户名或手机号")
	}

	var account model.Account
	query := db.WithContext(ctx).Where("password = ? AND del = ?", util.MD5Password(req.Password), false)
	if req.Phone != "" {
		query = query.Where("phone = ?", req.Phone)
	} else {
		query = query.Where("username = ?", req.Username)
	}

	err := query.First(&account).Error
	if err == gorm.ErrRecordNotFound {
		return "", errors.New("用户名或密码错误")
	}
	if err != nil {
		return "", err
	}

	return util.GenerateLoginToken(account.ID, account.Username), nil
}

func (s *AccountService) UploadAvatar(ctx context.Context, file *multipart.FileHeader, accountID int64) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	url, err := getStoreEngine().UploadAvatar(ctx, src, file.Filename)
	if err != nil {
		return "", err
	}

	db := config.GetDB()
	if err := db.WithContext(ctx).Model(&model.Account{}).
		Where("id = ?", accountID).
		Update("avatar_url", url).Error; err != nil {
		return "", err
	}

	return url, nil
}

func (s *AccountService) UploadAvatarFile(ctx context.Context, file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	return getStoreEngine().UploadAvatar(ctx, src, file.Filename)
}

func (s *AccountService) GetAccountDetail(ctx context.Context, accountID int64) (*model.AccountDTO, error) {
	db := config.GetDB()

	var account model.Account
	if err := db.WithContext(ctx).Where("id = ? AND del = 0", accountID).First(&account).Error; err != nil {
		return nil, err
	}

	var storage model.Storage
	db.WithContext(ctx).Where("account_id = ?", accountID).First(&storage)

	var rootFolder model.AccountFile
	db.WithContext(ctx).Where("account_id = ? AND is_dir = 1 AND parent_id = 0 AND del = 0", accountID).
		Order("id ASC").First(&rootFolder)

	dto := &model.AccountDTO{
		ID:          account.ID,
		Username:    account.Username,
		AvatarURL:   account.AvatarURL,
		Phone:       account.Phone,
		Role:        account.Role,
		Del:         account.Del,
		GmtCreate:   account.GmtCreate.Format("2006-01-02 15:04:05"),
		GmtModified: account.GmtModified.Format("2006-01-02 15:04:05"),
		RootFileID:  rootFolder.ID,
		RootFileName: rootFolder.FileName,
	}
	if storage.ID != 0 {
		dto.Storage = &model.StorageDTO{
			ID:          storage.ID,
			AccountID:   storage.AccountID,
			UsedSize:    storage.UsedSize,
			TotalSize:   storage.TotalSize,
			GmtCreate:   storage.GmtCreate.Format("2006-01-02 15:04:05"),
			GmtModified: storage.GmtModified.Format("2006-01-02 15:04:05"),
		}
	}

	return dto, nil
}

func (s *AccountService) UpdateAccount(ctx context.Context, accountID int64, req *model.AccountUpdateReq) error {
	db := config.GetDB()
	updates := map[string]interface{}{}
	if req.Username != "" {
		updates["username"] = req.Username
	}
	if req.AvatarURL != "" {
		updates["avatar_url"] = req.AvatarURL
	}
	if len(updates) == 0 {
		return nil
	}
	return db.WithContext(ctx).Model(&model.Account{}).
		Where("id = ?", accountID).Updates(updates).Error
}
