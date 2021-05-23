package api

import (
	"as/global"
	"as/internal/model"
	"as/internal/service"
	"as/pkg/app"
	"as/pkg/convert"
	"as/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Store struct {
}

func (s Store) Update(c *gin.Context) {
	response := app.NewResponse(c)
	var store model.Store
	err := c.ShouldBind(&store)
	if err != nil {
		response.ToErrorResponse(errcode.InvilidParams.WithDetails(err.Error()))
		return
	}
	param := c.Param("id")
	if param == "" {
		response.ToErrorResponse(errcode.InvilidParams.WithDetails("参数错误"))
		return
	}
	res, _ := convert.StrTo(param).Uint64()
	m := new(model.Store)
	m.Model = new(model.Model)
	m.Model.ID = res
	s2 := service.New(c.Request.Context())
	unique := s2.CheckStoreNameIsUnique(store.Name)
	global.DBEngine.First(m)
	if !unique && m.Name != store.Name {
		response.ToErrorResponse(errcode.InvilidParams.WithDetails("已经存在该门店"))
		return
	}

	global.DBEngine.Model(&m).Omit("id", "created", "updated_at").Update(&store)
	response.ToResponse("更新成功")
}

func (s Store) Index(c *gin.Context) {
	var data []*model.Store
	find := global.DBEngine.Find(&data).Error
	response := app.NewResponse(c)
	if find != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails(find.Error()))
		return
	}
	response.ToResponse(data)
}

func (s Store) Get(c *gin.Context) {
	param := c.Param("id")
	response := app.NewResponse(c)
	if param == "" {
		response.ToErrorResponse(errcode.InvilidParams.WithDetails("参数错误"))
		return
	}
	res, _ := convert.StrTo(param).Uint64()
	m := new(model.Store)
	m.Model = new(model.Model)
	m.Model.ID = res
	get, err := m.Get(global.DBEngine)
	if err != nil {
		response.ToErrorResponse(errcode.NotFound)
		return
	}
	response.ToResponse(get)
}

func (s Store) Store(c *gin.Context) {
	response := app.NewResponse(c)
	var store model.Store
	err := c.ShouldBind(&store)
	if err != nil {
		response.ToErrorResponse(errcode.InvilidParams.WithDetails(err.Error()))
		return
	}
	s2 := service.New(c.Request.Context())
	unique := s2.CheckStoreNameIsUnique(store.Name)
	if !unique {
		response.ToErrorResponse(errcode.InvilidParams.WithDetails("已经存在该门店"))
		return
	}
	err = global.DBEngine.Create(&store).Error
	if err != nil {
		response.ToErrorResponse(errcode.NewError(10008, "创建失败"))
		return
	}
	response.ToResponse("创建成功")

}
func NewStore() *Store {
	return &Store{}
}

func (s Store) Delete(c *gin.Context) {
	param := c.Param("id")
	response := app.NewResponse(c)
	if param == "" {
		response.ToErrorResponse(errcode.InvilidParams.WithDetails("参数错误"))
		return
	}
	res, _ := convert.StrTo(param).Uint64()
	m := new(model.Store)
	m.Model = new(model.Model)
	m.Model.ID = res
	err := m.Delete(global.DBEngine)
	if err != nil {
		response.ToErrorResponse(errcode.InvilidParams)
		return
	}
	response.ToResponse("删除成功")
}
