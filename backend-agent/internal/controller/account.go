package controller

import (
	"net/http"

	"github.com/aqi/AqiCloud-AgentPan-Go/internal/model"
	"github.com/aqi/AqiCloud-AgentPan-Go/internal/service"
	"github.com/gin-gonic/gin"
)

type AccountController struct {
	svc *service.AccountService
}

func NewAccountController() *AccountController {
	return &AccountController{svc: service.NewAccountService()}
}

func (c *AccountController) Register(r *gin.RouterGroup) {
	group := r.Group("/api/account/v1")
	group.POST("/register", c.register)
	group.POST("/login", c.login)
	group.POST("/upload_avatar", c.uploadAvatar)
	group.GET("/detail", c.detail)
	group.POST("/update/my", c.update)
}

// register 用户注册
// @Summary      用户注册
// @Description  注册新账号并自动分配10GB存储空间
// @Tags         账号管理
// @Accept       json
// @Produce      json
// @Param        body  body     model.AccountRegisterReq  true  "注册信息"
// @Success      200  {object}  model.JsonData
// @Router       /api/account/v1/register [post]
func (c *AccountController) register(ctx *gin.Context) {
	var req model.AccountRegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, model.Error("参数错误", model.CodeParamError))
		return
	}
	if err := c.svc.Register(ctx.Request.Context(), &req); err != nil {
		ctx.JSON(http.StatusOK, model.Error(err.Error(), model.CodeAccountExists))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(nil))
}

// login 用户登录
// @Summary      用户登录
// @Description  登录并获取JWT令牌
// @Tags         账号管理
// @Accept       json
// @Produce      json
// @Param        body  body     model.AccountLoginReq  true  "登录信息"
// @Success      200  {object}  model.JsonData
// @Router       /api/account/v1/login [post]
func (c *AccountController) login(ctx *gin.Context) {
	var req model.AccountLoginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, model.Error("参数错误", model.CodeParamError))
		return
	}
	token, err := c.svc.Login(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Error(err.Error(), model.CodeParamError))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(token))
}

// uploadAvatar 上传头像
// @Summary      上传头像
// @Description  上传用户头像图片
// @Tags         账号管理
// @Accept       multipart/form-data
// @Produce      json
// @Param        file  formData  file  true  "头像图片"
// @Success      200   {object}  model.JsonData
// @Router       /api/account/v1/upload_avatar [post]
func (c *AccountController) uploadAvatar(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusOK, model.Error("上传文件为空", model.CodeFileUploadError))
		return
	}
	accountID, exists := ctx.Get("account_id")
	if exists {
		url, err := c.svc.UploadAvatar(ctx.Request.Context(), file, accountID.(int64))
		if err != nil {
			ctx.JSON(http.StatusOK, model.Error(err.Error(), model.CodeFileUploadError))
			return
		}
		ctx.JSON(http.StatusOK, model.Success(url))
	} else {
		url, err := c.svc.UploadAvatarFile(ctx.Request.Context(), file)
		if err != nil {
			ctx.JSON(http.StatusOK, model.Error(err.Error(), model.CodeFileUploadError))
			return
		}
		ctx.JSON(http.StatusOK, model.Success(url))
	}
}

// detail 获取用户详情
// @Summary      获取用户详情
// @Description  获取当前登录用户的详细信息（含存储空间）
// @Tags         账号管理
// @Produce      json
// @Security     Token
// @Success      200  {object}  model.JsonData
// @Router       /api/account/v1/detail [get]
func (c *AccountController) detail(ctx *gin.Context) {
	accountID, _ := ctx.Get("account_id")
	dto, err := c.svc.GetAccountDetail(ctx.Request.Context(), accountID.(int64))
	if err != nil {
		ctx.JSON(http.StatusOK, model.Error("获取用户信息失败", model.CodeAccountNotExists))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(dto))
}

// update 更新用户信息
// @Summary      更新用户信息
// @Description  更新当前用户的昵称和头像
// @Tags         账号管理
// @Accept       json
// @Produce      json
// @Param        body  body     model.AccountUpdateReq  true  "更新信息"
// @Security     Token
// @Success      200   {object}  model.JsonData
// @Router       /api/account/v1/update/my [post]
func (c *AccountController) update(ctx *gin.Context) {
	accountID, _ := ctx.Get("account_id")
	var req model.AccountUpdateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, model.Error("参数错误", model.CodeParamError))
		return
	}
	if err := c.svc.UpdateAccount(ctx.Request.Context(), accountID.(int64), &req); err != nil {
		ctx.JSON(http.StatusOK, model.Error(err.Error(), model.CodeAccountNotExists))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(nil))
}
