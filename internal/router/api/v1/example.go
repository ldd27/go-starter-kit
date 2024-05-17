package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/ldd27/go-starter-kit/internal/service"
	"github.com/ldd27/go-starter-kit/internal/types"
	"github.com/ldd27/go-starter-kit/internal/types/serialize"
)

type ExampleController struct {
	exampleSvc *service.ExampleSvc
}

func NewExampleController(exampleSvc *service.ExampleSvc) *ExampleController {
	return &ExampleController{
		exampleSvc: exampleSvc,
	}
}

// GetList
//
//	@Summary		获取example列表
//	@Description	获取example列表
//	@Tags			example
//	@Success		200	{object}	types.Res{data=[]types.Example} "response"
//	@Security		BearerToken
//	@Router			/example [GET]
func (r *ExampleController) GetList(c echo.Context) (err error) {
	ctrl := parseCtx(c)
	ctrl.logger.Info("example.GetList")

	result, err := r.exampleSvc.GetList(ctrl.ctx)
	if err != nil {
		return err
	}

	return ctrl.JsonRes(serialize.Example2Apis(result))
}

// GetPageList
//
//	@Summary		分页获取example列表
//	@Description	分页获取example列表
//	@Tags			example
//	@Param			request	query		types.GetExamplePageListReq		true "request"
//	@Success		200	{object}	types.Res{data=types.PageRes{data=[]types.Example}} "response"
//	@Security		BearerToken
//	@Router			/example/page [GET]
func (r *ExampleController) GetPageList(c echo.Context) (err error) {
	ctrl := parseCtx(c)
	ctrl.logger.Info("example.GetPageList")

	req := types.GetExamplePageListReq{}
	if err = ctrl.BindValidate(&req); err != nil {
		return err
	}

	result, total, err := r.exampleSvc.GetPageList(ctrl.ctx, req.PageIndex, req.PageSize)
	if err != nil {
		return err
	}

	return ctrl.JsonPageRes(total, serialize.Example2Apis(result))
}

// GetCursorPageList
//
//	@Summary		根据cursor分页获取example列表
//	@Description	根据cursor分页获取example列表
//	@Tags			example
//	@Param			request	query		types.GetExampleCursorPageListReq		true "request"
//	@Success		200	{object}	types.Res{data=types.CursorPageRes{data=[]types.Example}} "response"
//	@Security		BearerToken
//	@Router			/example/cursor [GET]
func (r *ExampleController) GetCursorPageList(c echo.Context) (err error) {
	ctrl := parseCtx(c)
	ctrl.logger.Info("example.GetCursorPageList")

	req := types.GetExampleCursorPageListReq{}
	if err = ctrl.BindValidate(&req); err != nil {
		return err
	}

	result, err := r.exampleSvc.GetCursorPageList(ctrl.ctx, req.Cursor, 20)
	if err != nil {
		return err
	}

	res := pkgCursorRes(req.Cursor, serialize.Example2Apis(result), func(item types.Example) int { return item.ID })
	return ctrl.JsonRes(res)
}

// Create
//
//	@Summary		创建example
//	@Description	创建example
//	@Tags			example
//	@Param			request	body		types.CreateExampleReq		true "request"
//	@Success		200	{object}	types.Res{data=interface{}} "response"
//	@Security		BearerToken
//	@Router			/example [POST]
func (r *ExampleController) Create(c echo.Context) (err error) {
	ctrl := parseCtx(c)
	ctrl.logger.Info("example.Create")

	req := types.CreateExampleReq{}
	if err = ctrl.BindValidate(&req); err != nil {
		return err
	}

	err = r.exampleSvc.Create(ctrl.ctx, req.Name)
	if err != nil {
		return err
	}

	return ctrl.JsonRes(nil)
}

// Update
//
//	@Summary		编辑example
//	@Description	编辑example
//	@Tags			example
//	@Param			request	body		types.UpdateExampleReq		true "request"
//	@Success		200	{object}	types.Res{data=interface{}} "response"
//	@Security		BearerToken
//	@Router			/example [PUT]
func (r *ExampleController) Update(c echo.Context) (err error) {
	ctrl := parseCtx(c)
	ctrl.logger.Info("example.Create")

	req := types.UpdateExampleReq{}
	if err = ctrl.BindValidate(&req); err != nil {
		return err
	}

	err = r.exampleSvc.Update(ctrl.ctx, req.ID, req.Name)
	if err != nil {
		return err
	}

	return ctrl.JsonRes(nil)
}

// Delete
//
//	@Summary		删除example
//	@Description	删除example
//	@Tags			example
//	@Param			request	body		types.DeleteExampleReq		true "request"
//	@Success		200	{object}	types.Res{data=interface{}} "response"
//	@Security		BearerToken
//	@Router			/example [DELETE]
func (r *ExampleController) Delete(c echo.Context) (err error) {
	ctrl := parseCtx(c)
	ctrl.logger.Info("example.Delete")

	req := types.DeleteExampleReq{}
	if err = ctrl.BindValidate(&req); err != nil {
		return err
	}

	err = r.exampleSvc.Delete(ctrl.ctx, req.ID)
	if err != nil {
		return err
	}

	return ctrl.JsonRes(nil)
}
