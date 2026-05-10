package controller

import (
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/aqi/AqiCloud-AgentPan-Go/internal/model"
	"github.com/aqi/AqiCloud-AgentPan-Go/internal/service"
	"github.com/gin-gonic/gin"
)

type FileController struct {
	svc *service.FileService
}

func NewFileController() *FileController {
	return &FileController{svc: service.NewFileService()}
}

func (c *FileController) Register(r *gin.RouterGroup) {
	group := r.Group("/api/file/v1")
	group.GET("/list", c.list)
	group.POST("/create_folder", c.createFolder)
	group.POST("/rename_file", c.renameFile)
	group.GET("/folder/tree", c.folderTree)
	group.POST("/upload", c.upload)
	group.POST("/move_batch", c.moveBatch)
	group.POST("/del_batch", c.delBatch)
	group.POST("/copy_batch", c.copyBatch)
	group.POST("/second_upload", c.secondUpload)
	group.POST("/init_file_chunk_task", c.initChunkTask)
	group.GET("/get_chunk_upload_url/:identifier/:partNumber", c.getChunkURL)
	group.POST("/merge_file_chunk", c.mergeChunks)
	group.GET("/chunk_upload_progress/:identifier", c.getProgress)
	group.GET("/search", c.search)
	group.POST("/batch_download_url", c.batchDownloadURL)
	group.POST("/batch_download", c.batchDownload)
	group.Any("/download", c.download)
	group.GET("/preview", c.preview)
}

// list 查询文件列表
// @Summary      查询文件列表
// @Description  根据父文件夹ID查询文件列表
// @Tags         文件管理
// @Produce      json
// @Security     Token
// @Param        parent_id  query     int64  true  "父文件夹ID"
// @Success      200  {object}  model.JsonData
// @Router       /api/file/v1/list [get]
func (c *FileController) list(ctx *gin.Context) {
	accountID, _ := ctx.Get("account_id")
	parentID, _ := strconv.ParseInt(ctx.Query("parent_id"), 10, 64)
	files, err := c.svc.ListFiles(ctx.Request.Context(), accountID.(int64), parentID)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Error("获取文件列表失败", 500))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(files))
}

// createFolder 创建文件夹
// @Summary      创建文件夹
// @Description  创建新文件夹
// @Tags         文件管理
// @Accept       json
// @Produce      json
// @Security     Token
// @Param        body  body     model.FolderCreateReq  true  "文件夹信息"
// @Success      200   {object}  model.JsonData
// @Router       /api/file/v1/create_folder [post]
func (c *FileController) createFolder(ctx *gin.Context) {
	accountID, _ := ctx.Get("account_id")
	var req model.FolderCreateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, model.Error("参数错误", model.CodeParamError))
		return
	}
	req.AccountID = accountID.(int64)
	if err := c.svc.CreateFolder(ctx.Request.Context(), &req); err != nil {
		ctx.JSON(http.StatusOK, model.Error(err.Error(), model.CodeFileNameDuplicate))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(nil))
}

// renameFile 重命名文件
// @Summary      重命名文件
// @Description  重命名文件或文件夹
// @Tags         文件管理
// @Accept       json
// @Produce      json
// @Security     Token
// @Param        body  body     model.FileUpdateReq  true  "更新信息"
// @Success      200   {object}  model.JsonData
// @Router       /api/file/v1/rename_file [post]
func (c *FileController) renameFile(ctx *gin.Context) {
	accountID, _ := ctx.Get("account_id")
	var req model.FileUpdateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, model.Error("参数错误", model.CodeParamError))
		return
	}
	req.AccountID = accountID.(int64)
	if err := c.svc.RenameFile(ctx.Request.Context(), &req); err != nil {
		ctx.JSON(http.StatusOK, model.Error(err.Error(), model.CodeFileNotFound))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(nil))
}

// folderTree 获取文件夹树
// @Summary      获取文件夹树
// @Description  获取当前用户的文件夹树形结构
// @Tags         文件管理
// @Produce      json
// @Security     Token
// @Success      200  {object}  model.JsonData
// @Router       /api/file/v1/folder/tree [get]
func (c *FileController) folderTree(ctx *gin.Context) {
	accountIDVal, _ := ctx.Get("account_id")
	if accountIDVal == nil {
		ctx.JSON(http.StatusOK, model.Success([]*model.FolderTreeNodeDTO{}))
		return
	}
	tree, err := c.svc.GetFolderTree(ctx.Request.Context(), accountIDVal.(int64))
	if err != nil {
		ctx.JSON(http.StatusOK, model.Error("获取文件夹树失败", 500))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(tree))
}

