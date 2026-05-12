package controller

import (
	"net/http"
	"strconv"

	"github.com/aqi/AqiCloud-Agent/internal/middleware"
	"github.com/aqi/AqiCloud-Agent/internal/model"
	"github.com/aqi/AqiCloud-Agent/internal/service"
	"github.com/gin-gonic/gin"
)

type ShareController struct {
	svc *service.ShareService
}

func NewShareController() *ShareController {
	return &ShareController{svc: service.NewShareService()}
}

func (c *ShareController) Register(r *gin.RouterGroup) {
	group := r.Group("/api/share/v1")
	group.GET("/list", c.list)
	group.GET("/listShare", c.list)
	group.POST("/create", c.create)
	group.POST("/createShare", c.create)
	group.POST("/cancel", c.cancel)
	group.POST("/cancelShare", c.cancel)
	group.GET("/visit", c.visit)
	group.GET("/getShareSimpleDetail", c.visit)
	group.POST("/check_share_code", c.checkCode)

	// Routes requiring share-token header
	shareGroup := group.Group("")
	shareGroup.Use(middleware.ShareTokenMiddleware())
	{
		shareGroup.GET("/detail", c.detail)
		shareGroup.GET("/getShareDetail", c.detail)
		shareGroup.POST("/list_share_file", c.listFiles)
		shareGroup.POST("/listShareFile", c.listFiles)
		shareGroup.POST("/transfer", c.transfer)
	}
}

// list 获取分享列表
// @Summary      获取分享列表
// @Description  获取当前用户的分享列表
// @Tags         分享管理
// @Produce      json
// @Security     Token
// @Success      200  {object}  model.JsonData
// @Router       /api/share/v1/list [get]
func (c *ShareController) list(ctx *gin.Context) {
	accountID, _ := ctx.Get("account_id")
	list, err := c.svc.ListShare(ctx.Request.Context(), accountID.(int64))
	if err != nil {
		ctx.JSON(http.StatusOK, model.Error("获取分享列表失败", model.CodeShareNotExists))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(list))
}

// create 创建分享
// @Summary      创建分享
// @Description  创建文件分享链接，支持有码/无码分享
// @Tags         分享管理
// @Accept       json
// @Produce      json
// @Security     Token
// @Param        body  body     model.ShareCreateReq  true  "分享创建请求"
// @Success      200   {object}  model.JsonData
// @Router       /api/share/v1/create [post]
func (c *ShareController) create(ctx *gin.Context) {
	accountID, _ := ctx.Get("account_id")
	var req model.ShareCreateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, model.Error("参数错误", model.CodeParamError))
		return
	}
	req.AccountID = accountID.(int64)
	dto, err := c.svc.CreateShare(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Error(err.Error(), model.CodeShareNotExists))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(dto))
}

// cancel 取消分享
// @Summary      取消分享
// @Description  批量取消分享
// @Tags         分享管理
// @Accept       json
// @Produce      json
// @Security     Token
// @Param        body  body     model.ShareCancelReq  true  "取消分享请求"
// @Success      200   {object}  model.JsonData
// @Router       /api/share/v1/cancel [post]
func (c *ShareController) cancel(ctx *gin.Context) {
	accountID, _ := ctx.Get("account_id")
	var req model.ShareCancelReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, model.Error("参数错误", model.CodeParamError))
		return
	}
	req.AccountID = accountID.(int64)
	if err := c.svc.CancelShare(ctx.Request.Context(), &req); err != nil {
		ctx.JSON(http.StatusOK, model.Error(err.Error(), model.CodeCancelShareIllegal))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(nil))
}

// visit 访问分享
// @Summary      访问分享
// @Description  访问分享链接，NO_CODE类型返回shareToken
// @Tags         分享管理
// @Produce      json
// @Param        shareId  query     int64  true  "分享ID"
// @Success      200      {object}  model.JsonData
// @Router       /api/share/v1/visit [get]
func (c *ShareController) visit(ctx *gin.Context) {
	shareID, _ := strconv.ParseInt(ctx.Query("shareId"), 10, 64)
	dto, err := c.svc.VisitShare(ctx.Request.Context(), shareID)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Error(err.Error(), model.CodeShareNotExists))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(dto))
}

// checkCode 校验分享码
// @Summary      校验分享提取码
// @Description  输入分享提取码获取访问令牌
// @Tags         分享管理
// @Accept       json
// @Produce      json
// @Param        body  body     model.ShareCheckReq  true  "校验提取码请求"
// @Success      200   {object}  model.JsonData
// @Router       /api/share/v1/check_share_code [post]
func (c *ShareController) checkCode(ctx *gin.Context) {
	var req model.ShareCheckReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, model.Error("参数错误", model.CodeParamError))
		return
	}
	token, err := c.svc.CheckShareCode(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Error(err.Error(), model.CodeShareCodeIllegal))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(token))
}

// detail 分享详情（需share-token）
// @Summary      分享详情
// @Description  获取分享详情（含文件列表），需携带share-token
// @Tags         分享管理
// @Produce      json
// @Param        share-token  header    string  true  "分享访问令牌"
// @Success      200          {object}  model.JsonData
// @Router       /api/share/v1/detail [get]
func (c *ShareController) detail(ctx *gin.Context) {
	shareID, _ := ctx.Get("share_id")
	dto, err := c.svc.GetShareDetail(ctx.Request.Context(), shareID.(int64))
	if err != nil {
		ctx.JSON(http.StatusOK, model.Error(err.Error(), model.CodeShareNotExists))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(dto))
}

// listFiles 分享文件列表（需share-token）
// @Summary      分享文件列表
// @Description  浏览分享内的文件列表，需携带share-token
// @Tags         分享管理
// @Accept       json
// @Produce      json
// @Param        share-token  header  string                   true  "分享访问令牌"
// @Param        body         body    model.ShareFileQueryReq  true  "查询请求"
// @Success      200          {object}  model.JsonData
// @Router       /api/share/v1/list_share_file [post]
func (c *ShareController) listFiles(ctx *gin.Context) {
	shareID, _ := ctx.Get("share_id")
	var req model.ShareFileQueryReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, model.Error("参数错误", model.CodeParamError))
		return
	}
	req.ShareID = shareID.(int64)
	files, err := c.svc.ListShareFiles(ctx.Request.Context(), req.ShareID, req.ParentID)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Error(err.Error(), model.CodeShareFileIllegal))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(files))
}

// transfer 转存分享文件（需share-token）
// @Summary      转存分享文件
// @Description  将分享文件转存到自己的网盘，需携带share-token
// @Tags         分享管理
// @Accept       json
// @Produce      json
// @Param        share-token       header  string                   true  "分享访问令牌"
// @Param        body              body    model.ShareFileTransferReq  true  "转存请求"
// @Success      200               {object}  model.JsonData
// @Router       /api/share/v1/transfer [post]
func (c *ShareController) transfer(ctx *gin.Context) {
	accountID, _ := ctx.Get("account_id")
	if accountID == nil {
		ctx.JSON(http.StatusOK, model.Error("请先登录", model.CodeNotLogin))
		return
	}
	shareID, _ := ctx.Get("share_id")
	var req model.ShareFileTransferReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, model.Error("参数错误", model.CodeParamError))
		return
	}
	req.AccountID = accountID.(int64)
	req.ShareID = shareID.(int64)
	if err := c.svc.TransferToOwn(ctx.Request.Context(), req.ShareID, req.FileIDs, req.TargetParentID, req.AccountID); err != nil {
		ctx.JSON(http.StatusOK, model.Error(err.Error(), model.CodeBatchOpError))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(nil))
}
