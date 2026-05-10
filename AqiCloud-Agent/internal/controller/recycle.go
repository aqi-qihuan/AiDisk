package controller

import (
	"net/http"

	"github.com/aqi/AqiCloud-AgentPan-Go/internal/model"
	"github.com/aqi/AqiCloud-AgentPan-Go/internal/service"
	"github.com/gin-gonic/gin"
)

type RecycleController struct {
	svc *service.RecycleService
}

func NewRecycleController() *RecycleController {
	return &RecycleController{svc: service.NewRecycleService()}
}

func (c *RecycleController) Register(r *gin.RouterGroup) {
	group := r.Group("/api/recycle/v1")
	group.GET("/list", c.list)
	group.POST("/delete", c.delete)
	group.POST("/restore", c.restore)
}

// list 回收站列表
// @Summary      回收站列表
// @Description  获取回收站中的文件列表
// @Tags         回收站
// @Produce      json
// @Security     Token
// @Success      200  {object}  model.JsonData
// @Router       /api/recycle/v1/list [get]
func (c *RecycleController) list(ctx *gin.Context) {
	accountID, _ := ctx.Get("account_id")
	files, err := c.svc.ListRecycle(ctx.Request.Context(), accountID.(int64))
	if err != nil {
		ctx.JSON(http.StatusOK, model.Error("获取回收站列表失败", 500))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(files))
}

// delete 永久删除
// @Summary      永久删除
// @Description  从回收站中永久删除文件（不可恢复）
// @Tags         回收站
// @Accept       json
// @Produce      json
// @Security     Token
// @Param        body  body     model.RecycleDelReq  true  "删除请求"
// @Success      200   {object}  model.JsonData
// @Router       /api/recycle/v1/delete [post]
func (c *RecycleController) delete(ctx *gin.Context) {
	accountID, _ := ctx.Get("account_id")
	var req model.RecycleDelReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, model.Error("参数错误", model.CodeParamError))
		return
	}
	req.AccountID = accountID.(int64)
	if err := c.svc.PermanentDelete(ctx.Request.Context(), &req); err != nil {
		ctx.JSON(http.StatusOK, model.Error("删除失败", 500))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(nil))
}

// restore 还原文件
// @Summary      还原文件
// @Description  从回收站还原文件到原位置
// @Tags         回收站
// @Accept       json
// @Produce      json
// @Security     Token
// @Param        body  body     model.RecycleRestoreReq  true  "还原请求"
// @Success      200   {object}  model.JsonData
// @Router       /api/recycle/v1/restore [post]
func (c *RecycleController) restore(ctx *gin.Context) {
	accountID, _ := ctx.Get("account_id")
	var req model.RecycleRestoreReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, model.Error("参数错误", model.CodeParamError))
		return
	}
	req.AccountID = accountID.(int64)
	if err := c.svc.Restore(ctx.Request.Context(), &req); err != nil {
		ctx.JSON(http.StatusOK, model.Error("还原失败", 500))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(nil))
}