// upload 小文件上传
// @Summary      小文件上传
// @Description  直接上传小文件到存储服务器
// @Tags         文件管理
// @Accept       multipart/form-data
// @Produce      json
// @Security     Token
// @Param        filename   formData  string  true  "文件名"
// @Param        identifier formData  string  true  "文件MD5标识"
// @Param        parentId   formData  int64   true  "父文件夹ID"
// @Param        fileSize   formData  int64   true  "文件大小"
// @Param        file       formData  file    true  "文件"
// @Success      200  {object}  model.JsonData
// @Router       /api/file/v1/upload [post]
func (c *FileController) upload(ctx *gin.Context) {
	accountID, _ := ctx.Get("account_id")
	parentID, _ := strconv.ParseInt(ctx.PostForm("parentId"), 10, 64)
	fileSize, _ := strconv.ParseInt(ctx.PostForm("fileSize"), 10, 64)
	fileName := ctx.PostForm("filename")

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusOK, model.Error("上传文件为空", model.CodeFileUploadError))
		return
	}

	src, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusOK, model.Error("打开文件失败", model.CodeFileUploadError))
		return
	}
	defer src.Close()

	if err := c.svc.UploadFile(ctx.Request.Context(), accountID.(int64), parentID, fileName, fileSize, src); err != nil {
		ctx.JSON(http.StatusOK, model.Error(err.Error(), model.CodeFileUploadError))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(nil))
}

// moveBatch 批量移动
// @Summary      批量移动文件
// @Description  批量移动文件到目标文件夹
// @Tags         文件管理
// @Accept       json
// @Produce      json
// @Security     Token
// @Param        body  body     model.FileBatchReq  true  "批量移动请求"
// @Success      200   {object}  model.JsonData
// @Router       /api/file/v1/move_batch [post]
func (c *FileController) moveBatch(ctx *gin.Context) {
	accountID, _ := ctx.Get("account_id")
	var req model.FileBatchReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, model.Error("参数错误", model.CodeParamError))
		return
	}
	req.AccountID = accountID.(int64)
	if err := c.svc.BatchMove(ctx.Request.Context(), &req); err != nil {
		ctx.JSON(http.StatusOK, model.Error("移动失败", 500))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(nil))
}

// delBatch 批量删除
// @Summary      批量删除文件
// @Description  批量删除文件（移入回收站）
// @Tags         文件管理
// @Accept       json
// @Produce      json
// @Security     Token
// @Param        body  body     model.FileDelReq  true  "批量删除请求"
// @Success      200   {object}  model.JsonData
// @Router       /api/file/v1/del_batch [post]
func (c *FileController) delBatch(ctx *gin.Context) {
	accountID, _ := ctx.Get("account_id")
	var req model.FileDelReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, model.Error("参数错误", model.CodeParamError))
		return
	}
	req.AccountID = accountID.(int64)
	if err := c.svc.BatchDelete(ctx.Request.Context(), &req); err != nil {
		ctx.JSON(http.StatusOK, model.Error("删除失败", 500))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(nil))
}

// copyBatch 批量复制
// @Summary      批量复制文件
// @Description  批量复制文件到目标文件夹（递归）
// @Tags         文件管理
// @Accept       json
// @Produce      json
// @Security     Token
// @Param        body  body     model.FileBatchReq  true  "批量复制请求"
// @Success      200   {object}  model.JsonData
// @Router       /api/file/v1/copy_batch [post]
func (c *FileController) copyBatch(ctx *gin.Context) {
	accountID, _ := ctx.Get("account_id")
	var req model.FileBatchReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, model.Error("参数错误", model.CodeParamError))
		return
	}
	req.AccountID = accountID.(int64)
	if err := c.svc.BatchCopy(ctx.Request.Context(), &req); err != nil {
		ctx.JSON(http.StatusOK, model.Error(err.Error(), model.CodeBatchOpError))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(nil))
}

// secondUpload 秒传
// @Summary      文件秒传
// @Description  根据文件MD5标识秒传文件，若已存在则直接创建引用
// @Tags         文件管理
// @Accept       json
// @Produce      json
// @Security     Token
// @Param        body  body     model.FileSecondUploadReq  true  "秒传请求"
// @Success      200   {object}  model.JsonData{data=bool}
// @Router       /api/file/v1/second_upload [post]
func (c *FileController) secondUpload(ctx *gin.Context) {
	accountID, _ := ctx.Get("account_id")
	var req model.FileSecondUploadReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, model.Error("参数错误", model.CodeParamError))
		return
	}
	req.AccountID = accountID.(int64)
	hit, err := c.svc.SecondUpload(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Error(err.Error(), model.CodeFileUploadError))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(hit))
}

