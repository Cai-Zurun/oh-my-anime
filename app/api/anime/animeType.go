package anime

import (
	"github.com/gogf/gf/net/ghttp"
	"oh-my-anime_gf/app/model/anime_type"
	"oh-my-anime_gf/app/service/anime"
	"oh-my-anime_gf/library/response"
)

// @summary 动漫类型添加接口
// @tags    动漫类型
// @produce json
// @param   Type formData string true "动漫类型"
// @router  /anime/type/add [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func AddType(r *ghttp.Request) {
	var data *anime.AddTypeInput
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, response.FAIL, err.Error())
	}
	if err := anime.AddType(data.Type); err != nil {
		response.JsonExit(r, response.FAIL, err.Error())
	} else {
		response.JsonExit(r, response.SUCCESS, "动漫类型添加成功")
	}
}

// @summary 动漫类型获取接口
// @tags    动漫类型
// @produce json
// @router  /anime/type/get [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func GetType(r *ghttp.Request) {
	TypeArr, _ := anime_type.FindArray("type")
	response.JsonExit(r, response.SUCCESS, "动漫类型获取成功", TypeArr)
}

// @summary 动漫类型删除接口
// @tags    动漫类型
// @produce json
// @param   Type formData string true "动漫类型"
// @router  /anime/type/delete [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func DeleteType(r *ghttp.Request)  {
	var data *anime.DeleteTypeInput
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, response.FAIL, err.Error())
	}
	if err := anime.DeleteType(data); err != nil {
		response.JsonExit(r, response.FAIL, err.Error())
	} else {
		response.JsonExit(r, response.SUCCESS, "动漫类型删除成功")
	}
}

// @summary 动漫类型更新接口
// @tags    动漫类型
// @produce json
// @param   Type formData string true "动漫类型"
// @param   NewType formData string true "新动漫类型"
// @router  /anime/type/update [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func UpdateType(r *ghttp.Request)  {
	var data *anime.UpdateTypeInput
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, response.FAIL, err.Error())
	}
	if err := anime.UpdateType(data); err != nil {
		response.JsonExit(r, response.FAIL, err.Error())
	} else {
		response.JsonExit(r, response.SUCCESS, "动漫类型修改成功")
	}
}