package anime

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"oh-my-anime_gf/app/service/anime"
	"oh-my-anime_gf/library/response"
	//animeModel "oh-my-anime_gf/app/model/anime"
)

// @summary 动漫添加接口
// @tags    动漫
// @produce json
// @param   Name formData string true "动漫名字"
// @param   Link formData string true "动漫链接"
// @param   Type formData string true "动漫类型"
// @router  /anime/add [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func AddAnime(r *ghttp.Request) {
	var data *anime.AddAnimeInput
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, response.FAIL, err.Error())
	}
	data.Img = r.GetUploadFile("AnimeImg")
	if err := anime.AddAnime(data); err != nil {
		response.JsonExit(r, response.FAIL, err.Error())
	} else {
		response.JsonExit(r, response.SUCCESS, "动漫添加成功")
	}
}

//func AddAnimeImg(r *ghttp.Request)  {
//	Img := r.GetUploadFile("upload-img")
//
//}

// @summary 动漫获取接口
// @tags    动漫
// @produce json
// @param   Type formData string true "动漫类型"
// @router  /anime/get [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func GetAnime(r *ghttp.Request)  {
	var data *anime.GetAnimeInput
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, response.FAIL, err.Error())
	}
	AnimeArr, err := anime.GetAnime(data.Type)
	if err != nil {
		response.JsonExit(r, response.FAIL, err.Error())
	} else {
		response.JsonExit(r, response.SUCCESS, "动漫获取成功", AnimeArr)
	}
}

func GetAllAnime (r *ghttp.Request) {
	TypeArr := anime.GetType()
	AnimeArr := make([][]anime.AnimeOutput, len(TypeArr))
	for i := 0; i < len(TypeArr); i++ {
		AnimeArr[i], _ = anime.GetAnime(gconv.String(TypeArr[i]))
	}
	response.JsonExit(r, response.SUCCESS, "所有动漫获取成功", AnimeArr)
}

// @summary 动漫添加接口
// @tags    动漫
// @produce json
// @param   Type formData string true "动漫类型"
// @param   Name formData string true "动漫名字"
// @router  /anime/delete [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func DeleteAnime(r *ghttp.Request)  {
	var data *anime.DeleteAnimeInput
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, response.FAIL, err.Error())
	}
	if err := anime.DeleteAnime(data); err != nil {
		response.JsonExit(r, response.FAIL, err.Error())
	} else {
		response.JsonExit(r, response.SUCCESS, "动漫删除成功")
	}
}

// @summary 动漫添加接口
// @tags    动漫
// @produce json
// @param   Type formData string true "动漫类型"
// @param   Name formData string true "动漫名字"
// @param   NewName formData string true "新动漫名字"
// @param   NewLink formData string true "新动漫类型"
// @router  /anime/update [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func UpdateAnime(r *ghttp.Request)  {
	var data *anime.UpdateAnimeInput
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, response.FAIL, err.Error())
	}
	if err := anime.UpdateAnime(data); err != nil {
		response.JsonExit(r, response.FAIL, err.Error())
	} else {
		response.JsonExit(r, response.SUCCESS, "动漫更新成功")
	}
}
