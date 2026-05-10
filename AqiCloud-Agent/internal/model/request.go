package model

import (
	"encoding/json"
	"strconv"
)

type AccountLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone"`
}

type AccountRegisterReq struct {
	Username  string `json:"username"`
	Password  string `json:"password" binding:"required"`
	Phone     string `json:"phone"`
	AvatarURL string `json:"avatarUrl"`
}

type AccountUpdateReq struct {
	Username  string `json:"username"`
	AvatarURL string `json:"avatarUrl"`
}

type FolderCreateReq struct {
	FolderName string `json:"folderName" binding:"required"`
	ParentID   int64  `json:"parentId"`
	AccountID  int64  `json:"-"`
}

// stringInt64 is a helper type that unmarshals from JSON string or number
type stringInt64 int64

func (s *stringInt64) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	// Remove quotes if string
	if data[0] == '"' {
		var str string
		if err := json.Unmarshal(data, &str); err != nil {
			return err
		}
		if str == "" || str == "null" || str == "undefined" {
			*s = 0
			return nil
		}
		v, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return err
		}
		*s = stringInt64(v)
		return nil
	}
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*s = stringInt64(v)
	return nil
}

// stringInt64Slice unmarshals from JSON array of strings or numbers
type stringInt64Slice []int64

func (s *stringInt64Slice) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		return nil
	}
	// Try unmarshaling as []string first
	var strSlice []string
	if err := json.Unmarshal(data, &strSlice); err == nil {
		result := make([]int64, 0, len(strSlice))
		for _, str := range strSlice {
			if str == "" {
				continue
			}
			v, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				return err
			}
			result = append(result, v)
		}
		*s = result
		return nil
	}
	// Fallback: try []int64
	var intSlice []int64
	if err := json.Unmarshal(data, &intSlice); err != nil {
		return err
	}
	*s = intSlice
	return nil
}

func (s *FolderCreateReq) UnmarshalJSON(data []byte) error {
	type Alias FolderCreateReq
	aux := &struct {
		ParentID stringInt64 `json:"parentId"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	s.ParentID = int64(aux.ParentID)
	return nil
}

type FileUpdateReq struct {
	FileID      int64  `json:"fileId"`
	NewFileName string `json:"newFilename" binding:"required"`
	AccountID   int64  `json:"-"`
}

func (s *FileUpdateReq) UnmarshalJSON(data []byte) error {
	type Alias FileUpdateReq
	aux := &struct {
		FileID stringInt64 `json:"fileId"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	s.FileID = int64(aux.FileID)
	return nil
}

type FileBatchReq struct {
	FileIDs        stringInt64Slice `json:"fileIds" binding:"required"`
	TargetParentID int64            `json:"targetParentId"`
	AccountID      int64            `json:"-"`
}

func (s *FileBatchReq) UnmarshalJSON(data []byte) error {
	type Alias FileBatchReq
	aux := &struct {
		TargetParentID stringInt64 `json:"targetParentId"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	s.TargetParentID = int64(aux.TargetParentID)
	return nil
}

type FileDelReq struct {
	FileIDs   stringInt64Slice `json:"fileIds" binding:"required"`
	AccountID int64            `json:"-"`
}

type FileSecondUploadReq struct {
	FileName   string `json:"filename" binding:"required"`
	Identifier string `json:"identifier" binding:"required"`
	ParentID   int64  `json:"parentId"`
	AccountID  int64  `json:"-"`
}

func (s *FileSecondUploadReq) UnmarshalJSON(data []byte) error {
	type Alias FileSecondUploadReq
	aux := &struct {
		ParentID stringInt64 `json:"parentId"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	s.ParentID = int64(aux.ParentID)
	return nil
}

type FileChunkInitTaskReq struct {
	FileName   string      `json:"fileName" binding:"required"`
	TotalSize  stringInt64 `json:"totalSize"`
	ChunkSize  stringInt64 `json:"chunkSize"`
	Identifier string      `json:"identifier" binding:"required"`
	AccountID  int64       `json:"-"`
}

type FileChunkMergeReq struct {
	Identifier string `json:"identifier" binding:"required"`
	ParentID   int64  `json:"parentId"`
	AccountID  int64  `json:"-"`
}

func (s *FileChunkMergeReq) UnmarshalJSON(data []byte) error {
	type Alias FileChunkMergeReq
	aux := &struct {
		ParentID stringInt64 `json:"parentId"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	s.ParentID = int64(aux.ParentID)
	return nil
}

type FileDownloadReq struct {
	FileIDs   stringInt64Slice `json:"fileIds" binding:"required"`
	AccountID int64            `json:"-"`
}

type ShareCreateReq struct {
	ShareName    string           `json:"shareName" binding:"required"`
	ShareType    string           `json:"shareType" binding:"required"`
	ShareDayType int              `json:"shareDayType" binding:"gte=0"`
	FileIDs      stringInt64Slice `json:"fileIds" binding:"required"`
	AccountID    int64            `json:"-"`
}

type ShareCancelReq struct {
	ShareIDs  stringInt64Slice `json:"shareIds" binding:"required"`
	AccountID int64            `json:"-"`
}

type ShareCheckReq struct {
	ShareID   int64  `json:"shareId"`
	ShareCode string `json:"shareCode" binding:"required"`
}

func (s *ShareCheckReq) UnmarshalJSON(data []byte) error {
	type Alias ShareCheckReq
	aux := &struct {
		ShareID stringInt64 `json:"shareId"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	s.ShareID = int64(aux.ShareID)
	return nil
}

type ShareFileQueryReq struct {
	ParentID int64 `json:"parentId"`
	ShareID  int64 `json:"-"`
}

func (s *ShareFileQueryReq) UnmarshalJSON(data []byte) error {
	type Alias ShareFileQueryReq
	aux := &struct {
		ParentID stringInt64 `json:"parentId"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	s.ParentID = int64(aux.ParentID)
	return nil
}

type ShareFileTransferReq struct {
	ShareID        int64            `json:"-"`
	FileIDs        stringInt64Slice `json:"fileIds" binding:"required"`
	TargetParentID int64            `json:"parentId" binding:"required"`
	AccountID      int64            `json:"-"`
}

func (s *ShareFileTransferReq) UnmarshalJSON(data []byte) error {
	type Alias ShareFileTransferReq
	aux := &struct {
		TargetParentID stringInt64 `json:"parentId"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	s.TargetParentID = int64(aux.TargetParentID)
	return nil
}

type RecycleDelReq struct {
	FileIDs   stringInt64Slice `json:"fileIds" binding:"required"`
	AccountID int64            `json:"-"`
}

type RecycleRestoreReq struct {
	FileIDs   stringInt64Slice `json:"fileIds" binding:"required"`
	AccountID int64            `json:"-"`
}