// initChunkTask 初始化分片上传
// @Summary      初始化分片上传
// @Description  创建分片上传任务
// @Tags         文件管理
// @Accept       json
// @Produce      json
// @Security     Token
// @Param        body  body     model.FileChunkInitTaskReq  true  "分片上传初始化请求"
// @Success      200   {object}  model.JsonData
// @Router       /api/file/v1/init_file_chunk_task [post]
func (c *FileController) initChunkTask(ctx *gin.Context) {
	accountID, _ := ctx.Get("account_id")
	var req model.FileChunkInitTaskReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, model.Error("参数错误", model.CodeParamError))
		return
	}
	req.AccountID = accountID.(int64)
	chunkSvc := service.NewChunkService()
	dto, err := chunkSvc.InitChunkTask(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Error(err.Error(), model.CodeFileUploadError))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(dto))
}

// getChunkURL 获取分片上传地址
// @Summary      获取分片上传地址
// @Description  获取分片临时上传地址
// @Tags         文件管理
// @Produce      json
// @Security     Token
// @Param        identifier  path     string  true  "文件唯一标识"
// @Param        partNumber  path     int     true  "分片序号"
// @Success      200         {object}  model.JsonData
// @Router       /api/file/v1/get_chunk_upload_url/{identifier}/{partNumber} [get]
func (c *FileController) getChunkURL(ctx *gin.Context) {
	identifier := ctx.Param("identifier")
	accountID, _ := ctx.Get("account_id")
	partNumberStr := ctx.Param("partNumber")
	partNumber, err := strconv.Atoi(partNumberStr)
	if err != nil || partNumber < 1 {
		ctx.JSON(http.StatusOK, model.Error("分片号无效", model.CodeParamError))
		return
	}
	chunkSvc := service.NewChunkService()
	u, err := chunkSvc.GetPresignedURL(ctx.Request.Context(), identifier, accountID.(int64), int32(partNumber))
	if err != nil {
		ctx.JSON(http.StatusOK, model.Error(err.Error(), model.CodeChunkTaskNotExists))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(u))
}

// mergeChunks 合并分片
// @Summary      合并分片
// @Description  合并分片并完成文件上传
// @Tags         文件管理
// @Accept       json
// @Produce      json
// @Security     Token
// @Param        body  body     model.FileChunkMergeReq  true  "分片合并请求"
// @Success      200   {object}  model.JsonData
// @Router       /api/file/v1/merge_file_chunk [post]
func (c *FileController) mergeChunks(ctx *gin.Context) {
	accountID, _ := ctx.Get("account_id")
	var req model.FileChunkMergeReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, model.Error("参数错误", model.CodeParamError))
		return
	}
	req.AccountID = accountID.(int64)
	chunkSvc := service.NewChunkService()
	if err := chunkSvc.MergeChunks(ctx.Request.Context(), &req); err != nil {
		ctx.JSON(http.StatusOK, model.Error(err.Error(), model.CodeChunkCountInsufficient))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(nil))
}

// getProgress 查询分片进度
// @Summary      查询分片上传进度
// @Description  根据文件标识查询分片上传进度
// @Tags         文件管理
// @Produce      json
// @Security     Token
// @Param        identifier  path     string  true  "文件唯一标识"
// @Success      200         {object}  model.JsonData
// @Router       /api/file/v1/chunk_upload_progress/{identifier} [get]
func (c *FileController) getProgress(ctx *gin.Context) {
	identifier := ctx.Param("identifier")
	accountID, _ := ctx.Get("account_id")
	chunkSvc := service.NewChunkService()
	dto, err := chunkSvc.GetProgress(ctx.Request.Context(), accountID.(int64), identifier)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Error(err.Error(), model.CodeChunkTaskNotExists))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(dto))
}

// search 搜索文件
// @Summary      搜索文件
// @Description  根据文件名模糊搜索文件
// @Tags         文件管理
// @Produce      json
// @Security     Token
// @Param        search  query     string  true  "搜索关键词"
// @Success      200     {object}  model.JsonData
// @Router       /api/file/v1/search [get]
func (c *FileController) search(ctx *gin.Context) {
	accountID, _ := ctx.Get("account_id")
	keyword := ctx.Query("search")
	if len(keyword) > 100 {
		keyword = keyword[:100]
	}
	files, err := c.svc.SearchFiles(ctx.Request.Context(), accountID.(int64), keyword)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Error("搜索失败", 500))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(files))
}

// batchDownloadURL 获取批量下载URL
// @Summary      获取批量下载URL
// @Description  获取多个文件的临时下载链接
// @Tags         文件管理
// @Accept       json
// @Produce      json
// @Security     Token
// @Param        body  body     model.FileDownloadReq  true  "批量下载请求"
// @Success      200   {object}  model.JsonData
// @Router       /api/file/v1/batch_download_url [post]
func (c *FileController) batchDownloadURL(ctx *gin.Context) {
	accountID, _ := ctx.Get("account_id")
	var req model.FileDownloadReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, model.Error("参数错误", model.CodeParamError))
		return
	}
	req.AccountID = accountID.(int64)
	urls, err := c.svc.BatchDownloadURL(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Error("获取下载链接失败", 500))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(urls))
}

// download 单文件下载
// @Summary      单文件下载
// @Description  下载单个文件流
// @Tags         文件管理
// @Produce      octet-stream
// @Security     Token
// @Param        fileId  query     int64  true  "文件ID"
// @Success      200     {file}    binary
// @Router       /api/file/v1/download [get]
func (c *FileController) download(ctx *gin.Context) {
	fileIDStr := ctx.Query("fileId")
	if fileIDStr == "" {
		fileIDStr = ctx.PostForm("fileId")
	}
	if fileIDStr == "" {
		ctx.JSON(http.StatusOK, model.Error("缺少fileId", model.CodeFileNotFound))
		return
	}
	fileID, _ := strconv.ParseInt(fileIDStr, 10, 64)

	reader, fileName, size, err := c.svc.DownloadFile(ctx.Request.Context(), fileID)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Error(err.Error(), model.CodeFileNotFound))
		return
	}

	ctx.Header("Content-Disposition", "attachment; filename=\""+url.QueryEscape(fileName)+"\"")
	ctx.DataFromReader(http.StatusOK, size, "application/octet-stream", reader, nil)
}

// preview 文件预览
// @Summary      文件预览
// @Description  以 inline 模式返回文件流，用于浏览器预览（图片、PDF等）
// @Tags         文件管理
// @Produce      application/octet-stream
// @Security     Token
// @Param        fileId  query     int64  true  "文件ID"
// @Success      200     {file}    binary
// @Router       /api/file/v1/preview [get]
func (c *FileController) preview(ctx *gin.Context) {
	fileIDStr := ctx.Query("fileId")
	if fileIDStr == "" {
		ctx.JSON(http.StatusOK, model.Error("缺少fileId", model.CodeFileNotFound))
		return
	}
	fileID, _ := strconv.ParseInt(fileIDStr, 10, 64)

	reader, fileName, size, mimeType, err := c.svc.PreviewFile(ctx.Request.Context(), fileID)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Error(err.Error(), model.CodeFileNotFound))
		return
	}

	ctx.DataFromReader(http.StatusOK, size, mimeType, reader, map[string]string{
		"Content-Disposition": "inline; filename=\"" + url.QueryEscape(fileName) + "\"",
		"Cache-Control":       "no-cache",
	})
}

// batchDownload 批量ZIP下载
// @Summary      批量ZIP下载
// @Description  将多个文件打包为ZIP下载
// @Tags         文件管理
// @Produce      octet-stream
// @Security     Token
// @Param        fileIds  query     string  true  "文件ID列表，逗号分隔"
// @Success      200      {file}    binary
// @Router       /api/file/v1/batch_download [post]
func (c *FileController) batchDownload(ctx *gin.Context) {
	accountID, _ := ctx.Get("account_id")
	// 前端用 FormData 发送，字段名为 fileIdsStr
	fileIDsStr := ctx.PostForm("fileIdsStr")
	if fileIDsStr == "" {
		fileIDsStr = ctx.Query("fileIds")
	}
	if fileIDsStr == "" {
		fileIDsStr = ctx.PostForm("fileIds")
	}
	log.Printf("[batchDownload] fileIDsStr=%q", fileIDsStr)
	if fileIDsStr == "" {
		ctx.JSON(http.StatusOK, model.Error("缺少fileIds", model.CodeFileNotFound))
		return
	}

	var fileIDs []int64
	for _, part := range splitInt64s(fileIDsStr) {
		fileIDs = append(fileIDs, part)
	}
	log.Printf("[batchDownload] parsed fileIDs=%v accountID=%d", fileIDs, accountID)

	buf, err := c.svc.BatchDownloadZip(ctx.Request.Context(), fileIDs, accountID.(int64))
	if err != nil {
		log.Printf("[batchDownload] error: %v", err)
		ctx.JSON(http.StatusOK, model.Error(err.Error(), 500))
		return
	}

	log.Printf("[batchDownload] ZIP size=%d", buf.Len())
	ctx.Header("Content-Disposition", "attachment; filename=\"files.zip\"")
	ctx.Data(http.StatusOK, "application/zip", buf.Bytes())
}

func splitInt64s(s string) []int64 {
	var result []int64
	for _, part := range splitComma(s) {
		if v, err := strconv.ParseInt(part, 10, 64); err == nil {
			result = append(result, v)
		}
	}
	return result
}

func splitComma(s string) []string {
	var result []string
	start := 0
	for i, c := range s {
		if c == ',' {
			result = append(result, s[start:i])
			start = i + 1
		}
	}
	if start < len(s) {
		result = append(result, s[start:])
	}
	return result
}
